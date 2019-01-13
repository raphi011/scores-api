SELECT
	t.id,
	t.created_at,
	t.updated_at,
	t.format,
	t.start_date,
	t.end_date,
	t.name,
	t.league,
	t.link,
	t.entry_link,
	t.status,
	t.registration_open,
	t.location,
	t.html_notes,
	t.mode,
	t.max_points,
	t.min_teams,
	t.max_teams,
	t.end_registration,
	t.organiser,
	t.phone,
	t.email,
	t.website,
	t.current_points,
	t.live_scoring_link,
	t.loc_lat,
	t.loc_lon,
	t.season,
	t.signedup_teams
FROM tournaments t
WHERE
	t.format IN (?) AND
	t.league IN (?) AND
	t.season IN (?)