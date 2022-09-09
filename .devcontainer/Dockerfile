ARG VARIANT=ubuntu-22.04
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libbz2-dev \
    libdb-dev \
    libreadline-dev \
    libffi-dev \
    libgdbm-dev \
    liblzma-dev \
    libncursesw5-dev \
    libsqlite3-dev \
    libssl-dev \
    zlib1g-dev \
    uuid-dev \
    bat \
    curl \
    exa \
    fish \
    make \
    ripgrep \
    wget \
 && rm -rf /var/lib/apt/lists/*

# Install Starship
RUN curl -sS https://starship.rs/install.sh | sh -s -- --yes

WORKDIR /app

# Install Go
ARG GO_VERSION=1.19.1
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"

# 
ENV GOPATH /go
ENV PATH="$PATH:$GOPATH/bin"
RUN mkdir -p "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Install Go Tools
ENV GO111MODULE=auto
COPY Makefile /tmp/Makefile
RUN make -f /tmp/Makefile install-go-tools

USER vscode
