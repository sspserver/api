#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import sys


def generate_insert_statements(json_data):
    for entry in json_data:
        parent_id = 'NULL'
        yield f"""INSERT INTO type_browser (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{entry['name']}', '', '{entry['description']}', '{entry['match_name_exp']}', '{entry['match_ua_exp']}', '{entry['match_ver_min_exp']}', '{entry['match_ver_max_exp']}', {entry['year_release']}, {entry['year_end_support']}, '{entry['active']}', {parent_id});"""
        
        for version in entry.get("versions", []):
            yield f"""INSERT INTO type_browser (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{version['name']}', '{version['version']}', '{version['description']}', '{version['match_name_exp']}', '{version['match_ua_exp']}', '{version['match_ver_min_exp']}', '{version['match_ver_max_exp']}', {version['year_release']}, {version['year_end_support']}, '{version['active']}', (SELECT id FROM type_browser WHERE name = '{entry['name']}' LIMIT 1));"""

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python script.py <json_filename>")
        sys.exit(1)
    
    json_filename = sys.argv[1]
    
    with open(json_filename, "r", encoding="utf-8") as file:
        json_data = json.load(file)
    
    for sql in generate_insert_statements(json_data):
        print(sql)
