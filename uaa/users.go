package uaa

import (
	"strings"

	"github.com/xchapter7x/lo"
)

type Users struct {
	userMap map[string][]User
}

func (u *Users) Add(user User) {
	if u.userMap == nil {
		u.userMap = make(map[string][]User)
	}
	key := strings.Trim(strings.ToLower(user.Username), " ")
	existingUsers := u.userMap[key]
	existingUsers = append(existingUsers, user)
	u.userMap[key] = existingUsers
}

func (u *Users) List() []User {
	if u.userMap == nil {
		return nil
	}
	var result []User
	for key := range u.userMap {
		result = append(result, u.userMap[key]...)
	}
	return result
}

func (u *Users) Exists(userName string) bool {
	if u.userMap == nil {
		return false
	}
	_, ok := u.userMap[strings.ToLower(userName)]
	return ok
}

func (u *Users) GetByNameAndOrigin(userName, origin string) *User {
	if u.userMap == nil {
		lo.G.Debugf("GetByNameAndOrigin: No users found in UAA, returning NIL")
		return nil
	}
	userList := u.userMap[strings.ToLower(userName)]
	for _, user := range userList {
		lo.G.Debugf("GetByNameAndOrigin: Checking user's origin [%s] for user [%s] against origin [%s]", user.Origin, user.Username, origin)
		if strings.EqualFold(user.Origin, origin) {
			lo.G.Debugf("GetByNameAndOrigin: Found user [%s] with origin [%s]", user.Username, origin)
			return &user
		}
	}
	lo.G.Debugf("GetByNameAndOrigin: No user found for [%s] with origin [%s], returning NIL", userName, origin)
	return nil
}

func (u *Users) GetByID(ID string) *User {
	for _, user := range u.List() {
		if strings.EqualFold(user.GUID, ID) {
			return &user
		}
	}
	return nil
}

func (u *Users) GetByExternalID(externalID string) *User {
	var foundUsers []User
	for _, user := range u.List() {
		if strings.EqualFold(user.ExternalID, externalID) {
			foundUsers = append(foundUsers, user)
		}
	}
	if len(foundUsers) == 1 {
		return &foundUsers[0]
	} else {
		for _, user := range foundUsers {
			lo.G.Debugf("Multiple User [%s] found for externalID [%s]", user.Username, externalID)
		}
	}
	return nil

}
