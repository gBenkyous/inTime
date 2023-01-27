package entity

import (
	"time"
)

//エンティティ層
type User struct {
	//	ID        ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	//	Books     []ID
}

// 単体テストもこちらで設定
// ただし、entの場合テストが行われるかも見れるためいらんかも
