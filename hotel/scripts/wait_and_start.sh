#!/bin/bash

while ! nc -z $INNER_POSTGRES_HOST $POSTGRES_PORT; do
  sleep 0.1
done

echo "POSTGRES Started"

exec ./hotel --port=$HOTEL_REST_PORT --host=$HOTEL_HOST

exec "$@"