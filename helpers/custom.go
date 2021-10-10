package helpers

func GetEndRangeDate(month string) string {
	end := ""
	switch month {
	case "JAN":
		end = "31"
	case "FEB":
		end = "28"
	case "MAR":
		end = "31"
	case "APR":
		end = "30"
	case "MAY":
		end = "31"
	case "JUN":
		end = "30"
	case "JUL":
		end = "31"
	case "AUG":
		end = "31"
	case "SEP":
		end = "30"
	case "OCT":
		end = "31"
	case "NOV":
		end = "30"
	case "DEC":
		end = "31"
	}
	return end
}
