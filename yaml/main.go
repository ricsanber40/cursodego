package main

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


type T struct {
	Db struct {
		Src struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
			Host string `yaml:"host"`
		}
		Dst struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		}
	}

}

func main() {
	t := T{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	fmt.Println("Testando propriedade do objeto: " +t.Db.Dst.Name)

	/*d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))*/

	/*m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))*/


}
