#!/bin/bash

set -e

mkdir -p secrets

generate() {
  openssl rand -base64 32
}

echo "Generating secrets for TransactMesh..."

echo "$(generate)" > secrets/jwt_secret.txt
echo "$(generate)" > secrets/grafana_admin_password.txt

echo "Secrets generated successfully."
echo "⚠️ Do NOT commit the secrets/ directory"