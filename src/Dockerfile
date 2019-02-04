FROM golang:latest
MAINTAINER ahmedhosman@live.com
RUN mkdir /app
COPY . /app
RUN apt -y update && apt -y install git
RUN go get github.com/gosimple/slug
RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/russross/blackfriday.v2
RUN go get github.com/gin-gonic/contrib/sessions
WORKDIR /app