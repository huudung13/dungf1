package models

import (
	"fmt"
)

type (
	Account struct {
		ID          int `storm:"id,increment"`
		UserName    string
		DisplayName string
		Password    string
		Description string
		Email       string
		Level       int
	}
)

func AccountKhoiTao() {
	aaaaaaaaaaaaaaa := Account{
		ID:          999999,
		UserName:    "huudung13",
		DisplayName: "DoHuuDung",
		Password:    "123",
		Description: "toilahuudung",
		Level:       2,
	}
	db.Save(&aaaaaaaaaaaaaaa)
}

func GetAccount() (accounts []Account, err error) {
	err = db.Select().Find(&accounts)
	return
}

func (b *Account) Resign() error {
	return db.Save(b)
}

func Login(username string, password string) (accounts Account, err error) {
	var accountcheck Account
	err = db.One("UserName", username, &accountcheck)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("not have this account")
	} else if accountcheck.Password != password {
		fmt.Println("pass not correct")

	} else {
		fmt.Println("login OKE")
		db.One("UserName", username, &accounts)

	}
	return
}

func GetByUserName(username string) (account Account, err error) {
	err = db.One("UserName", username, &account)
	return
}

func GetByUserID(userid int, errr error) (account Account, err error) {
	err = db.One("ID", userid, &account)
	return
}

func (b *Account) UpdateAccount() error {
	return db.Save(b)
}

func DeleteAccount(getid int, errr error) {
	var acc Account
	err := db.One("ID", getid, &acc)
	if err == nil {
		db.DeleteStruct(&acc)
	}

}
