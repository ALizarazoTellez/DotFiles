#!/usr/bin/env fish

~/.config/hypr/theme.sh
ags run &
poweralertd &
hypridle &
wayland-pipewire-idle-inhibit &
hyprpaper &
~/.config/hypr/support wallpaper-service &
playerctld &
