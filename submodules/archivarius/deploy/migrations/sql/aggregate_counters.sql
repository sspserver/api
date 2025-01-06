
-- Remove day
INSERT INTO stats.counters_local (sign, datemark, status, counter, spent, event, project, campaign, ad)
  SELECT -1, datemark, status, counter, spent, event, project, campaign, ad
    FROM stats.counters_local WHERE datemark = toDate('2016-06-15');

-- Insert changes for the date
INSERT INTO stats.counters_local (sign, datemark, pricing_model, status, counter, event, project, campaign, ad, spent)
  SELECT 1, datemark, pricing_model, status, SUM(CASE WHEN event!='lead' OR (event='lead' AND merged = 1) THEN sign ELSE 0 END) AS counter,
    event, project, campaign, ad,
    SUM(CASE
      WHEN pricing_model = 1 AND event IN ('view', 'direct') THEN 
        price * sign
      WHEN pricing_model = 2 AND event IN ('click') THEN
        price * sign
      WHEN pricing_model = 3 AND merged = 1 AND event IN ('lead') THEN
        price * sign
      ELSE 0
    END) AS spent
  FROM stats.events_local
  WHERE datemark = toDate(today())
  GROUP BY datemark, pricing_model, status, event, project, campaign, ad;

SELECT status, event, project, pricing_model, campaign, ad, SUM(counter) AS counter, SUM(spent) AS spent
  FROM stats.counters_local
  GROUP BY status, event, project, campaign, ad
  ORDER BY campaign, ad;

SELECT campaign, ad,
    SUM(CASE WHEN event = 'impression' AND status = 0 THEN counter ELSE 0 END) AS pre_imps,
    SUM(CASE WHEN event = 'impression' AND status = 0 THEN spent ELSE 0 END) AS pre_imp_spent,
    SUM(CASE WHEN event = 'impression' AND status = 1 THEN counter ELSE 0 END) AS imps,
    SUM(CASE WHEN event = 'impression' AND status = 1 THEN spent ELSE 0 END) AS imp_spent,
    SUM(CASE WHEN event IN ('view', 'direct') THEN counter ELSE 0 END) AS views,
    SUM(CASE WHEN event IN ('view', 'direct') THEN spent ELSE 0 END) AS view_spent,
    SUM(CASE WHEN event = 'click' AND status = 0 THEN counter ELSE 0 END) AS pre_clicks,
    SUM(CASE WHEN event = 'click' AND status = 0 THEN spent ELSE 0 END) AS pre_click_spent,
    SUM(CASE WHEN event = 'click' AND status = 1 THEN counter ELSE 0 END) AS clicks,
    SUM(CASE WHEN event = 'click' AND status = 1 THEN spent ELSE 0 END) AS click_spent,
    SUM(CASE WHEN event = 'lead' THEN counter ELSE 0 END) AS leads,
    SUM(CASE WHEN event = 'lead' THEN spent ELSE 0 END) AS lead_spent,
    (CASE WHEN pricing_model = 1 THEN view_spent WHEN pricing_model = 2 THEN click_spent ELSE lead_spent END) AS target_spent
  FROM stats.counters_local
  GROUP BY campaign, ad, pricing_model
  ORDER BY campaign, ad;
