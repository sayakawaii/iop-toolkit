#!/usr/bin/env python3
"""iop-toolkit client — upload logs, poll status, download OMCI JSON.

All endpoints confirmed from api-catalog.json (2026-03-31).
Run `python toolkit_client.py --help` to see available subcommands.
"""

import argparse
import json
import os
import sys
import time
from urllib.error import HTTPError, URLError
from urllib.parse import urlencode, urljoin
from urllib.request import Request, urlopen

_HTTP_ERROR_FMT = "HTTP {} {}: {}"
_REACH_ERROR_FMT = "Cannot reach {}: {}"

DEFAULT_URL = "http://10.101.15.238:8080"
POLL_INTERVAL = 3   # seconds between status polls
POLL_TIMEOUT = 300  # seconds before giving up

# Confirmed endpoints from api-catalog.json (2026-03-31).
ENDPOINTS = {
    "upload":          "/api/omcianalyzer/request",
    "status":          "/api/omcianalyzer/progress",   # POST multipart
    "list_onus":       "/api/omcianalyzer/onus",
    "get_omci":        "/api/omcianalyzer/omci",
    "get_omci_detail": "/api/omcianalyzer/omci/detail",
    "diagram":         "/api/omcianalyzer/diagram",
    "plantuml":        "/api/omcianalyzer/plantuml",
}

# Confirmed multipart field name from api-catalog.json.
DEFAULT_UPLOAD_FIELD = "omcianalyzerFile"


def _require_endpoint(key):
    ep = ENDPOINTS.get(key)
    if ep is None:
        raise NotImplementedError(
            "Endpoint '{}' is not configured yet. "
            "Update ENDPOINTS in toolkit_client.py after the API is available "
            "(see iop-toolkit/references/toolkit-api.md).".format(key)
        )
    return ep


# ---------------------------------------------------------------------------
# HTTP helpers
# ---------------------------------------------------------------------------

def _http_get(base_url, path, params=None):
    url = urljoin(base_url.rstrip("/") + "/", path.lstrip("/"))
    if params:
        url = "{}?{}".format(url, urlencode(params))
    req = Request(url)
    try:
        resp = urlopen(req)
        return json.loads(resp.read().decode("utf-8"))
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(url, exc.reason))


def _http_post_json(base_url, path, payload):
    url = urljoin(base_url.rstrip("/") + "/", path.lstrip("/"))
    data = json.dumps(payload).encode("utf-8")
    req = Request(url, data=data, headers={"Content-Type": "application/json"})
    try:
        resp = urlopen(req)
        return json.loads(resp.read().decode("utf-8"))
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(url, exc.reason))


def _http_post_form(base_url, path, fields):
    """POST multipart/form-data with simple string key-value fields."""
    url = urljoin(base_url.rstrip("/") + "/", path.lstrip("/"))
    boundary = "---IopToolkitFormBoundary7MA4YWxk"
    body = b""
    for name, value in fields.items():
        body += (
            "--{}\r\nContent-Disposition: form-data; name=\"{}\"\r\n\r\n{}\r\n"
        ).format(boundary, name, value).encode("utf-8")
    body += "--{}--\r\n".format(boundary).encode("utf-8")
    req = Request(
        url,
        data=body,
        headers={"Content-Type": "multipart/form-data; boundary={}".format(boundary)},
    )
    try:
        resp = urlopen(req)
        return json.loads(resp.read().decode("utf-8"))
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(url, exc.reason))


def _http_upload(base_url, path, file_path, field="omcianalyzerFile"):
    """Multipart file upload. Returns parsed JSON response."""
    url = urljoin(base_url.rstrip("/") + "/", path.lstrip("/"))
    boundary = "---IopToolkitBoundary7MA4YWxkTrZu0gW"
    filename = os.path.basename(file_path)
    with open(file_path, "rb") as fh:
        file_bytes = fh.read()
    body = (
        "--{boundary}\r\n"
        "Content-Disposition: form-data; name=\"{field}\"; filename=\"{filename}\"\r\n"
        "Content-Type: application/octet-stream\r\n\r\n"
    ).format(boundary=boundary, field=field, filename=filename).encode("utf-8")
    body += file_bytes
    body += "\r\n--{}--\r\n".format(boundary).encode("utf-8")
    req = Request(
        url,
        data=body,
        headers={"Content-Type": "multipart/form-data; boundary={}".format(boundary)},
    )
    try:
        resp = urlopen(req)
        return json.loads(resp.read().decode("utf-8"))
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(url, exc.reason))


