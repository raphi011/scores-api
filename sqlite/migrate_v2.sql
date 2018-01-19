/* teams */
select
	created_at,
	name,
	player1_id,
	player2_id
from teams

/* users */
select
  id,
	created_at,
	updated_at,
	deleted_at,
	email,
	profile_image_url
 from users

/* players */
select 
	 p.id,
	 p.created_at,
	 p.updated_at,
	 p.deleted_at,
	 p.name,
	 u.id as user_id
from players p
left join users u on u.player_id = p.id

/* matches */
select
	m.id,
	m.created_at,
	m.updated_at,
	m.deleted_at,
	t1.player1_id as team1_player1_id,
	t1.player2_id as team1_player2_id,
	t2.player1_id as team2_player1_id,
	t2.player2_id as team2_player2_id,
	m.score_team1,
	m.score_team2,
	m.created_by_id as created_by_user_id
from matches m
join teams t1 on m.team1_id = t1.id
join teams t2 on m.team2_id = t2.id
where m.deleted_at is null