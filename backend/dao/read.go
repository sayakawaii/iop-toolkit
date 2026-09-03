/*
# ------------------------------------------------------------
# -- read.go
# --
# -- Huang Minghe
# -- 2022-8-2
# ------------------------------------------------------------
*/

package dao

import (
	"omciAnalyzer/global"
	"time"
)

// func Read(key string) string {
// 	var ctx = context.Background()

// 	val, err := global.RedisDB.Get(ctx, key).Result()
// 	if err == nil {
// 		return val
// 	} else if err.Error() == redis.Nil.Error() {
// 		utils.Log("user data not exist")
// 	} else {
// 		panic(err)
// 	}

// 	return ""
// }

func MysqlRecordDataRead(record any, key string, value any) error {
	result := global.MysqlDB.Model(record).Where(key+" = ? ", value).Find(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MysqlRecordDataReadAll(record any) error {
	result := global.MysqlDB.Model(record).Find(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MysqlRecordDataReadAllWithCond(record any, cond map[string]any) error {
	db := global.MysqlDB.Model(record)

	for key, val := range cond {
		db = db.Where(key, val)
	}

	result := db.Find(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func MysqlRecordDataCount(record any) (uint64, error) {
	var counter int64

	result := global.MysqlDB.Model(record).Count(&counter)
	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(counter), nil
}

func MysqlRecordDataCountToday(record any) (uint64, error) {
	var counter int64

	today := time.Now().Format("2006-01-02")

	result := global.MysqlDB.
		Model(record).
		Where("DATE(created_at) = ?", today).
		Count(&counter)

	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(counter), nil
}

func MysqlRecordDataReadFirst(record any, key string, value any) error {
	result := global.MysqlDB.Model(record).Where(key+" = ? ", value).First(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MysqlRecordDataReadLast(record any, key string, value any) error {
	result := global.MysqlDB.Model(record).Where(key+" = ? ", value).Last(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MysqlRecordDataReadMaxID(record any) error {
	result := global.MysqlDB.Last(record)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IOPLibraryRecordDataSearch(str string) []IOPLibraryRecord {
	cu := "customer like ? OR "
	rc := "rcr_num like ? OR "
	ve := "vendor like ? OR "
	on := "onu_type like ? OR "
	po := "pon_type like ? OR "
	rg := "rg_mode like ? OR "
	eq := "equip_id like ? OR "
	hw := "hw_version like ? OR "
	sw := "sw_version like ? OR "
	re := "ls_release like ? OR "
	co := "comments like ? OR "
	lo := "input_log like ? "
	searchStr := cu + rc + ve + on + po + rg + eq + hw + sw + re + co + lo
	var record []IOPLibraryRecord
	// global.MysqlDB.Find(&record, condition)
	global.MysqlDB.Where(searchStr, "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%").Find(&record)
	return record
}
