FROM python:3.12-slim

# Proxy is only needed while installing dependencies; it is scoped to the pip
# layer via build args so it never leaks into the running container (which must
# talk directly to OLTs / Kafka / MinIO on the internal network).
ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY
# Optional PyPI index mirror (e.g. https://pypi.tuna.tsinghua.edu.cn/simple) to
# speed up installs on slow networks. Leave empty to use the default PyPI.
ARG PIP_INDEX_URL

WORKDIR /app

COPY requirements.txt .
ENV PIP_DEFAULT_TIMEOUT=120 \
    PIP_RETRIES=10 \
    PIP_INDEX_URL=${PIP_INDEX_URL}
RUN http_proxy=${HTTP_PROXY} https_proxy=${HTTPS_PROXY} no_proxy=${NO_PROXY} \
    pip install -U pip && \
    http_proxy=${HTTP_PROXY} https_proxy=${HTTPS_PROXY} no_proxy=${NO_PROXY} \
    pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chmod +x run.sh

CMD ["./run.sh"]
