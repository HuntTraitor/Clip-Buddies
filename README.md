# Clip Buddies Portal

Clip Buddies Portal is a self-hosted media portal that provides a single login experience for internal users to access both a PeerTube and Jellyfin instance. The goal of the project is to remove friction and reliance on third-party platforms by allowing teams to manage and distribute clips, movies, and TV content internally through infrastructure they control.

The portal uses Authentik as a centralized identity provider, enabling a secure Single Sign-On (SSO) experience across all services in the stack.

---

## Architecture Overview

The system consists of four main components working together behind a reverse proxy.

Users authenticate once through Authentik and gain access to both PeerTube and Jellyfin without needing separate accounts or credentials.

Core components:

Caddy
Acts as the reverse proxy and TLS termination layer. It exposes the services through local domains and routes traffic to the appropriate containers.

Authentik
Handles identity management and authentication using OpenID Connect (OIDC). Authentik is responsible for user accounts, authentication flows, and issuing identity tokens used by PeerTube.

PeerTube
Provides a decentralized video platform used internally for hosting clips, recordings, and other shareable video content.

Jellyfin
Serves as the media server for movies and TV shows.

Docker Compose
Orchestrates all services and provides a reproducible development environment.

---

## Why This Project Exists

Many teams rely on external platforms for media distribution, which introduces several issues:

• Content ownership concerns
• Limited control over access policies
• Vendor lock-in
• Dependency on third-party infrastructure

Clip Buddies Portal addresses these issues by enabling organizations to host their own media ecosystem.

With this stack:

Internal users get unlimited access to clips and media without needing accounts on external platforms.
Organizations maintain full control over their content and authentication.
The infrastructure can scale and integrate with additional services over time.

---

## Authentication Flow

Authentication is handled through Authentik using OpenID Connect.

1. A user visits the portal and selects login.
2. The user is redirected to Authentik.
3. Authentik authenticates the user.
4. Authentik issues an OIDC token.
5. PeerTube uses the token to create or link the internal account.
6. The user is logged in automatically.

This creates a seamless Single Sign-On experience.

---

## Future Plans

This portal is intended to become a broader internal application hub.

Planned improvements include:

### Domain-based application routing

Additional applications will be exposed through subdomains behind the same reverse proxy.

Examples:

portal.example.com
auth.example.com
video.example.com
media.example.com
clips.example.com

Each service will integrate with Authentik so that users only authenticate once.

---

### Additional Application Integrations

Planned services include:

Internal tooling dashboards
Media processing pipelines
Clip management systems
Developer utilities
Analytics and monitoring tools

All new applications will integrate with Authentik via OpenID Connect to maintain a consistent authentication model.

---

### Unified Portal Interface

The long-term goal is to build a dedicated frontend portal application that presents users with a single dashboard containing links to all internal applications.

Instead of navigating between independent services, users will interact with a unified system.

---

## Development Setup

This repository uses Docker Compose to run the entire stack locally.

Services include:

Authentik
PeerTube
Jellyfin
Postgres
Redis
Caddy

To start the stack:

docker compose up -d

---

## TLS Certificates for Local Development

Local HTTPS is required for OpenID Connect to function properly.

This project uses mkcert to generate trusted development certificates.

Install mkcert and initialize the local certificate authority:

mkcert -install

Generate certificates for the local services:

mkcert -cert-file caddy/certs/localhost.pem
-key-file caddy/certs/localhost-key.pem
auth.localhost peertube.localhost jellyfin.localhost

After generating the certificates, start the stack:

docker compose up -d

Note that the caddy/certs directory is intentionally excluded from version control.
Each developer must generate their own certificates locally.

---

## Repository Structure

docker-compose.yml
Caddyfile
.env.example

peertube/
config/
production.yaml

caddy/
certs/ (ignored in git)

---

## Security Notes

The following files should never be committed to version control:

.env
TLS private keys
Generated certificates
Local root certificate authorities

These files contain secrets or machine-specific data and must remain local to each environment.

---

## License

This project is currently intended for internal development and experimentation. A license may be added in the future depending on distribution plans.

---
