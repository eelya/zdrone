FROM golang:1.12 as builder
WORKDIR /app
ADD . /app
RUN go mod download && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./... && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o runme ./cmd

FROM alpine:latest
COPY --from=builder /app/runme /bin 
EXPOSE 8989
CMD ["/bin/runme"]

