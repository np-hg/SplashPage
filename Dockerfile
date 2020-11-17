FROM golang:1.14-alpine

WORKDIR /app
COPY SplashServer/ .
RUN apk add --no-cache git
RUN go build -o /opt/splashserver

# Prometheus
EXPOSE 9100
# Application
EXPOSE 8080
ENTRYPOINT ["/opt/splashserver"]
