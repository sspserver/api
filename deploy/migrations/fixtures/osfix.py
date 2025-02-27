#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import sys


def generate_insert_statements(json_data):
    for os_entry in json_data:
        parent_id = 'NULL'
        yield f"""INSERT INTO type_os (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{os_entry['name']}', '', '{os_entry['description']}', '{os_entry['match_name_exp']}', '{os_entry['match_ua_exp']}', '{os_entry['match_ver_min_exp']}', '{os_entry['match_ver_max_exp']}', {os_entry['year_release']}, {os_entry['year_end_support']}, '{os_entry['active']}', {parent_id});"""
        
        for version in os_entry.get("versions", []):
            yield f"""INSERT INTO type_os (name, version, description, match_name_exp, match_ua_exp, match_ver_min_exp, match_ver_max_exp, year_release, year_end_support, active, parent_id) VALUES ('{version['name']}', '{version['version']}', '{version['description']}', '{version['match_name_exp']}', '{version['match_ua_exp']}', '{version['match_ver_min_exp']}', '{version['match_ver_max_exp']}', {version['year_release']}, {version['year_end_support']}, '{version['active']}', (SELECT id FROM type_os WHERE name = '{os_entry['name']}' LIMIT 1));"""

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python script.py <json_filename>")
        sys.exit(1)
    
    json_filename = sys.argv[1]
    
    with open(json_filename, "r", encoding="utf-8") as file:
        json_data = json.load(file)
    
    for sql in generate_insert_statements(json_data):
        print(sql)
