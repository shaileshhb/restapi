package model

type Student struct {
	Base        `structs:"id,omitnested"`
	Name        string  `gorm:"type:varchar(30)" json:"name"`
	Age         *int    `gorm:"type:int" json:"age"`
	RollNo      *int    `gorm:"type:int" json:"rollNo"`
	PhoneNumber *string `gorm:"type:varchar(10)" json:"phone"`
	Email       string  `gorm:"type:varchar(50)" json:"email"`
	IsMale      *bool   `json:"isMale"`
	Date        *string `gorm:"type:date" json:"date"`
	DateTime    *string `gorm:"type:datetime" json:"dateTime"`
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
