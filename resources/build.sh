#!/bin/bash
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi
go get github.com/adrg/libvlc-go/v3
go get github.com/gotk3/gotk3/gtk
cd $HOME
git clone https://github.com/bitijin/ipcamviewer
cd ipcamviewer/
go build
mv ipcamviewer /usr/bin/
mv resources/icon.png /usr/share/pixmaps
mv resources/ipcamviewer.desktop /usr/share/applications
echo "INSTALLED"
