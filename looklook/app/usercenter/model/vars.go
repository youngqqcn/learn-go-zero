package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var UserAuthTypeSystem string = "system"
var UserAuthTypeSmallWX string = "wxMini" //微信小程序
