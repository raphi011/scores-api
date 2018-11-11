package scores

type Services struct {
	Group    *GroupService
	Password *PBKDF2PasswordService
	User     *UserService

	Match     MatchRepository
	Statistic StatisticRepository
	Team      TeamRepository
	Player    PlayerRepository
	Volleynet VolleynetRepository
}
