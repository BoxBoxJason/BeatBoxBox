package db_model

import "gorm.io/gorm"

// DeleteDBRecord deletes a single record from the database	based on its id
func DeleteDBRecord(db *gorm.DB, model interface{}, id int) error {
	err := db.Delete(model, id).Error
	return err
}

// DeleteDBRecordNoFetch deletes a single record from the database based on the record itself
func DeleteDBRecordNoFetch(db *gorm.DB, record interface{}) error {
	err := db.Delete(record).Error
	return err
}

// DeleteDBRecords deletes multiple records from the database based on their ids
func DeleteDBRecords(db *gorm.DB, model interface{}, ids []int) error {
	err := db.Where("id IN ?", ids).Delete(model).Error
	return err
}

// DeleteDBRecordsNoFetch deletes multiple records from the database based on the records themselves
func DeleteDBRecordsNoFetch(db *gorm.DB, records []interface{}) error {
	err := db.Delete(records).Error
	return err
}

// DeleteDBRecordsByField deletes records from the database based on a specific field
func DeleteDBRecordsByField(db *gorm.DB, model interface{}, field string, value interface{}) error {
	err := db.Where(field+" = ?", value).Delete(model).Error
	return err
}

// DeleteDBRecordsByFields deletes records from the database based on multiple fields
func DeleteDBRecordsByFields(db *gorm.DB, model interface{}, fields map[string]interface{}) error {
	err := db.Where(fields).Delete(model).Error
	return err
}

// DeleteAllDBRecords deletes all records from the database
func DeleteAllDBRecords(db *gorm.DB, model interface{}) error {
	err := db.Delete(model).Error
	return err
}
