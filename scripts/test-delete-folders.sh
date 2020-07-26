#!/usr/bin/env bash
function check_configuration {
  if [[ -z "${PROJECT_RELATIVE_FOLDER}" ]]; then
    echo "Could not find required configuration [PROJECT_RELATIVE_FOLDER]. Exiting"
    exit 1
  fi
}

function delete_folder() {
  echo "Deleting test-folder..."
  rm -rf "${PROJECT_RELATIVE_FOLDER}"/test-folder
  echo "Deleted!"
  exit 0
}

check_configuration
delete_folder


