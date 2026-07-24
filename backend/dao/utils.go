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
	"log"
	"omciAnalyzer/global"
)

func MysqlAutoMigrate() error {
	err := global.MysqlDB.AutoMigrate(
		&OmciAnalyzerRequestRecord{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	return nil
}

func MysqlStartupRepair() {
	// 查询所有 progress != 100 的记录
	var records []OmciAnalyzerRequestRecord
	err := MysqlRecordDataReadAllWithCond(
		&records,
		map[string]any{"progress <> ?": 100},
	)
	if err != nil {
		log.Printf("startup repair query failed: %v", err)
		return
	}

	if len(records) == 0 {
		log.Println("startup repair: no incomplete records found")
		return
	}

	log.Printf("startup repair: found %d incomplete records", len(records))

	for _, rec := range records {
		MysqlRecordDataUpdateMany(
			&rec,
			map[string]any{
				"progress":  100,
				"status":    "error",
				"error_msg": "abnormal termination due to server restart",
			},
		)
	}
}
