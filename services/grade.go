package services

func CheckGrade(score int) string {
	switch {
	case score >= 80 && score <= 100:
		return "A"
	case score >= 70 && score < 80:
		return "B"
	case score >= 60 && score < 70:
		return "C"
	case score >= 50 && score < 60:
		return "D"
	default:
		return "F"
	}
}
