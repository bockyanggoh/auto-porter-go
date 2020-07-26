#!/usr/bin/env bash

function make_base_folders {
  echo "Creating base folder: test-folder..."
  mkdir -p test-folder/downloads
  mkdir -p test-folder/movies
  mkdir -p test-folder/tv
  mkdir -p test-folder/others
  mkdir -p test-folder/apptemp
  echo "Created!"
}

function delete_base_folders {
  echo "Deleting test-folder..."
  rm -rf test-folder
  echo "Deleted!"
}


delete_base_folders
make_base_folders

# Movie Sample 1
baseFolderName="Avengers.Endgame.2019.2160p.BluRay.x265.10bit.SDR.DTS-HD.MA.TrueHD.7.1.Atmos-SWTYBLZ"
mkdir -p test-folder/downloads/${baseFolderName}
mkdir -p test-folder/downloads/${baseFolderName}/Sample
touch test-folder/downloads/${baseFolderName}/Sample/Sample-AEBSM10.mkv
touch test-folder/downloads/${baseFolderName}/${baseFolderName}.mkv
touch test-folder/downloads/${baseFolderName}/RARBG.txt

# Movie Sample 2
baseFolderName="Sonic.the.Hedgehog.2020.2160p.UHD.BluRay.x265.10bit.HDR.DTS-HD.MA.TrueHD.7.1.Atmos-SWTYBLZ"
mkdir -p test-folder/downloads/${baseFolderName}
mkdir -p test-folder/downloads/${baseFolderName}/Sample
touch test-folder/downloads/${baseFolderName}/Sample/Sample-STHBHM10.mkv
touch test-folder/downloads/${baseFolderName}/${baseFolderName}.mkv
touch test-folder/downloads/${baseFolderName}/RARBG.txt

# Movie Sample 3 : No folder
touch "test-folder/downloads/Spirited.Away.2001.RERIP.1080p.BluRay.X264-AMIABLE[rarbg].mkv"