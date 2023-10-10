package repos

import "gorm.io/gorm"

type StatsRepo struct {
	db *gorm.DB
}

func (r *StatsRepo) GetUsersCount() (int, error) {
	return 0, nil
}
func (r *StatsRepo) GetTeamsCount() (int, error) {
	return 0, nil
}
func (r *StatsRepo) GetRoomsCount() (int, error) {
	return 0, nil
}
func (r *StatsRepo) GetLoginsCount() (int, error) {
	return 0, nil
}
