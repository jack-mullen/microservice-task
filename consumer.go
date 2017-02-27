package main

import (
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
	"microservice-task/queue"
	"encoding/json"
)

type Url2PdfJob struct{
	Name string //the slugged url used to create a subfolder for where to store the pdf
	Url string //the url to generate pdf from
	Pdf string //the basename of the pdf file to generate
}

type Url2PdfConsumer struct {
	outputDir string
	program   string
}

func NewUrl2PdfConsumer(outputDir string, program string) *Url2PdfConsumer{
	consumer := &Url2PdfConsumer{
		outputDir: outputDir,
		program: program,
	}

	if _, err := os.Stat(consumer.outputDir); err != nil {
		os.Mkdir(consumer.outputDir, os.ModePerm)
	}

	return consumer
}

func (c Url2PdfConsumer) Process(msg *queue.Message) bool {

	job := &Url2PdfJob{}
	json.Unmarshal([]byte(msg.Body), job)

	//directory where pdf file will be stored
	jobDir := filepath.Join(c.outputDir, job.Name)
	os.MkdirAll(jobDir, os.ModePerm)


	if _, err := os.Stat(job.Pdf); err != nil {
		//invoke html to pdf binary
		fmt.Println("Executing:", c.program, job.Url, job.Pdf)
		cmd := exec.Command(c.program, job.Url, job.Pdf)
		err := cmd.Run()
		if err != nil{
			return false
		}

	} else {
		fmt.Println("Job found in cache")
	}

	return true
}


