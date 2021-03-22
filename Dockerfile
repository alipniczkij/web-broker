FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o web-broker ./cmd/main.go

CMD ["./web-broker"]