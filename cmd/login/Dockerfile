FROM alpine:latest
WORKDIR /app
COPY ./cmd/management/management_server /app
COPY ./configs/config.yml /app/configs/config.yml
CMD ["/app/management_server"]
