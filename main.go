//
// main.go
// Copyright (C) 2019 reidmiller <reidmiller@Arrakis.local>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	services := parseYaml("services.yml")

	//fmt.Println(services)
	for i, s := range services {
		fmt.Println(s, "at", i)
		switch i {
		case "web":
			fmt.Println("Create Web Service")
			createWebNode()
		default:
			fmt.Println("Unknown Service type to deploy")
		}
	}
	//fmt.Println("=========")
	datastores := parseYaml("datastores.yml")
	//fmt.Println(datastores)
	for i, d := range datastores {
		fmt.Println(d, "at", i)
		switch i {
		case "redis":
			fmt.Println("Create Redis Service")
			createRedisService()
		default:
			fmt.Println("Unknown data store type")
		}
	}
	//fmt.Println("=========")
	//connections := parseYaml("connections.yml")
	//fmt.Println(connections)

	//fmt.Println("=========")
}

func parseYaml(fileName string) map[string]string {
	p := make(map[string]string)
	y, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to read input file: \n%v\n", err)
	}
	err2 := yaml.Unmarshal([]byte(y), &p)
	if err2 != nil {
		log.Fatalf("Error 2: Unable to Unmarshal YAML file file:\n%v\n", err2)
	}
	return p
}
