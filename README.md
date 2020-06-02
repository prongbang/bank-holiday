# Bank Holiday

### Install

```
go get github.com/prongbang/bank-holiday
```

### How to use

```go
	package main

import (
	"encoding/json"
	"log"

	"github.com/prongbang/bank-holiday/pkg/holiday"
)

func main() {
    utility := holiday.NewUtility()
    
	year := "2563"
	holidayData, _ := json.Marshal(utility.GetFinancialHoliday(year))
	fmt.Println(string(holidayData))
}
```

### Results

```json
{
   "april":{
      "6":"วันพระบาทสมเด็จพระพุทธยอดฟ้าจุฬาโลกมหาราช และวันที่ระลึกมหาจักรีบรมราชวงศ์"
   },
   "august":{
      "12":"วันเฉลิมพระชนมพรรษาสมเด็จพระนางเจ้าสิริกิติ์ พระบรมราชินีนาถ พระบรมราชชนนีพันปีหลวง และวันแม่แห่งชาติ"
   },
   "december":{
      "10":"วันพระราชทานรัฐธรรมนูญ",
      "31":"วันสิ้นปี",
      "7":"ชดเชยวันคล้ายวันพระบรมราชสมภพของพระบาทสมเด็จพระบรมชนกาธิเบศร มหาภูมิพลอดุลยเดชมหาราช บรมนาถบพิตร วันชาติ และวันพ่อแห่งชาติ (วันเสาร์ที่ 5 ธันวาคม 2563)"
   },
   "february":{
      "10":"ชดเชยวันมาฆบูชา"
   },
   "january":{
      "1":"วันขึ้นปีใหม่"
   },
   "july":{
      "28":"วันเฉลิมพระชนมพรรษาพระบาทสมเด็จพระเจ้าอยู่หัว",
      "6":"ชดเชยวันอาสาฬหบูชา (วันอาทิตย์ที่ 5 กรกฎาคม 2563)"
   },
   "june":{
      "3":"วันเฉลิมพระชนมพรรษาสมเด็จพระนางเจ้าสุทิดา พัชรสุธาพิมลลักษณ พระบรมราชินี"
   },
   "march":{

   },
   "may":{
      "1":"วันแรงงานแห่งชาติ",
      "4":"วันฉัตรมงคล",
      "6":"วันวิสาขบูชา"
   },
   "november":{

   },
   "october":{
      "13":"วันคล้ายวันสวรรคตพระบาทสมเด็จพระบรมชนกาธิเบศรมหาภูมิพลอดุลยเดชมหาราช บรมนาถบพิตร",
      "23":"วันปิยมหาราช"
   },
   "september":{

   }
}
```

### DataSource

- https://www.bot.or.th/Thai/FinancialInstitutions/FIholiday/Pages/HolidayCalendar.aspx