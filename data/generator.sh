#!/usr/bin/env bash

declare -r distinfo_path="https://sources.debian.org/data/main/d/distro-info-data"
declare -r version="$2"

# get list of distros to download
IFS=',' read -r -a distro_list <<< "$1"

if [[ -z "$version" || -z "$distro_list" ]] ; then
    echo "Usage: $0 <distro-list> <version>"
    exit 1
fi

for distro in "${distro_list[@]}"; do
  echo "prepare distro-info-data for ${distro}"
  # create temporary file for downloaded csv data
  tmp="$(mktemp)"

  # download csv data; error on download failure
  curl -s "${distinfo_path}/${version}/${distro}.csv" > "${tmp}"
  if [ $? -ne 0 ]; then
    echo "Failed to download distro-info-data for ${distro}"
    rm "${tmp}"
    exit 1
  fi

  # convert csv data to structured json
  python3 -c 'import csv, json, sys; print(json.dumps([dict(r) for r in csv.DictReader(sys.stdin)]))' < "${tmp}" | jq > "$(dirname $0)/${distro}.json"

  # remove temporary file
  rm "${tmp}"
done
