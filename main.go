package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func errorHandler(err error) {
	fmt.Println("Oops, something went wrong:", err)
}

func viewAttendance() {
	fmt.Println("View Attendance\n")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		fmt.Println("No attendance records available")
		return
	}

	file, err := os.Open("attendance.txt")
	if err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		errorHandler(err)
		log.Fatal(err)
	}
}

func resetAttendance() {
	fmt.Println("Reset Attendance\n")

	if _, err := os.Stat("attendance.txt"); os.IsNotExist(err) {
		fmt.Println("Attendance already cleared")
		return
	}

	if err := os.Remove("attendance.txt"); err != nil {
		errorHandler(err)
		log.Fatal(err)
	}

	fmt.Println("Attendance Cleared")
}

func getStudentInfo() (normtime, epochtime, name, roll, course string) {
	now := time.Now()
	fmt.Println("Time: ", now.Local(), "\n")
	epoch := now.Unix()
	norm := now.Local()
	epochtime = fmt.Sprint(epoch)
	normtime = fmt.Sprint(norm)
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the student name: ")
	name, _ = reader.ReadString('\n')

	fmt.Print("Enter the student roll number: ")
	roll, _ = reader.ReadString('\n')

	fmt.Print("Enter the course: ")
	course, _ = reader.ReadString('\n')

	return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}

func main() {
	fmt.Println("Attendance Management System\n")

	fmt.Println("Options:")
	fmt.Println("1. View attendance")
	fmt.Println("2. Log attendance")
	fmt.Println("3. Reset attendance")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		viewAttendance()
	case 2:
		normtime, epochtime, name, roll, course := getStudentInfo()
		record := []string{normtime, epochtime, name, roll, course}

		file, err := os.OpenFile("attendance.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			errorHandler(err)
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write(record); err != nil {
			errorHandler(err)
			log.Fatalf("%s", err)
		}
	case 3:
		resetAttendance()
	default:
		fmt.Println("Invalid option selected")
	}
}
