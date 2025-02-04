#!/bin/bash

FILE=/state/pg-migrations
if [[ -e $FILE ]]; then
  echo 'Already applied!'
else
  for file in blaze_migrations/*.up.sql; do
      if [ -n "$file" ] && [ -e "$file" ]; then
        echo "$file"
      fi
      psql -h postgres  -f $file
  done

  for file in migrations/*.up.sql; do
      if [ -n "$file" ] && [ -e "$file" ]; then
        echo "$file"
      fi
      psql -h postgres  -f $file
  done

  for file in billing_migrations/*.up.sql; do
      if [ -n "$file" ] && [ -e "$file" ]; then
        echo "$file"
      fi
      psql -h postgres  -f $file
  done
fi
touch /state/pg-migrations