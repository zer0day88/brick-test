FROM golang:1.21 AS builder
RUN mkdir -p /app
ENV TZ=Asia/Jakarta
RUN GOCACHE=OFF
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY . /app
RUN cd /app && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/pg ./cmd/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/bin/pg .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY config/config-docker.yml config/config.yml
ENV TZ=Asia/Jakarta

ENTRYPOINT ["/root/pg"]
EXPOSE 3000