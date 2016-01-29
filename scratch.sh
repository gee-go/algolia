#!/usr/bin/env bash

set -eE -o pipefail


algolia_req() {
  curl -X $1 \
    -H "X-Algolia-API-Key: ${ALGOLIA_API_KEY}" \
    -H "X-Algolia-Application-Id: ${ALGOLIA_APP_ID}" \
      "https://${ALGOLIA_APP_ID}-dsn.algolia.net/1/$2"
}

algolia_req DELETE "indexes/non_existant"