FROM golang:1.12-alpine as builder

RUN adduser -D -g 'splashuser' splashuser
WORKDIR /app
COPY SplashServer SplashServer
RUN apk add --no-cache git
RUN go build

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/opt/splashuser /opt/splashuser
VOLUME /data

# Our chosen default for Prometheus
EXPOSE 8080
USER splashuser
ENTRYPOINT ["/opt/splashuser"]