// storage
package models

import (
	"container/list"
	"fmt"
)

const (
	MAX_USER   = 100 // max user in this system
	MAX_PLAYER = 10  // max player in one court
)

// error state code
type ErrCode int

const (
	NO_ERR = iota
	ERR_MAX_USER
	ERR_ANY = -1
)

// define interactive messges with external
type MsgID int

const (
	USER_REGISTER = iota
	USER_LOGIN
	COURT_ADD
	COURT_REMOVE
	COURT_MODIFY
	PLAYER_ADD
	PLAYER_REMOVE
)

// usually we let court owner decided the courtType
type CourtType int

const (
	MULTI_BALLS = iota
	COMPETITION
	PRACTISE
)

type UserInfo struct {
	UserID   int
	NickName string
	Password string
	Phone    string
	Email    string
}

var userList = list.New()

func Register(user *UserInfo) ErrCode {
	if userList.Len() > MAX_USER {
		return ERR_MAX_USER
	}

	if user.UserID <= 0 || len(user.NickName) == 0 || len(user.Password) == 0 || len(user.Phone) == 0 {
		return ERR_ANY
	}

	// add user to list
	userList.PushBack(*user)

	return NO_ERR
}

func Login(user *UserInfo) ErrCode {
	if user.UserID <= 0 || len(user.Password) <= 0 {
		return ERR_ANY
	}

	for list := userList.Front(); list != nil; list = list.Next() {
		if user.UserID == list.Value.(UserInfo).UserID && user.Password == list.Value.(UserInfo).Password {
			return NO_ERR
		}
	}

	return ERR_ANY
}

func GetNickName(uID int) string {
	for list := userList.Front(); list != nil; list = list.Next() {
		if uID == list.Value.(UserInfo).UserID {
			return list.Value.(UserInfo).NickName
		}
	}

	return ""
}

func init() {
	fmt.Println("STORAGE Init add Test Data .... ")
	user1 := UserInfo{123, "ken", "kenchow", "15220090853", "buzz@boos.com"}
	Register(&user1)
}

type CourtInfo struct {
	Index      int
	Owner      int
	Number     int
	CourtType  CourtType
	Date       string
	Start_time int
	End_time   int
	PlayersID  []int
}

var courtList = list.New()

func AddCourt(court *CourtInfo) ErrCode {
	var vaild = int(0)

	// check the user list if the user had registered
	for users := userList.Front(); users != nil; users = users.Next() {
		if users.Value.(UserInfo).UserID == court.Owner {
			vaild = 1
		}
	}

	if vaild == 0 {
		return ERR_ANY
	}

	// in case court repeat adding
	for courts := courtList.Front(); courts != nil; courts = courts.Next() {
		if courts.Value.(CourtInfo).Number == court.Number && courts.Value.(CourtInfo).Date == court.Date && courts.Value.(CourtInfo).Start_time == court.Start_time {
			return ERR_ANY
		}
	}

	if courtList.Front() != nil {
		// apply unique index for the court
		court.Index = courtList.Back().Value.(CourtInfo).Index + 1
	}

	courtList.PushBack(*court)
	return NO_ERR
}

func RemoveCourt(court *CourtInfo) ErrCode {
	for list := courtList.Front(); list != nil; list = list.Next() {
		if list.Value.(CourtInfo).Index == court.Index {
			courtList.Remove(list)
			return NO_ERR
		}
	}
	return ERR_ANY
}

func ModifyCourt(court *CourtInfo) ErrCode {
	// TODO
	return ERR_ANY
}

func indexOfPlayer(players []int, UserID int) int {
	for i, v := range players {
		if v == UserID {
			return i
		}
	}

	return ERR_ANY
}

func AddPlayer(courtIndex int, UserID int) ErrCode {

	for list := courtList.Front(); list != nil; list = list.Next() {
		if list.Value.(CourtInfo).Index == courtIndex {
			tmp := list.Value.(CourtInfo)

			// verify user if he got another court at the same time
			// TODO

			// find space for the new player
			tmp.PlayersID = append(tmp.PlayersID, UserID)
			list.Value = tmp
		}
	}

	return NO_ERR
}

func RemovePlayer(courtIndex int, UserID int) ErrCode {
	for list := courtList.Front(); list != nil; list = list.Next() {
		if list.Value.(CourtInfo).Index == courtIndex {
			tmp := list.Value.(CourtInfo)

			index := indexOfPlayer(tmp.PlayersID, UserID)
			// delete index value
			tmp.PlayersID = append(tmp.PlayersID[:index], tmp.PlayersID[index+1:]...)

			list.Value = tmp
		}
	}

	return NO_ERR
}

type ScoreBoard struct {
	UserID     int
	joinTimes  int
	totalScore int
	average    int
}
