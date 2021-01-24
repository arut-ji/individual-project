package main

import (
	"fmt"
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v2"
	"reflect"
)

const file = `apiVersion: v1
kind: Pod
metadata:
  name: goproxy
  labels:
    app: goproxy
spec:
  containers:
  - name: goproxy
    image: k8s.gcr.io/goproxy:0.1
    ports:
    - containerPort: 8080
    livenessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 15
      periodSeconds: 20
    readinessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 1
      periodSeconds: 20
`

func getReadinessProbe(container interface{}) map[interface{}]interface{} {
	probe := container.(map[interface{}]interface{})["readinessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[interface{}]interface{})
}

func getLivenessProbe(container interface{}) map[interface{}]interface{} {
	probe := container.(map[interface{}]interface{})["livenessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[interface{}]interface{})
}

func getContainers(manifest map[interface{}]interface{}) []interface{} {
	var result []interface{}
	for key, value := range manifest {
		switch value.(type) {
		case map[interface{}]interface{}:
			result = getContainers(value.(map[interface{}]interface{}))
		case []interface{}:
			if key.(string) == "containers" {
				result = value.([]interface{})
			}
		}
	}
	return result
}

func main() {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(file), &t)
	if err != nil {
		panic(err)
	}
	containers := util.GetContainers(t)
	for _, container := range containers {
		livenessProbe := util.GetLivenessProbe(container)
		readinessProbe := util.GetReadinessProbe(container)
		fmt.Println(reflect.DeepEqual(readinessProbe, livenessProbe))
	}
}
