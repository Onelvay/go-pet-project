FROM golang:latest
WORKDIR app/
COPY ./ ./

ENV PORT=8080

RUN go build -o main .
CMD ["./main"]