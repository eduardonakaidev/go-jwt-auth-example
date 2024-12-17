FROM golang:1.19

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 3000
CMD ["./main"]
