package helper

func Response(msgVal string, dataVal interface{}, totalData int64) map[string]interface{} {
	resMsg := getResMsg(msgVal)
	resVal := map[string]interface{}{
		"message": resMsg,
		"data":    dataVal,
	}
	if totalData >= 0 {
		resVal["count"] = totalData
	}
	return resVal
}
func getResMsg(resCode string) string {
	resList := map[string]string{
		"DATA_UPDATED":        "Data have been updated",
		"INVALID_CSV":         "Invalid CSV were uploaded",
		"TOKEN_INVALID":       "Authorized access invalid",
		"TOKEN_EXPIRED":       "Authorized access expired",
		"INVALID_HEADER_TYPE": "Invalid request header",
		"CSV_FAIL":            "Failed to process the uploaded CSV",
	}
	val, ok := resList[resCode]
	if ok {
		return val
	}
	return resCode
}
