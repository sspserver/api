#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import sys


def generate_insert_statements(json_data):
    for ontry in json_data:
        parent_id = 'NULL'
        yield f"""INSERT INTO type_os (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{ontry['name']}', '', '{ontry['description']}', '{ontry['match_name_exp']}', '{ontry['match_ua_exp']}', '{ontry['match_ver_min_exp']}', '{ontry['match_ver_max_exp']}', {ontry['year_release']}, {ontry['year_end_support']}, '{ontry['active']}', {parent_id});"""
        
        for version in ontry.get("versions", []):
            yield f"""INSERT INTO type_os (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{version['name']}', '{version['version']}', '{version['description']}', '{version['match_name_exp']}', '{version['match_ua_exp']}', '{version['match_ver_min_exp']}', '{version['match_ver_max_exp']}', {version['year_release']}, {version['year_end_support']}, '{version['active']}', (SELECT id FROM type_os WHERE name = '{ontry['name']}' LIMIT 1));"""

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python script.py <json_filename>")
        sys.exit(1)
    
    json_filename = sys.argv[1]
    
    with open(json_filename, "r", encoding="utf-8") as file:
        json_data = json.load(file)
    
    for sql in generate_insert_statements(json_data):
        print(sql)
