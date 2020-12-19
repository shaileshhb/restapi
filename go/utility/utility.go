package utility

import "github.com/shaileshhb/restapi/model"

func ConvertDateTime(students *[]model.Student) {

	for i := 0; i < len(*students); i++ {
		(*students)[i].Date = (*students)[i].Date[:10]
	}

}
