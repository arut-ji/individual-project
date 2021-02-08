package util

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
	probe := container.(map[string]interface{})["readinessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[string]interface{})
}

func GetLivenessProbe(container interface{}) map[string]interface{} {
	probe := container.(map[string]interface{})["livenessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[string]interface{})
}
