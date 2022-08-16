FROM golang:1.17-alpine3.15 as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

##FROM gcr.io/distroless/static-debian11
##COPY --from=builder /app /app

CMD ["./main"]