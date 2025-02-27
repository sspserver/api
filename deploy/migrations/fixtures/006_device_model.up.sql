INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename) VALUES
  ('iphone', 'iPhone', 'Apple''s flagship smartphone series.', '$re$(?i)iPhone', 'active', 'apple', 'phone')
, ('ipad', 'iPad', 'Apple''s tablet series.', '$re$(?i)iPad', 'active', 'apple', 'tablet')
, ('pixel', 'Google Pixel', 'Google''s flagship smartphone series.', '$re$(?i)Pixel', 'active', 'google', 'phone')
, ('galaxy', 'Samsung Galaxy', 'Samsung''s flagship smartphone series.', '$re$(?i)Galaxy', 'active', 'samsung', 'phone')
, ('galaxy_tab', 'Samsung Galaxy Tab', 'Samsung''s flagship tablet series.', '$re$(?i)Galaxy Tab', 'active', 'samsung', 'tablet')
, ('huawei_p', 'Huawei P Series', 'Huawei''s flagship smartphone series focused on photography and performance.', '$re$(?i)Huawei P', 'active', 'huawei', 'phone')
, ('huawei_mate', 'Huawei Mate Series', 'Huawei''s high-end smartphone series focused on performance and business users.', '$re$(?i)Huawei Mate', 'active', 'huawei', 'phone')
, ('xiaomi_mi', 'Xiaomi Mi Series', 'Xiaomi''s premium smartphone series focused on innovation and performance.', '$re$(?i)Xiaomi Mi', 'active', 'xiaomi', 'phone')
, ('xiaomi_redmi', 'Xiaomi Redmi Series', 'Xiaomi''s budget-friendly smartphone series with strong performance.', '$re$(?i)Redmi Note', 'active', 'xiaomi', 'phone')
, ('xiaomi_poco', 'Xiaomi Poco Series', 'Xiaomi''s Poco series focuses on high performance at an affordable price.', '$re$(?i)Poco', 'active', 'xiaomi', 'phone')
, ('oneplus', 'OnePlus Series', 'OnePlus flagship smartphone series known for speed and performance.', '$re$(?i)OnePlus', 'active', 'oneplus', 'phone')
, ('oppo', 'Oppo Smartphones', 'Oppo''s flagship and mid-range smartphone series focused on innovation and design.', '$re$(?i)Oppo', 'active', 'oppo', 'phone')
, ('vivo', 'Vivo Smartphones', 'Vivo''s flagship and mid-range smartphone series focused on photography and innovation.', '$re$(?i)Vivo', 'active', 'vivo', 'phone')
, ('motorola', 'Motorola Smartphones', 'Motorola''s smartphone series, known for durability and near-stock Android experience.', '$re$(?i)Motorola', 'active', 'motorola', 'phone')
, ('sony_xperia', 'Sony Xperia Smartphones', 'Sony''s Xperia series, known for high-quality displays and camera technology.', '$re$(?i)Xperia', 'active', 'sony', 'phone')
, ('nokia', 'Nokia Smartphones', 'Nokia''s modern smartphones known for durability and stock Android experience.', '$re$(?i)Nokia', 'active', 'nokia', 'phone')
, ('asus', 'Asus Smartphones', 'Asus smartphone series including ROG gaming phones and Zenfone flagship devices.', '$re$(?i)Asus', 'active', 'asus', 'phone')
, ('galaxy_tab', 'Samsung Galaxy Tab Series', 'Samsung''s flagship tablet series, offering premium displays and performance.', '$re$(?i)Galaxy Tab', 'active', 'samsung', 'tablet')
, ('matepad', 'Huawei MatePad Series', 'Huawei''s premium tablet lineup featuring HarmonyOS and high-end hardware.', '$re$(?i)MatePad', 'active', 'huawei', 'tablet')
, ('xiaomi_pad', 'Xiaomi Pad Series', 'Xiaomi''s premium tablet lineup focused on performance and multimedia.', '$re$(?i)Xiaomi Pad', 'active', 'xiaomi', 'tablet')
, ('lenovo_tablets', 'Lenovo Tab Series', 'Lenovo''s lineup of Android tablets for entertainment and productivity.', '$re$(?i)Lenovo Tab', 'active', 'lenovo', 'tablet')
, ('microsoft_surface', 'Microsoft Surface Series', 'Microsoft''s premium Windows tablets with detachable keyboards and stylus support.', '$re$(?i)Surface', 'active', 'microsoft', 'tablet')
, ('amazon_fire', 'Amazon Fire Tablets', 'Amazon''s budget-friendly tablets optimized for media consumption and Alexa integration.', '$re$(?i)Fire', 'active', 'amazon', 'tablet')
, ('xbox', 'Xbox Series', 'Microsoft''s gaming console lineup with high-performance gaming capabilities.', '$re$(?i)Xbox', 'active', 'microsoft', 'connected')
, ('playstation', 'PlayStation Series', 'Sony''s gaming console lineup offering immersive gaming experiences.', '$re$(?i)PlayStation', 'active', 'sony', 'connected')
, ('nintendo', 'Nintendo Consoles', 'Nintendo''s gaming console lineup, known for innovation and family-friendly games.', '$re$(?i)Nintendo', 'active', 'nintendo', 'connected')
, ('apple_tv', 'Apple TV', 'Apple''s streaming media player with access to apps, games, and Apple services.', '$re$(?i)Apple TV', 'active', 'apple', 'connected')
, ('samsung_smart_tv', 'Samsung Smart TV', 'Samsung''s lineup of Smart TVs with Tizen OS and advanced display technology.', '$re$(?i)Samsung Smart TV', 'active', 'samsung', 'connected')
, ('lg_smart_tv', 'LG Smart TV', 'LG''s lineup of Smart TVs with webOS and OLED technology.', '$re$(?i)LG Smart TV', 'active', 'lg', 'connected')
  ON CONFLICT (codename) DO NOTHING;


