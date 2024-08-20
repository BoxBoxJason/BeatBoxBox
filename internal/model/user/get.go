package user_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// GetUsersFromFilters returns a list of users from the database
// Filters can be passed to filter the users
func GetUsersFromFilters(db *gorm.DB, filters map[string]interface{}) []db_tables.User {
	raw_users := db_model.GetRecordsByFields(db, &db_tables.User{}, filters)
	if raw_users == nil {
		return nil
	}
	users := make([]db_tables.User, len(raw_users))
	for i, user := range raw_users {
		users[i] = user.(db_tables.User)
	}

	return users
}

// GetUser returns a user from the database
// Selects the user with the given user_id
func GetUser(db *gorm.DB, user_id int) (db_tables.User, error) {
	user := db_model.GetRecordFromId(db, &db_tables.User{}, user_id)
	if user == nil {
		return db_tables.User{}, gorm.ErrRecordNotFound
	}
	return *user.(*db_tables.User), nil
}

// GetUsers returns a list of users from the database
// Selects the users with the given user_ids
func GetUsers(db *gorm.DB, user_ids []int) ([]db_tables.User, error) {
	raw_users := db_model.GetRecordsFromIds(db, &db_tables.User{}, user_ids)
	if raw_users == nil {
		return nil, gorm.ErrRecordNotFound
	}
	users := make([]db_tables.User, len(raw_users))
	for i, user := range raw_users {
		users[i] = user.(db_tables.User)
	}

	return users, nil
}

// GetUsersFromPartialPseudo returns a list of users from the database
func GetUsersFromPartialPseudo(db *gorm.DB, partial_pseudo string) ([]db_tables.User, error) {
	raw_users := db_model.GetRecordsByFieldsWithCondition(db, &db_tables.User{}, map[string]interface{}{}, "pseudo LIKE ?", "%"+partial_pseudo+"%")
	if raw_users == nil {
		return nil, gorm.ErrRecordNotFound
	}
	users := make([]db_tables.User, len(raw_users))
	for i, user := range raw_users {
		users[i] = user.(db_tables.User)
	}

	return users, nil
}

func UserAlreadyExists(db *gorm.DB, pseudo string, email string) (bool, []string, error) {
	user := db_tables.User{}
	err := db.Where("pseudo = ? OR email = ?", pseudo, email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, []string{}, nil
	} else if err != nil {
		return false, []string{}, err
	} else {
		var fields []string
		if user.Pseudo == pseudo {
			fields = append(fields, "pseudo")
		}
		if user.Email == email {
			fields = append(fields, "email")
		}
		return true, fields, nil
	}
}

func GetUsersFromPseudos(db *gorm.DB, pseudos []string) ([]db_tables.User, error) {
	var users []db_tables.User
	err := db.Where("pseudo IN ?", pseudos).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersFromPartialPseudos(db *gorm.DB, partial_pseudos []string) ([]db_tables.User, error) {
	var users []db_tables.User
	query := db.Model(&db_tables.User{})
	for _, partial_pseudo := range partial_pseudos {
		query = query.Or("pseudo LIKE ?", "%"+partial_pseudo+"%")
	}
	err := query.Find(&users).Error
	return users, err
}
