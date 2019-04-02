package football

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetTrainingsFromJSON() ([]Training, error) {
	file, err := ioutil.ReadFile("../training.json")
	if err != nil {
		fmt.Println(err)
	}

	data := Trainings{}
	_ = json.Unmarshal([]byte(file), &data)

	return data.Trainings, nil
}

func GetEquipmentsFromJSON() ([]Equipment, error) {
	file, err := ioutil.ReadFile("../equipment.json")
	if err != nil {
		fmt.Println(err)
	}

	data := Equipments{}
	_ = json.Unmarshal([]byte(file), &data)

	// for i := 0; i < len(data.Equipments); i++ {
	// 	fmt.Println("Equipment", data.Equipments[i])
	// }

	return data.Equipments, nil
}

// func ReadJsonFile(s string) map[string]interface{} {

// 	jsonFile, err := os.Open("../" + s)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully Opened json file")
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	var result map[string]interface{}
// 	err2 := json.Unmarshal([]byte(byteValue), &result)
// 	if err2 != nil {
// 		panic(err2)
// 	}

// 	// fmt.Println(result)

// 	return result
// }
