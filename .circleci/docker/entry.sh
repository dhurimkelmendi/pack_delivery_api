#!/bin/bash

# Check if a particular host/port is ready to accept TCP connection.
# This is used to ensure we wait for dependent services to be ready before
# starting this service itself.
waitForHost() {
  local host="$1" port="$2"

  while ! exec 6<>/dev/tcp/$host/$port; do
    echo "!!! Still waiting for host: $host:$port ..."
    sleep 5
  done
  exec 6>&-

  echo "!!! Host: $host:$port is now ready"
}

if [ -z "$DB_NAME" ]; then DB_NAME=story_builder_db; fi
if [ -z "$DB_ADDR" ]; then DB_ADDR=localhost; fi
if [ -z "$DB_PORT" ]; then DB_PORT=5432; fi
if [ -z "$DB_USERNAME" ]; then DB_USERNAME=story_builder_backend; fi
if [ -z "$DB_PASSWORD" ]; then DB_PASSWORD=""; fi
if [ -z "$HTTP_ADDR" ]; then HTTP_ADDR=":8888"; fi

# Wait for Postgres server
if ! $(exec 6<>/dev/tcp/$DB_ADDR/$DB_PORT); then
  echo "!!! Waiting for postgres server ..."
  waitForHost $DB_ADDR $DB_PORT
fi
echo "... Postgres server is ready."

if [ -f "/go/src/github.com/dhurimkelmendi/pack_delivery_api/main.go" ]; then
  # If source volume is mounted, then run from there ...
  echo "!!! Change directory to /go/src/github.com/dhurimkelmendi/pack_delivery_api ..."
  cd /go/src/github.com/dhurimkelmendi/pack_delivery_api

  if [ -n "$TEST_DB_NAME" ]; then
    echo "!!! Dropping test database (if exist) ..."
    psql "postgresql://$DB_USERNAME:$DB_PASSWORD@$DB_ADDR:$DB_PORT/$DB_NAME" -c "DROP DATABASE IF EXISTS \"$TEST_DB_NAME\";"

    echo "!!! Creating test database ..."
    psql "postgresql://$DB_USERNAME:$DB_PASSWORD@$DB_ADDR:$DB_PORT/$DB_NAME" -c "CREATE DATABASE \"$TEST_DB_NAME\";"

    echo "!!! Resetting test database migrations ..."
    KINZOO_ENV=test DB_NAME="$TEST_DB_NAME" go run main.go migrate reset
    KINZOO_ENV=test DB_NAME="$TEST_DB_NAME" go run main.go migrate up
  fi

  if [ "$RUN_FROM_SOURCE" == "true" ]; then
    if [ "$RESET_DB" == "true" ]; then
      echo "!!! Resetting migrations ..."
      go run main.go migrate reset
    fi

    echo "!!! Running migrations ..."
    go run main.go migrate up

    echo "!!! Running from source hot reload..."
    go run main.go
  else
    echo "!!! Building server binary ..."
    go build -o /pack_delivery_api
    cd /

    if [ "$RESET_DB" == "true" ]; then
      echo "!!! Resetting migrations ..."
      ./pack_delivery_api migrate reset
    fi

    echo "!!! Running migrations ..."
    ./pack_delivery_api migrate up

    echo "!!! Running API server binary ..."
    ./pack_delivery_api
  fi
else
  # ... else just do nothing but keep the server running.
  while true; do
    sleep 1
  done
fi
