package scores

type Services struct {
	Group    *GroupService
	Password *PBKDF2PasswordService

	Match     MatchRepository
	User      UserRepository
	Statistic StatisticRepository
	Team      TeamRepository
	Player    PlayerRepository
	Volleynet VolleynetRepository
}
