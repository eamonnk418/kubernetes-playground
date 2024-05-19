#!/usr/bin/env sh

while true; do
    touch /tmp/healthcheck.txt;
    sleep 5;
done

test find /tmp/healthcheck.txt -mmin -1
