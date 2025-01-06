FROM alpine:latest

ARG TARGETPLATFORM

EXPOSE 8080 9500 6060

LABEL maintainer="Maxim Bogdanov <muravlion@gmail.com>"
LABEL service.name=archivarius
LABEL service.weight=1
LABEL service.port=8080
LABEL service.check.http=http://{{address}}/healthcheck
LABEL service.check.interval=5s
LABEL service.check.timeout=2s
LABEL service.public=false

COPY .build/${TARGETPLATFORM}/archivarius /archivarius
COPY deploy/develop/config.yaml /config.yaml
COPY deploy/migrations/stats /deploy/migrations/stats

ENTRYPOINT [ "/archivarius","archivarius" ]

