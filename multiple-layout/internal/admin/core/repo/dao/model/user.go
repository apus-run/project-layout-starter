package model

import (
	"database/sql/driver"
	"encoding/json"

	"multiple-layout/internal/admin/core/domain/entity"
)

type User struct {
	ID   uint64 `gorm:"primaryKey,autoIncrement;comment:'用户ID'"`
	Name string `gorm:"type:varchar(20);not null;comment:'用户名'"`
}

// The TableName method returns the name of the data table that the struct is mapped to.
func (u *User) TableName() string {
	return "user"
}

// MarshalBinary ...
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary ...
func (u *User) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, u)
}

// Value ...
func (u *User) Value() (driver.Value, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

// Scan ...
func (u *User) Scan(input any) error {
	return json.Unmarshal(input.([]byte), u)
}

// ToEntity model to entity
func (u *User) ToEntity() entity.User {
	return entity.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

// FromEntity entity to model
func (u *User) FromEntity(userEntity entity.User) User {
	return User{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}
}
