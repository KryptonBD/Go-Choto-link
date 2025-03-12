FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /choto-link

EXPOSE 8080

CMD ["/choto-link"]

