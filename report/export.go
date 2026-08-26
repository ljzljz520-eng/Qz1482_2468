package report

import (
	"aerialfarm/domain"
	"encoding/json"
)

func EncodeSummary(s Summary) ([]byte, error)          { return json.Marshal(s) }
func EncodeRecords(rs []domain.Record) ([]byte, error) { return json.Marshal(rs) }
func StatusLabel(status string) string {
	switch status {
	case "processed":
		return "已处理"
	case "pending":
		return "待处理"
	case "archived":
		return "已归档"
	default:
		return "处理中"
	}
}
