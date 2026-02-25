FROM golang:latest

# Permite passar UID/GID na build
ARG USER_ID=1000
ARG GROUP_ID=1000

# Install required tools
RUN apt-get update && apt-get install -y \
    git \
	bash \
    vim \
    apache2-utils \
    && rm -rf /var/lib/apt/lists/* \
	&& groupadd -g $GROUP_ID felipecaue \
    && useradd -m -u $USER_ID -g felipecaue -s /bin/bash felipecaue \
    && mkdir -p /app \
    && chown -R felipecaue:felipecaue /app /go

# Install Go tools
RUN go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Set the working directory inside the container
WORKDIR /app

# Switch to non-root user
USER felipecaue

# Command to run the application
CMD ["tail", "-f", "/dev/null"]