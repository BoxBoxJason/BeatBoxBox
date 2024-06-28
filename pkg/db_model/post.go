package db_model

import (
	"reflect"

	"gorm.io/gorm"
)

// CreateDBRecord creates a new record in the database
func CreateDBRecord(db *gorm.DB, values map[string]interface{}, model interface{}) (int, error) {
	model_value := reflect.ValueOf(model).Elem()
	for key, value := range values {
		field := model_value.FieldByName(key)
		if field.IsValid() {
			field.Set(reflect.ValueOf(value))
		}
	}

	err := db.Create(model).Error
	if err != nil {
		return -1, err
	}
	return model_value.FieldByName("Id").Interface().(int), nil
}
