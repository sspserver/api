#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import sys


def generate_sql(json_data):
    yield f"INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename) VALUES"

    first = True
    for entry in json_data:
        codename = entry['codename'].replace("'", "''")
        name = entry['name'].replace("'", "''")
        description = entry['description'].replace("'", "''")
        match_exp = entry['match_exp'].replace("'", "''")
        active = entry['active'].replace("'", "''")
        maker_codename = entry['maker_codename'].replace("'", "''")
        type_codename = entry['type_codename'].replace("'", "''")

        if first:
            yield f"  ('{codename}', '{name}', '{description}', '{match_exp}', '{active}', '{maker_codename}', '{type_codename}')"
            first = False
        else:
            yield f", ('{codename}', '{name}', '{description}', '{match_exp}', '{active}', '{maker_codename}', '{type_codename}')"

    yield "  ON CONFLICT (codename) DO NOTHING;\n"

    for entry in json_data:
        codename = entry['codename'].replace("'", "''")
        name = entry['name'].replace("'", "''")
        description = entry['description'].replace("'", "''")
        match_exp = entry['match_exp'].replace("'", "''")
        active = entry['active'].replace("'", "''")
        maker_codename = entry['maker_codename'].replace("'", "''")
        type_codename = entry['type_codename'].replace("'", "''")

        if 'versions' in entry:
            first = True
            yield "\n-------------------------\n"
            yield f"INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES"

            for version in entry['versions']:
                version_codename = version['codename'].replace("'", "''")
                version_name = version['name'].replace("'", "''")
                version_description = version['description'].replace("'", "''")
                version_match_exp = version['match_exp'].replace("'", "''")
                version_active = version['active'].replace("'", "''")
                year_release = version.get('year_release', 0)
                if first:
                    yield f"  ('{version_codename}', '{version_name}', '{version_description}', '{version_match_exp}', '{version_active}', '{maker_codename}', '{type_codename}', (SELECT id FROM type_device_model WHERE codename = '{codename}'), {year_release})"
                    first = False
                else:
                    yield f", ('{version_codename}', '{version_name}', '{version_description}', '{version_match_exp}', '{version_active}', '{maker_codename}', '{type_codename}', (SELECT id FROM type_device_model WHERE codename = '{codename}'), {year_release})"

            yield "  ON CONFLICT (codename) DO NOTHING;"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: migrate.py <file.json>")
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
