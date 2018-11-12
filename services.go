package scores

type Services struct {
	Group     *GroupService
	Password  *PBKDF2Password
	User      *UserService
	Match     *MatchService
	Statistic *StatisticService
	Team      *TeamService
	Player    *PlayerService
	Volleynet VolleynetRepository
}
