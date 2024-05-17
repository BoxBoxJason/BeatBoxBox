package user_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// GetUsersFromFilters returns a list of users from the database
// Filters can be passed to filter the users
func GetUsersFromFilters(db *gorm.DB, filters map[string]interface{}) ([]db_model.User, error) {
	var users []db_model.User
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUser returns a user from the database
// Selects the user with the given user_id
func GetUser(db *gorm.DB, user_id int) (db_model.User, error) {
	var user db_model.User
	// Using `First` to retrieve the first record that matches the user_id
	result := db.Where("id = ?", user_id).First(&user)
	if result.Error != nil {
		return db_model.User{}, result.Error
	}
	return user, nil
}

// GetUsers returns a list of users from the database
// Selects the users with the given user_ids
func GetUsers(db *gorm.DB, user_ids []int) ([]db_model.User, error) {
	var users []db_model.User
	// Using `Find` to retrieve records with the IDs in user_ids slice
	result := db.Where("id IN ?", user_ids).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUsersFromPartialPseudo(db *gorm.DB, partial_pseudo string) ([]db_model.User, error) {
	var users []db_model.User
	// Using `Find` to retrieve records with the pseudo containing partial_pseudo
	result := db.Where("pseudo LIKE ?", "%"+partial_pseudo+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
