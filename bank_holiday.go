package bank_holiday

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	January   = "january"
	February  = "february"
	March     = "march"
	April     = "april"
	May       = "may"
	June      = "june"
	July      = "july"
	August    = "august"
	September = "september"
	October   = "october"
	November  = "november"
	December  = "december"
	url       = "https://www.bot.or.th/Thai/FinancialInstitutions/FIholiday/Pages/HolidayCalendar.aspx?y=%d"
	caleClass = ".cal_month"
	monthID   = "#ctl00_ctl73_g_2d5ecfc7_5e54_499d_8285_2fff313ce83f_ctl00_MonthLabel%d"
)

var Month = map[string]string{
	January:   "01",
	February:  "02",
	March:     "03",
	April:     "04",
	May:       "05",
	June:      "06",
	July:      "07",
	August:    "08",
	September: "09",
	October:   "10",
	November:  "11",
	December:  "12",
}

type Holidays map[string]map[string]string

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type Utility interface {
	GetFinancialHolidayList(year string) []Holiday
	GetFinancialHolidayMap(year string) Holidays
}

type utility struct {
}

// GetFinancialHolidayList implements Utility
func (u *utility) GetFinancialHolidayList(year string) []Holiday {
	hmap := u.GetFinancialHolidayMap(year)

	var date string
	hs := []Holiday{}
	for month, days := range hmap {
		for day, title := range days {
			m := Month[month]
			if len(day) > 1 {
				date = fmt.Sprintf("%s-%s-%s", year, m, day)
			} else {
				date = fmt.Sprintf("%s-%s-0%s", year, m, day)
			}
			hs = append(hs, Holiday{
				Name: title,
				Date: date,
			})
		}
	}

	return hs
}

// GetFinancialHolidayList implements Utility
func (u *utility) GetFinancialHolidayMap(year string) Holidays {
	y, _ := strconv.Atoi(year)
	holidayURL := fmt.Sprintf(url, y+543)
	fmt.Println(holidayURL)

	// Request the HTML page.
	res, err := http.Get(holidayURL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	calendarDom := doc.Find(caleClass)

	data := Holidays{}
	calendarDom.Find(".row").Each(func(i int, s *goquery.Selection) {
		s.Find(".col-xs-12 .col-sm-12 .col-md-4").Each(func(i int, col *goquery.Selection) {
			month := map[string]string{}
			title := col.Find(".cal_month_text").Text()
			col.Find("table > tbody > tr > td").Each(func(i int, td *goquery.Selection) {
				dayTitle := td.AttrOr("title", "")
				if dayTitle != "" {
					day := td.Text()
					month[day] = dayTitle
				}
			})
			data[monthMapping(title)] = month
		})
	})
	return data
}

func monthMapping(month string) string {
	switch month {
	case "มกราคม":
		return "january"
	case "กุมภาพันธ์":
		return "february"
	case "มีนาคม":
		return "march"
	case "เมษายน":
		return "april"
	case "พฤษภาคม":
		return "may"
	case "มิถุนายน":
		return "june"
	case "กรกฎาคม":
		return "july"
	case "สิงหาคม":
		return "august"
	case "กันยายน":
		return "september"
	case "ตุลาคม":
		return "october"
	case "พฤศจิกายน":
		return "november"
	case "ธันวาคม":
		return "december"
	default:
		return ""
	}
}

func New() Utility {
	return &utility{}
}
