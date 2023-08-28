package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// - DBにつないでUserの構造体からnameを引いてきて、最初の10文字を返す GetHeadUserName() メソッドがある
// - DBからUser構造体を返すパターンを以下の4つ用意する
//     - ちゃんとしたユーザ
//     - ユーザのnameが空文字
//     - ユーザのnameが100文字
// - テストパターンは以下の3つを用意する
//     - 本物のDBにつなぐ
//     - Fake DBを使う
//     - gomockを使う

type User struct {
	UserID int
	Name   string
}

type UserRepository interface {
	GetUserNameByID(userID int) (string, error)
}

type userRepository struct{}
type fakeUserRepository struct{}

// おまじない. userRepositoryがUserRepositoryを満たしているかどうかをコンパイルでチェックする
var _ UserRepository = (*userRepository)(nil)

// 最初の10文字を返す
func (up *userRepository) GetUserNameByID(userID int) (string, error) {
	db, err := DBConnect()
	if err != nil {
		log.Fatalf("failed to connect to the MySQL database: %v", err)
		return "", err
	}
	user := User{}
	err = db.QueryRow("SELECT id, name FROM users where id = ?", userID).Scan(&user.UserID, &user.Name)
	if err != nil {
		log.Fatalf("failed to query row: %v", err)
		return "", err
	}
	defer db.Close()

	return user.Name, nil
}

func (up *fakeUserRepository) GetUserNameByID(userID int) (string, error) {
	if userID == 1 {
		str := "20testtest"
		return str, nil
	}
	if userID == 2 {
		return "", nil
	}
	if userID == 3 {
		str := "100testtes"
		return str, nil
	}
	return "", fmt.Errorf("userIDが不正です")
}

// Fake DBを使う

func DBConnect() (db *sql.DB, err error) {
	// DBに接続する処理
	dsn := "root:@tcp(localhost:3306)/gomock?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to the MySQL database: %v", err)
		return nil, err
	}
	fmt.Println("success")
	return db, nil
}

type UserUsecase struct {
	Repository UserRepository
}

// GetHeadUserNameByID ユーザ名の先頭10文字を返す
func (uc *UserUsecase) GetHeadUserNameByID(userID int) (string, error) {
	name, err := uc.Repository.GetUserNameByID(userID)
	if err != nil {
		return "", err
	}

	if len(name) < 10 {
		return name, nil
	}
	return name[:10], nil
}

func main() {}
