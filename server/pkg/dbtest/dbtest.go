package dbtest

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // このファイルでは使わないが他のパッケージで使う場合「_」入れる
)

func DbTest() {
	//sql.Open("mysql", "user:password@/dbname")
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード] dockerからアクセスの際にはローカルのアドレスを直指定する必要がある
	db, err := sql.Open("mysql", "user:pass@tcp(192.168.3.4:8081)/sample01?charset=utf8mb4") // SQLドライバ接続
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}

	rows, err := db.Query("SELECT * FROM name") //
	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
