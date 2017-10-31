#!/bin/bash

function find_base_dir {
    local real_path=$(python -c "import os,sys;print os.path.realpath('$0')")
    local dir_name="$(dirname "$real_path")"
    BASEDIR="${dir_name}/.."
}

find_base_dir


for GOOS in darwin linux; do
	for GOARCH in amd64; do
	docker run \
		--rm \
		-it \
		-v "$BASEDIR":/usr/local/go/src/github.com/DigitalPacific/squirrel \
		-w /usr/local/go/src/github.com/DigitalPacific/squirrel \
		-e GOOS=$GOOS \
		-e GOARCH=$GOARCH \
		golang:1.8 bash -c 'go get . && go build -v -o releases/squirrel-"$GOOS"-"$GOARCH"'
	done
done

chmod 744 ${BASEDIR}/releases/*
