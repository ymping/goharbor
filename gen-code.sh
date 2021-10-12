#!/usr/bin/env bash

# go swagger install
# https://goswagger.io/install.html

SwagFile=swagger.yaml

curl -fsSL https://raw.githubusercontent.com/goharbor/harbor/master/api/v2.0/swagger.yaml -o ${SwagFile}

# before generate, need run "go mod init github.com/ymping/goharbor" if go mod not init
swagger generate client -f ${SwagFile}
