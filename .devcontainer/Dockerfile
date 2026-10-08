FROM golang:1.25

ENV DEBIAN_FRONTEND=noninteractive

# Устанавливаем git для работы с репозиторием и golangci-lint
RUN apt-get update && apt-get install -y git && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /usr/local/bin && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
