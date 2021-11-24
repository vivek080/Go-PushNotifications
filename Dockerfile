# Dockerfile for local development
# Setting official golang image tag 1.15.2 as baseimage 
FROM golang:1.15.2
VOLUME /app
RUN go get github.com/githubnemo/CompileDaemon
WORKDIR /app
EXPOSE 5000
CMD ["go","run","main.go"]
