package util

func GetContainers(manifest map[interface{}]interface{}) []interface{} {
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

func GetReadinessProbe(container interface{}) map[interface{}]interface{} {
	probe := container.(map[interface{}]interface{})["readinessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[interface{}]interface{})
}

func GetLivenessProbe(container interface{}) map[interface{}]interface{} {
	probe := container.(map[interface{}]interface{})["livenessProbe"]
	if probe == nil {
		return nil
	}
	return probe.(map[interface{}]interface{})
}
