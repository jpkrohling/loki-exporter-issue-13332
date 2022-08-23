FROM golang:1.18-alpine 
RUN apk --update add ca-certificates
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY gen .
RUN go build -v -o /otelcol-contrib ./...

COPY otelcol.yaml /etc/otelcol-contrib/config.yaml
ENTRYPOINT ["/otelcol-contrib"]
CMD ["--config", "/etc/otelcol-contrib/config.yaml"]
EXPOSE 4317
