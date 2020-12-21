package utility

import (
	"github.com/shaileshhb/restapi/model"
)

func ConvertDateTime(students *[]model.Student) {

	for i := 0; i < len(*students); i++ {
		if (*students)[i].Date == "" {
			continue
		}
		(*students)[i].Date = (*students)[i].Date[:10]
	}

}

func TrimDateTime(students *[]model.Student) {

	for i := 0; i < len(*students); i++ {
		if (*students)[i].DateTime == "" {
			continue
		}
		(*students)[i].DateTime = (*students)[i].DateTime[11:19] + " " + (*students)[i].DateTime[11:19]
	}
}
