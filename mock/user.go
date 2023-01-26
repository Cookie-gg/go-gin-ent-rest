package mock

import (
	"encoding/json"
	"go-gin-ent-rest/ent"
	"go-gin-ent-rest/util"
	"sort"
	"time"
)

var User = ent.User{
	ID:        1,
	Name:      "テスト田中",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Edges:     ent.UserEdges{Profile: &Profile},
}

var Profile = ent.Profile{
	ID:      1,
	Age:     18,
	Gender:  "male",
}

var userJsonByte, _ = json.Marshal(User)

var UserJson = string(userJsonByte)

var UserColumns = util.Filter(util.GetJsonKyes(userJsonByte), func(key string) bool {
	return key != "profile" && key != "edges"
})

var profileJsonByte, _ = json.Marshal(Profile)

var ProfileJson = string(profileJsonByte)

var profileColumns = util.Filter(util.GetJsonKyes(profileJsonByte), func(key string) bool {
	return key != "edges"
})
var ProfileColumns = append(profileColumns, "user_id")

func init() {
	sort.SliceStable(UserColumns, func(i, j int) bool {
		return UserColumns[i] < UserColumns[j]
	})
	sort.SliceStable(ProfileColumns, func(i, j int) bool {
		return ProfileColumns[i] < ProfileColumns[j]
	})
}
