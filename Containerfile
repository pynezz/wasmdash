# Web application Dockerfile
# Work in progress
#
# TODO:
#   - Building go binary
#   - Node: build assets
#   - Ports
#   - Environment variables
#   - Secrets management


#- Multi-stage build for theme asset optimization
FROM node:alpine AS css-builder
WORKDIR /app
COPY assets/css/ ./assets/css/
COPY tailwind.config.js ./
RUN npm install -D tailwindcss && \
    npx tailwindcss -i ./assets/css/base.css -o ./dist/styles.css --minify

FROM golang:alpine AS builder
COPY scripts/theme-config.go ./scripts/
RUN go run scripts/theme-config.go  # Generates theme at build time
# -

# - Web app stage
FROM ubi/ubi-9
COPY --from=css-builder /app/dist/styles.css /static/css/
COPY --from=go-builder /app/wasmdash /