# ---------------------------------------------------------------------------
# API calls
# ---------------------------------------------------------------------------

def api_upload(base_url, file_path, upload_field):
    """Upload a log file. Returns requestKey string."""
    ep = _require_endpoint("upload")
    resp = _http_upload(base_url, ep, file_path, field=upload_field)
    # Response may be a list (one entry) or dict
    if isinstance(resp, list):
        if not resp:
            raise RuntimeError("Upload response is empty list")
        resp = resp[0]
    request_key = resp.get("requestKey") or resp.get("request_key") or resp.get("key") or resp.get("id")
    if not request_key:
        raise RuntimeError("Upload response missing requestKey: {}".format(resp))
    return str(request_key)


def api_status(base_url, request_key):
    """Return status string: 'pending' | 'running' | 'done' | 'failed'.

    Progress API is POST multipart (requestKey as form field, repeatable).
    Response may be a list (one entry per key) or a single dict.
    """
    ep = _require_endpoint("status")
    resp = _http_post_form(base_url, ep, {"requestKey": request_key})
    if isinstance(resp, list):
        entry = resp[0] if resp else {}
    else:
        entry = resp
    status = entry.get("status") or entry.get("progressStatus") or entry.get("state") or "unknown"
    return str(status)


def api_list_onus(base_url, request_key):
    """Return list of ONU dicts from `/api/omcianalyzer/onus` response."""
    ep = _require_endpoint("list_onus")
    resp = _http_get(base_url, ep, {"requestKey": request_key})
    if isinstance(resp, list):
        return resp
    return resp.get("onus") or resp.get("data") or []


def api_get_omci(base_url, request_key, onu_name=None):
    """Return OMCI JSON for a request, optionally filtered by onuName."""
    ep = _require_endpoint("get_omci")
    params = {"requestKey": request_key}
    if onu_name:
        params["onuName"] = onu_name
    return _http_get(base_url, ep, params)


def _http_get_binary(base_url, path, params=None):
    """Fetch binary content (e.g., SVG diagram) and return raw bytes."""
    url = urljoin(base_url.rstrip("/") + "/", path.lstrip("/"))
    if params:
        url = "{}?{}".format(url, urlencode(params))
    req = Request(url)
    try:
        resp = urlopen(req)
        return resp.read()
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(url, exc.reason))


def _http_get_binary_from_absolute_url(absolute_url):
    """Fetch binary content from an absolute URL (no base_url transformation)."""
    req = Request(absolute_url)
    try:
        resp = urlopen(req)
        return resp.read()
    except HTTPError as exc:
        detail = exc.read().decode("utf-8", "replace")
        raise RuntimeError(_HTTP_ERROR_FMT.format(exc.code, absolute_url, detail))
    except URLError as exc:
        raise RuntimeError(_REACH_ERROR_FMT.format(absolute_url, exc.reason))


def _api_get_latest_binary(base_url, endpoint_key, request_key, onu_name=None, artifact_label="artifact"):
    """Fetch latest binary from an endpoint that returns [{path, category}, ...].

    This function:
      1. Fetches the list from configured endpoint
      2. Filters entries where category == 'latest'
      3. Downloads and returns binary from absolute URL in `path`
    Raises RuntimeError if no 'latest' entry is found.
    """
    ep = _require_endpoint(endpoint_key)
    params = {"requestKey": request_key}
    if onu_name:
        params["onuName"] = onu_name
    entries = _http_get(base_url, ep, params)
    if not isinstance(entries, list):
        raise RuntimeError(
            "{} endpoint returned unexpected type {}; expected JSON list.".format(
                artifact_label.capitalize(),
                type(entries).__name__,
            )
        )

    latest_entries = [e for e in entries if e.get("category") == "latest"]
    if not latest_entries:
        categories = list({e.get("category") for e in entries})
        raise RuntimeError(
            "No {} entry with category='latest' found. Available categories: {}. Total entries: {}.".format(
                artifact_label,
                categories,
                len(entries),
            )
        )

    file_url = latest_entries[0].get("path") or latest_entries[0].get("url")
    if not file_url:
        raise RuntimeError("Latest {} entry has no 'path' field: {}".format(artifact_label, latest_entries[0]))
    # The backend returns root-relative paths (e.g. /uploads/<...>); join them
    # with the base URL so they resolve against the same backend origin.
    if not file_url.lower().startswith(("http://", "https://")):
        file_url = base_url.rstrip("/") + "/" + file_url.lstrip("/")
    return _http_get_binary_from_absolute_url(file_url)


