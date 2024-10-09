#!/bin/bash

REPO_URL="https://github.com/networktocode/ntc-templates.git"
FOLDER="ntc_templates/templates"
TARGET_DIR="./tmp"

mkdir -p "$TARGET_DIR"
cd "$TARGET_DIR" || exit 1

git init

git remote add origin "$REPO_URL"

git config core.sparseCheckout true

echo "$FOLDER" > .git/info/sparse-checkout

git pull origin master

rm -rf .git

echo "Files downloaded to $TARGET_DIR/$FOLDER"