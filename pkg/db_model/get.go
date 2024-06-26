package db_model

import (
	"BeatBoxBox/pkg/logger"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"fmt"
	"gorm.io/gorm"
)

// RecordExistsFromId fetches a single record from the database based on its id
func RecordExistsFromId(db *gorm.DB, model interface{}, id int) error {
	err := db.First(model, id).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecordFromId fetches a single record from the database based on its id
func GetRecordFromId(db *gorm.DB, model interface{}, id int) interface{} {
	result := db.First(model, id)
	if result.Error != nil {
		return nil
	}
	return model
}

// RecordsExistFromIds fetches multiple records from the database based on their ids
func RecordsExistFromIds(db *gorm.DB, model interface{}, ids []int) error {
	err := db.Where("id IN ?", ids).Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecordsFromIds fetches multiple records from the database based on their ids
func GetRecordsFromIds(db *gorm.DB, model interface{}, ids []int) []interface{} {
	// Create a slice of the model type
	records_ptr := format_utils.CreateSliceOfAny(model)
	// Perform the query
	result := db.Model(model).Where("id IN ?", ids).Find(records_ptr)
	if result.Error != nil {
		logger.Error(result.Error)
		return nil
	}
	// Convert to []interface{}
	return format_utils.ConvertRecordsToInterfaceArray(records_ptr)
}

// RecordExistsByField fetches a single record from the database based on a specific field
func RecordExistsByField(db *gorm.DB, model interface{}, field string, value interface{}) error {
	err := db.Where(field+" = ?", value).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecordsByField fetches multiple records from the database based on a specific field
func GetRecordsByField(db *gorm.DB, model interface{}, field string, value interface{}) []interface{} {
	// Create a slice of the model type
	records_ptr := format_utils.CreateSliceOfAny(model)
	// Perform the query
	result := db.Model(model).Where(field+" = ?", value).Find(records_ptr)
	if result.Error != nil {
		logger.Error(result.Error)
		return nil
	}
	// Convert to []interface{}
	return format_utils.ConvertRecordsToInterfaceArray(records_ptr)
}

// GetRecordByField fetches a single record from the database based on a specific field
func GetRecordByField(db *gorm.DB, model interface{}, field string, value interface{}) interface{} {
	result := db.Where(field+" = ?", value).First(model)
	if result.Error != nil {
		logger.Error(result.Error)
		return nil
	}
	return model
}

// RecordExistsByFields fetches a single record from the database based on multiple fields
func RecordExistsByFields(db *gorm.DB, model interface{}, fields map[string]interface{}) error {
	err := db.Where(fields).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecordsByFields fetches multiple records from the database based on multiple fields
func GetRecordsByFields(db *gorm.DB, model interface{}, fields map[string]interface{}) []interface{} {
	// Create a slice of the model type
	records_ptr := format_utils.CreateSliceOfAny(model)
	// Perform the query
	result := db.Model(model).Where(fields).Find(records_ptr)
	if result.Error != nil {
		logger.Error(result.Error)
		return nil
	}
	// Convert to []interface{}
	return format_utils.ConvertRecordsToInterfaceArray(records_ptr)
}

// RecordExistsByFieldWithCondition fetches a single record from the database based on a specific field with a condition
func RecordExistsByFieldsWithCondition(db *gorm.DB, model interface{}, fields map[string]interface{}, condition string) error {
	err := db.Where(fields).Where(condition).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

// GetRecordsByFieldWithCondition fetches multiple records from the database based on a specific field with a condition
func GetRecordsByFieldsWithCondition(db *gorm.DB, model interface{}, fields map[string]interface{}, condition string, condition_value interface{}) []interface{} {
	// Create a slice of the model type
	records_ptr := format_utils.CreateSliceOfAny(model)
	// Perform the query
	query := db.Model(model).Where(fields)
	if condition != "" {
		query = query.Where(condition, condition_value)
	}
	result := query.Find(records_ptr)
	if result.Error != nil {
		logger.Error(result.Error)
		return nil
	}
	// Convert to []interface{}
	return format_utils.ConvertRecordsToInterfaceArray(records_ptr)
}

// RecordExistsByListFieldElement fetches a single record from the database based on a specific list field if the element exists in the field
func RecordExistsByListFieldElement(db *gorm.DB, model interface{}, field string, value interface{}) error {
	condition := fmt.Sprintf("%s @> ?", field)
	result := db.Where(condition, fmt.Sprintf("{%v}", value)).First(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetRecordsByListFieldElement fetches multiple records from the database based on a specific list field if the element exists in the field
func GetRecordsByListFieldElement(db *gorm.DB, model interface{}, field string, value interface{}) []interface{} {
	// Create a slice of the model type
	records_ptr := format_utils.CreateSliceOfAny(model)
	// Perform the query
	condition := fmt.Sprintf("%s @> ?", field)
	result := db.Where(condition, fmt.Sprintf("{%v}", value)).Find(model)
	if result.Error != nil {
		return nil
	}
	// Convert to []interface{}
	return format_utils.ConvertRecordsToInterfaceArray(records_ptr)
}
