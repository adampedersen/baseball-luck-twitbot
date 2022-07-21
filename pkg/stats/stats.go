package stats

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"strconv"
	"math"
)

func GetLuckRating(playerName string) string{
	m, err := CSVFileToMap("fangraphs-id-map.csv") // m will hold player ids
	if(err != nil){
		fmt.Println("error with CSVFileToMap")
	}
	id := m[strings.ToLower(playerName)]
	wOBAstr := FindStat(id,14)
	xwOBAstr := FindStat(id, 15)
	const bitSize = 64 
	wOBA, err := strconv.ParseFloat(wOBAstr, bitSize)
	xwOBA, err := strconv.ParseFloat(xwOBAstr, bitSize)
	if(err != nil){
		fmt.Println("error parsing floats")
		return ""
	}
	diff := wOBA - xwOBA
	ret := playerName
	if(diff > 0){
		ret += " has been lucky this year. His wOBA is greater than his xWOBA by: "

	}else{
		ret += " has been unlucky this year. His wOBA is less than his xWOBA by: "
	}
	longDiff := fmt.Sprintf("%f", math.Abs(diff))
	longDiff = longDiff[:5]
	ret += longDiff
	if (diff > 0){
		ret += " (" + wOBAstr +" - "
		ret +=  xwOBAstr +")"
	}else{
		ret += " (" + xwOBAstr +" - "
		ret +=  wOBAstr +")"
	}
	return ret
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

/*
stat number:
0 1 2  3  4  5  6  7   8  9    10   11  12  13   14   15   16   17  18  19  20
# G PA HR R RBI SB BB% K% ISO BABIP AVG OBP SLG wOBA xwOBA wRC+ BsR Off Def WAR
*/
//param: player's fangraphs id, statNumber (above)
func FindStat(id string, statNumber int) string {
	url := "https://www.fangraphs.com/leaders.aspx?pos=all&stats=bat&lg=all&qual=0&type=8&season=2022&month=0&season1=2022&ind=0&team=0&rost=0&age=0&filter=&players=" + id + "&startdate=&enddate="
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respBodyStr := string(html)
	delim := `align="right">`
	var stats []string
	for i := 0;i < 21; i++{
		i := strings.Index(respBodyStr, delim)
		j := i + len(delim)
		respBodyStr = respBodyStr[j:]
		j = strings.Index(respBodyStr, "<")
		stats = append(stats,respBodyStr[:j])
	}

	return stats[statNumber]
}
	
