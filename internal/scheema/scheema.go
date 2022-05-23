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
}

type ScheemaFields struct {
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
	Address string `json:"address"`
}

//TODO implement array of scheemas
func ShowScheema(scheemaName string, value string) {

	// open json file
	// TODO implement cache
	jsonFile, err := os.Open("./jsons/" + scheemaName + ".json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var scheema Scheema
	json.Unmarshal(byteValue, &scheema)

	// printing title
	fmt.Printf("\n\n[%s]", scheema.Name)

	// printing fields
	for _, field := range scheema.Fields {

		fieldValue := gjson.Get(value, field.Address)
		if json.Valid([]byte(fieldValue.String())) {
			if len(field.Scheema) > 0 {
				// print scheema
				ShowScheema(field.Scheema, fieldValue.String())
			} else {
				// unknown scheema
				fmt.Printf("\n%s: %s", field.Name, fieldValue.String())
			}
		} else {
			// print basic value
			fmt.Printf("\n%s: %s", field.Name, fieldValue.String())
		}

	}
	fmt.Printf("\n")

}
