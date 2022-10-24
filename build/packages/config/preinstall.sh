#!/bin/sh
set -e

printf "\033[32m Pre Install, creating user\033[0m\n"

# use newusers util to create user and group to assign subuid/subguid
# newusers pw_name:pw_passwd:pw_uid:pw_gid:pw_gecos:pw_dir:pw_shel
getent passwd pmm-server >/dev/null >/dev/null || echo "pmm-server:::pmm::/home/pmm-server:/bin/false" | newusers ||:

# on ubuntu can fail first time, need to run second time
getent passwd pmm-server >/dev/null >/dev/null || echo "pmm-server:::pmm::/home/pmm-server:/bin/false" | newusers ||:

loginctl enable-linger pmm-server

exit 0