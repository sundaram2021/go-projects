package main 

import "fmt"


type User struct {
	id int32
	name string
}

type Users struct {
	users []User
}

func CreateUser(id int32, name string) *User {
	return &User{id, name}
}

func (u *Users) DeleteUser(id int32) User {
	var usr User

	for i, ele := range u.users {
		if ele.id == id {
			usr = ele
			u.users = append(u.users[:i],u.users[i+1:]... )
			break
		}
	}

	return usr
}

func (u *Users) UpdateUser(user User) []User {
	var usr []User

	for i, ele := range u.users{
		if ele.id == user.id {
			u.users[i] = user
			usr = u.users
			break
		}
	}

	return usr
}

func (u *Users) GetUser(id int32) User {
	var usr User

	for _, ele := range u.users {
		if ele.id == id {
			usr = ele
			break
		}
	}
	return usr
}


func main() {
	user1 := User{id: 1, name: "Sundaram"}
	user2 := User{id: 2, name: "Aman"}

	var users = Users {
		users : []User {
			user1,
			user2,
		},
	}

	var get = users.GetUser(2)
	fmt.Println("User of 2 :", get)
	createdUser  := CreateUser(3, "Pawan")
	fmt.Println("created user : ", createdUser)

	newUser := CreateUser(3, "Ravi")
	users.users = append(users.users, *newUser)
	fmt.Println("After Adding a User:", users.users)

	deletedUser := users.DeleteUser(1)
	fmt.Println("Deleted User:", deletedUser)
	fmt.Println("After Deleting a User:", users.users)

	updatedUser := User{id: 2, name: "Aman Updated"}
	users.UpdateUser(updatedUser)
	fmt.Println("After Updating a User:", users.users)

	retrievedUser := users.GetUser(3)
	fmt.Println("Retrieved User:", retrievedUser)
}