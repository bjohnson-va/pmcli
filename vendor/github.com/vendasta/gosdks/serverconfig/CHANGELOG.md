# 2.0.0
- Removed unused and confusing `StartHTTPSServer` because we configure HTTP
via Google Cloud Console, not in-service.
- Breaking change: satisfy golint (naming)
    - renamed server_config package to serverconfig

# 1.0.0

- Initial release of dedicated `server_config` package (broken off from old
`utils` package)
