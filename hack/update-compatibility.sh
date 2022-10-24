#!/bin/bash

mapfile -t apiPkgs < <(go list -mod=vendor  -find  ./... | grep -v -E "$(go list)/(pkg|cmd|tools|hack)" | grep -E "$(go list)/[^/]+/[^/]+")
inputArg=$(printf ",%s" "${apiPkgs[@]}")
inputArg=${inputArg:1}

go run github.com/uccps-samples/api/cmd/uccp-compatibility-gen --input-dirs "$inputArg"
