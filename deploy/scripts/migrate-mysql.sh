#!/usr/bin/env bash
#
# migrate-mysql.sh - export the IOP_Library database so the dockerized MySQL can
# restore it automatically on first boot.
#
# The dump is written to deploy/mysql/init/IOP_Library.sql, which is mounted into
# the mysql container's /docker-entrypoint-initdb.d and executed on first init.
#
# Two source modes:
#   1) Remote (default): dump from the running MySQL on the 238 server over SSH.
#   2) Local: dump from a MySQL reachable on this host (SRC_SSH="").
#
# Usage:
#   ./migrate-mysql.sh                       # remote dump from 238 via SSH
#   SRC_SSH="" ./migrate-mysql.sh            # local dump (run on the server)
#
# Environment overrides:
#   SRC_SSH        SSH target (default: fnmst001@10.101.15.238); empty => local
#   SSH_PORT       SSH port (default: 22)
#   SRC_DB_USER    MySQL user   (default: iop)
#   SRC_DB_PASS    MySQL pass   (default: eonu#1234)
#   SRC_DB_NAME    Database     (default: IOP_Library)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUT_DIR="$(cd "${SCRIPT_DIR}/../mysql/init" && pwd)"
OUT_FILE="${OUT_DIR}/IOP_Library.sql"

SRC_SSH="${SRC_SSH-fnmst001@10.101.15.238}"
SSH_PORT="${SSH_PORT:-22}"
SRC_DB_USER="${SRC_DB_USER:-iop}"
SRC_DB_PASS="${SRC_DB_PASS:-eonu#1234}"
SRC_DB_NAME="${SRC_DB_NAME:-IOP_Library}"

# --single-transaction: consistent snapshot without locking (InnoDB).
# --no-tablespaces: avoids needing the PROCESS privilege (non-root dump).
# --databases: includes CREATE DATABASE / USE so restore is self-contained.
DUMP_CMD="mysqldump -u${SRC_DB_USER} -p'${SRC_DB_PASS}' \
  --single-transaction --routines --triggers --events \
  --no-tablespaces --databases ${SRC_DB_NAME}"

echo "[migrate] target output: ${OUT_FILE}"

if [ -n "${SRC_SSH}" ]; then
  echo "[migrate] dumping ${SRC_DB_NAME} from ${SRC_SSH} (port ${SSH_PORT}) ..."
  ssh -p "${SSH_PORT}" -o BatchMode=yes -o StrictHostKeyChecking=no \
    "${SRC_SSH}" "${DUMP_CMD}" > "${OUT_FILE}"
else
  echo "[migrate] dumping ${SRC_DB_NAME} from local MySQL ..."
  bash -c "${DUMP_CMD}" > "${OUT_FILE}"
fi

if [ ! -s "${OUT_FILE}" ]; then
  echo "[migrate] ERROR: dump file is empty" >&2
  exit 1
fi

BYTES=$(wc -c < "${OUT_FILE}")
TABLES=$(grep -c "^CREATE TABLE" "${OUT_FILE}" || true)
echo "[migrate] wrote ${OUT_FILE} (${BYTES} bytes, ${TABLES} tables)"
echo "[migrate] done. It will be restored automatically on first 'docker compose up'."
