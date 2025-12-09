package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name       string
	Roll       string
	Subjects   []Subject
	TotalMarks int
	Percentage float64 
	Grade      string
	Passed     bool
}

type Subject struct {
	Name  string
	Marks int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of Student: ")
	name := readLine(reader)

	fmt.Print("Roll number: ")
	roll := readLine(reader)

	fmt.Print("Number of subjects: ")
	nStr := readLine(reader)
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		fmt.Println("Invalid number of subjects. Using 5 by default.")
		n = 5
	}
// Loop : Loop is a term that is used to perform similar task multiple time in any programing language 
	subjects := make([]Subject, 0, n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Name of subject %d: ", i)
		sname := readLine(reader)
		if strings.TrimSpace(sname) == "" {
			sname = fmt.Sprintf("Subject %d", i)
		}

		fmt.Printf("Marks obtained in %s (out of 100): ", sname)
		marksStr := readLine(reader)
		marks, err := strconv.Atoi(marksStr)
		if err != nil || marks < 0 {
			fmt.Println("Invalid marks - setting to 0.")
			marks = 0
		}
		if marks > 100 {
			fmt.Println("Marks cannot exceed 100 - setting to 100.")
			marks = 100
		}
		subjects = append(subjects, Subject{Name: sname, Marks: marks})
	}

	st := Student{
		Name:     name,
		Roll:     roll,
		Subjects: subjects,
	}

	calcResult(&st)   
	printMarksheet(st)
}

func readLine(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func calcResult(st *Student) {
	total := 0
	pass := true
	for _, s := range st.Subjects {
		total += s.Marks
		if s.Marks < 40 { // passing threshold per-subject
			pass = false
		}
	}
	st.TotalMarks = total

	maxTotal := len(st.Subjects) * 100
	if maxTotal == 0 {
		st.Percentage = 0
	} else {
		st.Percentage = (float64(total) / float64(maxTotal)) * 100.0
	}

	st.Passed = pass
	st.Grade = gradeFromPercentage(st.Percentage, st.Passed)
}

func gradeFromPercentage(p float64, passed bool) string {
	if !passed {
		return "F"
	}
	switch {
	case p >= 80:
		return "A+"
	case p >= 70:
		return "A"
	case p >= 60:
		return "B"
	case p >= 50:
		return "C"
	case p >= 40:
		return "D"
	default:
		return "F"
	}
}

func printMarksheet(st Student) {
	fmt.Println("\n========================================")
	fmt.Println("               MARKSHEET OF STUDENT                ")
	fmt.Println("========================================")
	fmt.Printf("Name : %s\n", st.Name)
	fmt.Printf("Roll : %s\n", st.Roll)
	fmt.Println("----------------------------------------")
	fmt.Printf("%-3s  %-20s  %6s\n", "No", "Subject", "Marks")
	fmt.Println("----------------------------------------")
	for i, s := range st.Subjects {
		fmt.Printf("%-3d  %-20s  %6d\n", i+1, s.Name, s.Marks)
	}
	fmt.Println("----------------------------------------")
	maxTotal := len(st.Subjects) * 100
	fmt.Printf("Total Marks : %d / %d\n", st.TotalMarks, maxTotal)
	fmt.Printf("Percentage  : %.2f%%\n", st.Percentage)
	fmt.Printf("Result      : %s\n", passFailText(st.Passed))
	fmt.Printf("Grade       : %s\n", st.Grade)
	fmt.Println("========================================")
}

func passFailText(p bool) string {
	if p {
		return "PASS"
	}
	return "FAIL"
}
