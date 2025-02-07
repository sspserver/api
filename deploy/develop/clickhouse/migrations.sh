#!/bin/bash

apply_migrations() {
  FILE=/state/clickhouse-$2
  if [[ -e $FILE ]]; then
    echo 'Already applied!'
  else
    for file in "$1/*.up.sql"; do
        if [ -n "$file" ] && [ -e "$file" ]; then
          echo "$file"
        fi
        clickhouse-client --host clickhouse-server  --queries-file $file
    done
  fi
  touch /state/clickhouse-$2
}

apply_migrations /migrations migrations
apply_migrations /migrations-gen migrations-gen
