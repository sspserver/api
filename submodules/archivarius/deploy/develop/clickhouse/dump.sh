#!/bin/bash
FILE=/state/clickhouse-dump
if [[ -e $FILE ]]; then
  echo 'Already applied!'
else
  clickhouse-client --host clickhouse-server  -q "INSERT INTO stats.events_local FROM INFILE '/clickhouse/user_files/dump.csv' FORMAT CSVWithNames"
fi

touch /state/clickhouse-dump