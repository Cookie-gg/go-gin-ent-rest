package service

import (
	"go-gin-ent-rest/ent"
	profileQuery "go-gin-ent-rest/ent/profile"
	userQuery "go-gin-ent-rest/ent/user"
)

type ProfileService struct {
	client *ent.Client
}

func CreateProfileService(client *ent.Client) *ProfileService {
	return &ProfileService{client}
}

func (s *ProfileService) Create(profile *ent.Profile) *ent.ProfileCreate {
	return s.client.Profile.Create().SetAge(profile.Age).SetGender(profile.Gender)
}
func (s *ProfileService) Get(profile *ent.Profile) *ent.ProfileQuery {
	return s.client.Profile.Query().Where(profileQuery.HasUserWith(userQuery.ID(profile.Edges.User.ID)))
}
func (s *ProfileService) Update(profile *ent.Profile) *ent.ProfileUpdate {
	return s.client.Profile.Update().Where(profileQuery.HasUserWith(userQuery.ID(profile.Edges.User.ID))).SetAge(profile.Age).SetGender(profile.Gender)
}
