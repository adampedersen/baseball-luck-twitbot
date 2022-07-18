package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strings"
)

func main() {
	m, err := CSVFileToMap("fangraphs-url-map.csv")
	if(err != nil){
		fmt.Println("error")
	}
	fmt.Println(m["cody bellinger"])
}

//return map: key = player name, value = fangraphs link
func CSVFileToMap(filePath string) (returnMap map[string]string, err error) {
	returnMap = make(map[string]string)
	// read csv file
	file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		s := scanner.Text()
		i := strings.Index(s, ",")
		k := s[:i]
		v := s[(i+1):]
		returnMap[k] = v
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return returnMap,nil
}
	
