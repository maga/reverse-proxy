FROM golang:latest as builder
WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

COPY . ./
RUN make build

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

CMD ./main
EXPOSE 8080
