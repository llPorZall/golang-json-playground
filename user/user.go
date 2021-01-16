package user

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/llporzall/golang-json-playground/user/entity"
)

//PATH variable
const PATH = "../../user.json"

type datastore struct {
	connector
}

//GetAllUser function

func getUserData() ([]byte, error) {
	data := datastore{connector{PATH}}
	dataArr, err := data.readFile()
	if err != nil {
		return nil, errors.New("Covert file string to json error")
	}
	return dataArr, nil
}

func stringToJSON(data []byte) ([]entity.User, error) {
	var userList []entity.User
	err := json.Unmarshal(data, &userList)
	if err != nil {
		return nil, errors.New("Covert file json to string error")
	}
	return userList, nil
}

func jsontoString(userList []entity.User) ([]byte, error) {
	data, err := json.Marshal(userList)
	if err != nil {
		return nil, errors.New("Covert file string to json error")
	}
	return data, nil
}

//GetAllUser function
func GetAllUser() {
	data, err := getUserData()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}

//AddNewUser function
func AddNewUser(name string, age int) {
	data, err := getUserData()
	if err != nil {
		fmt.Println(err)
	} else {
		userList, err := stringToJSON(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		user := entity.User{Name: name, Age: age}
		user.ID = len(userList) + 1
		userList = append(userList, user)
		store := datastore{connector{PATH}}
		dataArr, err := jsontoString(userList)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = store.writeFile(dataArr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
