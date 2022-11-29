package helper

import (
	"github.com/faridlan/daily/test/model/domain"
	"github.com/faridlan/daily/test/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {

	userResponse := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return userResponse

}

func ToUserResponses(users []domain.User) []web.UserResponse {

	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
