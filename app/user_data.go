package app

import "github.com/jamesstocktonj1/forlater-core/proto"

type UserData struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type UserToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (u *UserData) toProto() *proto.UserRequest {
	return &proto.UserRequest{
		Username:     u.Username,
		Firstname:    u.Firstname,
		Lastname:     u.Lastname,
		PasswordHash: u.Password,
	}
}

func toUser(user *proto.UserResponse) UserData {
	return UserData{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Password:  "",
	}
}

func (u *UserToken) toProto() *proto.TokenRequest {
	return &proto.TokenRequest{
		Token: u.Token,
	}
}

func toToken(token *proto.TokenResponse) UserToken {
	return UserToken{
		Token: token.Token,
	}
}
