#!/usr/bin/env bash
#
# run_tests.sh - end-to-end verification for a running iop-toolkit stack.
#
# Covers:
#   1. Infrastructure health  (MySQL tables/data, MinIO bucket, Kafka topics)
#   2. Backend end-to-end     (upload -> parse -> OMCI JSON / PlantUML / diagram)
#   3. Frontend reverse proxy (SPA index + /api proxied to backend)
#
# Usage:
#   ./run_tests.sh
#
# Assumes `docker compose up -d` has completed and the .env in deploy/ is present.
set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DEPLOY_DIR="$(cd "${SCRIPT_DIR}/../.." && pwd)"
ROOT_DIR="$(cd "${DEPLOY_DIR}/.." && pwd)"

# Load deployment configuration (ports / credentials).
if [ -f "${DEPLOY_DIR}/.env" ]; then
  set -a; . "${DEPLOY_DIR}/.env"; set +a
fi

MYSQL_DATABASE="${MYSQL_DATABASE:-IOP_Library}"
MYSQL_USER="${MYSQL_USER:-iop}"
MYSQL_PASSWORD="${MYSQL_PASSWORD:-eonu#1234}"
MINIO_BUCKET="${MINIO_BUCKET:-olt-logs}"
BACKEND_PORT="${BACKEND_PORT:-8080}"
FRONTEND_PORT="${FRONTEND_PORT:-80}"
KAFKA_TOPIC_REQUEST="${KAFKA_TOPIC_REQUEST:-collectorRequest}"
KAFKA_TOPIC_RESPONSE="${KAFKA_TOPIC_RESPONSE:-collectorResponse}"

BACKEND_URL="http://localhost:${BACKEND_PORT}"
FRONTEND_URL="http://localhost:${FRONTEND_PORT}"
SAMPLE_LOG="${SAMPLE_LOG:-${ROOT_DIR}/backend/static/uploads/omciIsam.log}"
WORKDIR="$(mktemp -d)"

PASS=0
FAIL=0

pass() { echo "  [PASS] $1"; PASS=$((PASS + 1)); }
fail() { echo "  [FAIL] $1"; FAIL=$((FAIL + 1)); }

echo "==================================================================="
echo " iop-toolkit test suite"
echo "   backend : ${BACKEND_URL}"
echo "   frontend: ${FRONTEND_URL}"
echo "   sample  : ${SAMPLE_LOG}"
echo "==================================================================="

# ---------------------------------------------------------------------------
echo ""
echo "[1/3] Infrastructure health"
# ---------------------------------------------------------------------------

# MySQL: expect 7 tables and non-empty seed tables.
TABLE_COUNT=$(docker exec iop-mysql sh -c \
  "mysql -u${MYSQL_USER} -p'${MYSQL_PASSWORD}' -N -B -e \
   'SELECT COUNT(*) FROM information_schema.tables WHERE table_schema=\"${MYSQL_DATABASE}\";'" 2>/dev/null)
if [ "${TABLE_COUNT}" = "7" ]; then
  pass "MySQL ${MYSQL_DATABASE} has 7 tables"
else
  fail "MySQL ${MYSQL_DATABASE} table count = '${TABLE_COUNT}' (expected 7)"
fi

for tbl in vendors customers iop_library_records; do
  ROWS=$(docker exec iop-mysql sh -c \
    "mysql -u${MYSQL_USER} -p'${MYSQL_PASSWORD}' -N -B -e \
     'SELECT COUNT(*) FROM ${MYSQL_DATABASE}.${tbl};'" 2>/dev/null)
  if [ -n "${ROWS}" ] && [ "${ROWS}" -gt 0 ] 2>/dev/null; then
    pass "MySQL table ${tbl} has ${ROWS} rows"
  else
    fail "MySQL table ${tbl} is empty or missing (rows='${ROWS}')"
  fi
done

# MinIO: bucket directory exists inside the container.
if docker exec iop-minio sh -c "ls -d /data/${MINIO_BUCKET}" >/dev/null 2>&1; then
  pass "MinIO bucket ${MINIO_BUCKET} exists"
else
  fail "MinIO bucket ${MINIO_BUCKET} missing"
fi

# Kafka: required topics present.
TOPICS=$(docker exec iop-kafka kafka-topics --bootstrap-server localhost:9092 --list 2>/dev/null)
for topic in "${KAFKA_TOPIC_REQUEST}" "${KAFKA_TOPIC_RESPONSE}"; do
  if echo "${TOPICS}" | grep -qx "${topic}"; then
    pass "Kafka topic ${topic} exists"
  else
    fail "Kafka topic ${topic} missing"
  fi
done

# ---------------------------------------------------------------------------
echo ""
echo "[2/3] Backend end-to-end (toolkit_client pipeline)"
# ---------------------------------------------------------------------------
if [ ! -f "${SAMPLE_LOG}" ]; then
  fail "sample log not found: ${SAMPLE_LOG}"
else
  if python3 "${SCRIPT_DIR}/toolkit_client.py" --url "${BACKEND_URL}" pipeline \
      --file "${SAMPLE_LOG}" \
      --download-plantuml --require-plantuml --download-diagram \
      --output-dir "${WORKDIR}" > "${WORKDIR}/pipeline.log" 2>&1; then
    OMCI_COUNT=$(find "${WORKDIR}" -name 'omci_*.json' | wc -l | tr -d ' ')
    WSD_COUNT=$(find "${WORKDIR}" -name 'plantuml_*.wsd' | wc -l | tr -d ' ')
    SVG_COUNT=$(find "${WORKDIR}" -name 'diagram_*.svg' | wc -l | tr -d ' ')
    [ "${OMCI_COUNT}" -gt 0 ] && pass "OMCI JSON produced (${OMCI_COUNT} file(s))" || fail "no OMCI JSON produced"
    [ "${WSD_COUNT}" -gt 0 ] && pass "PlantUML WSD produced (${WSD_COUNT} file(s))" || fail "no PlantUML WSD produced"
    [ "${SVG_COUNT}" -gt 0 ] && pass "Diagram SVG produced (${SVG_COUNT} file(s))" || fail "no diagram SVG produced"
  else
    fail "pipeline failed (see below)"
    sed 's/^/      /' "${WORKDIR}/pipeline.log"
  fi
fi

# ---------------------------------------------------------------------------
echo ""
echo "[3/3] Frontend reverse proxy"
# ---------------------------------------------------------------------------
INDEX_CODE=$(curl -s -o "${WORKDIR}/index.html" -w '%{http_code}' "${FRONTEND_URL}/")
if [ "${INDEX_CODE}" = "200" ] && grep -qi 'id="root"\|<title' "${WORKDIR}/index.html"; then
  pass "Frontend SPA index served (HTTP ${INDEX_CODE})"
else
  fail "Frontend index not served correctly (HTTP ${INDEX_CODE})"
fi

PROXY_CODE=$(curl -s -o /dev/null -w '%{http_code}' "${FRONTEND_URL}/api/omcianalyzer/counters")
if [ "${PROXY_CODE}" = "200" ]; then
  pass "Frontend /api reverse proxy reaches backend (HTTP ${PROXY_CODE})"
else
  fail "Frontend /api proxy failed (HTTP ${PROXY_CODE})"
fi

# ---------------------------------------------------------------------------
rm -rf "${WORKDIR}"
echo ""
echo "==================================================================="
echo " RESULT: ${PASS} passed, ${FAIL} failed"
echo "==================================================================="
[ "${FAIL}" -eq 0 ]
