#!/bin/sh
# This script is used to quickstart application

# Get the arguments from the .env file
BEATBOXBOX_BUILD_ARGS=""
if [ -f .env ]; then
    while IFS= read -r line; do
        if [[ ! "$line" =~ ^# && "$line" =~ .*=.* ]]; then
        VAR_NAME=$(echo "$line" | cut -d '=' -f 1)
        BUILD_ARGS+="--build-arg $VAR_NAME=${!VAR_NAME} "
    fi
    done < .env
fi

# Get rid of the old containers
docker compose down
docker rmi -f beatboxbox:latest

# Build the new container
docker build . -t beatboxbox:latest $BEATBOXBOX_BUILD_ARGS
docker compose up
