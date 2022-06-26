package main

func main() {
	user := domain.CreateUser(0, "Kohei", "Okazaki")
	user.ChangeUserName(user.UserName.FirstName, "Noda")

}
