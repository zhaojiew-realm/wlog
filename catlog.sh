#!/usr/bin/bash
curl $1:8108 -d "$(cat $2)"