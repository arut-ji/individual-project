package missing_readiness_probes

import (
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v2"
)

func Scan(scripts string) (bool, error) {
	return hasMissingReadinessProbes(scripts)
}

func hasMissingReadinessProbes(script string) (bool, error) {
	t := make(map[interface{}]interface{}, 1)
	err := yaml.Unmarshal([]byte(script), &t)
	if err != nil {
		panic(err)
	}
	containers := util.GetContainers(t)
	for _, container := range containers {
		if probe := util.GetReadinessProbe(container); probe == nil {
			return true, nil
		}
	}
	return false, nil
}
