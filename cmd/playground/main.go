package main

import (
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
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
	containers := getContainers(t)
	for _, container := range containers {
		livenessProbe := getLivenessProbe(container)
		readinessProbe := getReadinessProbe(container)
		spew.Dump(livenessProbe)
		spew.Dump(readinessProbe)
	}
}
