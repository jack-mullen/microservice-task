#!/usr/bin/env bash

#install glide
curl https://glide.sh/get | sh

#install libx dependencies
apt-get update
apt-get install -y libxrender1 libfontconfig1 libxext6

cd /go/src/microservice-task

#download dependencies using glide
glide up

#build and install micro service
go install

#make html2pdf program global
cp bin/wkhtmltopdf /usr/bin/wkhtmltopdf
chmod +x /usr/bin/wkhtmltopdf

#run a single consumer
microservice-task consume &> /var/log/url2pdf.log &

#run the micro service in the background
microservice-task server &> /var/log/url2pdf.log &

tail -f  /var/log/url2pdf.log
