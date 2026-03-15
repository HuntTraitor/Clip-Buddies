FROM jellyfin/jellyfin:latest

USER root

COPY ./jellyfin/rootCA.pem /usr/local/share/ca-certificates/mkcert-rootCA.crt

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*