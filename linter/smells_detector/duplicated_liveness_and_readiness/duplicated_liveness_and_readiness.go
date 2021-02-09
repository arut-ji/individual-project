package duplicated_liveness_and_readiness

import (
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v3"
	"reflect"
)

func GetNumberOfInstances(script string) (int, error) {
	return countSmellInstances(script)
}

func countSmellInstances(script string) (int, error) {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(script), &t)
	if err != nil {
		return -1, err
	}
	containers := util.GetContainers(t)

	count := 0
	for _, container := range containers {
		livenessProbe := util.GetLivenessProbe(container)
		readinessProbe := util.GetReadinessProbe(container)
		if reflect.DeepEqual(readinessProbe, livenessProbe) {
			count += 1
		}
	}
	return count, nil
}
