package user_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// DeleteUser deletes an existing user from the database
func DeleteUser(db *gorm.DB, user_id int) error {
	return db_model.DeleteDBRecord(db, &db_tables.User{}, user_id)
}

// DeleteUserFromRecord deletes an existing user from the database
func DeleteUserFromRecord(db *gorm.DB, user *db_tables.User) error {
	return db_model.DeleteDBRecordNoFetch(db, user)
}

// DeleteUsers deletes existing users from the database
func DeleteUsers(db *gorm.DB, user_ids []int) error {
	return db_model.DeleteDBRecords(db, &db_tables.User{}, user_ids)
}

// DeleteUsersFromRecords deletes existing users from the database
func DeleteUsersFromRecords(db *gorm.DB, users []*db_tables.User) error {
	result := db.Delete(users)
	return result.Error
}
