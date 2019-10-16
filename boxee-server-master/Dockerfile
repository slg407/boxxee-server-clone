FROM golang:alpine AS builder
WORKDIR /src
RUN apk add --no-cache git
ADD go.mod /src/go.mod
ADD go.sum /src/go.sum
RUN go mod download
ADD . /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o boxee-server

FROM jpillora/dnsmasq

ENV BOXEE_SERVER_PORT=80
ENV EXTERNAL_IP=""
RUN apk add --no-cache gettext ca-certificates
ADD docker/start.sh /start.sh
ADD docker/dnsmasq.conf.template /etc/dnsmasq.conf.template
ADD docker/hosts.dnsmasq.template /etc/hosts.dnsmasq.template
COPY --from=builder /src/boxee-server /boxee-server
RUN chmod +x /start.sh

ENTRYPOINT ["/start.sh"]
