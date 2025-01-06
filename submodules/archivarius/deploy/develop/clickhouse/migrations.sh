#!/bin/bash
FILE=/state/clickhouse-migrations
if [[ -e $FILE ]]; then
  echo 'Already applied!'
else
  for file in migrations/*.up.sql; do
    if [ -n "$file" ] && [ -e "$file" ]; then
      echo "$file"
    fi
    clickhouse-client --host clickhouse-server  --queries-file $file
  done

fi

touch /state/clickhouse-migrations