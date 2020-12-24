package model

import "github.com/shaileshhb/restapi/model"

type Student struct {
	model.Base  `structs:"id,omitnested"`
	Name        string  `json:"name" struct:"name"`
	Age         *int    `json:"age" struct:"age"`
	RollNo      *int    `json:"rollNo" struct:"rollNo"`
	PhoneNumber *string `json:"phoneno" struct:"phoneno"`
	Email       string  `json:"email" struct:"email"`
	IsMale      *bool   `json:"isMale" struct:"isMale"`
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
