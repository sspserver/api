
SELECT ad, country, zone, spent, views, ecpm, ecpm2
  FROM (

    SELECT ad, country, zone,
      SUM(spent) AS spent, SUM(counter) AS views,
      (toFloat64(spent) / toFloat64(views)) / 1000000.0 AS ecpm
    FROM stats.counters_local
    WHERE datemark = today() AND status = 1 AND event IN ('direct', 'view')
    GROUP BY ad, country, zone

  ) ANY LEFT JOIN (

    SELECT ad, country, zone,
      SUM(spent) AS spent2, SUM(counter) AS views2,
      (toFloat64(spent2) / toFloat64(views2)) / 1000000.0 AS ecpm2
    FROM stats.counters_local
    WHERE datemark <= today()-1 AND status = 1 AND event IN ('direct', 'view')
    GROUP BY datemark, ad, country, zone
    ORDER BY datemark DESC
    LIMIT 1 BY ad, country, zone

  ) USING (ad, country, zone);
