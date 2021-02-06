package duplicated_liveness_and_readiness

import (
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v3"
	"reflect"
)

func Scan(script string) (bool, error) {
	return hasDuplicatedLivenessAndReadiness(script)
}

func hasDuplicatedLivenessAndReadiness(script string) (bool, error) {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(script), &t)
	if err != nil {
		panic(err)
	}
	containers := util.GetContainers(t)
	result := false
	for _, container := range containers {
		livenessProbe := util.GetLivenessProbe(container)
		readinessProbe := util.GetReadinessProbe(container)
		result = result || reflect.DeepEqual(readinessProbe, livenessProbe)
	}
	return false, nil
}
