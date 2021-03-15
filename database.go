package main

var users []User
var genders []Gender
var universities []University

type User struct {
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
	Gender Gender `json:"gender,omitempty"`
	University University `json:"university,omitempty"`
}

type University struct {
	ID int `json:"id"`
	Institution string `json:"institution,omitempty"`
}

type Gender struct {
	ID int `json:"id"`
	Gender string `json:"gender,omitempty"`
}

type UnsortedArray struct {
	Unsorted []int `json:"unsorted,omitempty"`
}

type SortedArray struct {
	Sorted []int `json:"sorted,omitempty"`
}

type ErrorMessage struct{
	Message string `json:"error,omitempty"`
}

// Initializing datas
func InitDatas() {
	//genders
	male := Gender{ ID:1, Gender: "Hombre"}
	female := Gender{ ID:2, Gender: "Mujer"}

	genders = append(genders,male,female)

	//universities
	unal := University{ ID:1, Institution: "UNAL"}
	udea := University{ ID:2, Institution: "UDEA"}

	universities = append(universities,unal,udea)

	//users
	userOne := User{ ID:1, Name: "Juan", Age: 24, Gender: genders[0], University: universities[0]}
	userTwo := User{ ID:2, Name: "Laura", Age: 19, Gender: genders[1], University: universities[1]}
	userThree := User{ ID:3, Name: "Luis", Age: 22, Gender: genders[0], University: universities[1]}

	users = append(users,userOne,userTwo,userThree)
}

func FindUserByID(ID int) User {
	var result User
	for _, user := range users {
		if user.ID == ID {
			result = user
			break
		}
	}
	return result
}