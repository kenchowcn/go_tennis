// storage_test.go
package models

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {

	for _, c := range []struct {
		in   UserInfo
		want ErrCode
	}{
		{UserInfo{123, "Ken1", "HelloWorld", "12345678901", "buzz@boos.com"}, NO_ERR}, // normal1
		{UserInfo{124, "Ken2", "HelloWorld", "12345678901", "buzz@boos.com"}, NO_ERR}, // normal2
		{UserInfo{125, "Ken3", "HelloWorld", "12345678901", "buzz@boos.com"}, NO_ERR}, // normal3
		{UserInfo{0, "Ken", "HelloWorld", "12345678901", "buzz@boos.com"}, ERR_ANY},   // username error
		{UserInfo{123, "", "HelloWorld", "12345678901", "buzz@boos.com"}, ERR_ANY},    // password error
		{UserInfo{123, "Ken", "", "12345678901", "buzz@boos.com"}, ERR_ANY},           // nickName error
		{UserInfo{123, "Ken", "HelloWorld", "", "buzz@boos.com"}, ERR_ANY},            // phone error
		{UserInfo{123, "Ken", "HelloWorld", "12345678901", ""}, NO_ERR},               // no email
	} {
		got := Register(&c.in)
		if got != c.want {
			t.Errorf("TestRegister(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLogin(t *testing.T) {
	for _, c := range []struct {
		in   UserInfo
		want ErrCode
	}{
		{UserInfo{123, "", "HelloWorld", "", ""}, NO_ERR},    // normal
		{UserInfo{1234, "", "HelloWorld", "", ""}, ERR_ANY},  // userID wrong
		{UserInfo{123, "", "HelloWorld1", "", ""}, ERR_ANY},  // password wrong
		{UserInfo{1234, "", "HelloWorld1", "", ""}, ERR_ANY}, // both wrong
	} {
		got := Login(&c.in)
		if got != c.want {
			t.Errorf("Login(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func printCourtList() {
	for courts := courtList.Front(); courts != nil; courts = courts.Next() {
		fmt.Println("COURT LIST: ", courts.Value.(CourtInfo), "LIST LEN: ", courtList.Len())
	}
}

func TestAddCourt(t *testing.T) {
	for _, c := range []struct {
		in   CourtInfo
		want ErrCode
	}{
		{CourtInfo{0, 123, 4, MULTI_BALLS, "2015-11-11", 6789, 7789, []int{}}, NO_ERR},   // normal1
		{CourtInfo{0, 123, 5, MULTI_BALLS, "2015-11-11", 6789, 7789, []int{}}, NO_ERR},   // normal2
		{CourtInfo{0, 123, 5, COMPETITION, "2015-11-12", 6789, 7789, []int{}}, NO_ERR},   // normal3
		{CourtInfo{0, 123, 5, COMPETITION, "2015-11-12", 6779, 7789, []int{}}, NO_ERR},   // normal4
		{CourtInfo{0, 1234, 4, MULTI_BALLS, "2015-11-11", 6789, 7789, []int{}}, ERR_ANY}, // owner is not a registered user
		{CourtInfo{0, 123, 4, MULTI_BALLS, "2015-11-11", 6789, 7789, []int{}}, ERR_ANY},  // repeat court
	} {
		got := AddCourt(&c.in)
		if got != c.want {
			t.Errorf("AddCourt(%q)==%q, want %q", c.in, got, c.want)
		}
	}

	i := 0

	for courts := courtList.Front(); courts != nil; courts = courts.Next() {
		if courts.Value.(CourtInfo).index != i {
			t.Errorf("AddCourt.index = %q, i = %q", courts.Value.(CourtInfo).index, i)
		}
		i = i + 1
	}
}

func TestAddPlayer(t *testing.T) {
	for _, c := range []struct {
		in1, in2 int
		want     ErrCode
	}{
		{2, 123, NO_ERR},
		{1, 123, NO_ERR},
		{0, 123, NO_ERR},
	} {
		got := AddPlayer(c.in1, c.in2)
		if got != c.want {
			t.Errorf("AddCourt(%d, %d)==%d, want %d", c.in1, c.in2, got, c.want)
		}
	}

	printCourtList()
}

func TestRemovePlayer(t *testing.T) {
	for _, c := range []struct {
		in1, in2 int
		want     ErrCode
	}{
		{0, 123, NO_ERR},
	} {
		got := RemovePlayer(c.in1, c.in2)
		if got != c.want {
			t.Errorf("RemovePlayer(%d, %d)==%d, want %d", c.in1, c.in2, c.want)
		}
	}

	printCourtList()
}

func TestRemoveCourt(t *testing.T) {
	for _, c := range []struct {
		in   CourtInfo
		want ErrCode
	}{
		{CourtInfo{1, 0, 0, COMPETITION, "", 0, 0, []int{}}, NO_ERR}, // index = 1
		{CourtInfo{3, 0, 0, COMPETITION, "", 0, 0, []int{}}, NO_ERR}, // index = 3
	} {
		got := RemoveCourt(&c.in)
		if got != c.want {
			t.Errorf("AddCourt(%q)==%q, want %q", c.in, got, c.want)
		}
	}

	printCourtList()
}
