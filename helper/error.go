package helper

func Error(err error, errCode string) map[string]interface{} {
	LoggerErr(err.Error())
	return Response(errCode, nil, -1)
}
