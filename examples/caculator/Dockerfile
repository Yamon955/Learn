FROM golang:alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o caculator .

FROM alpine:latest
RUN apk update && \
    apk add vim && \
    apk add net-tools
COPY --from=builder /app/caculator /app/trpc_go.yaml /app/
WORKDIR /app
CMD ["./caculator"]