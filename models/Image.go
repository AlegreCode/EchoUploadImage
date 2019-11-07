package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

type Image struct {
	Id   uuid.UUID `json:"id,omitempty" form:"id"`
	Name string    `json:"name,omitempty" form:"name"`
	Url  string    `json:"url,omitempty" form:"omitemptye"`
}

type Images struct {
	Images []Image `json:"images,omitempty"`
}

func (i *Image) Save() {
	path_data := path.Join("data", "images.json")
	data, _ := ioutil.ReadFile(path_data)
	datos := Images{}
	_ = json.Unmarshal([]byte(data), &datos)

	datos.Images = append(datos.Images, *i)

	// Destination
	_, err1 := os.Stat("data")

	if os.IsNotExist(err1) {
		errDir := os.MkdirAll("data", 0755)
		if errDir != nil {
			log.Fatal(err1)
		}

	}
	json, _ := json.MarshalIndent(datos, "", " ")
	_ = ioutil.WriteFile(path_data, json, 0644)
}

func (i *Image) Read() Images {
	path_data := path.Join("data", "images.json")
	data, _ := ioutil.ReadFile(path_data)
	datos := Images{}
	_ = json.Unmarshal([]byte(data), &datos)
	return datos
}

func (i *Image) Delete() {
	path_data := path.Join("data", "images.json")
	data, _ := ioutil.ReadFile(path_data)
	images := Images{}
	_ = json.Unmarshal([]byte(data), &images)
	for index, val := range images.Images {
		if i.Id == val.Id {
			// Remove data from slice and save in file json
			images.Images = append(images.Images[:index], images.Images[index+1:]...)
			json, _ := json.MarshalIndent(images, "", " ")
			_ = ioutil.WriteFile(path_data, json, 0644)
			// Remove file img from folder
			path_file := filepath.Join("assets", "uploads", val.Url)
			_ = os.Remove(path_file)
		}
	}

}
