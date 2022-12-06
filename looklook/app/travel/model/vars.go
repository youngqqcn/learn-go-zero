package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var HomestayActivityPreferredType = "preferredHomestay" //优选民宿
var HomestayActivityGoodBusiType = "goodBusiness"       //最佳房东

var HomestayActivityDownStatus int64 = 0 //上架
var HomestayActivityUpStatus int64 = 1   //上架

var HomestayCommentNotDeleted int64 = 0
