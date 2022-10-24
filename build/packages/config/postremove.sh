#!/bin/sh
set -e

upgrade() {

    exit 0
}

remove() {
    loginctl terminate-user pmm-server || true
    getent passwd pmm-server >/dev/null && userdel pmm-server
    getent group pmm >/dev/null && groupdel pmm

    exit 0
}


action="$1"
case "$action" in
  "0" | "remove" | "purge")
    printf "\033[32m Post Remove\033[0m\n"
    remove
    ;;
  "1" | "upgrade")
    printf "\033[32m Post Remove of an upgrade\033[0m\n"
    upgrade
    ;;
  *)
    remove
    ;;
esac

exit 0
