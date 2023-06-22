package libs

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// .envからdb情報を取得
	err := godotenv.Load("db.env")
	if err != nil {
		return fmt.Errorf("DB接続情報がありません:%v", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("BD_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("DB接続に失敗しました:%v", err)
	}

	// //db.Migrator().DropTable(&Test{}) // 既存のテーブルを削除
	// // テーブルがない場合は新規作成する
	// if !DB.Migrator().HasTable(&model.TbTUser{}) {
	// 	// log.Fatalln("テーブルを作成する") テーブル作成がなぜか失敗する
	// 	DB.Migrator().CreateTable(&model.TbTUser{})
	// }
	return nil
}
