#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import sys


def generate_sql(json_data):
    yield "INSERT INTO type_device_maker (codename, name, description, match_exp, active)"
    yield "VALUES"
    
    first = True
    for entry in json_data:
        codename = entry['codename'].replace("'", "''")
        name = entry['name'].replace("'", "''")
        description = entry['description'].replace("'", "''")
        match_exp = entry['match_exp'].replace("'", "''")
        active = entry['active'].replace("'", "''")
        
        if first:
            yield f"    ('{codename}', '{name}', '{description}', '{match_exp}', '{active}')"
            first = False
        else:
            yield f"    ,('{codename}', '{name}', '{description}', '{match_exp}', '{active}')"
    
    yield "ON CONFLICT (codename) DO NOTHING;"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: convert.py <file.json>")
        sys.exit(1)
    
    file_path = sys.argv[1]
    try:
        with open(file_path, "r", encoding="utf-8") as file:
            json_data = json.load(file)
        
        for line in generate_sql(json_data):
            print(line)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)