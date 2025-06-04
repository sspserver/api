# syntax=docker/dockerfile:1.2
FROM --platform=$TARGETPLATFORM scratch

EXPOSE 8080

ARG TARGETPLATFORM

LABEL maintainer="Dmitry Ponomarev <demdxx@gmail.com>"
LABEL service.name=adnet.api
LABEL service.weight=1
LABEL service.port=8080
LABEL service.check.http=http://{{address}}/healthcheck
LABEL service.check.interval=5s
LABEL service.check.timeout=2s
LABEL service.public=false

ENV SERVER_HTTP_LISTEN=:8080
ENV SERVER_GRPC_LISTEN=tcp://:8081
ENV SERVER_PROFILE_MODE=net
ENV SERVER_PROFILE_LISTEN=:8082

COPY .build/${TARGETPLATFORM}/api /api
COPY deploy/migrations /data/migrations
COPY submodules/blaze-api/migrations/initial /data/migrations/initial

ENTRYPOINT [ "/api"]
