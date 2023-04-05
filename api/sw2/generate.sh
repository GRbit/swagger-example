#!/usr/bin/env sh
# https://github.com/go-swagger/go-swagger
swagger generate server -f api/sw2/api.yaml -P model.Principal -m pkg/sw2/model -s internal/sw2
swagger generate client -f api/sw2/api.yaml -P model.Principal -m pkg/sw2/model -c pkg/sw2/client
