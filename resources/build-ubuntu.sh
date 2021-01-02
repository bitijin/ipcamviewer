#!/bin/bash
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi
apt -y install golang-go 
apt -y install libgtk-3-0 libgtk-3-dev libvlc-dev vlc-plugin-base vlc-plugin-video-output libcairo2-dev libglib2.0-dev 
curl -s -L https://raw.githubusercontent.com/bitijin/ipcamviewer/main/resources/build.sh | bash
