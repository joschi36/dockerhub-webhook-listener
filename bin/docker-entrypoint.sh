#!/bin/bash
set -e

export listen_addr=${listen_addr:-`localhost:8080`}

test $api_key
test $rancher_url
test $rancher_user_key
test $rancher_secret_key
test $rancher_service

/usr/local/bin/confd -onetime -backend env

exec "$@"
