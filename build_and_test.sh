#!/usr/bin/env bash

set -euo pipefail

run() {
  echo "> $@"
  "$@"
}

build_and_test() {
  local -r dir="$1"
  echo "=== $dir ==="
  cd "$dir"
  run go mod tidy
  run go fmt .
  run weaver generate .
  run go build .
  run staticcheck .
  run go test .
  run go test -race .
  run addlicense .
  cd - > /dev/null
  echo
}

main() {
  if [[ $# == 0 ]]; then
    for dir in */; do
      build_and_test "$dir"
    done
  else
    for dir in "$@"; do
      build_and_test "$dir"
    done
  fi
}

main "$@"
