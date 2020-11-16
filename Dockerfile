FROM golang:1.14-alpine as builder

RUN adduser -D -g 'splashuser' splashuser
WORKDIR /app
COPY SplashServer/ .
RUN apk add --no-cache git
RUN go build -o /opt/splashserver

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /opt/splashserver /opt/splashserver
VOLUME /config

# Prometheus
EXPOSE 9100
# Application
EXPOSE 8080
USER splashuser
ENTRYPOINT ["/opt/splashserver"]
