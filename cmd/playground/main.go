package main

import (
	"fmt"
	"github.com/arut-ji/individual-project/linter/smells_detector/duplicated_liveness_and_readiness"
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
    readinessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 14
      periodSeconds: 20
`

func main() {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(file), &t)
	if err != nil {
		panic(err)
	}
	result, err := duplicated_liveness_and_readiness.Scan(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
