FROM alpine:latest

WORKDIR /home/user/BeatBoxBox

# Install required packages
RUN apk add --no-cache \
    bash \
    curl \
    openssh \
    npm \
    go && \
    npm install -g npm@latest @vue/cli

# Copy the source code
COPY ./internal ./internal/
COPY ./frontend ./frontend/
COPY ./pkg ./pkg/
COPY ./go.mod \
    ./go.sum \
    ./
COPY ./cmd ./cmd/
COPY ./secret ./secret/

# Build frontend
RUN cd /home/user/BeatBoxBox/frontend && \
    npm install && \
    npm run build && \
    cd /home/user/BeatBoxBox && \
    go mod tidy

ENTRYPOINT ["go", "run", "/home/user/BeatBoxBox/cmd/server/main.go"]
