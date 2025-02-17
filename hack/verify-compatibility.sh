#!/bin/bash
#shellcheck source=lib/init.sh
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

TMP_ROOT=$(mktemp --directory)

cleanup() {
  rm -rf "${TMP_ROOT}"
}
trap "cleanup" EXIT SIGINT

V_ROOT="${TMP_ROOT}/src/github.com/openshift/api"
mkdir -p "$V_ROOT"
cp -a --no-preserve=timestamp -r "$SCRIPT_ROOT"/* "$V_ROOT"
(
  cd "$V_ROOT" || exit
  export GOPATH="$TMP_ROOT"
  rm -Rf _output
#  bash --init-file <(echo cd $V_ROOT)
  ./hack/update-compatibility.sh > /dev/null
)

mapfile -t GENERATED < <(find "${SCRIPT_ROOT}" -name "*.go" -printf "%P\n" | sort)
for g in "${GENERATED[@]}" ; do
  if ! diff --unified --text "$SCRIPT_ROOT/$g" "$V_ROOT/$g" ; then
    printf "\nopenshift_compatibility is out of date. Please run hack/update-compatibility.sh\n"
    exit 1
  fi
done
