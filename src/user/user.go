package user

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"mysql"
	_ "mysql"
	"net/http"
)

type User struct {
	OpenId		string	`json:"openId"`
	Nick		string	`json:"nick"`
	Level		string	`json:"level"`
	Avatar		string	`json:"avatar"`
}

func GetUserInfo(code string) (*User, error)  {

	opendId, err := GetOpenId(code)
	if err != nil {
		fmt.Println("getOpenIdErr", err)
		return nil, err
	}
	fmt.Println("通过openId去DB获取UserInfo", opendId)

	opendId = "0"
	user := User{}
	//
	stmt, err := mysql.SharedDb().Prepare("SELECT openId, nick, avatar, level FROM t_user WHERE openId = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	nullNick := sql.NullString{}
	ava := sql.NullString{}
	level := sql.NullString{}

	err = stmt.QueryRow(opendId).Scan(&user.OpenId, &nullNick, &ava, &level)
	if err != nil {
		panic(err.Error())
	}

	user.Nick = nullNick.String
	user.Avatar = ava.String
	user.Level = level.String

	return &user, nil
}


func GetOpenId(code string) (string, error) {

	client := http.Client{}

	appId := "wx48a9d80716d5e341"
	appSecrety := "2d9202f4b54f9be364d96d27e99f4037"

	url := "https://api.weixin.qq.com/sns/jscode2session?appid="
	url += appId
	url += "&secret="
	url += appSecrety
	url += "&js_code="
	url += code
	url += "&grant_type=authorization_code"

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body);
	fmt.Println(bodystr)
	return bodystr, nil;
}


func (u *User) Save() error {

	row := mysql.SharedDb().QueryRow("REPLACE INTO t_user(openId, nick, avatar, level) VALUES (?, ?, ?, ?)", u.OpenId, u.Nick, u.Avatar, u.Level)
	if row != nil {
		return nil;
	} else {
		return fmt.Errorf("save db fail")
	}

}

