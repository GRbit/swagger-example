#!/usr/bin/env sh
swagger generate server -f api/sw2/api.yaml -P model.Principal -m pkg/model -s internal/rest/server
swagger generate client -f api/sw2/api.yaml -P model.Principal -m pkg/model -c pkg/rest/client
