#builder
FROM golang:alpine AS builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /home
COPY . .
RUN go build -o invoice-challenge main.go

#final image
FROM alpine

COPY --from=builder /home/invoice-challenge .
EXPOSE 6005
CMD ./invoice-challenge