package user_model

import "gorm.io/gorm"

// GetUsersFromFilters returns a list of users from the database
// Filters can be passed to filter the users
func GetUsersFromFilters(db *gorm.DB, filters map[string]interface{}) ([]User, error) {
	var users []User
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUser returns a user from the database
// Selects the user with the given user_id
func GetUser(db *gorm.DB, user_id int) (User, error) {
	var user User
	// Using `First` to retrieve the first record that matches the user_id
	result := db.Where("Id = ?", user_id).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

// GetUsers returns a list of users from the database
// Selects the users with the given user_ids
func GetUsers(db *gorm.DB, user_ids []int) ([]User, error) {
	var users []User
	// Using `Find` to retrieve records with the IDs in user_ids slice
	result := db.Where("Id IN ?", user_ids).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
