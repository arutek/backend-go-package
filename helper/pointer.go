package helper

func DeferString(s *string) string {
	if s != nil {
		return *s
	} else {
		return ""
	}
}
func DeferInt(s *int) int {
	if s != nil {
		return *s
	} else {
		return 0
	}
}
