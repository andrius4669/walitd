package users

import "time"

// external user info exported to forum/news/poll subsystems
// include fields only useful for everyone, skip ones which aren't going to be used elsewhere

type UserInfo struct {
	UserName       string    // short login name
	AccountPicture string    // filename. additional function may be needed to get full path useful for serving this
	RealName       string    // name plus surname. we don't really need them separate
	RegisteredDate time.Time // day of registration. note: it's not exact timestamp, it's just day
	Role           uint32    // uint32 representation of user's privilege. additional functions may exist which can interept this
	Gender         string    // string representation of gender
	Description    string    // description
}

func GetUserInfo(id uint32) (*UserInfo) {
	return nil // TODO: pull data from sql
}

func IsRoleAdmin(r uint32) bool {
	return r > 1
}

func IsRoleSuperAdmin (r uint32) bool {
	return r > 2
}