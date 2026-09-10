#!/usr/bin/env bash
#
# fetch_yang.sh - pull the real Nokia LightSpan YANG modules for a board.
#
# yang/ is gitignored, so this script is the reproducible way to restore it.
# The docserver is internal-only; run this from the Nokia network.
#
# The lab LT is FWLT-C (FX platform), which is the board whose confd serves the
# ONU "mounted" YANG (bbf-fiber-onu-emulated-mount and friends). That mounted
# tree is the compile target, so fwlt-c is the default.
#
# Usage:
#   ./fetch_yang.sh              # fwlt-c
#   BOARD=lllt-a ./fetch_yang.sh # another board
set -euo pipefail

BOARD="${BOARD:-fwlt-c}"
DOCSERVER="${DOCSERVER:-http://docserver.fnbba.dyn.nesc.nokia.net/yang_model}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
YANG_DIR="$(cd "${SCRIPT_DIR}/../../yang" && pwd)"

TAR_URL="${DOCSERVER}/yang/${BOARD}/default/yang_device.tar"
TAR_FILE="${YANG_DIR}/${BOARD}_yang_device.tar"
OUT_DIR="${YANG_DIR}/${BOARD}"

echo "[fetch-yang] board=${BOARD}"
echo "[fetch-yang] GET ${TAR_URL}"

mkdir -p "${OUT_DIR}"
curl -fsS --max-time 600 -o "${TAR_FILE}" "${TAR_URL}"
tar -xf "${TAR_FILE}" -C "${OUT_DIR}"

COUNT=$(find "${OUT_DIR}" -name '*.yang' | wc -l)
echo "[fetch-yang] extracted ${COUNT} .yang modules into ${OUT_DIR}"

# The compile target must be present, otherwise the board is the wrong one.
if [ ! -f "${OUT_DIR}/bbf-fiber-onu-emulated-mount.yang" ]; then
  echo "[fetch-yang] ERROR: bbf-fiber-onu-emulated-mount.yang missing." >&2
  echo "[fetch-yang] Board ${BOARD} does not serve the ONU mounted tree." >&2
  exit 1
fi

echo "[fetch-yang] done."
