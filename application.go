package main

import (
	"fmt"
	"github.com/gosimple/slug"
	"io/ioutil"
	"microservice-task/utils"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
	"microservice-task/queue"
	"strconv"
)
type Application struct{
	QueueProvider queue.Provider
}

func (a *Application) Run(mode string){
	defer a.QueueProvider.Cleanup()

	switch mode{
	case "server":
		portInt, err := strconv.Atoi(port)
		utils.FailOnError(err, "Port number must be numeric")
		a.StartServer(portInt)
	case "consume":
		consumer := NewUrl2PdfConsumer(outputDir, program)
		a.QueueProvider.Consume(consumer)
	default:
		fmt.Println("Valid options are: [server | consume]")
	}
}

func (a *Application) SetQueueProvider(provider queue.Provider){
	a.QueueProvider = provider
}

func (a *Application) StartServer(port int) {
	fmt.Println("starting url2pdf microserver...")
	http.HandleFunc("/", a.urlToPdfService)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *Application) urlToPdfService(w http.ResponseWriter, r *http.Request) {
	jobUrl := r.FormValue("url")
	//check parameter passed is a valid url
	if x, err := url.Parse(jobUrl); err == nil {
		if len(x.Host) > 0 {
			fmt.Println("Queuing up " + jobUrl + " for pdf conversion")
			job, err := a.createJob(jobUrl)
			if err == nil {
				a.QueueProvider.Publish(job)

				//time PDF conversion. If it takes more than 30 seconds
				//bail out and show a failed message to user
				timer := time.NewTimer(30 * time.Second)
				for {
					select {
					case <-timer.C:
						w.Header().Set("Content-Type", "text/html")
						w.Write([]byte(jobFailed()))
						return
					default:
						pdfData, err := ioutil.ReadFile(job.Pdf)
						if err == nil {
							fmt.Println("Serving up pdf to browser")
							w.Header().Set("Content-Type", "application/pdf")
							w.Write(pdfData)
							return
						}

						fmt.Println(err.Error())
						time.Sleep(1 * time.Second)
					}
				}
			}

			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(jobFailed()))
		}
	}
}

func (a *Application) createJob(url string) (*Url2PdfJob, error) {
	//catch error and inform user
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			jobHash := utils.Md5Hash(string(body))
			jobName := slug.Make(url)
			jobPdfName := jobHash + ".pdf"
			tempJobPdfName := jobHash + ".pdfx"

			//full path to pdf file
			pdfFile := filepath.Join(outputDir, jobName, jobPdfName)

			fmt.Println("Job Hash: " + jobHash)
			fmt.Println("Job Output File: " + jobPdfName)

			job := &Url2PdfJob{
				Name: jobName,
				Url:  url,
				Pdf:  pdfFile,
				TPdf: tempJobPdfName,
			}

			return job, nil
		}
	}

	return nil, err
}

func jobFailed() string {
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Url2Pdf MicroService</title>
	</head>
	<body>

	<div style="text-align: center">
	<h1>Sorry PDF could not be generated at this time. Please try again later</h1>
	</div>

	</body>
	</html>
`
	return html
}
