# !/bin/bash

rm -Rf dist
mkdir dist

GOOS=linux GOARCH=amd64 go build -o dist/weather

cp -R content dist
