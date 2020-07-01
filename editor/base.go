package editor

import "api/provider"

var baseProvider = provider.Base

//基础editor模块
type IEditor interface {
	Save()
}

type Editor struct {
	User
}
