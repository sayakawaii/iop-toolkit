/*
# ------------------------------------------------------------
# -- create.go
# --
# -- Huang Minghe
# -- 2022-8-2
# ------------------------------------------------------------
*/
package dao

import (
	"omciAnalyzer/global"
)

// func Create(key string, value any) error {
// 	var ctx = context.Background()

// 	err := global.RedisDB.Set(ctx, key, value, 0).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func MysqlRecordDataInsert(record any) error {
	result := global.MysqlDB.Create(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
