// distribute.go

package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kenchowcn/go_tennis/models"
)

type UserEvent struct {
	msgID models.MsgID
	user  models.UserInfo
	conn  *websocket.Conn // only exist in user event
}

type CourtEvent struct {
	MsgID int `json: "MsgID"`
	court models.CourtInfo
}

type PlayerEvent struct {
	msgID      models.MsgID
	courtIndex int
	userID     int
}

var (
	// a chan for handling user process
	userHandle = make(chan UserEvent, 10)
	// a chan for handling court process
	courtHandle = make(chan CourtEvent, 10)
	// a chan for handling player process
	playerHandle = make(chan PlayerEvent, 10)
)

func userHndl(event *UserEvent) {
	switch event.msgID {
	case models.USER_REGISTER:
		models.Register(&event.user)
	case models.USER_LOGIN:
		models.Login(&event.user)
	default:
		fmt.Println("MSG ID ERROR.")
	}
}

func courtHndl(event *CourtEvent) {
	//	switch event.MsgID {
	//	case models.COURT_ADD:
	//		models.AddCourt(&event.court)
	//	case models.COURT_MODIFY:
	//		models.ModifyCourt(&event.court)
	//	case models.COURT_REMOVE:
	//		models.RemoveCourt(&event.court)
	//	default:
	//		fmt.Println("MSG ID ERROR.")
	//	}
}

func playerHndl(event *PlayerEvent) {
	switch event.msgID {
	case models.PLAYER_ADD:
		models.AddPlayer(event.courtIndex, event.userID)
	case models.PLAYER_REMOVE:
		models.RemovePlayer(event.courtIndex, event.userID)
	default:
		fmt.Println("MSG ID ERROR.")
	}
}

func distribute() {
	fmt.Println("Start Distributing .... ")

	for {
		select {
		case userE := <-userHandle:
			userHndl(&userE)
		case courtE := <-courtHandle:
			courtHndl(&courtE)
		case playerE := <-playerHandle:
			playerHndl(&playerE)
		}
	}

}

func init() {
	go distribute()
}
