package users

import(
	"fmt"
	"errors"
)

type User struct {
    FirstName	string	`json:"FirstName"`
    LastName	string	`json:"LastName"`
}


var Test []User

func init(){
	Test = append(Test, User{"A", "B"})
	Test = append(Test, User{"C", "D"})
	fmt.Printf("Users have been initialized.\n")
}

func GetUsers() (users []User, err error){
	if (len(Test) > 0){
		users = Test
	}else{
		return users, errors.New("There no users.")
	}

	return 
}

func GetUser(id int) (user User, err error){
	if (id >= 0 && id < len(Test)){
		user = Test[id]
	}else{
		return user, errors.New("There no such user.")
	}

	return
}

func CreateUser(FirstName string, LastName string){
	Test = append(Test, User{FirstName, LastName})
}

func UpdateUser(id int, FirstName string, LastName string){
	Test[id] = User{FirstName, LastName}
}

func DeleteUser(id int){
	Test = append(Test[:id], Test[id+1:]...)
}