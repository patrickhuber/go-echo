FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /echoapp

FROM ubuntu:latest

COPY --from=build /echoapp /echoapp

EXPOSE 8080

ENTRYPOINT [ "/echoapp" ]