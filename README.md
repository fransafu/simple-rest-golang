# simple-rest-golang

This is a simple REST API in Golang

# Get started

## Run in your computer

* go get -d -v ./...
* go install -v ./...
* go run main.go

## Run with docker

The Dockerfile contains configurations like the port, if you want change de port, you must go to the Dockerfile and change

* docker build -t fransafu/simple-rest-golang .
* docker run -it --rm --name simple-rest-golang fransafu/simple-rest-golang

Repository for Golang images: https://hub.docker.com/_/golang

Maybe you are a corious and you want looking at into Docker container, just run the follow command line:

* docker run -it --rm --name simple-rest-golang fransafu/simple-rest-golang bash

Now you are into a docker. Run the any command for example `ls`
