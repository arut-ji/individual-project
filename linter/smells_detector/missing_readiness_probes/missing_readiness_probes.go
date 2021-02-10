package missing_readiness_probes

import (
	"bytes"
	"github.com/arut-ji/individual-project/util"
	"gopkg.in/yaml.v3"
)

func GetNumberOfInstances(scripts string) (int, error) {
	return countMissingReadinessProbes(scripts)
}

func countMissingReadinessProbes(script string) (int, error) {
	dec := yaml.NewDecoder(bytes.NewReader([]byte(script)))
	var t map[interface{}]interface{}
	count := 0
	for dec.Decode(&t) == nil {
		containers := util.GetContainers(t)
		for _, container := range containers {
			if probe := util.GetReadinessProbe(container); probe == nil {
				count += 1
			}
		}
	}
	return count, nil
}
