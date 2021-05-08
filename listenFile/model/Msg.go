package model

import "time"

type Msg struct {
	Content    string
	Len        int
	ModifyTime time.Time
}