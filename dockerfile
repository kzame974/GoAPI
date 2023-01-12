FROM golang:1.20 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/kzame974/GoAPI
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/GoAPI .

# Second stage
FROM alpine
COPY --from=builder /go/bin/GoAPI /go/bin/GoAPI

EXPOSE 8083

CMD ["/go/bin/GoAPI"]