package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type file struct {
	filePath string
}

func CreateFile(filePath string) *file {
	return &file{
		filePath: filePath,
	}
}

func (f file) ReadJson(obj interface{}) error {
	if content, err := ioutil.ReadFile(f.filePath); err != nil {
		log.Println("on ReadJson an error occured: ", err)
		return err
	} else {
		if err := json.Unmarshal(content, &obj); err != nil {
			log.Println("on ReadJson an error occured: ", err)
			return err
		}
	}
	return nil
}
