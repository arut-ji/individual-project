package duplicated_liveness_and_readiness

import (
	"bytes"
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v3"
	"reflect"
)

func GetNumberOfInstances(script string) (int, error) {
	return countSmellInstances(script)
}

func countSmellInstances(script string) (int, error) {
	dec := yaml.NewDecoder(bytes.NewReader([]byte(script)))
	var t map[interface{}]interface{}
	count := 0
	for dec.Decode(&t) == nil {
		containers := util.GetContainers(t)
		for _, container := range containers {
			livenessProbe := util.GetLivenessProbe(container)
			readinessProbe := util.GetReadinessProbe(container)
			if reflect.DeepEqual(readinessProbe, livenessProbe) {
				count += 1
			}
		}
	}
	return count, nil
}
