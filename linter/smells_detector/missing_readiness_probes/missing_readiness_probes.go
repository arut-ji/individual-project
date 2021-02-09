package missing_readiness_probes

import (
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v2"
)

func GetNumberOfInstances(scripts string) (int, error) {
	return countMissingReadinessProbes(scripts)
}

func countMissingReadinessProbes(script string) (int, error) {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(script), &t)
	if err != nil {
		panic(err)
	}
	containers := util.GetContainers(t)
	count := 0
	for _, container := range containers {
		if probe := util.GetReadinessProbe(container); probe == nil {
			count += 1
		}
	}
	return count, nil
}
