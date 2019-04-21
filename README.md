# Walle

## Experimental reverse proxy written in go

### Features

1. Small Configuration (Currently working on it)
2. SSL & Letsencrypt (Automatic SSL)
3. Custom Middlewares
4. Advanced Features (Caching, Client config, Server config)
5. Service discovery (Registering services with names, and exposing an api to capture those)


## Foreword

A little introduction of what is what. This may change over time. For now we are talking about "Servers" and "Location"
The server holds the information of what domain, port we are talking about. Locations are path mappings, with configuration.
