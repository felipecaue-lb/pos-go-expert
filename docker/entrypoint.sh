#!/bin/bash
set -e

# Detecta UID/GID do dono do /app (volume montado do host)
HOST_UID=$(stat -c '%u' /app)
HOST_GID=$(stat -c '%g' /app)

USERNAME=devuser

# Cria/ajusta grupo
if getent group "$HOST_GID" > /dev/null 2>&1; then
    GROUP_NAME=$(getent group "$HOST_GID" | cut -d: -f1)
else
    groupadd -g "$HOST_GID" "$USERNAME"
    GROUP_NAME="$USERNAME"
fi

# Cria/ajusta usuário
if id "$USERNAME" > /dev/null 2>&1; then
    usermod -u "$HOST_UID" -g "$HOST_GID" "$USERNAME"
else
    useradd -m -u "$HOST_UID" -g "$GROUP_NAME" -s /bin/bash "$USERNAME"
fi

# Garante permissões do GOPATH
chown -R "$HOST_UID:$HOST_GID" /go

# Garante home com permissões corretas
HOME_DIR="/home/$USERNAME"
mkdir -p "$HOME_DIR"
chown -R "$HOST_UID:$HOST_GID" "$HOME_DIR"

# Executa o comando como o usuário correto
export HOME="$HOME_DIR"
exec gosu "$USERNAME" "$@"
