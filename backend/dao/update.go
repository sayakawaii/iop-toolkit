/*
# ------------------------------------------------------------
# -- update.go
# --
# -- Huang Minghe
# -- 2022-8-2
# ------------------------------------------------------------
*/

package dao

import (
	"omciAnalyzer/global"
)

func MysqlRecordDataUpdate(record any, key string, value any) error {
	result := global.MysqlDB.Model(record).Update(key, value)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MysqlRecordDataUpdateMany(record any, updates map[string]any) error {
	result := global.MysqlDB.Model(record).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func MysqlRecordDataUpdateByRecord(old any, new any) error {
	result := global.MysqlDB.Model(old).Updates(new)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
