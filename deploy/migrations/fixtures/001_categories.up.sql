BEGIN;
INSERT INTO adv_category (name, description, iab_code, position)
VALUES
  ('Arts & Entertainment', 'This category includes all content related to visual arts, music, theater, movies, and other forms of entertainment. It covers topics such as performance art, exhibitions, and cultural events.', 'IAB1', 1),
  ('Automotive', 'This category encompasses all aspects of the automotive industry including cars, trucks, parts, repairs, and innovations in vehicle technology.', 'IAB2', 2),
  ('Business', 'This category covers business-related topics including corporate news, management strategies, market trends, and economic insights relevant to entrepreneurs and corporations.', 'IAB3', 3),
  ('Careers', 'This category provides resources and advice for career planning, job search strategies, resume writing, and professional development across various industries.', 'IAB4', 4),
  ('Education', 'This category includes topics related to educational systems, academic advice, learning resources, and higher education, ranging from K-12 to college and beyond.', 'IAB5', 5),
  ('Family & Parenting', 'This category covers content related to family life and parenting, including child care, family relationships, and guidance for raising children and managing household dynamics.', 'IAB6', 6),
  ('Health & Fitness', 'This category features health advice, fitness tips, medical news, wellness practices, and information on nutrition, exercise, and lifestyle improvements.', 'IAB7', 7),
  ('Food & Drink', 'This category includes culinary content such as recipes, restaurant reviews, cooking techniques, and discussions about various cuisines and beverages.', 'IAB8', 8),
  ('Hobbies & Interests', 'This category covers a wide range of personal interests and leisure activities including arts, crafts, gaming, and other recreational pursuits.', 'IAB9', 9),
  ('Home & Garden', 'This category provides content related to home improvement, interior design, gardening tips, and DIY projects to enhance living spaces.', 'IAB10', 10),
  ('Law, Government, & Politics', 'This category focuses on legal issues, governmental policies, political analysis, and the impact of public policy on society.', 'IAB11', 11),
  ('News', 'This category delivers current events, breaking news, and in-depth reporting on local, national, and international affairs.', 'IAB12', 12),
  ('Personal Finance', 'This category offers advice on managing personal finances, investing, budgeting, retirement planning, and navigating financial markets.', 'IAB13', 13),
  ('Society', 'This category covers social issues, cultural trends, relationship dynamics, and discussions on societal norms and values.', 'IAB14', 14),
  ('Science', 'This category explores topics in science, including research breakthroughs, technological innovations, and studies in fields such as biology, physics, and chemistry.', 'IAB15', 15),
  ('Pets', 'This category is dedicated to pet care, animal health, and advice on raising and training pets, as well as information on different pet species.', 'IAB16', 16),
  ('Sports', 'This category includes coverage of sporting events, athletic news, game analyses, and discussions on various sports and related activities.', 'IAB17', 17),
  ('Style & Fashion', 'This category covers trends in fashion, beauty, personal style, and the latest in clothing, accessories, and lifestyle aesthetics.', 'IAB18', 18),
  ('Technology & Computing', 'This category features content on technology trends, computing advancements, software reviews, hardware innovations, and digital culture.', 'IAB19', 19),
  ('Travel', 'This category provides travel guides, destination reviews, tips for planning trips, and insights into various travel experiences around the world.', 'IAB20', 20),
  ('Real Estate', 'This category covers topics in real estate including property buying and selling, market trends, and advice on home investments.', 'IAB21', 21),
  ('Shopping', 'This category offers insights into consumer shopping trends, product reviews, discounts, and comparisons across various retail sectors.', 'IAB22', 22),
  ('Religion & Spirituality', 'This category explores topics related to religious beliefs, spiritual practices, and philosophical discussions across different faiths.', 'IAB23', 23),
  ('Uncategorized', 'This category is used for content that does not fit into any other predefined category and remains uncategorized.', 'IAB24', 24),
  ('Non-Standard Content', 'This category includes content that deviates from traditional standards, such as user-generated content and non-conventional media.', 'IAB25', 25),
  ('Illegal Content', 'This category covers content that is prohibited or falls under illegal activities, including copyright infringement and other unlawful material.', 'IAB26', 26)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Arts & Entertainment (IAB1)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Books & Literature', 'This subcategory under Arts & Entertainment focuses on literature, including books and literary reviews.', 'IAB1-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 1),
  ('Celebrity Fan/Gossip', 'This subcategory covers celebrity news and gossip within the Arts & Entertainment domain.', 'IAB1-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 2),
  ('Fine Art', 'This subcategory is dedicated to fine art, including painting, sculpture, and art exhibitions under Arts & Entertainment.', 'IAB1-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 3),
  ('Humor', 'This subcategory highlights comedic content and humor in the context of Arts & Entertainment.', 'IAB1-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 4),
  ('Movies', 'This subcategory focuses on films and cinematic content, reviews, and movie industry news.', 'IAB1-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 5),
  ('Music', 'This subcategory covers music genres, industry news, and related entertainment topics.', 'IAB1-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 6),
  ('Television', 'This subcategory encompasses TV shows, broadcast news, and other television-related content.', 'IAB1-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB1'), 7)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Automotive (IAB2)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Auto Parts', 'This subcategory under Automotive focuses on auto parts, components, and accessories.', 'IAB2-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 1),
  ('Auto Repair', 'This subcategory covers repair services, maintenance tips, and troubleshooting for vehicles.', 'IAB2-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 2),
  ('Buying/Selling Cars', 'This subcategory provides information on buying and selling vehicles, including market trends and advice.', 'IAB2-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 3),
  ('Car Culture', 'This subcategory explores the lifestyle, trends, and cultural aspects surrounding automobiles.', 'IAB2-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 4),
  ('Certified Pre-Owned', 'This subcategory focuses on certified pre-owned vehicles and related buying advice.', 'IAB2-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 5),
  ('Convertible', 'This subcategory is dedicated to convertible cars, their features, and market reviews.', 'IAB2-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 6),
  ('Coupe', 'This subcategory covers coupe-style vehicles and their performance features.', 'IAB2-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 7),
  ('Crossover', 'This subcategory focuses on crossover vehicles, blending SUV and car characteristics.', 'IAB2-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 8),
  ('Diesel', 'This subcategory covers diesel-powered vehicles, their efficiency, and performance.', 'IAB2-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 9),
  ('Electric Vehicle', 'This subcategory focuses on electric vehicles, including technology, innovations, and market trends.', 'IAB2-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 10),
  ('Hatchback', 'This subcategory covers hatchback vehicles and their design features.', 'IAB2-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 11),
  ('Hybrid', 'This subcategory is dedicated to hybrid vehicles and their environmental and performance benefits.', 'IAB2-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 12),
  ('Luxury', 'This subcategory focuses on luxury automobiles, premium features, and high-end market trends.', 'IAB2-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 13),
  ('Minivan', 'This subcategory covers minivans, family vehicles, and related market advice.', 'IAB2-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 14),
  ('Motorcycles', 'This subcategory focuses on motorcycles, including news, reviews, and maintenance tips.', 'IAB2-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 15),
  ('Off-Road Vehicles', 'This subcategory covers off-road vehicles and related adventure and utility topics.', 'IAB2-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 16),
  ('Performance Vehicles', 'This subcategory is dedicated to high-performance vehicles and their engineering.', 'IAB2-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 17),
  ('Pickup', 'This subcategory focuses on pickup trucks, including features, reviews, and market trends.', 'IAB2-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 18),
  ('Road-Side Assistance', 'This subcategory provides information on roadside assistance services and emergency repair.', 'IAB2-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 19),
  ('Sedan', 'This subcategory covers sedan vehicles and their design, performance, and market positioning.', 'IAB2-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 20),
  ('Trucks & Accessories', 'This subcategory focuses on trucks and related accessories, including modifications and aftermarket parts.', 'IAB2-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 21),
  ('Vintage Cars', 'This subcategory covers vintage cars, classic models, and restoration projects.', 'IAB2-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 22),
  ('Wagon', 'This subcategory is dedicated to wagon-style vehicles and their practical features.', 'IAB2-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB2'), 23)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Business (IAB3)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Advertising', 'This subcategory under Business focuses on advertising strategies, trends, and industry news.', 'IAB3-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 1),
  ('Agriculture', 'This subcategory covers agricultural business, market trends, and farming innovations.', 'IAB3-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 2),
  ('Biotech/Biomedical', 'This subcategory focuses on the biotech and biomedical sectors, including research and industry developments.', 'IAB3-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 3),
  ('Business Software', 'This subcategory covers software solutions and technology trends in business management.', 'IAB3-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 4),
  ('Construction', 'This subcategory focuses on construction industry news, projects, and market trends.', 'IAB3-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 5),
  ('Forestry', 'This subcategory covers forestry, lumber, and related business topics in natural resources.', 'IAB3-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 6),
  ('Government', 'This subcategory focuses on government-related business issues and regulatory news.', 'IAB3-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 7),
  ('Green Solutions', 'This subcategory covers sustainable business practices and green technology initiatives.', 'IAB3-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 8),
  ('Human Resources', 'This subcategory provides insights into human resource management, recruitment, and workplace culture.', 'IAB3-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 9),
  ('Logistics', 'This subcategory focuses on logistics, supply chain management, and transportation solutions.', 'IAB3-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 10),
  ('Marketing', 'This subcategory covers marketing strategies, digital marketing trends, and brand management.', 'IAB3-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 11),
  ('Metals', 'This subcategory focuses on the metals industry, including market trends and raw material analysis.', 'IAB3-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB3'), 12)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Careers (IAB4)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Career Planning', 'This subcategory under Careers provides guidance and strategies for career planning and development.', 'IAB4-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 1),
  ('College', 'This subcategory focuses on college-related topics including admissions, campus life, and academic advice.', 'IAB4-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 2),
  ('Financial Aid', 'This subcategory covers financial aid, scholarships, and funding options for education and careers.', 'IAB4-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 3),
  ('Job Fairs', 'This subcategory provides information on job fairs, networking events, and career expos.', 'IAB4-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 4),
  ('Job Search', 'This subcategory focuses on job search strategies, resume tips, and interview preparation.', 'IAB4-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 5),
  ('Resume Writing/Advice', 'This subcategory offers advice on crafting effective resumes and cover letters.', 'IAB4-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 6),
  ('Nursing', 'This subcategory covers career opportunities and advice specifically for the nursing field.', 'IAB4-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 7),
  ('Scholarships', 'This subcategory focuses on scholarships and financial support options for students and professionals.', 'IAB4-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 8),
  ('Telecommuting', 'This subcategory covers remote work opportunities and telecommuting best practices.', 'IAB4-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 9),
  ('U.S. Military', 'This subcategory provides information on career opportunities and guidance within the U.S. military.', 'IAB4-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 10),
  ('Career Advice', 'This subcategory offers general career advice, mentoring, and professional development tips.', 'IAB4-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB4'), 11)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Education (IAB5)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('7-12 Education', 'This subcategory under Education focuses on primary and secondary education topics.', 'IAB5-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 1),
  ('Adult Education', 'This subcategory covers continuing education and learning opportunities for adults.', 'IAB5-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 2),
  ('Art History', 'This subcategory focuses on art history as an academic subject and cultural study.', 'IAB5-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 3),
  ('College Administration', 'This subcategory covers topics related to college administration, policies, and management.', 'IAB5-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 4),
  ('College Life', 'This subcategory provides insights into college life, student experiences, and campus activities.', 'IAB5-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 5),
  ('Distance Learning', 'This subcategory focuses on distance learning and online education methods.', 'IAB5-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 6),
  ('English as a 2nd Language', 'This subcategory covers resources and tips for learning English as a second language.', 'IAB5-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 7),
  ('Language Learning', 'This subcategory provides information and resources for learning various languages.', 'IAB5-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 8),
  ('Graduate School', 'This subcategory focuses on graduate school admissions, experiences, and academic advice.', 'IAB5-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 9),
  ('Homeschooling', 'This subcategory covers homeschooling techniques, curricula, and family education strategies.', 'IAB5-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 10),
  ('Homework/Study Tips', 'This subcategory offers tips and strategies for effective studying and homework completion.', 'IAB5-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 11),
  ('K-6 Educators', 'This subcategory is designed for educators teaching K-6 and provides resources and teaching strategies.', 'IAB5-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 12),
  ('Private School', 'This subcategory covers topics related to private schooling, admissions, and educational standards.', 'IAB5-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 13),
  ('Special Education', 'This subcategory focuses on special education, inclusive teaching practices, and support strategies.', 'IAB5-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 14),
  ('Studying Business', 'This subcategory provides resources for studying business, including coursework and industry insights.', 'IAB5-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB5'), 15)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Family & Parenting (IAB6)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Adoption', 'This subcategory under Family & Parenting covers topics related to adoption and related legal processes.', 'IAB6-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 1),
  ('Babies & Toddlers', 'This subcategory focuses on the care, development, and needs of babies and toddlers.', 'IAB6-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 2),
  ('Daycare/Pre School', 'This subcategory covers daycare options, pre-school education, and early childhood care.', 'IAB6-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 3),
  ('Family Internet', 'This subcategory discusses safe internet practices and digital resources for families.', 'IAB6-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 4),
  ('Parenting - K-6 Kids', 'This subcategory provides advice and resources for parenting young children in kindergarten to grade 6.', 'IAB6-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 5),
  ('Parenting teens', 'This subcategory focuses on the challenges and advice related to parenting teenagers.', 'IAB6-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 6),
  ('Pregnancy', 'This subcategory covers topics related to pregnancy, prenatal care, and expecting parents.', 'IAB6-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 7),
  ('Special Needs Kids', 'This subcategory provides resources and support for families with special needs children.', 'IAB6-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 8),
  ('Eldercare', 'This subcategory focuses on care and support for elderly family members.', 'IAB6-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB6'), 9)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Health & Fitness (IAB7)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Exercise', 'This subcategory under Health & Fitness focuses on physical exercise routines and workout tips.', 'IAB7-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 1),
  ('ADD', 'This subcategory covers Attention Deficit Disorder, its management, and related therapies.', 'IAB7-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 2),
  ('AIDS/HIV', 'This subcategory provides information and updates on AIDS/HIV research and treatment.', 'IAB7-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 3),
  ('Allergies', 'This subcategory covers information on allergies, symptoms, and treatments.', 'IAB7-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 4),
  ('Alternative Medicine', 'This subcategory focuses on alternative medicine practices and holistic healing.', 'IAB7-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 5),
  ('Arthritis', 'This subcategory covers arthritis, its management, and treatment options.', 'IAB7-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 6),
  ('Asthma', 'This subcategory provides information on asthma, its triggers, and management techniques.', 'IAB7-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 7),
  ('Autism/PDD', 'This subcategory focuses on autism and pervasive developmental disorders, including support and resources.', 'IAB7-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 8),
  ('Bipolar Disorder', 'This subcategory covers bipolar disorder, its treatment, and management strategies.', 'IAB7-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 9),
  ('Brain Tumor', 'This subcategory provides information on brain tumors, research, and treatment options.', 'IAB7-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 10),
  ('Cancer', 'This subcategory focuses on cancer research, treatment options, and patient support.', 'IAB7-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 11),
  ('Cholesterol', 'This subcategory covers cholesterol management, diet tips, and related health advice.', 'IAB7-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 12),
  ('Chronic Fatigue Syndrome', 'This subcategory provides information on chronic fatigue syndrome and its management.', 'IAB7-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 13),
  ('Chronic Pain', 'This subcategory covers chronic pain management, therapies, and patient resources.', 'IAB7-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 14),
  ('Cold & Flu', 'This subcategory focuses on common illnesses such as colds and flu, including prevention and treatment.', 'IAB7-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 15),
  ('Deafness', 'This subcategory covers deafness, hearing loss, and related assistive technologies.', 'IAB7-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 16),
  ('Dental Care', 'This subcategory provides information on dental health, care routines, and treatments.', 'IAB7-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 17),
  ('Depression', 'This subcategory focuses on depression, mental health support, and treatment options.', 'IAB7-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 18),
  ('Dermatology', 'This subcategory covers skin care, dermatological treatments, and cosmetic procedures.', 'IAB7-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 19),
  ('Diabetes', 'This subcategory focuses on diabetes management, diet, and treatment strategies.', 'IAB7-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 20),
  ('Epilepsy', 'This subcategory covers epilepsy, seizure management, and patient support resources.', 'IAB7-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 21),
  ('GERD/Acid Reflux', 'This subcategory focuses on GERD and acid reflux, including dietary tips and treatments.', 'IAB7-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 22),
  ('Headaches/Migraines', 'This subcategory covers headaches and migraines, including causes and treatment options.', 'IAB7-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 23),
  ('Heart Disease', 'This subcategory provides information on heart disease, risk factors, and preventive measures.', 'IAB7-24', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 24),
  ('Herbs for Health', 'This subcategory focuses on the use of herbs in promoting health and wellness.', 'IAB7-25', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 25),
  ('Holistic Healing', 'This subcategory covers holistic healing practices and alternative health therapies.', 'IAB7-26', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 26),
  ('IBS/Crohn''s Disease', 'This subcategory focuses on digestive disorders like IBS and Crohn''s disease, including management tips.', 'IAB7-27', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 27),
  ('Incest/Abuse Support', 'This subcategory provides support and resources for survivors of abuse.', 'IAB7-28', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 28),
  ('Incontinence', 'This subcategory covers incontinence issues and treatment options.', 'IAB7-29', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 29),
  ('Infertility', 'This subcategory focuses on infertility, treatment options, and support resources.', 'IAB7-30', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 30),
  ('Men''s Health', 'This subcategory covers men''s health issues, wellness tips, and medical advice.', 'IAB7-31', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 31),
  ('Nutrition', 'This subcategory focuses on nutrition, dietary advice, and healthy eating practices.', 'IAB7-32', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 32),
  ('Orthopedics', 'This subcategory covers orthopedic health, joint care, and related treatments.', 'IAB7-33', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 33),
  ('Panic/Anxiety Disorders', 'This subcategory focuses on panic and anxiety disorders, providing support and treatment options.', 'IAB7-34', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 34),
  ('Pediatrics', 'This subcategory covers pediatric health, child wellness, and medical care for children.', 'IAB7-35', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 35),
  ('Physical Therapy', 'This subcategory focuses on physical therapy, rehabilitation, and recovery programs.', 'IAB7-36', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 36),
  ('Psychology/Psychiatry', 'This subcategory covers mental health, psychology, and psychiatric treatment options.', 'IAB7-37', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 37),
  ('Senior Health', 'This subcategory focuses on health issues related to senior citizens and aging.', 'IAB7-38', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 38),
  ('Sexuality', 'This subcategory covers topics related to sexuality, sexual health, and relationships.', 'IAB7-39', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 39),
  ('Sleep Disorders', 'This subcategory focuses on sleep disorders, their causes, and treatment options.', 'IAB7-40', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 40),
  ('Smoking Cessation', 'This subcategory provides resources and tips for quitting smoking.', 'IAB7-41', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 41),
  ('Substance Abuse', 'This subcategory covers substance abuse issues, treatment programs, and recovery support.', 'IAB7-42', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 42),
  ('Thyroid Disease', 'This subcategory focuses on thyroid health, disorders, and treatment options.', 'IAB7-43', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 43),
  ('Weight Loss', 'This subcategory covers weight loss strategies, diets, and fitness programs.', 'IAB7-44', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 44),
  ('Women''s Health', 'This subcategory focuses on women''s health issues, wellness, and medical care.', 'IAB7-45', (SELECT id FROM adv_category WHERE iab_code = 'IAB7'), 45)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Food & Drink (IAB8)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('American Cuisine', 'This subcategory under Food & Drink focuses on American culinary traditions and recipes.', 'IAB8-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 1),
  ('Barbecues & Grilling', 'This subcategory covers barbecue techniques, grilling recipes, and outdoor cooking.', 'IAB8-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 2),
  ('Cajun/Creole', 'This subcategory focuses on Cajun and Creole cuisines, highlighting their unique flavors and dishes.', 'IAB8-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 3),
  ('Chinese Cuisine', 'This subcategory covers Chinese culinary traditions, recipes, and cooking techniques.', 'IAB8-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 4),
  ('Cocktails/Beer', 'This subcategory focuses on cocktail recipes, beer reviews, and mixology tips.', 'IAB8-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 5),
  ('Coffee/Tea', 'This subcategory covers topics related to coffee and tea, including brewing techniques and product reviews.', 'IAB8-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 6),
  ('Cuisine-Specific', 'This subcategory explores cuisine-specific topics, highlighting regional and ethnic culinary traditions.', 'IAB8-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 7),
  ('Desserts & Baking', 'This subcategory focuses on desserts, baking techniques, and sweet culinary creations.', 'IAB8-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 8),
  ('Dining Out', 'This subcategory covers restaurant reviews, dining experiences, and food critique.', 'IAB8-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 9),
  ('Food Allergies', 'This subcategory provides information on food allergies, dietary restrictions, and safe eating practices.', 'IAB8-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 10),
  ('French Cuisine', 'This subcategory focuses on French culinary traditions, recipes, and cooking techniques.', 'IAB8-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 11),
  ('Health/Low-Fat Cooking', 'This subcategory covers healthy cooking methods, low-fat recipes, and nutritional advice.', 'IAB8-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 12),
  ('Italian Cuisine', 'This subcategory focuses on Italian culinary traditions, pasta dishes, and regional specialties.', 'IAB8-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 13),
  ('Japanese Cuisine', 'This subcategory covers Japanese culinary arts, sushi, ramen, and traditional recipes.', 'IAB8-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 14),
  ('Mexican Cuisine', 'This subcategory focuses on Mexican food, recipes, and culinary traditions.', 'IAB8-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 15),
  ('Vegan', 'This subcategory covers vegan cooking, recipes, and plant-based dietary tips.', 'IAB8-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 16),
  ('Vegetarian', 'This subcategory focuses on vegetarian cuisine, recipes, and healthy eating practices.', 'IAB8-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 17),
  ('Wine', 'This subcategory covers wine reviews, tasting notes, and vineyard news.', 'IAB8-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB8'), 18)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Hobbies & Interests (IAB9)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Art/Technology', 'This subcategory under Hobbies & Interests focuses on the intersection of art and technology.', 'IAB9-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 1),
  ('Arts & Crafts', 'This subcategory covers various arts and crafts activities, DIY projects, and creative hobbies.', 'IAB9-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 2),
  ('Beadwork', 'This subcategory focuses on beadwork techniques, jewelry making, and decorative crafts.', 'IAB9-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 3),
  ('Bird-Watching', 'This subcategory covers bird-watching tips, guides, and nature observation techniques.', 'IAB9-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 4),
  ('Board Games/Puzzles', 'This subcategory focuses on board games, puzzles, and related recreational activities.', 'IAB9-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 5),
  ('Candle & Soap Making', 'This subcategory covers techniques for making candles and soaps as creative hobbies.', 'IAB9-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 6),
  ('Card Games', 'This subcategory focuses on card games, strategies, and related entertainment.', 'IAB9-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 7),
  ('Chess', 'This subcategory covers chess strategies, tournaments, and learning resources.', 'IAB9-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 8),
  ('Cigars', 'This subcategory focuses on cigars, including reviews, history, and smoking culture.', 'IAB9-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 9),
  ('Collecting', 'This subcategory covers various collecting hobbies, from stamps to antiques.', 'IAB9-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 10),
  ('Comic Books', 'This subcategory focuses on comic books, graphic novels, and related art forms.', 'IAB9-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 11),
  ('Drawing/Sketching', 'This subcategory covers drawing techniques, sketching, and artistic expression.', 'IAB9-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 12),
  ('Freelance Writing', 'This subcategory focuses on freelance writing, creative writing tips, and literary pursuits.', 'IAB9-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 13),
  ('Genealogy', 'This subcategory covers genealogy research, family trees, and ancestral history.', 'IAB9-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 14),
  ('Getting Published', 'This subcategory provides tips and advice for writers looking to get published.', 'IAB9-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 15),
  ('Guitar', 'This subcategory focuses on guitar playing, techniques, and music creation.', 'IAB9-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 16),
  ('Home Recording', 'This subcategory covers home recording techniques, equipment, and music production.', 'IAB9-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 17),
  ('Investors & Patents', 'This subcategory focuses on creative investments, patents, and intellectual property in hobbies.', 'IAB9-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 18),
  ('Jewelry Making', 'This subcategory covers jewelry making techniques, design, and craft.', 'IAB9-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 19),
  ('Magic & Illusion', 'This subcategory focuses on magic tricks, illusions, and performance art in hobbies.', 'IAB9-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 20),
  ('Needlework', 'This subcategory covers needlework, embroidery, and related textile crafts.', 'IAB9-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 21),
  ('Painting', 'This subcategory focuses on painting techniques, art styles, and creative expression.', 'IAB9-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 22),
  ('Photography', 'This subcategory covers photography tips, techniques, and equipment reviews.', 'IAB9-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 23),
  ('Radio', 'This subcategory focuses on radio broadcasting, podcasting, and audio entertainment.', 'IAB9-24', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 24),
  ('Roleplaying Games', 'This subcategory covers roleplaying games, game mechanics, and community discussions.', 'IAB9-25', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 25),
  ('Sci-Fi & Fantasy', 'This subcategory focuses on science fiction and fantasy genres within hobbies and interests.', 'IAB9-26', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 26),
  ('Scrapbooking', 'This subcategory covers scrapbooking techniques, materials, and creative layouts.', 'IAB9-27', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 27),
  ('Screenwriting', 'This subcategory focuses on screenwriting, script development, and film writing.', 'IAB9-28', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 28),
  ('Stamps & Coins', 'This subcategory covers stamp and coin collecting, valuation, and history.', 'IAB9-29', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 29),
  ('Video & Computer Games', 'This subcategory focuses on video games, computer games, and gaming culture.', 'IAB9-30', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 30),
  ('Woodworking', 'This subcategory covers woodworking projects, tools, and techniques for hobbyists.', 'IAB9-31', (SELECT id FROM adv_category WHERE iab_code = 'IAB9'), 31)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Home & Garden (IAB10)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Appliances', 'This subcategory under Home & Garden focuses on home appliances, reviews, and buying guides.', 'IAB10-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 1),
  ('Entertaining', 'This subcategory covers tips and ideas for entertaining guests at home.', 'IAB10-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 2),
  ('Environmental Safety', 'This subcategory focuses on environmental safety in the home and garden.', 'IAB10-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 3),
  ('Gardening', 'This subcategory covers gardening tips, plant care, and landscape design.', 'IAB10-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 4),
  ('Home Repair', 'This subcategory provides advice and tips on home repair and maintenance.', 'IAB10-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 5),
  ('Home Theater', 'This subcategory focuses on home theater systems, reviews, and setup guides.', 'IAB10-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 6),
  ('Interior Decorating', 'This subcategory covers interior decorating ideas, trends, and design tips.', 'IAB10-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 7),
  ('Landscaping', 'This subcategory focuses on landscaping ideas, garden design, and outdoor aesthetics.', 'IAB10-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 8),
  ('Remodeling & Construction', 'This subcategory covers remodeling projects, construction advice, and renovation ideas.', 'IAB10-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB10'), 9)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Law, Government, & Politics (IAB11)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Immigration', 'This subcategory under Law, Government, & Politics focuses on immigration policies and legal issues.', 'IAB11-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB11'), 1),
  ('Legal Issues', 'This subcategory covers legal issues, court cases, and regulatory news.', 'IAB11-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB11'), 2),
  ('U.S. Government Resources', 'This subcategory provides information on U.S. government resources and services.', 'IAB11-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB11'), 3),
  ('Politics', 'This subcategory covers political news, analysis, and discussions.', 'IAB11-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB11'), 4),
  ('Commentary', 'This subcategory focuses on political commentary, opinions, and analysis.', 'IAB11-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB11'), 5)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for News (IAB12)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('International News', 'This subcategory under News covers global events and international news.', 'IAB12-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB12'), 1),
  ('National News', 'This subcategory focuses on national news and events within a country.', 'IAB12-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB12'), 2),
  ('Local News', 'This subcategory covers local news, community events, and regional updates.', 'IAB12-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB12'), 3)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Personal Finance (IAB13)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Beginning Investing', 'This subcategory under Personal Finance focuses on introductory investing strategies and tips.', 'IAB13-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 1),
  ('Credit/Debt & Loans', 'This subcategory covers topics related to credit, debt management, and loan advice.', 'IAB13-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 2),
  ('Financial News', 'This subcategory provides updates on financial markets and economic news.', 'IAB13-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 3),
  ('Financial Planning', 'This subcategory focuses on financial planning, budgeting, and long-term wealth management.', 'IAB13-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 4),
  ('Hedge Fund', 'This subcategory covers hedge funds, investment strategies, and market analysis.', 'IAB13-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 5),
  ('Insurance', 'This subcategory provides information on insurance products and financial protection strategies.', 'IAB13-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 6),
  ('Investing', 'This subcategory covers various investment opportunities and strategies.', 'IAB13-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 7),
  ('Mutual Funds', 'This subcategory focuses on mutual funds, portfolio management, and investment analysis.', 'IAB13-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 8),
  ('Options', 'This subcategory covers options trading, strategies, and market insights.', 'IAB13-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 9),
  ('Retirement Planning', 'This subcategory provides guidance on retirement planning and saving for the future.', 'IAB13-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 10),
  ('Stocks', 'This subcategory focuses on stock market trends, analysis, and trading tips.', 'IAB13-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 11),
  ('Tax Planning', 'This subcategory covers tax planning strategies and updates on tax regulations.', 'IAB13-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB13'), 12)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Society (IAB14)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Dating', 'This subcategory under Society focuses on dating advice, relationship tips, and social interactions.', 'IAB14-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 1),
  ('Divorce Support', 'This subcategory covers divorce support, legal advice, and coping strategies.', 'IAB14-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 2),
  ('Gay Life', 'This subcategory focuses on issues, culture, and support within the LGBTQ+ community.', 'IAB14-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 3),
  ('Marriage', 'This subcategory covers topics related to marriage, relationships, and wedding planning.', 'IAB14-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 4),
  ('Senior Living', 'This subcategory focuses on senior living, retirement communities, and elder care.', 'IAB14-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 5),
  ('Teens', 'This subcategory covers topics and issues relevant to teenagers and youth culture.', 'IAB14-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 6),
  ('Weddings', 'This subcategory focuses on wedding planning, bridal trends, and matrimonial advice.', 'IAB14-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 7),
  ('Ethnic Specific', 'This subcategory covers ethnic-specific cultural topics and community issues.', 'IAB14-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB14'), 8)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Science (IAB15)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Astrology', 'This subcategory under Science focuses on astrology, horoscopes, and celestial studies.', 'IAB15-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 1),
  ('Biology', 'This subcategory covers biology, life sciences, and related research topics.', 'IAB15-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 2),
  ('Chemistry', 'This subcategory focuses on chemistry, chemical research, and laboratory studies.', 'IAB15-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 3),
  ('Geology', 'This subcategory covers geology, earth sciences, and related research.', 'IAB15-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 4),
  ('Paranormal Phenomena', 'This subcategory focuses on paranormal phenomena, unexplained events, and related scientific investigations.', 'IAB15-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 5),
  ('Physics', 'This subcategory covers physics, theoretical studies, and scientific discoveries.', 'IAB15-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 6),
  ('Space/Astronomy', 'This subcategory focuses on space exploration, astronomy, and cosmic research.', 'IAB15-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 7),
  ('Geography', 'This subcategory covers geography, maps, and spatial analysis.', 'IAB15-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 8),
  ('Botany', 'This subcategory focuses on botany, plant sciences, and horticulture.', 'IAB15-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 9),
  ('Weather', 'This subcategory covers weather, meteorology, and climate studies.', 'IAB15-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB15'), 10)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Pets (IAB16)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Aquariums', 'This subcategory under Pets focuses on aquariums, fish care, and aquatic pet maintenance.', 'IAB16-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 1),
  ('Birds', 'This subcategory covers pet birds, care tips, and avian health.', 'IAB16-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 2),
  ('Cats', 'This subcategory focuses on cats, including care, behavior, and health issues.', 'IAB16-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 3),
  ('Dogs', 'This subcategory covers dog care, training, and health tips.', 'IAB16-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 4),
  ('Large Animals', 'This subcategory focuses on large pets such as horses and livestock, including care and management.', 'IAB16-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 5),
  ('Reptiles', 'This subcategory covers reptile care, habitat setup, and health information.', 'IAB16-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 6),
  ('Veterinary Medicine', 'This subcategory focuses on veterinary medicine, pet health, and animal care best practices.', 'IAB16-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB16'), 7)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Sports (IAB17)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Auto Racing', 'This subcategory under Sports focuses on auto racing events, news, and analysis.', 'IAB17-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 1),
  ('Baseball', 'This subcategory covers baseball news, games, and analysis.', 'IAB17-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 2),
  ('Bicycling', 'This subcategory focuses on bicycling, cycling events, and equipment reviews.', 'IAB17-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 3),
  ('Bodybuilding', 'This subcategory covers bodybuilding, fitness competitions, and training tips.', 'IAB17-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 4),
  ('Boxing', 'This subcategory focuses on boxing events, fighters, and match analyses.', 'IAB17-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 5),
  ('Canoeing/Kayaking', 'This subcategory covers canoeing and kayaking activities, equipment, and events.', 'IAB17-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 6),
  ('Cheerleading', 'This subcategory focuses on cheerleading, routines, and competitive events.', 'IAB17-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 7),
  ('Climbing', 'This subcategory covers climbing, bouldering, and mountaineering activities.', 'IAB17-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 8),
  ('Cricket', 'This subcategory focuses on cricket news, matches, and team analysis.', 'IAB17-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 9),
  ('Figure Skating', 'This subcategory covers figure skating events, competitions, and athlete profiles.', 'IAB17-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 10),
  ('Fly Fishing', 'This subcategory focuses on fly fishing techniques, gear, and locations.', 'IAB17-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 11),
  ('Football', 'This subcategory covers football news, games, and team strategies.', 'IAB17-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 12),
  ('Freshwater Fishing', 'This subcategory focuses on freshwater fishing tips, gear, and locations.', 'IAB17-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 13),
  ('Game & Fish', 'This subcategory covers hunting, fishing, and outdoor game activities.', 'IAB17-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 14),
  ('Golf', 'This subcategory focuses on golf news, tips, and tournament updates.', 'IAB17-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 15),
  ('Horse Racing', 'This subcategory covers horse racing events, betting, and industry news.', 'IAB17-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 16),
  ('Horses', 'This subcategory focuses on horse-related sports, riding, and care.', 'IAB17-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 17),
  ('Hunting/Shooting', 'This subcategory covers hunting and shooting sports, equipment, and safety tips.', 'IAB17-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 18),
  ('Inline Skating', 'This subcategory focuses on inline skating, gear, and recreational tips.', 'IAB17-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 19),
  ('Martial Arts', 'This subcategory covers martial arts disciplines, training, and competition news.', 'IAB17-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 20),
  ('Mountain Biking', 'This subcategory focuses on mountain biking trails, gear, and adventure rides.', 'IAB17-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 21),
  ('NASCAR Racing', 'This subcategory covers NASCAR racing events, drivers, and race analysis.', 'IAB17-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 22),
  ('Olympics', 'This subcategory focuses on the Olympics, including news, events, and athlete profiles.', 'IAB17-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 23),
  ('Paintball', 'This subcategory covers paintball sports, equipment, and game strategies.', 'IAB17-24', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 24),
  ('Power & Motorcycles', 'This subcategory focuses on power sports and motorcycle events and reviews.', 'IAB17-25', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 25),
  ('Pro Basketball', 'This subcategory covers professional basketball news, games, and analysis.', 'IAB17-26', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 26),
  ('Pro Ice Hockey', 'This subcategory focuses on professional ice hockey, including games, teams, and player news.', 'IAB17-27', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 27),
  ('Rodeo', 'This subcategory covers rodeo events, competitions, and related sports.', 'IAB17-28', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 28),
  ('Rugby', 'This subcategory focuses on rugby news, matches, and team strategies.', 'IAB17-29', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 29),
  ('Running/Jogging', 'This subcategory covers running and jogging tips, events, and training programs.', 'IAB17-30', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 30),
  ('Sailing', 'This subcategory focuses on sailing, boat racing, and maritime activities.', 'IAB17-31', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 31),
  ('Saltwater Fishing', 'This subcategory covers saltwater fishing techniques, gear, and locations.', 'IAB17-32', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 32),
  ('Scuba Diving', 'This subcategory focuses on scuba diving, underwater exploration, and safety tips.', 'IAB17-33', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 33),
  ('Skateboarding', 'This subcategory covers skateboarding, tricks, gear, and cultural aspects.', 'IAB17-34', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 34),
  ('Skiing', 'This subcategory focuses on skiing, equipment, resorts, and winter sports.', 'IAB17-35', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 35),
  ('Snowboarding', 'This subcategory covers snowboarding techniques, gear, and mountain resorts.', 'IAB17-36', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 36),
  ('Surfing/Body-Boarding', 'This subcategory focuses on surfing and bodyboarding, including tips and destination guides.', 'IAB17-37', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 37),
  ('Swimming', 'This subcategory covers swimming, aquatic sports, and training techniques.', 'IAB17-38', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 38),
  ('Table Tennis/Ping-Pong', 'This subcategory focuses on table tennis, including game strategies and equipment reviews.', 'IAB17-39', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 39),
  ('Tennis', 'This subcategory covers tennis news, matches, and player analysis.', 'IAB17-40', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 40),
  ('Volleyball', 'This subcategory focuses on volleyball, including game rules, teams, and tournaments.', 'IAB17-41', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 41),
  ('Walking', 'This subcategory covers walking as a fitness activity and leisure pursuit.', 'IAB17-42', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 42),
  ('Waterski/Wakeboard', 'This subcategory focuses on waterskiing and wakeboarding, including techniques and equipment.', 'IAB17-43', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 43),
  ('World Soccer', 'This subcategory covers international soccer news, matches, and analysis.', 'IAB17-44', (SELECT id FROM adv_category WHERE iab_code = 'IAB17'), 44)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Style & Fashion (IAB18)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Beauty', 'This subcategory under Style & Fashion focuses on beauty trends, skincare, and makeup tips.', 'IAB18-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 1),
  ('Body Art', 'This subcategory covers body art, tattoos, and other forms of personal expression through art.', 'IAB18-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 2),
  ('Fashion', 'This subcategory focuses on fashion trends, designer news, and style advice.', 'IAB18-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 3),
  ('Jewelry', 'This subcategory covers jewelry trends, design, and accessories.', 'IAB18-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 4),
  ('Clothing', 'This subcategory focuses on clothing styles, brands, and fashion advice.', 'IAB18-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 5),
  ('Accessories', 'This subcategory covers fashion accessories, including bags, scarves, and other items.', 'IAB18-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB18'), 6)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Technology & Computing (IAB19)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('3-D Graphics', 'This subcategory under Technology & Computing focuses on 3-D graphics, modeling, and animation.', 'IAB19-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 1),
  ('Animation', 'This subcategory covers animation techniques, software, and industry news.', 'IAB19-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 2),
  ('Antivirus Software', 'This subcategory focuses on antivirus software, security solutions, and cybersecurity tips.', 'IAB19-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 3),
  ('C/C++', 'This subcategory covers programming in C and C++, including tutorials and development tips.', 'IAB19-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 4),
  ('Cameras & Camcorders', 'This subcategory focuses on cameras, camcorders, and related technology reviews.', 'IAB19-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 5),
  ('Cell Phones', 'This subcategory covers cell phone technology, reviews, and mobile trends.', 'IAB19-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 6),
  ('Computer Certification', 'This subcategory focuses on computer certifications, training, and career advice in IT.', 'IAB19-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 7),
  ('Computer Networking', 'This subcategory covers networking technology, protocols, and infrastructure.', 'IAB19-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 8),
  ('Computer Peripherals', 'This subcategory focuses on computer peripherals, accessories, and hardware reviews.', 'IAB19-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 9),
  ('Computer Reviews', 'This subcategory covers reviews of computers, laptops, and related technology.', 'IAB19-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 10),
  ('Data Centers', 'This subcategory focuses on data centers, server technology, and cloud computing infrastructure.', 'IAB19-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 11),
  ('Databases', 'This subcategory covers database technology, management systems, and data solutions.', 'IAB19-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 12),
  ('Desktop Publishing', 'This subcategory focuses on desktop publishing software, design, and layout techniques.', 'IAB19-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 13),
  ('Desktop Video', 'This subcategory covers desktop video editing, production software, and tutorials.', 'IAB19-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 14),
  ('Email', 'This subcategory focuses on email technology, communication tools, and best practices.', 'IAB19-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 15),
  ('Graphics Software', 'This subcategory covers graphics software, design tools, and creative applications.', 'IAB19-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 16),
  ('Home Video/DVD', 'This subcategory focuses on home video technology, DVD reviews, and media players.', 'IAB19-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 17),
  ('Internet Technology', 'This subcategory covers internet technology, web development, and digital trends.', 'IAB19-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 18),
  ('Java', 'This subcategory focuses on Java programming, development tips, and software updates.', 'IAB19-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 19),
  ('JavaScript', 'This subcategory covers JavaScript programming, web development, and interactive design.', 'IAB19-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 20),
  ('Mac Support', 'This subcategory focuses on support and resources for Mac computers and software.', 'IAB19-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 21),
  ('MP3/MIDI', 'This subcategory covers MP3 and MIDI technology, audio software, and music production tools.', 'IAB19-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 22),
  ('Net Conferencing', 'This subcategory focuses on net conferencing tools, online collaboration, and communication platforms.', 'IAB19-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 23),
  ('Net for Beginners', 'This subcategory provides guidance and resources for beginners in internet technology.', 'IAB19-24', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 24),
  ('Network Security', 'This subcategory covers network security, cybersecurity, and protection measures.', 'IAB19-25', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 25),
  ('Palmtops/PDAs', 'This subcategory focuses on palmtop computers, PDAs, and mobile computing devices.', 'IAB19-26', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 26),
  ('PC Support', 'This subcategory covers support resources, troubleshooting, and maintenance for PCs.', 'IAB19-27', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 27),
  ('Portable', 'This subcategory focuses on portable computing devices, including tablets and laptops.', 'IAB19-28', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 28),
  ('Entertainment', 'This subcategory covers entertainment technology, including multimedia systems and home entertainment.', 'IAB19-29', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 29),
  ('Shareware/Freeware', 'This subcategory focuses on shareware and freeware software, including reviews and downloads.', 'IAB19-30', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 30),
  ('Unix', 'This subcategory covers Unix operating systems, commands, and related technology.', 'IAB19-31', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 31),
  ('Visual Basic', 'This subcategory focuses on Visual Basic programming and application development.', 'IAB19-32', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 32),
  ('Web Clip Art', 'This subcategory covers web clip art, digital images, and graphic resources for web design.', 'IAB19-33', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 33),
  ('Web Design/HTML', 'This subcategory focuses on web design, HTML coding, and website development.', 'IAB19-34', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 34),
  ('Web Search', 'This subcategory covers web search technologies, search engines, and online research.', 'IAB19-35', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 35),
  ('Windows', 'This subcategory focuses on Windows operating systems, updates, and software reviews.', 'IAB19-36', (SELECT id FROM adv_category WHERE iab_code = 'IAB19'), 36)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Travel (IAB20)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Adventure Travel', 'This subcategory under Travel focuses on adventure travel experiences and destinations.', 'IAB20-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 1),
  ('Africa', 'This subcategory covers travel information, guides, and tips for destinations in Africa.', 'IAB20-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 2),
  ('Air Travel', 'This subcategory focuses on air travel tips, airline reviews, and flight experiences.', 'IAB20-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 3),
  ('Australia & New Zealand', 'This subcategory covers travel guides and tips for Australia and New Zealand.', 'IAB20-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 4),
  ('Bed & Breakfasts', 'This subcategory focuses on bed and breakfast accommodations and travel experiences.', 'IAB20-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 5),
  ('Budget Travel', 'This subcategory covers budget travel tips, affordable destinations, and cost-saving strategies.', 'IAB20-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 6),
  ('Business Travel', 'This subcategory focuses on travel related to business, including corporate travel tips and services.', 'IAB20-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 7),
  ('By US Locale', 'This subcategory covers travel information categorized by US locales and regions.', 'IAB20-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 8),
  ('Camping', 'This subcategory focuses on camping tips, gear, and destination guides.', 'IAB20-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 9),
  ('Canada', 'This subcategory covers travel guides and tips for destinations in Canada.', 'IAB20-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 10),
  ('Caribbean', 'This subcategory focuses on travel to Caribbean destinations, including resorts and cultural experiences.', 'IAB20-11', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 11),
  ('Cruises', 'This subcategory covers cruise travel, itineraries, and tips for cruising.', 'IAB20-12', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 12),
  ('Eastern Europe', 'This subcategory focuses on travel in Eastern Europe, including destinations and cultural insights.', 'IAB20-13', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 13),
  ('Europe', 'This subcategory covers travel guides and tips for various European destinations.', 'IAB20-14', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 14),
  ('France', 'This subcategory focuses on travel information and tips for France, including culture and attractions.', 'IAB20-15', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 15),
  ('Greece', 'This subcategory covers travel guides for Greece, including historical sites and island destinations.', 'IAB20-16', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 16),
  ('Honeymoons/Getaways', 'This subcategory focuses on travel ideas for honeymoons and romantic getaways.', 'IAB20-17', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 17),
  ('Hotels', 'This subcategory covers hotel reviews, booking tips, and accommodation guides.', 'IAB20-18', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 18),
  ('Italy', 'This subcategory focuses on travel information and guides for Italy, including cultural highlights.', 'IAB20-19', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 19),
  ('Japan', 'This subcategory covers travel guides and tips for Japan, including cultural experiences and technology.', 'IAB20-20', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 20),
  ('Mexico & Central America', 'This subcategory focuses on travel in Mexico and Central American countries, including resorts and cultural insights.', 'IAB20-21', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 21),
  ('National Parks', 'This subcategory covers national parks, outdoor adventures, and nature travel.', 'IAB20-22', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 22),
  ('South America', 'This subcategory focuses on travel guides and tips for South American destinations.', 'IAB20-23', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 23),
  ('Spas', 'This subcategory covers spa resorts, wellness retreats, and relaxation travel experiences.', 'IAB20-24', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 24),
  ('Theme Parks', 'This subcategory focuses on theme parks, attractions, and family travel experiences.', 'IAB20-25', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 25),
  ('Traveling with Kids', 'This subcategory covers travel tips, destinations, and activities for families with children.', 'IAB20-26', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 26),
  ('United Kingdom', 'This subcategory focuses on travel guides and tips for the United Kingdom, including cultural insights.', 'IAB20-27', (SELECT id FROM adv_category WHERE iab_code = 'IAB20'), 27)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Real Estate (IAB21)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Apartments', 'This subcategory under Real Estate focuses on apartments, rental listings, and related advice.', 'IAB21-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB21'), 1),
  ('Architects', 'This subcategory covers topics related to architects, design, and building trends.', 'IAB21-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB21'), 2),
  ('Buying/Selling Homes', 'This subcategory focuses on the process of buying and selling homes, including market trends and advice.', 'IAB21-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB21'), 3)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Shopping (IAB22)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Contests & Freebies', 'This subcategory under Shopping covers contests, giveaways, and freebies.', 'IAB22-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB22'), 1),
  ('Couponing', 'This subcategory focuses on couponing strategies, discounts, and savings tips.', 'IAB22-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB22'), 2),
  ('Comparison', 'This subcategory covers product comparisons, reviews, and shopping guides.', 'IAB22-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB22'), 3),
  ('Engines', 'This subcategory focuses on shopping for engines and related automotive components.', 'IAB22-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB22'), 4)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Religion & Spirituality (IAB23)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Alternative Religions', 'This subcategory under Religion & Spirituality focuses on alternative religious beliefs and practices.', 'IAB23-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 1),
  ('Atheism/Agnosticism', 'This subcategory covers atheism, agnosticism, and secular perspectives.', 'IAB23-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 2),
  ('Buddhism', 'This subcategory focuses on Buddhism, its teachings, and cultural impact.', 'IAB23-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 3),
  ('Catholicism', 'This subcategory covers Catholicism, including traditions, teachings, and community.', 'IAB23-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 4),
  ('Christianity', 'This subcategory focuses on Christianity, its denominations, and religious practices.', 'IAB23-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 5),
  ('Hinduism', 'This subcategory covers Hinduism, its traditions, and cultural significance.', 'IAB23-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 6),
  ('Islam', 'This subcategory focuses on Islam, its teachings, and cultural impact.', 'IAB23-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 7),
  ('Judaism', 'This subcategory covers Judaism, its traditions, and cultural heritage.', 'IAB23-8', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 8),
  ('Latter-Day Saints', 'This subcategory focuses on the beliefs and practices of Latter-Day Saints.', 'IAB23-9', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 9),
  ('Pagan/Wiccan', 'This subcategory covers Pagan and Wiccan practices, beliefs, and cultural aspects.', 'IAB23-10', (SELECT id FROM adv_category WHERE iab_code = 'IAB23'), 10)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Non-Standard Content (IAB25)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Unmoderated UGC', 'This subcategory under Non-Standard Content focuses on unmoderated user-generated content.', 'IAB25-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 1),
  ('Extreme Graphic/Explicit Violence', 'This subcategory covers content featuring extreme graphic or explicit violence.', 'IAB25-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 2),
  ('Pornography', 'This subcategory focuses on adult content and pornography.', 'IAB25-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 3),
  ('Profane Content', 'This subcategory covers content that includes profanity and explicit language.', 'IAB25-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 4),
  ('Hate Content', 'This subcategory focuses on hate content and inflammatory material.', 'IAB25-5', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 5),
  ('Under Construction', 'This subcategory indicates content that is under construction or in development.', 'IAB25-6', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 6),
  ('Incentivized', 'This subcategory covers content that is incentivized or rewards-based.', 'IAB25-7', (SELECT id FROM adv_category WHERE iab_code = 'IAB25'), 7)
ON CONFLICT DO NOTHING;
COMMIT;

-- Child categories for Illegal Content (IAB26)
BEGIN;
INSERT INTO adv_category (name, description, iab_code, parent_id, position)
VALUES
  ('Illegal Content', 'This subcategory under Illegal Content focuses on content that is deemed illegal or prohibited.', 'IAB26-1', (SELECT id FROM adv_category WHERE iab_code = 'IAB26'), 1),
  ('Warez', 'This subcategory covers warez, including pirated software and unauthorized content sharing.', 'IAB26-2', (SELECT id FROM adv_category WHERE iab_code = 'IAB26'), 2),
  ('Spyware/Malware', 'This subcategory focuses on spyware, malware, and related cybersecurity threats.', 'IAB26-3', (SELECT id FROM adv_category WHERE iab_code = 'IAB26'), 3),
  ('Copyright Infringement', 'This subcategory covers copyright infringement issues and unauthorized content distribution.', 'IAB26-4', (SELECT id FROM adv_category WHERE iab_code = 'IAB26'), 4)
ON CONFLICT DO NOTHING;
COMMIT;

UPDATE adv_category SET active = 'active' WHERE iab_code != '';
