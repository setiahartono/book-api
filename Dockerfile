FROM golang:1.18
WORKDIR /code
ADD . ./
RUN go mod tidy
RUN go mod download