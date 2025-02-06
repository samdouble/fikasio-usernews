#!/bin/bash
#
# Development helper

set -euo pipefail

USAGE="Usage: app COMMAND [SUBCOMMAND]

Common usage:
  app dev
    start the development containers
  app enter
    enter the specified container"

if [[ "$#" == "0" ]]; then
  echo "No argument supplied"
  cmd="help"
else
  cmd="$1"
  shift
fi

# cd to the dir of this script
cd "${0%/*}"

case "${cmd}" in
  dev)
    sudo docker-compose build
    sudo docker-compose up
    ;;
  enter)
    if [[ -z "${1-}" ]]; then
      sudo docker exec -it fikasio-usernews /bin/bash
    elif [[ "$@" =~ ^test.* ]]; then
      sudo docker exec -it fikasio-usernews /bin/bash -c "npm run $@"
    else
      sudo docker exec -it fikasio-usernews /bin/bash -c "npm run $@"
    fi
    ;;
  help)
    echo "$USAGE"
    ;;
  *)
    echo "Unexpected command '${cmd}'" >&2
    exit 1
    ;;
esac

exit 0