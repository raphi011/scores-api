package sql

import (
	"testing"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/test"
)

func TestCreateSetting(t *testing.T) {
	db := SetupDB(t)
	settingRepo := &settingRepository{DB: db}
	_ = CreateUsers(t, db, U{ID: 1})

	setting := &scores.Setting{Key: "FILTER_LEAGUE", UserID: 1}

	_, err := settingRepo.Create(setting)
	test.Check(t, "settingRepository.New(), err: %v", err)

	persistedSetting, err := settingRepo.ByUserID(1)

	test.Check(t, "settingRepo.Get() failed: %v", err)
	test.Compare(t, "setting is not equal:\n%s", persistedSetting[0], setting)
}

func TestUpdateSetting(t *testing.T) {
	db := SetupDB(t)
	settingRepo := &settingRepository{DB: db}
	_ = CreateUsers(t, db, U{ID: 1})

	setting := &scores.Setting{
		Key:    "FILTER_LEAGUE",
		Value:  "amateur-league",
		UserID: 1,
	}

	_, err := settingRepo.Create(setting)
	test.Check(t, "settingRepository.New(), err: %v", err)

	err = settingRepo.Update(&scores.Setting{
		Key:    "FILTER_LEAGUE",
		Value:  "junior-league",
		UserID: 1,
	})

	test.Check(t, "settingRepo.Update() failed: %v", err)

	persistedSetting, err := settingRepo.ByUserID(1)
	updatedSetting := persistedSetting[0]

	test.Assert(t,
		"setting should be %q but is %q",
		updatedSetting.Value == "junior-league",
		setting.Value,
		updatedSetting.Value,
	)
}
