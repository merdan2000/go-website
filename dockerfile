FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make entrypoint.sh executable
RUN chmod +x entrypoint.sh

# build go app
RUN go mod download
RUN go build -o innowise_web ./cmd/main.go

CMD ./entrypoint.sh