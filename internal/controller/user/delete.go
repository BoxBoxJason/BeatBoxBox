package user_controller

import "fmt"

// DeleteUser deletes a user by its id
func DeleteUser(user_id int) error {
	user_exists := UserExists(user_id)
	if !user_exists {
		return fmt.Errorf("user with id %d does not exist", user_id)
	}
	return DeleteUser(user_id)
}

// DeleteUsers deletes users by their ids
func DeleteUsers(user_ids []int) error {
	users_exists := UsersExist(user_ids)
	if !users_exists {
		return fmt.Errorf("at least one of the ids %v does not exist", user_ids)
	}
	return DeleteUsers(user_ids)
}
