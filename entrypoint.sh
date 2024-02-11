#!/bin/sh

if [ -f /vault/secrets/config ]; then
  . /vault/secrets/config # Inject vault credentials
  echo "Using config from vault"
elif [ -f ./.env ]; then
  echo "Using config from .env"

  export $(grep -Ev '^#|^$' .env | xargs)
else
  echo "Using config from default .env (sample)"

  cp .env-sample .env
  export $(grep -Ev '^#|^$' .env | xargs)
fi

if [[ -z $ENV ]]; then
  echo "Deploy ENV is not set"
  
  exit 1
fi

./app start $ENV