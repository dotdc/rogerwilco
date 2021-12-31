# Builder
FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o rogerwilco .
# Image
FROM alpine:latest
RUN adduser -S -D -H -h /app app
USER app
COPY --from=builder /build/rogerwilco /app/
WORKDIR /app
CMD ["./rogerwilco"]
