#!/usr/bin/env bash

set -e

BRANCH="main"
COMMIT_MSG="auto commit $(date '+%Y-%m-%d %H:%M:%S')"

while true; do
  echo "==> Running go program..."
  go run main.go || echo "go run failed, continue..."

  echo "==> Checking git status..."
  if [[ -n "$(git status --porcelain)" ]]; then
    echo "==> Changes detected, committing..."
    git add .
    git commit -m "$COMMIT_MSG"
    git push origin "$BRANCH"
  else
    echo "==> No changes, skip commit."
  fi

  echo "==> Sleeping 1800 seconds..."
  sleep 1800
done
