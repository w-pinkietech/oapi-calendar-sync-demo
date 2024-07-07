#!/usr/bin/env bash

RUN_NAME=oapi_calendar_sync_demo

mkdir -p output/bin output/conf
cp -r conf/* output/conf
export GO111MODULE=on

go build -o output/bin/$RUN_NAME *.go