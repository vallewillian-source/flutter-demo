package scheema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

type Scheema struct {
	Name       string          `json:"name"`
	PrimaryKey string          `json:"primary_key"`
	Fields     []ScheemaFields `json:"fields"`
	Result     map[string]interface{}
}

type ScheemaFields struct {
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
	Address string `json:"address"`
}

//TODO implement array of scheemas
func GenerateScheema(serviceName string, scheemaName string, value string, cache *map[string]Scheema) (Scheema, error) {

	// trying to get scheema from cache
	var scheema Scheema
	scheema, ok := (*cache)[scheemaName]
	if !ok {
		// not found on cache. getting from file..
		jsonFile, err := os.Open("./jsons/services/" + serviceName + "/scheemas/" + scheemaName + ".json")
		if err != nil {
			fmt.Println(err)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		defer jsonFile.Close()

		// convert to struct
		json.Unmarshal(byteValue, &scheema)

		// saving to cache
		(*cache)[scheemaName] = scheema

	}

	// initializing scheema result
	scheema.Result = make(map[string]interface{})

	// getting fields
	for _, field := range scheema.Fields {

		fieldValue := gjson.Get(value, field.Address)
		if json.Valid([]byte(fieldValue.String())) {
			if len(field.Scheema) > 0 {
				// insert scheema
				var err error
				scheema.Result[field.Name], err = GenerateScheema(serviceName, field.Scheema, fieldValue.String(), cache)
				if err != nil {
					scheema.Result[field.Name] = "Err"
				}
			} else {
				// unknown scheema
				scheema.Result[field.Name] = fieldValue.String()
			}
		} else {
			// insert basic value
			scheema.Result[field.Name] = fieldValue.String()
		}

	}

	return scheema, nil

}
