#!/bin/bash

#
if ! command -v php &> /dev/null
then
    echo "Installing PHP-CLI..." "📦"

    if [[ -n "$(command -v apt-get)" ]]; then
        sudo apt-get update
        sudo DEBIAN_FRONTEND=noninteractive apt-get install -y php-cli
    elif [[ -n "$(command -v yum)" ]]; then
        sudo yum install -y php-cli
    elif [[ -n "$(command -v dnf)" ]]; then
        sudo dnf install -y php-cli
    elif [[ -n "$(command -v pacman)" ]]; then
        sudo pacman -Syu php-cli
    elif [[ -n "$(command -v brew)" ]]; then
        brew install php
    else
        echo "Unknown package manager. Exiting."
        exit 1
    fi
    echo "PHP-CLI installation completed!"
fi

exit 0