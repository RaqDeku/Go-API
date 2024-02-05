package services

import "fmt"

// User Struct
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Initialise slice
var userDb []User

// Sign Up Fuction
func (u User) SignUp() string {
	user := User{
		Email:    u.Email,
		Password: u.Password,
	}

	userDb = append(userDb, user)

	return "Sign Up Successful"
}

// Login Function
func (u User) Login() string {
	for _, v := range userDb {
		if u.Email == v.Email && u.Password == v.Password {
			return "Login Suggesstion"
		}
		return "Invalid Email or Password"
	}
	return "User not found"
}

// Logout Function
func (u User) LogOut() string {
	fmt.Println(userDb)
	for i, v := range userDb {
		if u.Email == v.Email {
			userDb = append(userDb[:i], userDb[i+1:]...)
			fmt.Println(userDb)
			return "Logout Successfully"
		}
	}
	return "User not found"
}
