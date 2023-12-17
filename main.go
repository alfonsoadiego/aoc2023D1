package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Advent of Code day 1 (I now, llego algo tarde)")

  file, error := os.Open("input.txt")
  if error != nil {
    fmt.Println("Error reading the file")
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var result int = 0

  for scanner.Scan(){
    line := scanner.Text()
    //fmt.Printf("%c%c\n",getFirstdigit(line),getLastDigit(line))

    intValue, err := strconv.Atoi(string(getFirstdigit(line))+string(getLastDigit(line)))

    if err != nil {
      fmt.Println("Error parsing")
    }else{
      fmt.Printf("%d\t",intValue)
    }

    result+=intValue
  }

  fmt.Printf("\nCalibration vals:%d", result)

}

func getFirstdigit(cadena string) rune {
	for _, caracter := range cadena {
		if unicode.IsDigit(caracter) {
			return caracter
		}
	}

	return '0'
}
func getLastDigit(cadena string) rune {
  return getFirstdigit(reverse(cadena))
}

func reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}
