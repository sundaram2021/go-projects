package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	// "google.golang.org/genproto/protobuf/field_mask"
)

type Person struct {
	Name string `json:name`
	Age  uint   `json:age`
}

func CreateData(a []Person) {
	file, err := os.Create("person.json")

	if err != nil {
		fmt.Println("Error in creating file: ", err)
		return
	}

	defer file.Close()

	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println("error in marshal : ", err)
		return
	}
	_, err = file.Write(data)

	if err != nil {
		fmt.Println("error in writing data into the file: ", err)
		return
	}

	fmt.Println("file is created successfully")
}

func ReadData() {
	file, err := os.Open("person.json")

	if err != nil {
		fmt.Println("error in opening file: ", err)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error in reading file: ", err)
		return
	}

	fmt.Println("file data: ",string(data))
}


func UpdateData(a []Person) {
	file, err := os.OpenFile("person.json",  os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("error in opening file: ", err)
		return
	}

	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println("error in marshal : ", data)
	}

	_, err = file.Write([]byte(data))


	if err != nil {
		fmt.Println("error in writing file: ", err)
		return
	}

	defer file.Close()
	fmt.Println("file is updated sucessfully")

}

func DeletData() {
	err := os.Remove("person.json")

	if err != nil {
		fmt.Println("Error in deleting filee: ", err)
		return
	}
	fmt.Println("file is deleted sucessfully")
}

func main() {
	a := []Person{
		{Name: "Sundaram", Age: 22},
		{Name: "Aman", Age: 22},
	}

	b := []Person {
		{Name: "Sundaram", Age: 22},
		{Name: "Dheeraj", Age: 24},
	}


	CreateData(a)
	ReadData()
	UpdateData(b)
	ReadData()
	DeletData()
}
