#!/usr/bin/env bash

function check_configuration {
  if [[ -z "${PROJECT_RELATIVE_FOLDER}" ]]; then
    echo "Could not find required configuration [PROJECT_RELATIVE_FOLDER]. Exiting"
    exit 1
  fi
}

function make_base_folders {
  echo "Creating base folder: test-folder..."
  mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads
  mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/movies
  mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/tv
  mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/others
  mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/apptemp
  echo "Created!"
}

function delete_base_folders {
  echo "Deleting test-folder..."
  rm -rf "${PROJECT_RELATIVE_FOLDER}"/test-folder
  echo "Deleted!"
}

check_configuration
delete_base_folders
make_base_folders

# Movie Sample 1
baseFolderName="Avengers.Endgame.2019.2160p.BluRay.x265.10bit.SDR.DTS-HD.MA.TrueHD.7.1.Atmos-SWTYBLZ"
mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}
mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/Sample
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/Sample/Sample-AEBSM10.mkv
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/${baseFolderName}.mkv
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/RARBG.txt

# Movie Sample 2
baseFolderName="Sonic.the.Hedgehog.2020.2160p.UHD.BluRay.x265.10bit.HDR.DTS-HD.MA.TrueHD.7.1.Atmos-SWTYBLZ"
mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}
mkdir -p "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/Sample
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/Sample/Sample-STHBHM10.mkv
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/${baseFolderName}.mkv
touch "${PROJECT_RELATIVE_FOLDER}"/test-folder/downloads/${baseFolderName}/RARBG.txt

# Movie Sample 3 : No folder
touch "${PROJECT_RELATIVE_FOLDER}"/"test-folder/downloads/Spirited.Away.2001.RERIP.1080p.BluRay.X264-AMIABLE[rarbg].mkv"