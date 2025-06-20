#!/bin/bash

cmd="Install Dependencies"
description="Install dependencies"
author="thibaud.willm@dotworld.ch"

source $UTILS_DIR/functions.sh

# Install Migrate CLI tool

## Detect architecture
architecture="$(dpkg --print-architecture)"

print_message "Installing migrate" "📦"
curl -L -o migrate.deb https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-${architecture}.deb
dpkg -i migrate.deb
rm migrate.deb

## Check migrate
echo "Checking migrate version"
migrate -version