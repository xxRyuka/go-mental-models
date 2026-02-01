package main

import "fmt"

func Logger(level string, message ...string) {
	var fullMessage string = level + ": "

	for i, msg := range message {
		if i < len(message)-1 {
			fullMessage += msg + " | - | "
		} else {
			fullMessage += msg
		}
	}

	fmt.Println(fullMessage)
}

func main() {
	Logger("INFO", "Application started")
	Logger("ERROR", "File not found", "Check the path", "Operation aborted")
	dbErrors := []string{"Connection failed", "Timeout occurred", "Invalid credentials"}

	//Logger("DB_ERROR", dbErrors)  slicenin yanına variadic operatorunu eklemeliyiz
	Logger("DB_ERROR", dbErrors...)
}