def api_get_diagram(base_url, request_key, onu_name=None):
    """Fetch the SVG diagram with category=latest for a request."""
    return _api_get_latest_binary(base_url, "diagram", request_key, onu_name, "diagram")


def api_get_plantuml(base_url, request_key, onu_name=None):
    """Fetch the PlantUML source (.wsd) with category=latest for a request."""
    return _api_get_latest_binary(base_url, "plantuml", request_key, onu_name, "plantuml")


# ---------------------------------------------------------------------------
# Subcommand implementations
# ---------------------------------------------------------------------------

def cmd_upload(args):
    print("Uploading {} ...".format(args.file))
    request_key = api_upload(args.url, args.file, args.upload_field)
    print("Uploaded. requestKey={}".format(request_key))
    return request_key


def cmd_status(args):
    status = api_status(args.url, args.request_key)
    print("requestKey={} status={}".format(args.request_key, status))
    return status


def cmd_list_onus(args):
    onus = api_list_onus(args.url, args.request_key)
    if not onus:
        print("No ONUs found for requestKey={}".format(args.request_key))
        return
    for onu in onus:
        print("onuName={onu_name}  swVersion={sw}  hwVersion={hw}  status={status}".format(
            onu_name=onu.get("onuName", onu.get("name", "")),
            sw=onu.get("swVersion", ""),
            hw=onu.get("hwVersion", ""),
            status=onu.get("status", ""),
        ))
    return onus


def cmd_get_omci(args):
    onu_name = getattr(args, "onu_name", None)
    output_path = args.output or ("omci_{}.json".format(onu_name) if onu_name else "omci_all.json")
    print("Downloading OMCI data (onuName={}) ...".format(onu_name or "<all>"))
    data = api_get_omci(args.url, args.request_key, onu_name)
    os.makedirs(os.path.dirname(os.path.abspath(output_path)), exist_ok=True)
    with open(output_path, "w", encoding="utf-8") as fh:
        json.dump(data, fh, indent=2, ensure_ascii=True)
    print("Written to {}".format(output_path))
    return output_path


def cmd_get_diagram(args):
    onu_name = getattr(args, "onu_name", None)
    output_path = args.output or ("diagram_{}.svg".format(onu_name) if onu_name else "diagram.svg")
    print("Downloading OMCI diagram (category=latest, onuName={}) ...".format(onu_name or "<all>"))
    svg_data = api_get_diagram(args.url, args.request_key, onu_name)
    os.makedirs(os.path.dirname(os.path.abspath(output_path)), exist_ok=True)
    with open(output_path, "wb") as fh:
        fh.write(svg_data)
    size_kb = len(svg_data) // 1024
    print("Written to {} ({} KB)".format(output_path, size_kb))
    return output_path


def cmd_get_plantuml(args):
    onu_name = getattr(args, "onu_name", None)
    output_path = args.output or ("plantuml_{}.wsd".format(onu_name) if onu_name else "plantuml.wsd")
    print("Downloading OMCI PlantUML (category=latest, onuName={}) ...".format(onu_name or "<all>"))
    wsd_data = api_get_plantuml(args.url, args.request_key, onu_name)
    os.makedirs(os.path.dirname(os.path.abspath(output_path)), exist_ok=True)
    with open(output_path, "wb") as fh:
        fh.write(wsd_data)
    size_kb = len(wsd_data) // 1024
    print("Written to {} ({} KB)".format(output_path, size_kb))
    return output_path


