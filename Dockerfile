# Prepare
FROM golang:1.13-alpine as baseimg

RUN apk --no-cache upgrade && apk --no-cache add git make

# First only download the dependencies, so thid layer can be cahced before we copy the code
COPY ./go.mod ./go.sum ./Makefile /app/
WORKDIR /app/

# Build
FROM baseimg as builder

COPY . ./
RUN make build

# Run
FROM alpine

COPY --from=builder /app/maga-golang-test /opt/
WORKDIR /opt/
ARG ENV
EXPOSE 8080
CMD ["./maga-golang-test"]
