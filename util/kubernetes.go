package util

import (
	"regexp"
	"strings"
)

func GetContainers(manifest interface{}) []interface{} {
	var result []interface{}
	switch manifest.(type) {
	case map[interface{}]interface{}:
		manifest := manifest.(map[interface{}]interface{})
		for key, value := range manifest {
			if key.(string) == "containers" {
				return value.([]interface{})
			} else {
				result = append(result, GetContainers(value)...)
			}
		}
	case map[string]interface{}:
		manifest := manifest.(map[string]interface{})
		for key, value := range manifest {
			if key == "containers" {
				return value.([]interface{})
			} else {
				result = append(result, GetContainers(value)...)
			}
		}
	}
	return result
}

func GetReadinessProbe(container interface{}) map[string]interface{} {
	result := make(map[string]interface{}, 1)
	switch container.(type) {
	case map[string]interface{}:
		probe := container.(map[string]interface{})["readinessProbe"]
		if probe == nil {
			return nil
		}
		for k, v := range probe.(map[string]interface{}) {
			result[k] = v
		}
	case map[interface{}]interface{}:
		probe := container.(map[interface{}]interface{})["readinessProbe"]
		if probe == nil {
			return nil
		}
		for k, v := range probe.(map[interface{}]interface{}) {
			result[k.(string)] = v
		}
	}
	return result
}

func GetLivenessProbe(container interface{}) map[string]interface{} {
	probe := container.(map[string]interface{})["livenessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[string]interface{})
}

func GetNumberOfResources(script string) int {
	count := 1
	lines := strings.Split(script, "\n")
	for _, line := range lines {
		match, _ := regexp.MatchString("---", line)
		if match {
			count += 1
		}
	}
	return count
}
