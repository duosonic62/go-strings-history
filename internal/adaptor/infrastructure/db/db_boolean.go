package db

// bool値をdbに保存する形に変換
func ConvertToDBBoolean(bool bool) string {
	if bool {
		return "1"
	}
	return "0"
}