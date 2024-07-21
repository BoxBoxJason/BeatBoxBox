package role_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"testing"
)

// POST FUNCTIONS
func TestPostRole(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	_, err = CreateRole(db, "Test Role 1", "Test Description 1")
	if err != nil {
		t.Error(err)
	}
}

// DELETE FUNCTIONS
func TestDeleteRole(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 2",
		Description: "Test Description 2",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	role_id := role.Id
	err = DeleteRoleFromRecord(db, &role)
	if err != nil {
		t.Error(err)
	}
	err = db.First(&db_tables.Role{}, role_id).Error
	if err == nil {
		t.Error("Role not deleted")
	}
}

func TestDeleteRoles(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 3",
		Description: "Test Description 3",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	role_id := role.Id
	err = DeleteRolesFromRecords(db, []*db_tables.Role{&role})
	if err != nil {
		t.Error(err)
	}
	err = db.First(&db_tables.Role{}, role_id).Error
	if err == nil {
		t.Error("Role not deleted")
	}
}

// PUT FUNCTIONS
func TestUpdateRole(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 4",
		Description: "Test Description 4",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	err = UpdateRole(db, &role, map[string]interface{}{"name": "Test Role 4 updated"})
	if err != nil {
		t.Error(err)
	} else if role.Name != "Test Role 4 updated" {
		t.Error("Role not updated")
	}
}

// GET FUNCTIONS
func TestRoleExistsFromName(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 5",
		Description: "Test Description 5",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	if !RoleExistsFromName(db, role.Name) {
		t.Error("Role not found")
	}
}

func TestGetRoleByName(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 6",
		Description: "Test Description 6",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	role_id := role.Id
	role, err = GetRoleByName(db, role.Name)
	if err != nil {
		t.Error(err)
	} else if role.Id != role_id {
		t.Error("Role not found")
	}
}

func TestGetRoleIdByName(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 7",
		Description: "Test Description 7",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	role_id, err := GetRoleIdByName(db, role.Name)
	if err != nil {
		t.Error(err)
	} else if role_id != role.Id {
		t.Error("Role not found")
	}
}

func TestGetRolesByName(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 8",
		Description: "Test Description 8",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	roles, err := GetRolesByName(db, []string{role.Name})
	if err != nil {
		t.Error(err)
	} else if len(roles) != 1 {
		t.Error("Roles not found")
	}
}

func TestRolesExistFromNames(t *testing.T) {
	db, err := db_model.OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db_model.CloseDB(db)
	role := db_tables.Role{
		Name:        "Test Role 9",
		Description: "Test Description 9",
	}
	err = db.Create(&role).Error
	if err != nil {
		t.Error(err)
	}
	if !RolesExistFromNames(db, []string{role.Name}) {
		t.Error("Roles not found")
	}
}
