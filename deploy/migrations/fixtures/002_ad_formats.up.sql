BEGIN;

INSERT INTO adv_format
    (id, codename, type, title, active, width, height, min_width, min_height, config)
VALUES
    (1, 'direct', 'direct', 'Direct', 'active', NULL, NULL, NULL, NULL, '{}'::jsonb),

    (2, 'proxy', 'proxy', 'Proxy Stretch', 'active', 0, 0, 10, 10, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),

    (3, 'video', 'video', 'Video', 'active', NULL, NULL, NULL, NULL,
        $json$
        {
          "assets": [
            {
              "id": 1,
              "required": true,
              "name": "main",
              "adjust_size": true,
              "width": 1500,
              "height": 1500,
              "min_width": 150,
              "min_height": 150,
              "animated": true,
              "sound": true,
              "thumbs": ["300x", "500x"],
              "allowed_types": ["video/mp4", "video/webm"]
            },
            {
              "id": 2,
              "required": false,
              "name": "preview",
              "width": 1500,
              "height": 1500,
              "min_width": 150,
              "min_height": 150,
              "animated": false,
              "allowed_types": ["image/jpeg", "image/png"]
            },
            {
              "id": 3,
              "required": false,
              "name": "logo",
              "width": 100,
              "height": 100,
              "min_width": 50,
              "min_height": 50,
              "animated": false,
              "sound": false,
              "allowed_types": ["image/jpeg", "image/png"]
            }
          ],
          "fields": [
            {
              "id": 101,
              "required": true,
              "title": "Title",
              "name": "title",
              "type": "string",
              "min": 3,
              "max": 150,
              "multilang": true
            },
            {
              "id": 102,
              "required": false,
              "title": "Display start",
              "description": "When to show the companion overlay",
              "name": "start",
              "type": "string",
              "select": [
                {"title": "Start", "value": "start"},
                {"title": "First Quartile", "value": "first_quartile"},
                {"title": "Midpoint", "value": "midpoint"},
                {"title": "Third Quartile", "value": "third_quartile"},
                {"title": "Complete", "value": "complete"}
              ]
            },
            {
              "id": 103,
              "required": false,
              "title": "Display on specific time",
              "description": "Show companion at this second",
              "name": "start_on",
              "exclude": ["start"],
              "type": "int"
            }
          ]
        }
        $json$::jsonb),

    (4, 'native', 'native', 'Native', 'active', NULL, NULL, NULL, NULL,
        $json$
        {
          "assets": [
            {
              "id": 1,
              "required": true,
              "name": "main",
              "adjust_size": true,
              "width": 1500,
              "height": 1500,
              "min_width": 50,
              "min_height": 50,
              "animated": false,
              "sound": false,
              "thumbs": ["250x", "350x", "500x"],
              "allowed_types": ["image/jpeg", "image/png", "video/mp4", "video/webm"]
            },
            {
              "id": 2,
              "required": false,
              "name": "logo",
              "width": 100,
              "height": 100,
              "min_width": 50,
              "min_height": 50,
              "animated": false,
              "sound": false,
              "allowed_types": ["image/jpeg", "image/png"]
            }
          ],
          "fields": [
            {
              "id": 101,
              "required": true,
              "title": "Title",
              "name": "title",
              "type": "string",
              "min": 5,
              "max": 40,
              "multilang": true
            },
            {
              "id": 102,
              "required": true,
              "title": "Description",
              "description": "Body text shown with the ad",
              "name": "description",
              "type": "string",
              "min": 5,
              "max": 80,
              "multiline": 3,
              "multilang": true
            },
            {
              "id": 103,
              "required": false,
              "title": "Brandname",
              "name": "brandname",
              "type": "string",
              "max": 30,
              "multilang": true
            },
            {
              "id": 104,
              "required": false,
              "title": "Phone",
              "name": "phone",
              "type": "phone"
            },
            {
              "id": 105,
              "required": false,
              "title": "Promotion URL",
              "description": "Click-through landing page URL",
              "name": "url",
              "type": "url"
            }
          ]
        }
        $json$::jsonb),

    (5, 'proxy_250x250', 'proxy', 'Proxy (Square)', 'active', 250, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (6, 'proxy_200x200', 'proxy', 'Proxy (Small Square)', 'active', 200, 200, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (7, 'proxy_468x60', 'proxy', 'Proxy (Banner)', 'active', 468, 60, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (8, 'proxy_728x90', 'proxy', 'Proxy (Leaderboard)', 'active', 728, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (9, 'proxy_300x250', 'proxy', 'Proxy (Inline Rectangle)', 'active', 300, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (10, 'proxy_336x280', 'proxy', 'Proxy (Large Rectangle)', 'active', 336, 280, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (11, 'proxy_120x600', 'proxy', 'Proxy (Skyscraper)', 'active', 120, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (12, 'proxy_160x600', 'proxy', 'Proxy (Wide Skyscraper)', 'active', 160, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (13, 'proxy_300x600', 'proxy', 'Proxy (Half-Page Ad)', 'active', 300, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (14, 'proxy_970x90', 'proxy', 'Proxy (Large Leaderboard)', 'active', 970, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),
    (15, 'proxy_320x50', 'proxy', 'Proxy (Mobile Leaderboard)', 'active', 320, 50, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "allowed_types": ["text/html"]
          }
        ],
        "fields": [
          {
            "id": 1001,
            "title": "HTML Content",
            "name": "content",
            "type": "html",
            "multiline": 10
          }
        ]
      }$json$::jsonb),

    (16, 'banner_250x250', 'banner', 'Square', 'active', 250, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 250,
            "height": 250,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (17, 'banner_200x200', 'banner', 'Small Square', 'active', 200, 200, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 200,
            "height": 200,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (18, 'banner_468x60', 'banner', 'Banner', 'active', 468, 60, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 468,
            "height": 60,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (19, 'banner_728x90', 'banner', 'Leaderboard', 'active', 728, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 728,
            "height": 90,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (20, 'banner_300x250', 'banner', 'Inline Rectangle', 'active', 300, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 300,
            "height": 250,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (21, 'banner_336x280', 'banner', 'Large Rectangle', 'active', 336, 280, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 336,
            "height": 280,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (22, 'banner_120x600', 'banner', 'Skyscraper', 'active', 120, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 120,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (23, 'banner_160x600', 'banner', 'Wide Skyscraper', 'active', 160, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 160,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (24, 'banner_300x600', 'banner', 'Half-Page Ad', 'active', 300, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 300,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (25, 'banner_970x90', 'banner', 'Large Leaderboard', 'active', 970, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 970,
            "height": 90,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb),
    (26, 'banner_320x50', 'banner', 'Mobile Leaderboard', 'active', 320, 50, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 320,
            "height": 50,
            "allowed_types": ["image/jpeg", "image/png"]
          }
        ]
      }$json$::jsonb)
ON CONFLICT (codename) DO UPDATE
SET
    type = EXCLUDED.type,
    title = EXCLUDED.title,
    active = EXCLUDED.active,
    width = EXCLUDED.width,
    height = EXCLUDED.height,
    min_width = EXCLUDED.min_width,
    min_height = EXCLUDED.min_height,
    config = EXCLUDED.config;

COMMIT;