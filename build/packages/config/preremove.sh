#!/bin/sh
set -e

remove() {
    systemctl --no-reload disable --now pmm-server.service

    exit 0
}

upgrade() {

    exit 0
}

action="$1"
case "$action" in
  "0" | "remove" | "purge")
    printf "\033[32m Pre Remove\033[0m\n"
    remove
    ;;
  "1" | "upgrade")
    printf "\033[32m Pre Remove of an upgrade\033[0m\n"
    upgrade
    ;;
  *)
    remove
    ;;
esac

exit 0
