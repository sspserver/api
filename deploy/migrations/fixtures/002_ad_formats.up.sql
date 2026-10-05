BEGIN;

INSERT INTO adv_format
    (id, codename, type, title, description, active, width, height, min_width, min_height, config)
VALUES
    (1, 'direct', 'direct', 'Direct', 'Direct placement with no creative template (Popunder, Direct, Tab-click, etc.)', 'active', NULL, NULL, NULL, NULL, '{}'::jsonb),

    (2, 'proxy', 'proxy', 'Proxy Stretch', 'Stretchable HTML creative that fills the placement. Minimum size 10×10.', 'active', 0, 0, 10, 10, $json${
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

    (3, 'video', 'video', 'Video', 'Video creative with optional preview image, logo, and title.', 'active', NULL, NULL, NULL, NULL,
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
              "allowed_types": ["image/jpeg", "image/png", "image/webp"]
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
              "allowed_types": ["image/jpeg", "image/png", "image/webp"]
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

    (4, 'native', 'native', 'Native', 'Native ad with image or video, title, body text, and optional brand, phone, and landing URL.', 'active', NULL, NULL, NULL, NULL,
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
              "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
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
              "allowed_types": ["image/jpeg", "image/png", "image/webp"]
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
              "type": "phone",
              "multilang": true
            },
            {
              "id": 105,
              "required": false,
              "title": "Promotion URL",
              "description": "Click-through landing page URL",
              "name": "url",
              "type": "url",
              "editable": false
            }
          ]
        }
        $json$::jsonb),

    (5, 'proxy_250x250', 'proxy', 'Proxy (Square)', 'Fixed-size HTML creative, 250×250 (Square).', 'active', 250, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (6, 'proxy_200x200', 'proxy', 'Proxy (Small Square)', 'Fixed-size HTML creative, 200×200 (Small Square).', 'active', 200, 200, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (7, 'proxy_468x60', 'proxy', 'Proxy (Banner)', 'Fixed-size HTML creative, 468×60 (Banner).', 'active', 468, 60, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (8, 'proxy_728x90', 'proxy', 'Proxy (Leaderboard)', 'Fixed-size HTML creative, 728×90 (Leaderboard).', 'active', 728, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (9, 'proxy_300x250', 'proxy', 'Proxy (Inline Rectangle)', 'Fixed-size HTML creative, 300×250 (Inline Rectangle).', 'active', 300, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (10, 'proxy_336x280', 'proxy', 'Proxy (Large Rectangle)', 'Fixed-size HTML creative, 336×280 (Large Rectangle).', 'active', 336, 280, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (11, 'proxy_120x600', 'proxy', 'Proxy (Skyscraper)', 'Fixed-size HTML creative, 120×600 (Skyscraper).', 'active', 120, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (12, 'proxy_160x600', 'proxy', 'Proxy (Wide Skyscraper)', 'Fixed-size HTML creative, 160×600 (Wide Skyscraper).', 'active', 160, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (13, 'proxy_300x600', 'proxy', 'Proxy (Half-Page Ad)', 'Fixed-size HTML creative, 300×600 (Half-Page Ad).', 'active', 300, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (14, 'proxy_970x90', 'proxy', 'Proxy (Large Leaderboard)', 'Fixed-size HTML creative, 970×90 (Large Leaderboard).', 'active', 970, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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
    (15, 'proxy_320x50', 'proxy', 'Proxy (Mobile Leaderboard)', 'Fixed-size HTML creative, 320×50 (Mobile Leaderboard).', 'active', 320, 50, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "name": "main",
            "required": false,
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

    (16, 'banner_250x250', 'banner', 'Square', 'Image or video creative, 250×250 (Square).', 'active', 250, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 250,
            "height": 250,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (17, 'banner_200x200', 'banner', 'Small Square', 'Image or video creative, 200×200 (Small Square).', 'active', 200, 200, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 200,
            "height": 200,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (18, 'banner_468x60', 'banner', 'Banner', 'Image or video creative, 468×60 (Banner).', 'active', 468, 60, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 468,
            "height": 60,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (19, 'banner_728x90', 'banner', 'Leaderboard', 'Image or video creative, 728×90 (Leaderboard).', 'active', 728, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 728,
            "height": 90,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (20, 'banner_300x250', 'banner', 'Inline Rectangle', 'Image or video creative, 300×250 (Inline Rectangle).', 'active', 300, 250, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 300,
            "height": 250,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (21, 'banner_336x280', 'banner', 'Large Rectangle', 'Image or video creative, 336×280 (Large Rectangle).', 'active', 336, 280, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 336,
            "height": 280,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (22, 'banner_120x600', 'banner', 'Skyscraper', 'Image or video creative, 120×600 (Skyscraper).', 'active', 120, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 120,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (23, 'banner_160x600', 'banner', 'Wide Skyscraper', 'Image or video creative, 160×600 (Wide Skyscraper).', 'active', 160, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 160,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (24, 'banner_300x600', 'banner', 'Half-Page Ad', 'Image or video creative, 300×600 (Half-Page Ad).', 'active', 300, 600, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 300,
            "height": 600,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (25, 'banner_970x90', 'banner', 'Large Leaderboard', 'Image or video creative, 970×90 (Large Leaderboard).', 'active', 970, 90, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 970,
            "height": 90,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb),
    (26, 'banner_320x50', 'banner', 'Mobile Leaderboard', 'Image or video creative, 320×50 (Mobile Leaderboard).', 'active', 320, 50, NULL, NULL, $json${
        "assets": [
          {
            "id": 1,
            "required": true,
            "name": "main",
            "adjust_size": true,
            "width": 320,
            "height": 50,
            "allowed_types": ["image/jpeg", "image/png", "image/webp", "video/mp4", "video/webm"]
          }
        ]
      }$json$::jsonb)
ON CONFLICT (codename) DO UPDATE
SET
    type = EXCLUDED.type,
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    active = EXCLUDED.active,
    width = EXCLUDED.width,
    height = EXCLUDED.height,
    min_width = EXCLUDED.min_width,
    min_height = EXCLUDED.min_height,
    config = EXCLUDED.config;

COMMIT;