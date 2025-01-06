FROM alpine:latest

ARG TARGETPLATFORM

EXPOSE 8080 9500 6060

LABEL maintainer="Maxim Bogdanov <muravlion@gmail.com>"
LABEL service.name=archivarius

COPY .build/${TARGETPLATFORM}/archivarius /archivarius
COPY deploy/develop/config.yaml /config.yaml
COPY deploy/migrations/stats /deploy/migrations/stats

ENTRYPOINT [ "/archivarius","archivarius" ]
