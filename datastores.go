package main

//
// datastores.go
// Copyright (C) 2019 reidmiller <reidmiller@Arrakis.local>
//
// Distributed under terms of the MIT license.
//
import (
	"fmt"
	"io/ioutil"

	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/cloudformation/resources"
)

func createRedisService() {
	template := cloudformation.NewTemplate()

	template.Resources["redisNode"] = &resources.AWSElastiCacheCacheCluster{
		Engine:        "redis",
		CacheNodeType: "cache.t2.micro",
		ClusterName:   "testing-cluster",
		//		EngineVersion: "",
		NumCacheNodes: 2,
	}

	y, err := template.YAML()
	if err != nil {
		fmt.Println("Unable to render datasource template: %s", err)
	}

	ioutil.WriteFile("output-datasources.yaml", []byte(string(y)), 0644)
}
