#!/bin/bash
FILE=/state/pg-dump
if [[ -e $FILE ]]; then
  echo 'Already applied!'
else
  psql -h postgres  < /dump.sql
fi

touch /state/pg-dump