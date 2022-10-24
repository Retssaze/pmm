#!/bin/sh
set -e

cleanInstall() {

    printf "\033[32m Reload the service unit from disk\033[0m\n"
    systemctl daemon-reload ||:
    printf "\033[32m Unmask the service\033[0m\n"
    systemctl unmask pmm-server.service ||:
    printf "\033[32m Set the preset flag for the service unit\033[0m\n"
    systemctl preset pmm-server.service ||:
    printf "\033[32m Set the enabled flag for the service unit\033[0m\n"
    systemctl enable pmm-server.service ||:
    systemctl restart pmm-server.service ||:

    exit 0

}

upgrade() {
    systemctl restart pmm-server.service ||:

    exit 0
}

action="$1"
case "$action" in
  "1" | "install" | "configure")
    printf "\033[32m Post Install of an install\033[0m\n"
    cleanInstall
    ;;
  "2" | "upgrade")
    printf "\033[32m Post Install of an upgrade\033[0m\n"
    upgrade
    ;;
  *)
    # $1 == version being installed
    printf "\033[32m Post Install unknown option\033[0m"
    exit 0
    ;;
esac

exit 0
