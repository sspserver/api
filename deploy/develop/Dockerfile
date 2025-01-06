FROM alpine:latest

ARG TARGETPLATFORM

EXPOSE 8080 6060

LABEL maintainer="Dmitry Ponomarev <demdxx@gmail.com>"
LABEL service.name=adnet.api

ENV SERVER_HTTP_LISTEN=:8080
ENV SERVER_GRPC_LISTEN=tcp://:8081
ENV SERVER_PROFILE_MODE=net
ENV SERVER_PROFILE_LISTEN=8082

COPY .build/${TARGETPLATFORM}/api /api
COPY deploy/migrations /data/migrations
COPY submodules/blaze-api/migrations/initial /data/migrations/initial

ENTRYPOINT [ "/api" ]
