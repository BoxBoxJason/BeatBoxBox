package db_model

import (
	"reflect"

	"gorm.io/gorm"
)

// EditRecordField edits a single field of a record in the database
func EditRecordField(db *gorm.DB, record interface{}, field string, value interface{}) error {
	result := db.Model(record).Update(field, value)
	return result.Error
}

// EditRecordFields edits multiple fields of a record in the database
func EditRecordFields(db *gorm.DB, record interface{}, fields map[string]interface{}) error {
	result := db.Model(record).Updates(fields)
	return result.Error
}

// AddElementToRecordListField adds an element to a list field of a record in the database
func AddElementToRecordListField(db *gorm.DB, record interface{}, field string, value interface{}) error {
	record_value := reflect.ValueOf(record).Elem()
	field_value := record_value.FieldByName(field)

	if !field_value.IsValid() || field_value.Kind() != reflect.Slice {
		return gorm.ErrInvalidField
	}

	field_value.Set(reflect.Append(field_value, reflect.ValueOf(value)))
	return db.Save(record).Error
}

// RemoveElementFromRecordListField removes an element from a list field of a record in the database
func RemoveElementFromRecordListField(db *gorm.DB, record interface{}, field string, value interface{}) error {
	record_value := reflect.ValueOf(record).Elem()
	field_value := record_value.FieldByName(field)

	if !field_value.IsValid() || field_value.Kind() != reflect.Slice {
		return gorm.ErrInvalidField
	}

	for i := 0; i < field_value.Len(); i++ {
		if field_value.Index(i).Interface() == value {
			field_value.Set(reflect.AppendSlice(field_value.Slice(0, i), field_value.Slice(i+1, field_value.Len())))
			return db.Save(record).Error
		}
	}
	return gorm.ErrRecordNotFound
}

// RemoveElementsFromAssociation removes elements from an association of a record in the database
func RemoveElementsFromAssociation(db *gorm.DB, record interface{}, association string, association_records interface{}) error {
	return db.Model(record).Association(association).Delete(association_records)
}

// AddElementsToAssociation adds elements to an association of a record in the database
func AddElementsToAssociation(db *gorm.DB, record interface{}, association string, association_records interface{}) error {
	return db.Model(record).Association(association).Append(association_records)
}
