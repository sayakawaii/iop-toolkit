/*
# ------------------------------------------------------------
# -- delete.go
# --
# -- Huang Minghe
# -- 2022-8-2
# ------------------------------------------------------------
*/

package dao

import (
	"omciAnalyzer/global"
)

// func Delete(key string) error {
// 	var ctx = context.Background()

// 	err := global.RedisDB.Del(ctx, key).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func MysqlRecordDataDelete(record any) error {
	result := global.MysqlDB.Model(record).Delete(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
