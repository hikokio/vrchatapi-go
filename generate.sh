#!/bin/bash

npm install @openapitools/openapi-generator-cli

rm -rf docs/


./node_modules/\@openapitools/openapi-generator-cli/main.js generate \
-g go \
--additional-properties=packageName=vrchatapi,enumClassPrefix=true \
--git-user-id=vrchatapi \
--git-repo-id=vrchatapi-go \
-o ./src \
-i https://raw.githubusercontent.com/vrchatapi/specification/gh-pages/openapi.yaml \
--http-user-agent="vrchatapi-go"

rm -rf api/ test/
go mod tidy
# Remove messily pasted markdown at top of every file

#Fix openapi-generator not importing "time"
sed -i '/.*net\/url.*/a "time"' ./api_system.go 
go fmt api_system.go