def _download_onu_artifacts(args, request_key, onu_name, output_dir):
    """Download OMCI JSON and optional diagram/plantuml artifacts for one ONU.

    Returns a tuple: (omci_path or '', diagram_path or '', plantuml_path or '').
    """
    omci_path = ""
    diagram_path = ""
    plantuml_path = ""

    out_path = os.path.join(output_dir, "omci_{}.json".format(onu_name))
    try:
        data = api_get_omci(args.url, request_key, onu_name)
        with open(out_path, "w", encoding="utf-8") as fh:
            json.dump(data, fh, indent=2, ensure_ascii=True)
        print("  [OK] {} -> {}".format(onu_name, out_path))
        omci_path = out_path
    except RuntimeError as exc:
        print("  [FAIL] {}: {}".format(onu_name, exc))

    if args.download_plantuml:
        plantuml_name = "{}_{}.wsd".format(args.plantuml_prefix, onu_name)
        plantuml_path = os.path.join(output_dir, plantuml_name)
        try:
            wsd_data = api_get_plantuml(args.url, request_key, onu_name)
            with open(plantuml_path, "wb") as fh:
                fh.write(wsd_data)
            print("  [OK] {} plantuml -> {}".format(onu_name, plantuml_path))
        except RuntimeError as exc:
            print("  [WARN] {} plantuml: {}".format(onu_name, exc))
            plantuml_path = ""

    if args.download_diagram:
        diagram_name = "{}_{}.svg".format(args.diagram_prefix, onu_name)
        diagram_path = os.path.join(output_dir, diagram_name)
        try:
            svg_data = api_get_diagram(args.url, request_key, onu_name)
            with open(diagram_path, "wb") as fh:
                fh.write(svg_data)
            print("  [OK] {} diagram -> {}".format(onu_name, diagram_path))
        except RuntimeError as exc:
            print("  [WARN] {} diagram: {}".format(onu_name, exc))
            diagram_path = ""

    return omci_path, diagram_path, plantuml_path


def _wait_until_done(base_url, request_key):
    print("Waiting for parsing to complete ...", end="", flush=True)
    elapsed = 0
    while elapsed < POLL_TIMEOUT:
        time.sleep(POLL_INTERVAL)
        elapsed += POLL_INTERVAL
        status = api_status(base_url, request_key)
        print(".", end="", flush=True)
        status_lower = status.lower()
        if status_lower in {"done", "completed", "success", "finished"}:
            print(" done.")
            return
        if status_lower in {"failed", "error"}:
            print()
            raise RuntimeError("Parsing failed for requestKey={}".format(request_key))

    print()
    raise RuntimeError("Timed out waiting for requestKey={}".format(request_key))


def _print_onu_list(onus):
    print("Found {} ONU(s):".format(len(onus)))
    for onu in onus:
        print("  onuName={}  swVersion={}  hwVersion={}  status={}".format(
            onu.get("onuName", onu.get("name", "")),
            onu.get("swVersion", ""),
            onu.get("hwVersion", ""),
            onu.get("status", ""),
        ))


def cmd_pipeline(args):
    """Full pipeline: upload -> wait -> list ONUs -> download OMCI per ONU.

    When --download-diagram/--download-plantuml are enabled, also downloads
    latest SVG and/or WSD artifacts for each ONU.
    """
    output_dir = args.output_dir or "."
    os.makedirs(output_dir, exist_ok=True)

    # Upload
    print("Uploading {} ...".format(args.file))
    request_key = api_upload(args.url, args.file, args.upload_field)
    print("requestKey={}".format(request_key))

    # Poll
    _wait_until_done(args.url, request_key)

    # List ONUs
    onus = api_list_onus(args.url, request_key)
    if not onus:
        print("No ONUs found. Nothing to extract.")
        return []

    _print_onu_list(onus)

    # Download OMCI per ONU
    output_files = []
    diagram_files = []
    plantuml_files = []
    missing_plantuml_for_onus = []
    for onu in onus:
        onu_name = onu.get("onuName") or onu.get("name") or "unknown"
        omci_path, diagram_path, plantuml_path = _download_onu_artifacts(args, request_key, onu_name, output_dir)
        if omci_path:
            output_files.append(omci_path)
        if diagram_path:
            diagram_files.append(diagram_path)
        if plantuml_path:
            plantuml_files.append(plantuml_path)
        elif args.require_plantuml:
            missing_plantuml_for_onus.append(onu_name)

    if args.require_plantuml and missing_plantuml_for_onus:
        raise RuntimeError(
            "Mandatory PlantUML contract failed: missing plantuml_<onu>.wsd for ONU(s): {}. "
            "Please re-run pipeline from raw log."
            .format(", ".join(missing_plantuml_for_onus))
        )

    print("\n{} OMCI file(s) ready. Feed each to iop-auto-diagnosis:".format(len(output_files)))
    if args.download_diagram:
        print("{} diagram file(s) downloaded.".format(len(diagram_files)))
    if args.download_plantuml:
        print("{} plantuml file(s) downloaded.".format(len(plantuml_files)))
    for path in output_files:
        print("  python iop-auto-diagnosis/scripts/normalize_toolkit_output.py --input {} --output <canonical.json> --onu-name <ONU_NAME>".format(path))
        print("  python iop-auto-diagnosis/scripts/extract_features.py --input <canonical.json> --pretty")

    return output_files


