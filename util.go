package main

func ooState(boolean bool) string {
	if boolean {
		return "ON"
	}
	return "OFF"
}

func ooBoolean(state string) bool {
	if state == "ON" {
		return true
	}
	return false
}

func ynState(boolean bool) string {
	if boolean {
		return "YES"
	}
	return "NO"
}
