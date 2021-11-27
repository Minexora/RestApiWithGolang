FROM golang:1.17-alpine3.13

ENV PYTHONUNBUFFERED 1
RUN mkdir /code
WORKDIR /code
COPY ./RestApiForGo/ /code/
RUN go get github.com/Sirupsen/logrus
RUN gom install 