# ---------------------------------------------------------------------------
# CLI
# ---------------------------------------------------------------------------

def main():
    parser = argparse.ArgumentParser(
        description="iop-toolkit client — parse OMCI logs via http://10.101.15.238:8080"
    )
    parser.add_argument("--url", default=DEFAULT_URL, help="iop-toolkit base URL")
    parser.add_argument("--upload-field", default=DEFAULT_UPLOAD_FIELD, help="multipart form field name for upload API")
    sub = parser.add_subparsers(dest="command", required=True)

    s_upload = sub.add_parser("upload", help="Upload log file and start parsing")
    s_upload.add_argument("--file", required=True, help="Path to log file")

    s_status = sub.add_parser("status", help="Check parsing status")
    s_status.add_argument("--request-key", required=True)

    s_list = sub.add_parser("list-onus", help="List ONUs in a parsed task")
    s_list.add_argument("--request-key", required=True)

    s_omci = sub.add_parser("get-omci", help="Download OMCI JSON (all ONUs or filtered by name)")
    s_omci.add_argument("--request-key", required=True)
    s_omci.add_argument("--onu-name", help="Filter by ONU name (optional; returns all if omitted)")
    s_omci.add_argument("--output", help="Output JSON path (default: omci_all.json or omci_<name>.json)")

    s_diagram = sub.add_parser("get-diagram", help="Download OMCI diagram (SVG, category=latest)")
    s_diagram.add_argument("--request-key", required=True)
    s_diagram.add_argument("--onu-name", help="Filter by ONU name (optional; returns all if omitted)")
    s_diagram.add_argument("--output", help="Output SVG path (default: diagram.svg or diagram_<name>.svg)")

    s_plantuml = sub.add_parser("get-plantuml", help="Download OMCI PlantUML source (WSD, category=latest)")
    s_plantuml.add_argument("--request-key", required=True)
    s_plantuml.add_argument("--onu-name", help="Filter by ONU name (optional; returns all if omitted)")
    s_plantuml.add_argument("--output", help="Output WSD path (default: plantuml.wsd or plantuml_<name>.wsd)")

    s_pipe = sub.add_parser("pipeline", help="Full pipeline: upload → wait → download all ONUs")
    s_pipe.add_argument("--file", required=True, help="Path to log file")
    s_pipe.add_argument("--output-dir", default=".", help="Directory for output JSON files")
    s_pipe.add_argument(
        "--download-diagram",
        action="store_true",
        help="Also download latest OMCI diagram SVG per ONU",
    )
    s_pipe.add_argument(
        "--download-plantuml",
        action="store_true",
        help="Also download latest OMCI PlantUML WSD per ONU",
    )
    s_pipe.add_argument(
        "--require-plantuml",
        action="store_true",
        help="Fail pipeline if any ONU does not produce plantuml_<onu>.wsd",
    )
    s_pipe.add_argument(
        "--diagram-prefix",
        default="diagram",
        help="Output prefix for diagram file names (default: diagram)",
    )
    s_pipe.add_argument(
        "--plantuml-prefix",
        default="plantuml",
        help="Output prefix for plantuml file names (default: plantuml)",
    )

    args = parser.parse_args()
    dispatch = {
        "upload": cmd_upload,
        "status": cmd_status,
        "list-onus": cmd_list_onus,
        "get-omci": cmd_get_omci,
        "get-diagram": cmd_get_diagram,
        "get-plantuml": cmd_get_plantuml,
        "pipeline": cmd_pipeline,
    }
    try:
        dispatch[args.command](args)
    except RuntimeError as exc:
        sys.stderr.write("Error: {}\n".format(exc))
        sys.exit(1)


if __name__ == "__main__":
    main()
