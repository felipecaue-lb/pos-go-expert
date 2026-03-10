FROM golang:latest

# Install required tools
RUN apt-get update && apt-get install -y \
    git \
    bash \
    vim \
    apache2-utils \
    sqlite3 \
    libsqlite3-dev \
    protobuf-compiler \
    gosu \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /app

# Copy entrypoint script
COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Install Go tools
RUN go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/ktr0731/evans@latest && \
    go install github.com/spf13/cobra-cli@latest

# Install migrate CLI e SqlC
RUN go install -tags 'postgres mysql sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install Wire
RUN go install github.com/google/wire/cmd/wire@latest

# Set the working directory inside the container
WORKDIR /app

ENTRYPOINT ["entrypoint.sh"]
CMD ["tail", "-f", "/dev/null"]