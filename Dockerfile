# Builder
FROM golang:alpine as builder
ARG APP_VERSION
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -ldflags "-X main.AppVersion=$APP_VERSION" -o rogerwilco .
# Image
FROM alpine:latest
RUN adduser -S -D -H -h /app app
USER app
COPY --from=builder /build/static /app/static
COPY --from=builder /build/templates /app/templates
COPY --from=builder /build/rogerwilco /app/
WORKDIR /app
CMD ["./rogerwilco"]
