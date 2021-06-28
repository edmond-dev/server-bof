##
## Build
##

FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

#Run stage
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 7001/tcp

CMD [ "/app/main" ]