#!/usr/bin/env bash
set -euo pipefail

# Default values keep the container runnable out of the box (docker-compose network names).
export SERVER_RUN_MODE="${SERVER_RUN_MODE:-release}"
export SERVER_HTTP_PORT="${SERVER_HTTP_PORT:-8080}"
export SERVER_ADDR="${SERVER_ADDR:-0.0.0.0:8080}"

export MINIO_ENDPOINT="${MINIO_ENDPOINT:-minio:9000}"
export MINIO_ACCESS_KEY="${MINIO_ACCESS_KEY:-eonu}"
export MINIO_SECRET_KEY="${MINIO_SECRET_KEY:-eonu#1234}"
export MINIO_SECURE="${MINIO_SECURE:-false}"
export MINIO_BUCKET="${MINIO_BUCKET:-olt-logs}"

export MYSQL_USER="${MYSQL_USER:-iop}"
export MYSQL_ADDR="${MYSQL_ADDR:-mysql:3306}"
export MYSQL_PASSWORD="${MYSQL_PASSWORD:-eonu#1234}"
export MYSQL_DATABASE="${MYSQL_DATABASE:-IOP_Library}"

export KAFKA_ADDR="${KAFKA_ADDR:-kafka:9092}"
export KAFKA_TOPIC_REQUEST="${KAFKA_TOPIC_REQUEST:-collectorRequest}"
export KAFKA_GROUP_REQUEST="${KAFKA_GROUP_REQUEST:-group-collector-request}"
export KAFKA_TOPIC_RESPONSE="${KAFKA_TOPIC_RESPONSE:-collectorResponse}"
export KAFKA_GROUP_RESPONSE="${KAFKA_GROUP_RESPONSE:-group-collector-response}"

TEMPLATE="./config/config.yaml.template"
TARGET="./config/config.yaml"

envsubst < "${TEMPLATE}" > "${TARGET}"

echo "[entrypoint] generated ${TARGET}:"
cat "${TARGET}"

# Optionally wait for MySQL/Kafka to be reachable before starting.
if [ "${WAIT_FOR_DEPS:-true}" = "true" ]; then
  mysql_host="${MYSQL_ADDR%%:*}"
  mysql_port="${MYSQL_ADDR##*:}"
  kafka_host="${KAFKA_ADDR%%:*}"
  kafka_port="${KAFKA_ADDR##*:}"
  for target in "${mysql_host}:${mysql_port}" "${kafka_host}:${kafka_port}"; do
    host="${target%%:*}"
    port="${target##*:}"
    echo "[entrypoint] waiting for ${host}:${port} ..."
    for _ in $(seq 1 60); do
      if (echo > "/dev/tcp/${host}/${port}") >/dev/null 2>&1; then
        echo "[entrypoint] ${host}:${port} is up"
        break
      fi
      sleep 2
    done
  done
fi

exec ./omciAnalyzer
