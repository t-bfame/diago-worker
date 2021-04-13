FROM golang:1.15 as builder

WORKDIR /src/diago-worker

COPY go.mod go.sum ./

RUN go mod verify

COPY cmd cmd
COPY pkg pkg
COPY proto-gen proto-gen

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /src/diago-worker/main .
ENTRYPOINT ["./main"]