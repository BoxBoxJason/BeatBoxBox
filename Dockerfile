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
    bash \
    curl \
    openssh \
    npm \
    go && \
    npm install -g npm@latest @vue/cli

# Copy the source code
COPY ./pkg ./pkg/
COPY ./go.mod \
    ./go.sum \
    ./.env \
    ./
COPY ./cmd ./cmd/
COPY ./secret ./secret/
COPY ./internal ./internal/
COPY ./frontend ./frontend/

# Build frontend & install go dependencies
RUN cd /home/user/beatboxbox/frontend && \
    npm install && \
    npm run build && \
    cd /home/user/beatboxbox && \
    go mod tidy

ENTRYPOINT ["go", "run", "/home/user/beatboxbox/cmd/server/main.go"]
#ENTRYPOINT ["go","test","./...","-parallel","4","-v"]
