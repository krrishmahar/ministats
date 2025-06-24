#!/bin/bash
set -e

BINARY_NAME="ministats"
MAN_SECTION="1"

echo "[*] Building $BINARY_NAME..."
go build -o $BINARY_NAME

echo "[*] Moving binary to /usr/local/bin/..."
sudo mv $BINARY_NAME /usr/local/bin/

echo "[*] Generating man page..."
go run main.go
sudo mv man/$BINARY_NAME.$MAN_SECTION /usr/share/man/man$MAN_SECTION/

echo "[*] Updating man database..."
sudo mandb

echo "[+] Installed $BINARY_NAME successfully."
echo "    Try running: $BINARY_NAME --help or man $BINARY_NAME"
