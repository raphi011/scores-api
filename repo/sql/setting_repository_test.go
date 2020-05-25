// +build repository

package sql

import (
	"testing"

	"github.com/raphi011/scores-backend"
	"github.com/raphi011/scores-backend/test"
)

func TestCreateSetting(t *testing.T) {
	db := SetupDB(t)
	settingRepo := &settingRepository{DB: db}
	users := CreateUsers(t, db, U{})

	setting := &scores.Setting{Key: "FILTER_LEAGUE", UserID: users[0].ID}

	_, err := settingRepo.Create(setting)
	test.Check(t, "settingRepository.New(), err: %v", err)

	persistedSetting, err := settingRepo.ByUserID(users[0].ID)

	test.Check(t, "settingRepo.Get() failed: %v", err)
	test.Compare(t, "setting is not equal:\n%s", persistedSetting[0], setting)
}

func TestUpdateSetting(t *testing.T) {
	db := SetupDB(t)
	settingRepo := &settingRepository{DB: db}
	users := CreateUsers(t, db, U{})

	setting := &scores.Setting{
		Key:    "FILTER_LEAGUE",
		Value:  "amateur-league",
		UserID: users[0].ID,
	}

	_, err := settingRepo.Create(setting)
	test.Check(t, "settingRepository.New(), err: %v", err)

	err = settingRepo.Update(&scores.Setting{
		Key:    "FILTER_LEAGUE",
		Value:  "junior-league",
		UserID: users[0].ID,
	})

	test.Check(t, "settingRepo.Update() failed: %v", err)

	persistedSetting, err := settingRepo.ByUserID(users[0].ID)

	test.Check(t, "settingRepository.ByUserID(), err: %v", err)

	updatedSetting := persistedSetting[0]

	test.Assert(t,
		"setting should be %q but is %q",
		updatedSetting.Value == "junior-league",
		setting.Value,
		updatedSetting.Value,
	)
}
