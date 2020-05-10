package pkgJson

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const (
	UID_HEADER = "UID_00"
)

const (
	GOLANG = iota + 1
	PYTHON
	JAVA
)

type ConvertToUserInfo []struct {
	ID       string `json:"Id"`
	Password string `json:"Password"`
	Age      int    `json:"Age"`
}

type UserInfo struct {
	Id       string
	Password string
	Age      int
}

type UserInfoList []*UserInfo

var (
	userInfo  UserInfo
	userInfos []UserInfo
)

func InitUserInfo(id, password string, age int) (u *UserInfo) {
	u = new(UserInfo)
	u.Id = id
	u.Password = password
	u.Age = age

	return u
}

func JsonPackageSummary() {
	var users UserInfoList
	for idx := 1; idx <= 5; idx++ {
		switch idx {
		case GOLANG:
			users = append(users, InitUserInfo(UID_HEADER+strconv.Itoa(idx), "golangUser", 2009))
		case PYTHON:
			users = append(users, InitUserInfo(UID_HEADER+strconv.Itoa(idx), "pythonUser", 1990))
		default:
		}
	}
	for _, u := range users {
		fmt.Println(u)
	}
	rowJson := PayLoadToJsonData(users)
	PayLoadToStructData(rowJson)
}

func PayLoadToJsonData(u UserInfoList) string {
	responseByte, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(responseByte))
	return string(responseByte)
}

func PayLoadToStructData(jsonData string) {
	c := new(ConvertToUserInfo)
	fmt.Println(c)
	if err := json.Unmarshal([]byte(jsonData), &c); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(c)
}

func (u *UserInfo) NewUserInfo(id, password string, age int) *UserInfo {
	u.Id = id
	u.Password = password
	u.Age = age

	return u
}

/*func FixTest() {
	u := make([]UserInfo, 5)

	for idx, uinfo := range u {
		switch idx {
		case GOLANG:
			uinfo.NewUserInfo(UID_HEADER+strconv.Itoa(idx), "golangUser", 2009)
		case PYTHON:
			uinfo.NewUserInfo(UID_HEADER+strconv.Itoa(idx), "pythonUser", 1990)
		case JAVA:
			uinfo.NewUserInfo(UID_HEADER+strconv.Itoa(idx), "javaUser", 1991)
		default:
		}
		//TODO
		if idx != 0 {
			PayLoadToJsonData(uinfo)
		}
	}
	fmt.Println(u)

}*/
