package presenters

// resonse handlersによって生成されたデータのフォーマットを担当する。
// 下記のようにjson形式に直す

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
