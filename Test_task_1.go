package main

import( "fmt"
        "encoding/json"
        "io/ioutil"
	      "log"
	      "net/http"
	      "time"
        "strconv"
)

func main(){
  Year := time.Now().Year()
  t := time.Now().YearDay()
  const DateFormat = "2006-01-02"
  var txt string

  hol := ListOfHoliday(Year)
  for i, val := range hol{
    h, _ := time.Parse(DateFormat, val.Date)
    _, m, d := h.Date()
    if h.YearDay() < t && i + 1 == len(hol) {
      NewYear := ListOfHoliday(Year+1)
      FirstHol, _ := time.Parse(DateFormat, NewYear[0].Date)
      txt += fmt.Sprint("\nThe next holiday is ", val.Name, "\nIt will be at the next year: ")
      txt += fmt.Sprint(FirstHol.Date())
      txt += fmt.Sprint(" (", FirstHol.Weekday(),")\n")
      txt += fmt.Sprint(CheckWeekend(FirstHol))
     }
    if h.YearDay() == t{
      txt += fmt.Sprint("\nToday is a holiday ", val.Name, "\n")
      txt += fmt.Sprint(CheckWeekend(h))
      break
    }
    if h.YearDay() > t {
      txt += fmt.Sprint("\nThe next holiday is ", val.Name, "\nIt will be on ", m, d, " (", h.Weekday(),")\n")
      txt += fmt.Sprint(CheckWeekend(h))
      break
    }
  }
  fmt.Println(txt)
}

type Holiday struct{
  Date string `json:"date"`
  Name string `json:"name"`
}

func CheckWeekend(t time.Time)string{
  _, m, d := t.Date()
  txt := "The weekend will last "
  switch t.Weekday() {
  case time.Weekday(5):txt += "3 days: " + fmt.Sprint(m,d," - ",m,d+2)
  case time.Weekday(6):txt += "3 days: " + fmt.Sprint(m,d," - ",m,d+2)
  case time.Weekday(0):txt += "3 days: " + fmt.Sprint(m,d-1," - ",m,d+1)
  case time.Weekday(1):txt += "3 days: " + fmt.Sprint(m,d-2," - ",m,d)
  default: txt += "1 day"
  }
  return txt
}

func ListOfHoliday(Year int) []Holiday{
  json_doc := DownloadPage("https://date.nager.at/api/v2/publicholidays/"+ strconv.Itoa(Year) +"/UA")
  var hol []Holiday
  json.Unmarshal(json_doc, &hol)
  return hol
}

func DownloadPage(url string) []uint8{
	res, err := http.Get(url)
	if err != nil { log.Fatal(err) }

	body, readErr := ioutil.ReadAll(res.Body)
  res.Body.Close()
	if readErr != nil { log.Fatal(readErr) }

  return body
}
