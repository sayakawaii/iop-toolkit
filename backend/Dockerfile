# ---------- Stage 1: build the Go binary ----------
FROM golang:1.24-bookworm AS builder

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY
# GOPROXY is overridable so the build works both on open networks and behind a
# corporate proxy (pass GOPROXY as a build arg from deploy/.env if needed).
ARG GOPROXY=https://proxy.golang.org,direct
ENV http_proxy=${HTTP_PROXY} \
    https_proxy=${HTTPS_PROXY} \
    no_proxy=${NO_PROXY} \
    GOPROXY=${GOPROXY} \
    GOFLAGS=-mod=mod

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o omciAnalyzer .

# ---------- Stage 2: minimal runtime ----------
FROM debian:bookworm-slim

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY
# Optional Debian apt mirror (e.g. https://mirrors.tuna.tsinghua.edu.cn) to
# speed up package downloads on networks where deb.debian.org is slow. Leave
# empty to use the default mirror.
ARG APT_MIRROR
ENV http_proxy=${HTTP_PROXY} \
    https_proxy=${HTTPS_PROXY} \
    no_proxy=${NO_PROXY} \
    DEBIAN_FRONTEND=noninteractive

# graphviz -> `dot` for PlantUML rendering; default-jre-headless -> run plantuml.jar;
# gettext-base -> `envsubst` used by the entrypoint to template config.yaml.
RUN set -eux; \
    if [ -n "${APT_MIRROR}" ]; then \
        sed -i "s|http://deb.debian.org|${APT_MIRROR}|g" /etc/apt/sources.list.d/debian.sources 2>/dev/null || true; \
        sed -i "s|http://deb.debian.org|${APT_MIRROR}|g" /etc/apt/sources.list 2>/dev/null || true; \
    fi; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        graphviz \
        default-jre-headless \
        gettext-base \
        ca-certificates; \
    rm -rf /var/lib/apt/lists/*

# Drop proxy so it does not leak into the running service.
ENV http_proxy="" https_proxy="" no_proxy=""

WORKDIR /omcianalyzer

COPY --from=builder /src/omciAnalyzer ./omciAnalyzer
COPY resource ./resource
COPY static ./static
COPY env ./env
COPY config/config.yaml.template ./config/config.yaml.template
COPY docker-entrypoint.sh ./docker-entrypoint.sh

RUN chmod +x ./omciAnalyzer ./docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./docker-entrypoint.sh"]