-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('iphone_3g', 'iPhone 3G', 'Second generation iPhone with 3G connectivity.', '$re$(?i)iPhone 3G', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2008)
, ('iphone_4', 'iPhone 4', 'Fourth generation iPhone with a new design and Retina display.', '$re$(?i)iPhone 4', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2010)
, ('iphone_5', 'iPhone 5', 'Fifth generation iPhone with a taller screen and Lightning connector.', '$re$(?i)iPhone 5', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2012)
, ('iphone_6', 'iPhone 6', 'Sixth generation iPhone with a larger display.', '$re$(?i)iPhone 6', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2014)
, ('iphone_7', 'iPhone 7', 'Seventh generation iPhone with improved cameras and no headphone jack.', '$re$(?i)iPhone 7', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2016)
, ('iphone_8', 'iPhone 8', 'Eighth generation iPhone with wireless charging.', '$re$(?i)iPhone 8', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2017)
, ('iphone_x', 'iPhone X', 'Tenth anniversary iPhone with an edge-to-edge OLED display.', '$re$(?i)iPhone X', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2017)
, ('iphone_11', 'iPhone 11', 'iPhone with improved cameras and A13 Bionic chip.', '$re$(?i)iPhone 11', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2019)
, ('iphone_12', 'iPhone 12', 'iPhone with 5G and Ceramic Shield front cover.', '$re$(?i)iPhone 12', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2020)
, ('iphone_13', 'iPhone 13', 'iPhone with improved battery life and A15 Bionic chip.', '$re$(?i)iPhone 13', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2021)
, ('iphone_14', 'iPhone 14', 'Latest generation iPhone with enhanced camera and safety features.', '$re$(?i)iPhone 14', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2022)
, ('iphone_15', 'iPhone 15', 'Newest iPhone with USB-C and upgraded processor.', '$re$(?i)iPhone 15', 'active', 'apple', 'phone', (SELECT id FROM type_device_model WHERE codename = 'iphone'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('ipad_1', 'iPad (1st generation)', 'The first-generation iPad released by Apple.', '$re$(?i)iPad 1', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2010)
, ('ipad_2', 'iPad 2', 'Second generation iPad with a thinner design and front camera.', '$re$(?i)iPad 2', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2011)
, ('ipad_3', 'iPad (3rd generation)', 'Third generation iPad with Retina display.', '$re$(?i)iPad 3', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2012)
, ('ipad_4', 'iPad (4th generation)', 'Fourth generation iPad with Lightning connector.', '$re$(?i)iPad 4', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2012)
, ('ipad_air', 'iPad Air', 'Thinner and lighter iPad with a new design.', '$re$(?i)iPad Air', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2013)
, ('ipad_air_2', 'iPad Air 2', 'Second generation iPad Air with improved performance.', '$re$(?i)iPad Air 2', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2014)
, ('ipad_pro', 'iPad Pro', 'High-performance iPad with Apple Pencil support.', '$re$(?i)iPad Pro', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2015)
, ('ipad_5', 'iPad (5th generation)', 'Fifth generation iPad with A9 chip.', '$re$(?i)iPad 5', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2017)
, ('ipad_6', 'iPad (6th generation)', 'Sixth generation iPad with Apple Pencil support.', '$re$(?i)iPad 6', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2018)
, ('ipad_7', 'iPad (7th generation)', 'Seventh generation iPad with a larger 10.2-inch display.', '$re$(?i)iPad 7', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2019)
, ('ipad_8', 'iPad (8th generation)', 'Eighth generation iPad with A12 Bionic chip.', '$re$(?i)iPad 8', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2020)
, ('ipad_9', 'iPad (9th generation)', 'Ninth generation iPad with improved front camera.', '$re$(?i)iPad 9', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2021)
, ('ipad_10', 'iPad (10th generation)', 'Tenth generation iPad with USB-C.', '$re$(?i)iPad 10', 'active', 'apple', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'ipad'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('pixel_1', 'Pixel 1', 'First generation Pixel phone by Google.', '$re$(?i)Pixel 1', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2016)
, ('pixel_2', 'Pixel 2', 'Second generation Pixel phone with improved camera.', '$re$(?i)Pixel 2', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2017)
, ('pixel_3', 'Pixel 3', 'Third generation Pixel phone with Night Sight mode.', '$re$(?i)Pixel 3', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2018)
, ('pixel_4', 'Pixel 4', 'Fourth generation Pixel with Motion Sense.', '$re$(?i)Pixel 4', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2019)
, ('pixel_5', 'Pixel 5', 'Fifth generation Pixel with 5G connectivity.', '$re$(?i)Pixel 5', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2020)
, ('pixel_6', 'Pixel 6', 'Sixth generation Pixel with Google''s custom Tensor chip.', '$re$(?i)Pixel 6', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2021)
, ('pixel_7', 'Pixel 7', 'Seventh generation Pixel with advanced AI features.', '$re$(?i)Pixel 7', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2022)
, ('pixel_8', 'Pixel 8', 'Eighth generation Pixel with improved camera and battery life.', '$re$(?i)Pixel 8', 'active', 'google', 'phone', (SELECT id FROM type_device_model WHERE codename = 'pixel'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('galaxy_s', 'Samsung Galaxy S', 'First Galaxy S smartphone released by Samsung.', '$re$(?i)Galaxy S$', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2010)
, ('galaxy_s2', 'Samsung Galaxy S II', 'Second generation Galaxy S smartphone.', '$re$(?i)Galaxy S II', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2011)
, ('galaxy_s3', 'Samsung Galaxy S III', 'Third generation Galaxy S smartphone.', '$re$(?i)Galaxy S III', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2012)
, ('galaxy_s4', 'Samsung Galaxy S4', 'Fourth generation Galaxy S smartphone.', '$re$(?i)Galaxy S4', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2013)
, ('galaxy_s5', 'Samsung Galaxy S5', 'Fifth generation Galaxy S smartphone.', '$re$(?i)Galaxy S5', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2014)
, ('galaxy_s6', 'Samsung Galaxy S6', 'Sixth generation Galaxy S smartphone with a metal design.', '$re$(?i)Galaxy S6', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2015)
, ('galaxy_s7', 'Samsung Galaxy S7', 'Seventh generation Galaxy S smartphone with water resistance.', '$re$(?i)Galaxy S7', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2016)
, ('galaxy_s8', 'Samsung Galaxy S8', 'Eighth generation Galaxy S smartphone with an edge-to-edge display.', '$re$(?i)Galaxy S8', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2017)
, ('galaxy_s9', 'Samsung Galaxy S9', 'Ninth generation Galaxy S smartphone with improved camera.', '$re$(?i)Galaxy S9', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2018)
, ('galaxy_s10', 'Samsung Galaxy S10', 'Tenth generation Galaxy S smartphone with Infinity-O display.', '$re$(?i)Galaxy S10', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2019)
, ('galaxy_s20', 'Samsung Galaxy S20', 'Eleventh generation Galaxy S smartphone with 5G connectivity.', '$re$(?i)Galaxy S20', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2020)
, ('galaxy_s21', 'Samsung Galaxy S21', 'Twelfth generation Galaxy S smartphone with improved performance.', '$re$(?i)Galaxy S21', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2021)
, ('galaxy_s22', 'Samsung Galaxy S22', 'Thirteenth generation Galaxy S smartphone with enhanced camera features.', '$re$(?i)Galaxy S22', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2022)
, ('galaxy_s23', 'Samsung Galaxy S23', 'Latest Galaxy S smartphone with upgraded technology.', '$re$(?i)Galaxy S23', 'active', 'samsung', 'phone', (SELECT id FROM type_device_model WHERE codename = 'galaxy'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('galaxy_tab_a', 'Samsung Galaxy Tab A', 'Affordable Galaxy Tab model for general use.', '$re$(?i)Galaxy Tab A', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2015)
, ('galaxy_tab_s6', 'Samsung Galaxy Tab S6', 'High-performance tablet with S Pen support.', '$re$(?i)Galaxy Tab S6', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2019)
, ('galaxy_tab_s7', 'Samsung Galaxy Tab S7', 'Premium Samsung tablet with 120Hz display and S Pen.', '$re$(?i)Galaxy Tab S7', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2020)
, ('galaxy_tab_s8', 'Samsung Galaxy Tab S8', 'Latest Samsung premium tablet with improved multitasking capabilities.', '$re$(?i)Galaxy Tab S8', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2022)
, ('galaxy_tab_s9', 'Samsung Galaxy Tab S9', 'Newest high-end Samsung tablet with upgraded features and S Pen support.', '$re$(?i)Galaxy Tab S9', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('huawei_p10', 'Huawei P10', 'Huawei P10 with dual-camera system and Leica optics.', '$re$(?i)Huawei P10', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2017)
, ('huawei_p20', 'Huawei P20', 'Huawei P20 with AI-powered camera features.', '$re$(?i)Huawei P20', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2018)
, ('huawei_p30', 'Huawei P30', 'Huawei P30 with improved zoom and night photography capabilities.', '$re$(?i)Huawei P30', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2019)
, ('huawei_p40', 'Huawei P40', 'Huawei P40 featuring advanced AI photography.', '$re$(?i)Huawei P40', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2020)
, ('huawei_p50', 'Huawei P50', 'Huawei P50 with HarmonyOS and top-tier camera technology.', '$re$(?i)Huawei P50', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2021)
, ('huawei_p60', 'Huawei P60', 'Latest Huawei P60 with cutting-edge mobile photography features.', '$re$(?i)Huawei P60', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_p'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('huawei_mate10', 'Huawei Mate 10', 'Huawei Mate 10 with AI-driven processor and dual-camera system.', '$re$(?i)Huawei Mate 10', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_mate'), 2017)
, ('huawei_mate20', 'Huawei Mate 20', 'Huawei Mate 20 with triple-camera system and high-performance chipset.', '$re$(?i)Huawei Mate 20', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_mate'), 2018)
, ('huawei_mate30', 'Huawei Mate 30', 'Huawei Mate 30 featuring 5G connectivity and enhanced AI photography.', '$re$(?i)Huawei Mate 30', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_mate'), 2019)
, ('huawei_mate40', 'Huawei Mate 40', 'Huawei Mate 40 with improved camera technology and fast charging.', '$re$(?i)Huawei Mate 40', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_mate'), 2020)
, ('huawei_mate50', 'Huawei Mate 50', 'Huawei Mate 50 with HarmonyOS and top-tier hardware.', '$re$(?i)Huawei Mate 50', 'active', 'huawei', 'phone', (SELECT id FROM type_device_model WHERE codename = 'huawei_mate'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('xiaomi_mi8', 'Xiaomi Mi 8', 'Xiaomi Mi 8 with AI-powered dual cameras and infrared face unlock.', '$re$(?i)Xiaomi Mi 8', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2018)
, ('xiaomi_mi9', 'Xiaomi Mi 9', 'Xiaomi Mi 9 with Snapdragon 855 and triple-camera setup.', '$re$(?i)Xiaomi Mi 9', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2019)
, ('xiaomi_mi10', 'Xiaomi Mi 10', 'Xiaomi Mi 10 with 5G support and a 108MP main camera.', '$re$(?i)Xiaomi Mi 10', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2020)
, ('xiaomi_mi11', 'Xiaomi Mi 11', 'Xiaomi Mi 11 with Snapdragon 888 and 120Hz AMOLED display.', '$re$(?i)Xiaomi Mi 11', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2021)
, ('xiaomi_mi12', 'Xiaomi Mi 12', 'Xiaomi Mi 12 with advanced camera system and faster charging.', '$re$(?i)Xiaomi Mi 12', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2022)
, ('xiaomi_mi13', 'Xiaomi Mi 13', 'Latest Xiaomi Mi 13 with improved AI processing and sleek design.', '$re$(?i)Xiaomi Mi 13', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_mi'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('redmi_note_7', 'Redmi Note 7', 'Redmi Note 7 with a 48MP camera and Snapdragon 660.', '$re$(?i)Redmi Note 7', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2019)
, ('redmi_note_8', 'Redmi Note 8', 'Redmi Note 8 with quad-camera setup and Snapdragon 665.', '$re$(?i)Redmi Note 8', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2019)
, ('redmi_note_9', 'Redmi Note 9', 'Redmi Note 9 with MediaTek Helio G85 and large battery.', '$re$(?i)Redmi Note 9', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2020)
, ('redmi_note_10', 'Redmi Note 10', 'Redmi Note 10 with AMOLED display and Snapdragon 678.', '$re$(?i)Redmi Note 10', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2021)
, ('redmi_note_11', 'Redmi Note 11', 'Redmi Note 11 with 90Hz display and Snapdragon 680.', '$re$(?i)Redmi Note 11', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2022)
, ('redmi_note_12', 'Redmi Note 12', 'Latest Redmi Note 12 with 120Hz AMOLED display and improved camera.', '$re$(?i)Redmi Note 12', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_redmi'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('poco_f1', 'Poco F1', 'Poco F1 with Snapdragon 845, offering flagship performance on a budget.', '$re$(?i)Poco F1', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_poco'), 2018)
, ('poco_f2', 'Poco F2 Pro', 'Poco F2 Pro with Snapdragon 865 and a full-screen design.', '$re$(?i)Poco F2', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_poco'), 2020)
, ('poco_x3', 'Poco X3', 'Poco X3 with 120Hz display and Snapdragon 732G.', '$re$(?i)Poco X3', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_poco'), 2020)
, ('poco_x4', 'Poco X4 Pro', 'Poco X4 Pro with 120Hz AMOLED and 108MP camera.', '$re$(?i)Poco X4', 'active', 'xiaomi', 'phone', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_poco'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('oneplus_5', 'OnePlus 5', 'OnePlus 5 with dual cameras and Snapdragon 835.', '$re$(?i)OnePlus 5', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2017)
, ('oneplus_6', 'OnePlus 6', 'OnePlus 6 with a notched display and Snapdragon 845.', '$re$(?i)OnePlus 6', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2018)
, ('oneplus_7', 'OnePlus 7', 'OnePlus 7 with pop-up camera and AMOLED display.', '$re$(?i)OnePlus 7', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2019)
, ('oneplus_8', 'OnePlus 8', 'OnePlus 8 with 5G support and 120Hz Fluid AMOLED display.', '$re$(?i)OnePlus 8', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2020)
, ('oneplus_9', 'OnePlus 9', 'OnePlus 9 with Hasselblad camera and Snapdragon 888.', '$re$(?i)OnePlus 9', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2021)
, ('oneplus_10', 'OnePlus 10', 'OnePlus 10 with improved camera and high refresh rate display.', '$re$(?i)OnePlus 10', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2022)
, ('oneplus_11', 'OnePlus 11', 'Latest OnePlus 11 with next-gen performance and design.', '$re$(?i)OnePlus 11', 'active', 'oneplus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oneplus'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('oppo_find_x', 'Oppo Find X', 'Oppo Find X with a motorized pop-up camera and bezel-less design.', '$re$(?i)Oppo Find X', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2018)
, ('oppo_find_x2', 'Oppo Find X2', 'Oppo Find X2 with 120Hz display and Snapdragon 865.', '$re$(?i)Oppo Find X2', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2020)
, ('oppo_find_x3', 'Oppo Find X3', 'Oppo Find X3 with Microlens camera and 1 billion colors display.', '$re$(?i)Oppo Find X3', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2021)
, ('oppo_find_x5', 'Oppo Find X5', 'Oppo Find X5 with advanced camera stabilization and AI photography.', '$re$(?i)Oppo Find X5', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2022)
, ('oppo_reno3', 'Oppo Reno 3', 'Oppo Reno 3 with quad-camera setup and AMOLED display.', '$re$(?i)Oppo Reno 3', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2019)
, ('oppo_reno4', 'Oppo Reno 4', 'Oppo Reno 4 with AI-enhanced photography and fast charging.', '$re$(?i)Oppo Reno 4', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2020)
, ('oppo_reno5', 'Oppo Reno 5', 'Oppo Reno 5 with 5G connectivity and AI-powered portrait video.', '$re$(?i)Oppo Reno 5', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2021)
, ('oppo_reno6', 'Oppo Reno 6', 'Oppo Reno 6 with Bokeh Flare Portrait and 65W fast charging.', '$re$(?i)Oppo Reno 6', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2021)
, ('oppo_reno7', 'Oppo Reno 7', 'Oppo Reno 7 with Sony IMX709 sensor and AI facial recognition.', '$re$(?i)Oppo Reno 7', 'active', 'oppo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'oppo'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('vivo_x50', 'Vivo X50', 'Vivo X50 with gimbal stabilization and AI-enhanced camera.', '$re$(?i)Vivo X50', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2020)
, ('vivo_x60', 'Vivo X60', 'Vivo X60 with ZEISS optics and advanced night photography.', '$re$(?i)Vivo X60', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2021)
, ('vivo_x70', 'Vivo X70', 'Vivo X70 with improved gimbal stabilization and AI enhancements.', '$re$(?i)Vivo X70', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2021)
, ('vivo_x80', 'Vivo X80', 'Vivo X80 with Dimensity 9000 and cinematic video features.', '$re$(?i)Vivo X80', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2022)
, ('vivo_x90', 'Vivo X90', 'Vivo X90 with ZEISS camera enhancements and next-gen performance.', '$re$(?i)Vivo X90', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2023)
, ('vivo_y50', 'Vivo Y50', 'Vivo Y50 with large battery and AI-powered quad camera.', '$re$(?i)Vivo Y50', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2020)
, ('vivo_y70', 'Vivo Y70', 'Vivo Y70 with AMOLED display and fast charging.', '$re$(?i)Vivo Y70', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2020)
, ('vivo_y91', 'Vivo Y91', 'Vivo Y91 with dual-camera setup and AI face unlock.', '$re$(?i)Vivo Y91', 'active', 'vivo', 'phone', (SELECT id FROM type_device_model WHERE codename = 'vivo'), 2019)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('motorola_edge_20', 'Motorola Edge 20', 'Motorola Edge 20 with 144Hz OLED display and Snapdragon 778G.', '$re$(?i)Motorola Edge 20', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2021)
, ('motorola_edge_30', 'Motorola Edge 30', 'Motorola Edge 30 with ultra-slim design and 50MP camera.', '$re$(?i)Motorola Edge 30', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2022)
, ('motorola_edge_40', 'Motorola Edge 40', 'Motorola Edge 40 with IP68 water resistance and curved OLED display.', '$re$(?i)Motorola Edge 40', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2023)
, ('motorola_g50', 'Motorola G50', 'Motorola G50 with 5G connectivity and 5000mAh battery.', '$re$(?i)Motorola G50', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2021)
, ('motorola_g100', 'Motorola G100', 'Motorola G100 with Snapdragon 870 and Ready For desktop mode.', '$re$(?i)Motorola G100', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2021)
, ('motorola_razr_2019', 'Motorola Razr (2019)', 'First foldable Motorola Razr with nostalgic design and modern tech.', '$re$(?i)Motorola Razr 2019', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2019)
, ('motorola_razr_2022', 'Motorola Razr (2022)', 'Latest foldable Motorola Razr with improved hinge and performance.', '$re$(?i)Motorola Razr 2022', 'active', 'motorola', 'phone', (SELECT id FROM type_device_model WHERE codename = 'motorola'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('xperia_xz', 'Sony Xperia XZ', 'Sony Xperia XZ with Motion Eye camera and water resistance.', '$re$(?i)Xperia XZ', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2016)
, ('xperia_1', 'Sony Xperia 1', 'Sony Xperia 1 with 4K HDR OLED display and triple-camera system.', '$re$(?i)Xperia 1', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2019)
, ('xperia_5', 'Sony Xperia 5', 'Sony Xperia 5 with compact design and cinematic display.', '$re$(?i)Xperia 5', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2019)
, ('xperia_10', 'Sony Xperia 10', 'Sony Xperia 10 with ultra-wide 21:9 display for immersive viewing.', '$re$(?i)Xperia 10', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2019)
, ('xperia_1_ii', 'Sony Xperia 1 II', 'Sony Xperia 1 II with Alpha camera technology and 5G support.', '$re$(?i)Xperia 1 II', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2020)
, ('xperia_5_ii', 'Sony Xperia 5 II', 'Sony Xperia 5 II with 120Hz display and gaming features.', '$re$(?i)Xperia 5 II', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2020)
, ('xperia_1_iii', 'Sony Xperia 1 III', 'Sony Xperia 1 III with real-time tracking autofocus and 120Hz 4K OLED display.', '$re$(?i)Xperia 1 III', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2021)
, ('xperia_5_iii', 'Sony Xperia 5 III', 'Sony Xperia 5 III with triple-camera and 120Hz HDR OLED display.', '$re$(?i)Xperia 5 III', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2021)
, ('xperia_1_iv', 'Sony Xperia 1 IV', 'Sony Xperia 1 IV with true optical zoom and 4K HDR 120Hz display.', '$re$(?i)Xperia 1 IV', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2022)
, ('xperia_5_iv', 'Sony Xperia 5 IV', 'Sony Xperia 5 IV with compact flagship performance and 4K HDR OLED.', '$re$(?i)Xperia 5 IV', 'active', 'sony', 'phone', (SELECT id FROM type_device_model WHERE codename = 'sony_xperia'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('nokia_6', 'Nokia 6', 'Nokia 6 with aluminum body and pure Android experience.', '$re$(?i)Nokia 6', 'active', 'nokia', 'phone', (SELECT id FROM type_device_model WHERE codename = 'nokia'), 2017)
, ('nokia_7', 'Nokia 7 Plus', 'Nokia 7 Plus with Zeiss optics and Android One.', '$re$(?i)Nokia 7 Plus', 'active', 'nokia', 'phone', (SELECT id FROM type_device_model WHERE codename = 'nokia'), 2018)
, ('nokia_8', 'Nokia 8', 'Nokia 8 with Dual-Sight mode and OZO Audio.', '$re$(?i)Nokia 8', 'active', 'nokia', 'phone', (SELECT id FROM type_device_model WHERE codename = 'nokia'), 2017)
, ('nokia_9_pureview', 'Nokia 9 PureView', 'Nokia 9 PureView with penta-camera system and HDR10 support.', '$re$(?i)Nokia 9 PureView', 'active', 'nokia', 'phone', (SELECT id FROM type_device_model WHERE codename = 'nokia'), 2019)
, ('nokia_xr20', 'Nokia XR20', 'Nokia XR20 with rugged design and long-term software updates.', '$re$(?i)Nokia XR20', 'active', 'nokia', 'phone', (SELECT id FROM type_device_model WHERE codename = 'nokia'), 2021)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('rog_phone_3', 'Asus ROG Phone 3', 'Asus ROG Phone 3 with 144Hz display and gaming triggers.', '$re$(?i)ROG Phone 3', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2020)
, ('rog_phone_5', 'Asus ROG Phone 5', 'Asus ROG Phone 5 with enhanced cooling and 6000mAh battery.', '$re$(?i)ROG Phone 5', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2021)
, ('rog_phone_6', 'Asus ROG Phone 6', 'Asus ROG Phone 6 with Snapdragon 8+ Gen 1 and AirTrigger controls.', '$re$(?i)ROG Phone 6', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2022)
, ('rog_phone_7', 'Asus ROG Phone 7', 'Asus ROG Phone 7 with improved cooling and advanced gaming features.', '$re$(?i)ROG Phone 7', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2023)
, ('zenfone_7', 'Asus Zenfone 7', 'Asus Zenfone 7 with motorized flip camera and AMOLED display.', '$re$(?i)Zenfone 7', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2020)
, ('zenfone_8', 'Asus Zenfone 8', 'Asus Zenfone 8 with compact design and Snapdragon 888.', '$re$(?i)Zenfone 8', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2021)
, ('zenfone_9', 'Asus Zenfone 9', 'Asus Zenfone 9 with gimbal-stabilized camera and flagship performance.', '$re$(?i)Zenfone 9', 'active', 'asus', 'phone', (SELECT id FROM type_device_model WHERE codename = 'asus'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('galaxy_tab_a', 'Samsung Galaxy Tab A', 'Affordable Samsung tablet for general use.', '$re$(?i)Galaxy Tab A', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2015)
, ('galaxy_tab_s6', 'Samsung Galaxy Tab S6', 'Samsung Galaxy Tab S6 with S Pen support and AMOLED display.', '$re$(?i)Galaxy Tab S6', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2019)
, ('galaxy_tab_s7', 'Samsung Galaxy Tab S7', 'Samsung Galaxy Tab S7 with 120Hz display and high-performance chipset.', '$re$(?i)Galaxy Tab S7', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2020)
, ('galaxy_tab_s8', 'Samsung Galaxy Tab S8', 'Samsung Galaxy Tab S8 with enhanced multitasking features.', '$re$(?i)Galaxy Tab S8', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2022)
, ('galaxy_tab_s9', 'Samsung Galaxy Tab S9', 'Samsung Galaxy Tab S9 with upgraded features and S Pen support.', '$re$(?i)Galaxy Tab S9', 'active', 'samsung', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'galaxy_tab'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('matepad_pro', 'Huawei MatePad Pro', 'Huawei MatePad Pro with stylus support and high-resolution display.', '$re$(?i)MatePad Pro', 'active', 'huawei', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'matepad'), 2019)
, ('matepad_10_4', 'Huawei MatePad 10.4', 'Huawei MatePad 10.4 with mid-range performance and stylus support.', '$re$(?i)MatePad 10.4', 'active', 'huawei', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'matepad'), 2020)
, ('matepad_11', 'Huawei MatePad 11', 'Huawei MatePad 11 with HarmonyOS and 120Hz display.', '$re$(?i)MatePad 11', 'active', 'huawei', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'matepad'), 2021)
, ('matepad_12_6', 'Huawei MatePad 12.6', 'Huawei MatePad 12.6 with OLED display and PC-like experience.', '$re$(?i)MatePad 12.6', 'active', 'huawei', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'matepad'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('mi_pad_4', 'Xiaomi Mi Pad 4', 'Xiaomi Mi Pad 4 with Snapdragon 660 and 8-inch display.', '$re$(?i)Mi Pad 4', 'active', 'xiaomi', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_pad'), 2018)
, ('mi_pad_5', 'Xiaomi Mi Pad 5', 'Xiaomi Mi Pad 5 with 120Hz display and Snapdragon 860.', '$re$(?i)Mi Pad 5', 'active', 'xiaomi', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_pad'), 2021)
, ('xiaomi_pad_6', 'Xiaomi Pad 6', 'Xiaomi Pad 6 with improved performance and 144Hz display.', '$re$(?i)Xiaomi Pad 6', 'active', 'xiaomi', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'xiaomi_pad'), 2023)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('lenovo_tab_p11', 'Lenovo Tab P11', 'Lenovo Tab P11 with 2K display and stylus support.', '$re$(?i)Lenovo Tab P11', 'active', 'lenovo', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'lenovo_tablets'), 2021)
, ('lenovo_tab_p12', 'Lenovo Tab P12 Pro', 'Lenovo Tab P12 Pro with AMOLED display and premium build.', '$re$(?i)Lenovo Tab P12', 'active', 'lenovo', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'lenovo_tablets'), 2022)
, ('lenovo_tab_m10', 'Lenovo Tab M10', 'Lenovo Tab M10 with mid-range performance and kids mode.', '$re$(?i)Lenovo Tab M10', 'active', 'lenovo', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'lenovo_tablets'), 2020)
, ('lenovo_tab_m8', 'Lenovo Tab M8', 'Lenovo Tab M8 with compact design and multimedia focus.', '$re$(?i)Lenovo Tab M8', 'active', 'lenovo', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'lenovo_tablets'), 2019)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('surface_go', 'Microsoft Surface Go', 'Compact and affordable Surface tablet for everyday use.', '$re$(?i)Surface Go', 'active', 'microsoft', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'microsoft_surface'), 2018)
, ('surface_pro_7', 'Microsoft Surface Pro 7', 'Microsoft Surface Pro 7 with 10th Gen Intel processors and USB-C.', '$re$(?i)Surface Pro 7', 'active', 'microsoft', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'microsoft_surface'), 2019)
, ('surface_pro_8', 'Microsoft Surface Pro 8', 'Microsoft Surface Pro 8 with larger display and improved performance.', '$re$(?i)Surface Pro 8', 'active', 'microsoft', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'microsoft_surface'), 2021)
, ('surface_pro_9', 'Microsoft Surface Pro 9', 'Latest Surface Pro 9 with ARM and Intel processor options.', '$re$(?i)Surface Pro 9', 'active', 'microsoft', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'microsoft_surface'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('fire_7', 'Amazon Fire 7', 'Compact 7-inch Fire tablet with Alexa integration.', '$re$(?i)Fire 7', 'active', 'amazon', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'amazon_fire'), 2019)
, ('fire_hd_8', 'Amazon Fire HD 8', '8-inch HD Fire tablet with Dolby Audio and long battery life.', '$re$(?i)Fire HD 8', 'active', 'amazon', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'amazon_fire'), 2020)
, ('fire_hd_10', 'Amazon Fire HD 10', '10-inch Full HD Fire tablet with powerful performance and Alexa hands-free.', '$re$(?i)Fire HD 10', 'active', 'amazon', 'tablet', (SELECT id FROM type_device_model WHERE codename = 'amazon_fire'), 2021)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('xbox_one', 'Xbox One', 'Xbox One with multimedia integration and cloud gaming features.', '$re$(?i)Xbox One', 'active', 'microsoft', 'connected', (SELECT id FROM type_device_model WHERE codename = 'xbox'), 2013)
, ('xbox_one_s', 'Xbox One S', 'Xbox One S with HDR support and compact design.', '$re$(?i)Xbox One S', 'active', 'microsoft', 'connected', (SELECT id FROM type_device_model WHERE codename = 'xbox'), 2016)
, ('xbox_one_x', 'Xbox One X', 'Xbox One X with 4K gaming and enhanced graphics capabilities.', '$re$(?i)Xbox One X', 'active', 'microsoft', 'connected', (SELECT id FROM type_device_model WHERE codename = 'xbox'), 2017)
, ('xbox_series_s', 'Xbox Series S', 'Xbox Series S with next-gen performance in a compact form factor.', '$re$(?i)Xbox Series S', 'active', 'microsoft', 'connected', (SELECT id FROM type_device_model WHERE codename = 'xbox'), 2020)
, ('xbox_series_x', 'Xbox Series X', 'Xbox Series X with powerful GPU and ray-tracing support.', '$re$(?i)Xbox Series X', 'active', 'microsoft', 'connected', (SELECT id FROM type_device_model WHERE codename = 'xbox'), 2020)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('ps3', 'PlayStation 3', 'PlayStation 3 with Blu-ray support and online gaming capabilities.', '$re$(?i)PlayStation 3', 'active', 'sony', 'connected', (SELECT id FROM type_device_model WHERE codename = 'playstation'), 2006)
, ('ps4', 'PlayStation 4', 'PlayStation 4 with enhanced graphics and VR support.', '$re$(?i)PlayStation 4', 'active', 'sony', 'connected', (SELECT id FROM type_device_model WHERE codename = 'playstation'), 2013)
, ('ps4_pro', 'PlayStation 4 Pro', 'PlayStation 4 Pro with 4K gaming and improved hardware.', '$re$(?i)PlayStation 4 Pro', 'active', 'sony', 'connected', (SELECT id FROM type_device_model WHERE codename = 'playstation'), 2016)
, ('ps5', 'PlayStation 5', 'PlayStation 5 with ultra-fast SSD and ray-tracing technology.', '$re$(?i)PlayStation 5', 'active', 'sony', 'connected', (SELECT id FROM type_device_model WHERE codename = 'playstation'), 2020)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('wii', 'Nintendo Wii', 'Nintendo Wii with motion controls and a family-friendly game library.', '$re$(?i)Wii', 'active', 'nintendo', 'connected', (SELECT id FROM type_device_model WHERE codename = 'nintendo'), 2006)
, ('wii_u', 'Nintendo Wii U', 'Nintendo Wii U with a gamepad controller and HD gaming.', '$re$(?i)Wii U', 'active', 'nintendo', 'connected', (SELECT id FROM type_device_model WHERE codename = 'nintendo'), 2012)
, ('switch', 'Nintendo Switch', 'Nintendo Switch with hybrid handheld and console gaming.', '$re$(?i)Switch', 'active', 'nintendo', 'connected', (SELECT id FROM type_device_model WHERE codename = 'nintendo'), 2017)
, ('switch_lite', 'Nintendo Switch Lite', 'Nintendo Switch Lite, a handheld-only version of the Switch.', '$re$(?i)Switch Lite', 'active', 'nintendo', 'connected', (SELECT id FROM type_device_model WHERE codename = 'nintendo'), 2019)
, ('switch_oled', 'Nintendo Switch OLED Model', 'Nintendo Switch OLED Model with a larger, more vibrant display.', '$re$(?i)Switch OLED', 'active', 'nintendo', 'connected', (SELECT id FROM type_device_model WHERE codename = 'nintendo'), 2021)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('apple_tv_hd', 'Apple TV HD', 'Apple TV HD with 1080p resolution and access to Apple services.', '$re$(?i)Apple TV HD', 'active', 'apple', 'connected', (SELECT id FROM type_device_model WHERE codename = 'apple_tv'), 2015)
, ('apple_tv_4k_1st_gen', 'Apple TV 4K (1st generation)', 'First-generation Apple TV 4K with HDR and Dolby Vision support.', '$re$(?i)Apple TV 4K 1st Gen', 'active', 'apple', 'connected', (SELECT id FROM type_device_model WHERE codename = 'apple_tv'), 2017)
, ('apple_tv_4k_2nd_gen', 'Apple TV 4K (2nd generation)', 'Second-generation Apple TV 4K with improved processor and HDMI 2.1.', '$re$(?i)Apple TV 4K 2nd Gen', 'active', 'apple', 'connected', (SELECT id FROM type_device_model WHERE codename = 'apple_tv'), 2021)
, ('apple_tv_4k_3rd_gen', 'Apple TV 4K (3rd generation)', 'Latest Apple TV 4K with A15 Bionic chip and USB-C remote charging.', '$re$(?i)Apple TV 4K 3rd Gen', 'active', 'apple', 'connected', (SELECT id FROM type_device_model WHERE codename = 'apple_tv'), 2022)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('samsung_qled_4k', 'Samsung QLED 4K', 'Samsung QLED 4K TV with Quantum Dot technology and HDR support.', '$re$(?i)Samsung QLED 4K', 'active', 'samsung', 'connected', (SELECT id FROM type_device_model WHERE codename = 'samsung_smart_tv'), 2017)
, ('samsung_qled_8k', 'Samsung QLED 8K', 'Samsung QLED 8K TV with AI upscaling and superior picture quality.', '$re$(?i)Samsung QLED 8K', 'active', 'samsung', 'connected', (SELECT id FROM type_device_model WHERE codename = 'samsung_smart_tv'), 2018)
, ('samsung_neo_qled', 'Samsung Neo QLED', 'Samsung Neo QLED with Mini LED technology for improved contrast and brightness.', '$re$(?i)Samsung Neo QLED', 'active', 'samsung', 'connected', (SELECT id FROM type_device_model WHERE codename = 'samsung_smart_tv'), 2021)
, ('samsung_the_frame', 'Samsung The Frame', 'Samsung The Frame, a lifestyle TV that doubles as a digital art display.', '$re$(?i)Samsung The Frame', 'active', 'samsung', 'connected', (SELECT id FROM type_device_model WHERE codename = 'samsung_smart_tv'), 2017)
, ('samsung_uhd_4k', 'Samsung UHD 4K', 'Samsung UHD 4K Smart TV with Tizen OS and HDR support.', '$re$(?i)Samsung UHD 4K', 'active', 'samsung', 'connected', (SELECT id FROM type_device_model WHERE codename = 'samsung_smart_tv'), 2019)
  ON CONFLICT (codename) DO NOTHING;

-------------------------

INSERT INTO type_device_model (codename, name, description, match_exp, active, maker_codename, type_codename, parent_id, year_release) VALUES
  ('lg_oled_c1', 'LG OLED C1', 'LG OLED C1 with self-lit pixels and NVIDIA G-Sync support.', '$re$(?i)LG OLED C1', 'active', 'lg', 'connected', (SELECT id FROM type_device_model WHERE codename = 'lg_smart_tv'), 2021)
, ('lg_oled_c2', 'LG OLED C2', 'LG OLED C2 with next-gen Evo panel for improved brightness.', '$re$(?i)LG OLED C2', 'active', 'lg', 'connected', (SELECT id FROM type_device_model WHERE codename = 'lg_smart_tv'), 2022)
, ('lg_qned', 'LG QNED Mini LED', 'LG QNED Mini LED TV with Quantum Dot and NanoCell technology.', '$re$(?i)LG QNED', 'active', 'lg', 'connected', (SELECT id FROM type_device_model WHERE codename = 'lg_smart_tv'), 2021)
, ('lg_nano', 'LG NanoCell', 'LG NanoCell Smart TV with IPS panel and AI ThinQ.', '$re$(?i)LG NanoCell', 'active', 'lg', 'connected', (SELECT id FROM type_device_model WHERE codename = 'lg_smart_tv'), 2019)
, ('lg_uhd_4k', 'LG UHD 4K', 'LG UHD 4K Smart TV with webOS and AI upscaling.', '$re$(?i)LG UHD 4K', 'active', 'lg', 'connected', (SELECT id FROM type_device_model WHERE codename = 'lg_smart_tv'), 2020)
  ON CONFLICT (codename) DO NOTHING;
