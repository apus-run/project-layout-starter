package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// User 用户 Entity
type User struct {
	ID    uint64 // 用户ID
	Name  string // 用户名
	Email string // 邮箱

	CreatedTime time.Time  // 创建时间
	UpdatedTime time.Time  // 更新时间
	DeletedTime *time.Time // 删除时间

	ChangeTracker
}

// 实体的赋值方法
// ------------------------------------------------------------------------

func (u *User) SetID(id uint64) *User {
	u.change()
	u.ID = id
	return u
}

func (u *User) SetName(name string) *User {
	u.change()
	u.Name = name
	return u
}

func (u *User) SetEmail(email string) *User {
	u.change()
	u.Email = email
	return u
}

// 实体 JSON 序列化和反序列化
// ------------------------------------------------------------------------

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, u)
}

func (u *User) Value() (driver.Value, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func (u *User) Scan(input any) error {
	return json.Unmarshal(input.([]byte), u)
}
