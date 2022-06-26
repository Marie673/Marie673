package domain

type User struct {
	UserId   UserId
	UserName *UserName
}

func CreateUser(userId UserId, firstname FirstName, lastName LastName) User {
	// idの被り
	username, err := newUserName(firstname, lastName)
	if err != nil {
		panic(err)
	}

	return User{UserId: userId, UserName: username}
}

func (user *User) ChangeUserName(firstname FirstName, lastName LastName) {
	username, err := newUserName(firstname, lastName)
	if err != nil {
		panic(err)
	}
	user.UserName = username
}
