# syntax=docker/dockerfile:1

FROM golang:1.18

# Set destination for COPY
WORKDIR /app

COPY . ./

RUN go mod download


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /pack_delivery_api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

FROM postgres
ENV POSTGRES_PASSWORD pack_delivery_api_password
ENV POSTGRES_DB pack_delivery_api_db
COPY world.sql /docker-entrypoint-initdb.d/

# Run
CMD ["/pack_delivery_api"]