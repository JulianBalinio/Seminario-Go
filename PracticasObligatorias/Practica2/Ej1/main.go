package main

import (
	"fmt"

	"Ej1/student"
	"Ej1/list"
)

func main() {

	studentList := list.New()

	list.PushBack(&studentList, student.NewStudent("Julian", "Baliño", "Bariloche", 1997, 2, 23, true, "APU"))
	list.PushBack(&studentList, student.NewStudent("Geddy", "Lee", "Bariloche", 1996, 5, 30, true, "LS"))
	list.PushBack(&studentList, student.NewStudent("Neil", "Peart", "La Plata", 1996, 6, 30, false, "APU"))
	list.PushBack(&studentList, student.NewStudent("Alex", "Lifeson", "Bariloche", 1996, 9, 30, true, "APU"))
	list.PushBack(&studentList, student.NewStudent("Tom", "Sawyer", "Bariloche", 1998, 12, 21, true, "APU"))

	birthMap := make(map[int]int)
	careerMap := make(map[string]int)

	careerMap["APU"] = 0
	careerMap["LI"] = 0
	careerMap["LS"] = 0

	list.Iterate(studentList, func(std *student.Student) {
		if std.City() == "Bariloche" {
			fmt.Printf("Ingresante %s nacido en Bariloche.\n", std.ToString())
		}

		if !std.HasDiploma() {
			_, err := list.Remove(&studentList)
			if err != nil {
				fmt.Println("Error al eliminar el estudiante:", err)
			}
			return //salgo de la iteracion
		}
		
		//actualizo maps
		birthMap[std.Date().Year()]++
		careerMap[std.CourseCode()]++

	})

	maxYear, maxNac := maxYear(birthMap)
	maxCod, maxStudents := maxCourse(careerMap)

	fmt.Printf("\nSe registraron %d alumnos en la carrera %s. Siendo la carrera con mas inscriptos.\n", maxStudents, maxCod)
	fmt.Printf("\nNacieron %d en el año %d. Siendo el año con mas nacimientos registrados.\n", maxNac, maxYear)
}

func maxYear(birthMap map[int]int) (int, int) {
	maxNac := 0
	max := 0
	for año, nacimientos := range birthMap {
		if nacimientos > maxNac {
			maxNac = nacimientos
			max = año
		}
	}
	return max, maxNac
}

func maxCourse(careerMap map[string]int) (string, int) {
	maxStudents := -1
	maxCod := ""
	for cod, students := range careerMap {
		if students > maxStudents {
			maxStudents = students
			maxCod = cod
		}
	}
	return maxCod, maxStudents
}
