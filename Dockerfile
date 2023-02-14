FROM golang:alpine as builder

WORKDIR /GoCensorSvc

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine

WORKDIR /GoCensorSvc

COPY --from=builder /GoCensorSvc/main /GoCensorSvc/main
COPY --from=builder /GoCensorSvc/pkg/config/envs/*.env /GoCensorSvc/

RUN chmod +x /GoCensorSvc/main

CMD ["./main"]
