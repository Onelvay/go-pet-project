FROM golang:latest
WORKDIR bookstoreass/
COPY ./ ./

ENV PORT=8080

RUN go build -o main .
CMD ["./main"]