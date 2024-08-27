package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Student struct {
	name      string
	rank      string
	math      float32
	physical  float32
	chemistry float32
	avg       float32
	age       uint8
}

type ManagerStudent struct {
	students []Student
}

func (mS *ManagerStudent) findStudentByName(name string) (Student, error) {
	for _, student := range mS.students {
		if student.name == name {
			return student, nil
		}
	}

	return Student{}, errors.New("could not find any student")
}

func (mS *ManagerStudent) getGoodStudents() []Student {
	result := make([]Student, 0)

	for _, student := range mS.students {
		if student.rank == "GIOI" {
			result = append(result, student)
		}
	}

	return result
}

func (mS *ManagerStudent) sort() {
	sort.Slice(mS.students, func(i, j int) bool {
		if mS.students[i].name != mS.students[j].name {
			return mS.students[i].name < mS.students[j].name
		} else {
			return mS.students[i].age < mS.students[j].age
		}
	})
}

func main() {
	mS := new(ManagerStudent)

	file, err := os.Open("./students.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records")
	}

	for _, eachRecord := range records {
		parsedAge, err := strconv.ParseUint(eachRecord[1], 0, 8)
		if err != nil {
			log.Fatal("Error reading record [1]")
		}
		parsedMath, err := strconv.ParseFloat(eachRecord[2], 32)
		if err != nil {
			log.Print(err)
			log.Fatal("Error reading record [2]")
		}
		parsedPhysical, err := strconv.ParseFloat(eachRecord[3], 32)
		if err != nil {
			log.Fatal("Error reading record [3]")
		}
		parsedChemistry, err := strconv.ParseFloat(eachRecord[4], 32)
		if err != nil {
			log.Fatal("Error reading record [4]")
		}

		student := Student{
			name:      eachRecord[0],
			age:       uint8(parsedAge),
			math:      float32(parsedMath),
			physical:  float32(parsedPhysical),
			chemistry: float32(parsedChemistry),
		}

		mS.students = append(mS.students, student)
	}

	for i, student := range mS.students {
		avg := (student.math + student.physical + student.chemistry) / 3
		var rank string

		if avg >= 8.0 {
			rank = "GIOI"
		} else if avg >= 6.5 {
			rank = "KHA"
		} else if avg >= 5.0 {
			rank = "TB"
		} else {
			rank = "YEU"
		}
		mS.students[i].avg = avg
		mS.students[i].rank = rank
	}

	var query string
	fmt.Print("Enter the name you want to find: ")
	fmt.Scan(&query)
	student, error := mS.findStudentByName(query)
	if error != nil {
		log.Fatal("could not found any student")
	}
	fmt.Println(student)

	fmt.Println("Good rank students")
	students := mS.getGoodStudents()

	for _, student := range students {
		fmt.Println(student)
	}

	mS.sort()
	fmt.Println("Sorted students")
	for _, student := range students {
		fmt.Println(student)
	}
}
