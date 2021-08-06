#!/bin/sh
set -e

if [[ ! -f go.mod ]];
then
  go mod init $PROJECT
fi

go mod tidy

if [ "${1#-}" != "${1}" ] || [ -z "$(command -v "${1}")" ];
  then
    set -- $@
fi

sleep 10

exec $@