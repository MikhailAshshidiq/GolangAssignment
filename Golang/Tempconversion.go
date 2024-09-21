package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func celsiusToFahrenheit(celsius float64) float64 {
        return (celsius * 9/5) + 32
}

func celsiusToKelvin(celsius float64) float64 
{
        return celsius + 273.15
}

func fahrenheitToCelsius(fahrenheit float64) float64 {return (fahrenheit - 32) * 5/9
}

func fahrenheitToKelvin(fahrenheit float64) float64 {
     return (fahrenheit - 32) * 5/9 + 273.15
}

func kelvinToCelsius(kelvin float64)float64 {
      return kelvin - 273.15
}

func kelvinToFahrenheit(kelvin float64) float64 {
    return (kelvin - 273.15) * 9/5 + 32
}

func main() {reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter the temperature: ")
        tempStr, _ := reader.ReadString('\n')
        tempStr = strings.TrimSpace(tempStr)

        fmt.Print("Enter the unit (C, F, or K): ")
        unitStr, _ := reader.ReadString('\n')
        unitStr = strings.TrimSpace(unitStr)

        temp, err := strconv.ParseFloat(tempStr, 64)
        if err != nil {
                fmt.Println("Invalid temperature input")
                return
        }

        switch unitStr {
        case "C":
                fmt.Printf("%.2f C = %.2f F\n", temp, celsiusToFahrenheit(temp))
                fmt.Printf("%.2f C = %.2f K\n", temp, celsiusToKelvin(temp))
        case "F":
                fmt.Printf("%.2f F = %.2f C\n", temp, fahrenheitToCelsius(temp))
                fmt.Printf("%.2f F = %.2f K\n", temp, fahrenheitToKelvin(temp))
        case "K":
                fmt.Printf("%.2f K = %.2f C\n", temp, kelvinToCelsius(temp))
                fmt.Printf("%.2f K = %.2f F\n", temp, kelvinToFahrenheit(temp))
        default:
                fmt.Println("Invalid unit input")
        }
}