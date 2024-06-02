FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /plugin-registry cmd/registry/main.go

EXPOSE 8081

CMD ["/plugin-registry"]
