package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var HomestayActivityPreferredType = "preferredHomestay" //优选民宿

var HomestayActivityUpStatus int64 = 1 //上架
