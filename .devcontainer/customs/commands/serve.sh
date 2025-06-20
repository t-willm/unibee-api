#!/bin/bash

cmd="Serve"
description="Launch Air"
author="thibaud.willm@dotworld.ch"

source $UTILS_DIR/functions.sh

print_message "Launching Air" "📦"

/root/go/bin/air -c $WORKSPACE_DIR/.air.toml
