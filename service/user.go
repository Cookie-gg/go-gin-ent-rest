package service

import (
	"go-gin-ent-rest/ent"
	userQuery "go-gin-ent-rest/ent/user"
)

type UserService struct {
	client *ent.Client
}

func CreateUserService(client *ent.Client) *UserService {
	return &UserService{client}
}

func (s *UserService) Create(user *ent.User) *ent.UserCreate {
	return s.client.User.Create().SetName(user.Name)
}
func (s *UserService) Get(user *ent.User) *ent.UserQuery {
	return s.client.User.Query().Where(userQuery.ID(user.ID)).WithProfile()
}
func (s *UserService) Update(user *ent.User) *ent.UserUpdateOne {
	return s.client.User.UpdateOneID(user.ID).SetName(user.Name)
}
func (s *UserService) Delete(user *ent.User) *ent.UserDeleteOne {
	return s.client.User.DeleteOneID(user.ID)
}
