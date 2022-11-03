package main

func ooState(boolean bool) string {
	if boolean {
		return "ON"
	}
	return "OFF"
}

func ynState(boolean bool) string {
	if boolean {
		return "YES"
	}
	return "NO"
}
