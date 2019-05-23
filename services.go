package main

import (
	"fmt"
	"io/ioutil"

	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/cloudformation/resources"
)

//
// services.go
// Copyright (C) 2019 reidmiller <reidmiller@Arrakis.local>
//
// Distributed under terms of the MIT license.
//

func createWebNode() {
	template := cloudformation.NewTemplate()

	template.Resources["webNode"] = &resources.AWSEC2Instance{
		ImageId:            "ami-03762dc82325bb34e",
		IamInstanceProfile: "web_autosign_instance_profile",
		InstanceType:       "t2.micro",
		KeyName:            "dumpling",
		SubnetId:           "subnet-0fec52d859d2da287",
		//		UserData:           "test data",
	}

	// and also the YAML AWS CloudFormation template
	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}

	ioutil.WriteFile("output-services.yaml", []byte(string(y)), 0644)
}
