# Get the arguments from the .env file
$BEATBOXBOX_BUILD_ARGS = ""

if (Test-Path .env) {
    Get-Content .env | ForEach-Object {
        $line = $_
        if ($line -notmatch "^#" -and $line -match "=") {
            $VAR_NAME, $VAR_VALUE = $line -split "=", 2
            $BEATBOXBOX_BUILD_ARGS += "--build-arg $VAR_NAME=$VAR_VALUE "
        }
    }
}

# Get rid of the old containers
docker-compose down
docker rmi -f beatboxbox:latest

# Build the new container
docker build . -t beatboxbox:latest $BEATBOXBOX_BUILD_ARGS
docker-compose up
