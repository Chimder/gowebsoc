#!/bin/bash

# Load environment variables from .env file
source .env


if [ -f .env ]; then
  export $(cat .env | grep -v '#' | awk -F= '{print $1}')
fi

# Ensure DB_URL is set
if [ -z "$DB_URL" ]; then
  echo "DB_URL is not set. Make sure it is defined in your .env file."
  exit 1
fi

DEV_URL="docker://postgres"
MODEL_PATH="file://sqlc/schema"
EXCLUDES=(
    "auth"
    "extensions"
    "graphql"
    "graphql_public"
    "pgbouncer"
    "pgsodium"
    "pgsodium_masks"
    "realtime"
    "storage"
    "vault"
    "atlas_schema_revisions"
)

EXCLUDE_FLAGS=""
for EXCLUDE in "${EXCLUDES[@]}"; do
    EXCLUDE_FLAGS+="--exclude $EXCLUDE "
done

# Apply schema
atlas schema apply \
    --url "$DB_URL" \
    --dev-url "$DEV_URL" \
    --to "$MODEL_PATH" \

# Check if schema apply was successful
if [ $? -eq 0 ]; then
    echo "Schema applied successfully."
else
    echo "Error applying schema."
    exit 1
fi

# Log end of script
echo "Schema application completed."