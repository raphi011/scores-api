
select 
p.id,
p.name,
m.created_at,
case when (t1.player1_id = 1 or t1.player2_id = 1) and m.score_team1 > m.score_team2 then 1 else  0 end as won,
case when t1.player1_id = 1 or t1.player2_id = 1 then m.score_team1 else m.score_team2 end as pointsWon,
case when t1.player1_id = 1 or t1.player2_id = 1 then m.score_team2 else m.score_team1 end as pointsLost
from matches m	
join teams t1 on m.team1_id = t1.id
join teams t2 on m.team2_id = t2.id
join players p on t1.player1_id = p.id or t1.player2_id = p.id or t2.player1_id = p.id or t2.player2_id = p.id
where m.deleted_at is null
order by p.id
