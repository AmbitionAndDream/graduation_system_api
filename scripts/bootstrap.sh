#!/bin/bash
CUR_DIR=$(
  cd $(dirname $0)
  pwd
)

BinaryName="graduation_system_api"

if [ "X$1" != "X" ]; then
  RUNTIME_ROOT=$1
else
  RUNTIME_ROOT=${CUR_DIR}
fi


CONF_FILE=${CUR_DIR}/conf/config.toml

args="-config=$CONF_FILE"

echo "${CUR_DIR}/bin/${BinaryName} ${args}"
exec "${CUR_DIR}"/bin/${BinaryName} "${args}"

