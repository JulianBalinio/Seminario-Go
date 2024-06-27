package main

import "fmt"


const (
	HighIncorrectLimit = 50
	LowIncorrectLimit = 20
	HighRange = 37.5
	LowRange = 36
	Limit = 10
)

type Celsius float64
type Farenheit float64

func main() {
	temperatureCelsius := readTemperatures(Limit)
	temperatureFarenheit := toFahrenheit(temperatureCelsius)
	report := analyzeTemperatures(temperatureCelsius)
	printReport(report)

	average := Average(temperatureCelsius)
	fmt.Printf("Promedio entre maximo y minimo: %0.2f",  float64(average))

	for i:= 0; i<len(temperatureFarenheit); i++ {
		fmt.Println(temperatureFarenheit[i])
	}
}

//Lectura y almacenamiento de temperaturas, retorna un slice de float64
func readTemperatures(limit int) []Celsius {
	temperatures := make([]Celsius, limit)
	for i := 0; i < limit; i++ {
		var temp Celsius
		fmt.Print("Temp: ")
		fmt.Scanln(&temp)
		temperatures[i] = temp
	}
	return temperatures
}


//Leidas y almacenadas las temperaturas se procesan y se devuelve un mapeo con la cantidad de pacientes en cada rango
func analyzeTemperatures(temperatures []Celsius) map[string]Celsius {
	report := make(map[string]Celsius)
	report["incorrect"] = 0
	report["high"] = 0
	report["normal"] = 0
	report["low"] = 0
	
	for _, temp := range temperatures {
		if isHigh(temp) {
			report["high"]++
		} else if isNormal(temp) {
			report["normal"]++
		} else if isLow(temp) {
			report["low"]++
		} else {
			report["incorrect"]++
		}
	}

	total := Celsius(Limit)
	for key, quantity := range report {
		report[key] = (quantity / total) * 100
	}

	return report
}

//Se imprime el mapa creado
func printReport(report map[string]Celsius) {
	fmt.Printf("\nPorcentajes de pacientes\n")
	for key, quantity := range report {
		fmt.Printf("%s temp: %.2f%%\n", key, quantity)
	}
}

//Se busca el maximo y el minimo para calcular el promedio entre estos
func Average(temperatures []Celsius) Celsius{
	max, min := findMinMax(temperatures)
	return (max + min) / Celsius(2)
}

func findMinMax(temperatures []Celsius) (max Celsius, min Celsius) {
	max, min = temperatures[0], temperatures[0]
	for _, temp := range temperatures {
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}
	return max, min
}

//Conversor de Celsius a Farenheit
func toFahrenheit(c []Celsius) []Farenheit {
	var temperature []Farenheit
	for _, temp:= range c{
		tempF:= (Farenheit(temp)*9/5) + 32
		temperature = append(temperature, tempF)
	}
    return  temperature
}

//Se evalua la temperatura con los valores constantes
func isHigh(t Celsius) bool {
	return (t > HighRange && t <= HighIncorrectLimit)
}

func isNormal(t Celsius) bool {
	return (t <= HighRange && t >= LowRange)
}

func isLow(t Celsius) bool {
	return (t < LowRange && t >= LowIncorrectLimit)
}


