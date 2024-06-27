package student

import (
	"errors"
	"time"
)

/* 
En Go, los métodos que modifican los campos de una estructura generalmente se definen utilizando punteros a la estructura. 
Hay algunas razones clave para esto:

Eficiencia de la memoria: Al pasar la estructura por referencia (utilizando un puntero), se evita la copia completa de la estructura, 
lo que puede ser costoso en términos de memoria y rendimiento, especialmente para estructuras más grandes.
Mutabilidad: Las estructuras en Go se pasan por valor de forma predeterminada, lo que significa que una copia de la estructura se pasa a la función o método. 
Si la función o método modifica la estructura pasada, solo se modificará la copia y no el original. 
Sin embargo, al utilizar un puntero a la estructura, la función o método puede modificar directamente el original.

Consistencia: Utilizar punteros en métodos que modifican la estructura es coherente con otros métodos de 
la misma estructura que también pueden necesitar acceder o modificar sus campos.

Conveniencia: Al utilizar un puntero, se evita tener que devolver explícitamente la estructura 
modificada en el método, ya que la modificación se reflejará en el original.
*/

type Student struct {
	lastName string
	firstName string
	city string
	date time.Time
	hasDiploma bool
	courseCode string
}

type StudentInt interface {
	LastName() string
	FirstName() string
	City() string
	Date() time.Time
	SchoolDiploma() bool
	CourseCode() string
}


var ErrorInvalidMonth = errors.New("Invalid month number")
var ErrorInvalidCourse = errors.New("Invalid course. Expected 'APU', 'LI' or 'LS'")

var monthMap = map[int]time.Month{
	1:  time.January,
	2:  time.February,
	3:  time.March,
	4:  time.April,
	5:  time.May,
	6:  time.June,
	7:  time.July,
	8:  time.August,
	9:  time.September,
	10: time.October,
	11: time.November,
	12: time.December,
}

func NewStudent(firstName, lastName, city string, year, month, day int, hasDiploma bool, code string) *Student {
	student := &Student{}
	student.setFirstName(firstName)
	student.setLastName(lastName)
	student.setCity(city)
	student.setSchoolDiploma(hasDiploma)
	student.setCourseCode(code)
	student.setDateOfBirth(year, month, day)
	return student
}

//SETTERS
func (s *Student) setLastName(lastName string) {
	s.lastName = lastName
}

func (s *Student) setFirstName(firstName string) {
	s.firstName = firstName
}

func (s *Student) setCity(city string) {
	s.city = city
}

func (s *Student) setDateOfBirth(year, month, day int) error {
	if month < 1 || month > 12 {
		return ErrorInvalidMonth
	}
	s.date = time.Date(year, monthMap[month], day, 0, 0, 0, 0, time.UTC)
	return nil
}

func (s *Student) setSchoolDiploma(hasDiploma bool) {
	s.hasDiploma = hasDiploma
}


func (s *Student) setCourseCode(code string) error {
	if code != "APU" && code != "LI" && code != "LS" {
		return ErrorInvalidCourse
	}

	s.courseCode = code
	return nil	
}

//GETTERS
func (s *Student) LastName() string {
	return s.lastName
}

func (s *Student) FirstName() string {
	return s.firstName
}

func (s *Student) City() string {
	return s.city
}

func (s *Student) Date() time.Time {
	return s.date
}

func (s *Student) HasDiploma() bool {
	return s.hasDiploma
}

func (s *Student) CourseCode() string {
	return s.courseCode
}

//String
func (s *Student) ToString() string {
	return s.FirstName() + " " + s.LastName()
}