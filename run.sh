#!/bin/sh

ARGS=$1

if [ "$ARGS" = "hot-serve" ]; then
    swag init -g main.go 
	air start
elif [ "$ARGS" == "start" ]; then
    if [ -z $ENV ]; then
        export ENV="dev"
    fi

    ./app start $ENV
elif [ "$ARGS" == "build" ]; then
    swag init -g main.go
    go build -v -o app main.go
elif [ "$ARGS" = "setup" ]; then
    echo "Installing dependency"
    go mod download
    go install github.com/swaggo/swag/cmd/swag@latest

    echo "Done..."
else
    echo "Usage: $0 <setup|build|start [dev|prod]>"
fi