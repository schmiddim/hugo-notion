#!/bin/sh
set -e

# Assign inputs to variables
NOTION_DB_ID=$1
NOTION_TOKEN=$2
PATH=$3

# Export environment variables
export NOTION_DB_ID
export NOTION_TOKEN

# Run the Go program with the provided path
echo path $PATH
/usr/local/bin/hugo-notion -p "$PATH"