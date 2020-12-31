#!/bin/bash
sudo wget "https://github.com/bitijin/ipcamviewer/releases/download/0.0.1/ipcamviewer" -P /usr/bin/
sudo chmod +x /usr/bin/ipcamviewer
wget "https://raw.githubusercontent.com/bitijin/ipcamviewer/main/resources/ipcamviewer.desktop" -P "$HOME"/.local/share/applications/
wget "https://raw.githubusercontent.com/bitijin/ipcamviewer/main/resources/icon.png" -P "$HOME"/.local/share/ipcamviewer/
echo "$HOME"/.local/share/ipcamviewer/icon.png >> "$HOME"/.local/share/applications/ipcamviewer.desktop
