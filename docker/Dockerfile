FROM golang:1.15.15

LABEL maintainer="Vinicius Renó <vinicius.vbreno@gmail.com>"

ARG GID=1000
ARG UID=1000

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y upgrade \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install \
    curl \
    locales \
    && apt-get -y autoclean autoremove \
    && rm -rf /var/lib/apt/lists/*

RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen \
    && locale-gen

# Create a non-privileged user for the application.
RUN groupadd --non-unique --gid $GID app \
    && useradd --non-unique --shell /bin/bash --gid app --uid $UID --home-dir /app app

# Create the app directory.
RUN mkdir /app /app/bin 

# Copy entrypoints.
COPY *-entrypoint.bash /app

# Set ownership of files to the app user and set default permissions.
RUN find /app /go | xargs chown $GID:$UID
RUN find /app /go -type f | xargs chmod 0640
RUN find /app /go -type d | xargs chmod 0750

RUN chmod +x /app/*-entrypoint.bash

USER app

ENV CGO_ENABLED=0

ENV PATH=/app/bin:$PATH

# Install air: https://github.com/cosmtrek/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && mv /go/bin/air /app/bin

# Install godotenv: https://github.com/joho/godotenv
RUN cd /go && go get github.com/joho/godotenv/cmd/godotenv && mv /go/bin/godotenv /app/bin

WORKDIR /app