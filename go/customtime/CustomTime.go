package customtime

import (
	"fmt"
	"strings"
	"time"
)

// type CustomTime struct {
// 	time.Time `gorm:"type:datetime" json:"customTime" time:"2006-01-02"`
// }

type CustomTime time.Time

// const dateFormat = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(data []byte) error {

	// field, ok := reflect.TypeOf(ct).Elem().FieldByName("customtime.CustomTime")
	// if !ok {
	// 	return errors.NewValidationError("tag not found!!")
	// }

	// fmt.Println(getStructTag(field))

	// dateFormat := getStructTag(field)

	strDate := strings.Replace(string(data), "\"", "", -1)

	// fmt.Println(" string(data) -> ", string(data))
	// fmt.Println(" strDate -> ", strDate)
	// fmt.Println(" dateFormat -> ", dateFormat)

	t, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		fmt.Println(" error in time.Parse")
		return err
	}

	*ct = CustomTime(t)
	// ct.Time = t

	return nil
}

// func getStructTag(f reflect.StructField) string {
// 	fmt.Println("f ->", f.Tag.Get("time"))
// 	return string(f.Tag.Get("time"))
// }
