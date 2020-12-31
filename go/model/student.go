package model

// Student represents the student for this application
//
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger: model studentModel
type Student struct {
	Base `structs:"id,omitnested"`

	// the name for this student
	//
	// required: true
	Name string `gorm:"type:varchar(30)" json:"name"`

	// the age for this student
	//
	// required: false
	Age *int `gorm:"type:int" json:"age"`

	// the rollno for this student
	//
	// required: false
	RollNo *int `gorm:"type:int" json:"rollNo"`

	// the phone number for this student
	//
	// required: false
	PhoneNumber *string `gorm:"type:varchar(10)" json:"phone"`

	// the email for this student
	//
	// required: true
	// example: user@provider.net
	Email string `gorm:"type:varchar(50)" json:"email"`

	// the ismale for this student
	//
	// required: false
	IsMale *bool `gorm:"type:tinyint" json:"isMale"`

	// the date for this student
	//
	// required: false
	Date *string `gorm:"type:date" json:"date"`

	// the datetime for this student
	//
	// required: false
	// DateTime *string `gorm:"type:datetime" json:"dateTime"`

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
