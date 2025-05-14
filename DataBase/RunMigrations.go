package DataBase

import "MPT-Schedule/Models"

func RunMigrations() error {
	err := DB.AutoMigrate(
		&Models.User{},
		&Models.PinCode{},
	)
	if err != nil {
		return err
	}
	return nil
}
