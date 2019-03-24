SELECT
	p.id,
	p.created_at,
	p.updated_at,
	p.first_name,
	p.last_name,
	p.birthday,
	p.gender,
	p.total_points,
	p.ladder_rank,
	p.club,
	p.country_union,
	p.license
FROM
(SELECT partners.player_id, max(partners.last_played) as last_played FROM
    (SELECT
        CASE 
            WHEN tt.player_1_id = :player_id THEN tt.player_2_id 
            ELSE tt.player_1_id
        END AS player_id,
        t.start_date AS last_played
    FROM tournament_teams tt
    JOIN tournaments t ON tt.tournament_id = t.id
    WHERE tt.player_1_id = :player_id OR tt.player_2_id = :player_id
    ORDER BY t.start_date DESC) AS partners
GROUP BY partners.player_id) AS distinct_partners
JOIN players p on distinct_partners.player_id = p.id