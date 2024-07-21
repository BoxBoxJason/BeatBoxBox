FROM alpine:latest

ARG DB_HOST=beatboxbox-db
ARG DB_PORT=5432
ARG DB_USER=admin
ARG DB_PASSWORD=admin
ARG DB_NAME=beatboxbox
ARG DB_SSLMODE=disable
ARG SECRET_JWT_KEY=secret

ENV DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASSWORD} \
    DB_NAME=${DB_NAME} \
    DB_SSLMODE=${DB_SSLMODE} \
    SECRET_JWT_KEY=${SECRET_JWT_KEY} \
    BEATBOXBOX_ROOT_DIR="/home/user/beatboxbox"

WORKDIR /home/user/beatboxbox

# Install required packages
RUN apk add --no-cache \
    npm \
    go && \
    rm -rf /var/cache/apk/* && \
    npm install --no-cache -g npm@latest @vue/cli

# Build frontend
COPY ./frontend ./frontend/
RUN cd /home/user/beatboxbox/frontend && \
    npm --no-cache --legacy-peer-deps install && \
    npm cache clean --force && \
    npm run build && \
    find dist -mindepth 1 -maxdepth 1 ! -name 'dist' -exec rm -rf {} +

# Copy the source code
COPY ./pkg ./pkg/
COPY ./go.mod \
    ./go.sum \
    ./.env \
    ./
COPY ./cmd ./cmd/
COPY ./secret ./secret/
COPY ./internal ./internal/

# Build backend
RUN cd /home/user/beatboxbox && \
    go mod tidy

# ENTRYPOINT ["go", "run", "/home/user/beatboxbox/cmd/server/main.go"]
ENTRYPOINT ["go","test","./...","-parallel","4","-v"]
