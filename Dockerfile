FROM golang:latest as builder
LABEL maintainer="rohitha.k95@gmail.com"
ENV TASKLOGGER_CONF_FILE=go/src/github.com/tb/task-logger/common-packages/conf
# Copy the source from the current directory to the Working Directory inside the container
COPY . /go/src/github.com/tb/task-logger/task-app-1/
WORKDIR /go/src/github.com/tb/task-logger/task-app-1/
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
ENV PORT 8098
RUN go build -o /go/src/github.com/tb/task-logger/task-app-1/main.go .
# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build_base /tmp/task-logger/taskapp /go/src/github.com/tb/task-logger/task-app-1
# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
CMD ["/go/src/github.com/tb/task-logger/taskapp"]