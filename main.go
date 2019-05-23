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

type ProcoreSpec struct {
	Description    string             `yaml:description`
	Infrastructure InfrastructureSpec `yaml:infrastructure`
	//Monitoring MonitoringSpec `yaml:monitoring`
	//Scaling ScalingSpec `yaml:scaling`
	//Deployment DeploymentSpec `yaml:deployment`
}

type InfrastructureSpec struct {
	Services    []ServiceSpec    `yaml:services`
	Datastores  []DatastoreSpec  `yaml:datastores`
	Connections []ConnectionSpec `yaml:connections`
}

type ServiceSpec struct {
	Services []string `yaml:services`
}

type DatastoreSpec struct {
	Datastores []string `yaml:datastores`
}

type Connection struct {
	From string `yaml:from`
	To   string `yaml:to`
	Port int    `yaml:port`
}

type ConnectionSpec struct {
	Connections []Connection `yaml:connections`
}

//struct MonitoringSpec {
//	Monitors string `yaml:monitoring`
//}

//struct ScalingSpec {
//	Web
//}
func main() {
	p := make(map[interface{}]interface{})
	y, err := ioutil.ReadFile("p.yml")
	if err != nil {
		log.Fatalf("Unable to read input file: \n%v\n", err)
	}
	err2 := yaml.Unmarshal([]byte(y), &p)
	if err2 != nil {
		log.Fatalf("Error 2: Unable to Unmarshal YAML file file:\n%v\n", err2)
	}
	for key, value := range p {
		fmt.Println("Key:", key, "Value:", value)
	}

	infra_value := p["infrastructure"]
	services := infra_value.(map[interface{}]interface{})["services"]
	fmt.Println(services)
	fmt.Println("=========")
	datastores := infra_value.(map[interface{}]interface{})["datastores"]
	fmt.Println(datastores)
	fmt.Println("=========")
	connections := infra_value.(map[interface{}]interface{})["connections"]
	fmt.Println(connections)
	fmt.Println("=========")
}

/*
func parseMap(aMap map[interface{}]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[interface{}]interface{}:
			parseMap(val.(map[interface{}]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[interface{}]interface{}:
			parseMap(val.(map[interface{}]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
			fmt.Println(i, ":", concreteVal)
		}
	}
}
*/
