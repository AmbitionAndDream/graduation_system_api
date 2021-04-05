#!/bin/bash
RUN_NAME="graduation_system_api"


CUR_DIR=$(cd $(dirname "$0");pwd)
cd "${CUR_DIR}";cd ../

if [ -d "output" ]; then
  rm -rf output
  mkdir output
fi

go build -a -o output/bin/${RUN_NAME} cmd/server/server.go


list="dev prod"

for env in $list; do
  mkdir -p output/"$env"
  mkdir -p output/"$env"/conf

  cp scripts/bootstrap.sh output/"$env"/
  cp configs/"$env"/config.toml output/"$env"/conf/config.toml

  chmod +x output/"$env"/bootstrap.sh

  cp -r output/bin output/"$env"/
done

rm -rf output/bin

