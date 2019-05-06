#!/usr/bin/env bash
cd $(dirname "${BASH_SOURCE[0]}")
set -ex

go generate ./...
cp ./frontend/index.html ../../slimsag.github.io/index.html
bash -c "cd ../../slimsag.github.io && git add . && git commit -m 'generate' && git push"
