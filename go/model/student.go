package model

type Student struct {
	Base
	Name     string `json:"name"`
	Age      int    `json:"age"`
	RollNo   int    `json:"rollNo"`
	Date     string `gorm:"type:date" json:"date"`
	DateTime string `gorm:"type:datetime" json:"dateTime"`
	Email    string `json:"email"`
	IsMale   *bool  `json:"isMale"`
}

// type DateTimestamp struct {
// 	time.Time
// }

// func (sd *DateTimestamp) UnmarshalJSON(input []byte) error {

// 	fmt.Println(string([]byte(input)))

// 	strInput := string(input)
// 	strInput = strings.Trim(strInput, `"`)
// 	newTime, err := time.Parse("2006-01-02 15:04:05", strInput)
// 	if err != nil {
// 		return err
// 	}

// 	sd.Time = newTime
// 	return nil
// }

// func (sd Student) MarshalJSON() ([]byte, error) {
// 	return []byte(sd.String()), nil
// }
