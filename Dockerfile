# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk --update upgrade
RUN apk add bash wget git
RUN apk add build-base gcc
RUN apk add postgresql-client

# Source volumes ==================================================================================
RUN rm -fr /go/src/github.com/dhurimkelmendi/pack_delivery_api-api
RUN mkdir -p /go/src/github.com/dhurimkelmendi/pack_delivery_api-api
VOLUME /go/src/github.com/dhurimkelmendi/pack_delivery_api-api

# Copy entry/support files ========================================================================
COPY .circleci/docker/entry.sh /entry.sh
RUN chmod +x /entry.sh

# Cleanup =========================================================================================
RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["/entry.sh"]