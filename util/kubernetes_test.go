package util

import "testing"

func TestGetContainers(t *testing.T) {
	manifest := map[interface{}]interface{}{
		"spec": map[interface{}]interface{}{
			"containers": []interface{}{
				map[interface{}]interface{}{
					"name":  "liveness",
					"image": "k8s.grc.io/busybox",
				},
				map[interface{}]interface{}{
					"name":  "liveness",
					"image": "k8s.grc.io/busybox",
				},
				map[interface{}]interface{}{
					"name":  "liveness",
					"image": "k8s.grc.io/busybox",
				},
			},
		},
	}
	result := GetContainers(manifest)
	if len(result) != 3 {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", len(result), 3)
	}
}

func TestGetReadinessProbe(t *testing.T) {
	manifest := map[string]interface{}{
		"name":  "liveness",
		"image": "k8s.grc.io/busybox",
		"livenessProbe": map[string]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
		"readinessProbe": map[string]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
	}
	result := GetReadinessProbe(manifest)
	if result["exec"] == nil {
		t.Error("Detection returned errors")
	}
}

func TestGetLivenessProbe(t *testing.T) {
	manifest := map[string]interface{}{
		"name":  "liveness",
		"image": "k8s.grc.io/busybox",
		"livenessProbe": map[string]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
		"readinessProbe": map[string]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
	}
	result := GetLivenessProbe(manifest)
	if result["exec"] == nil {
		t.Error("Detection returned errors")
	}
}

func loadFixture(name string) string {
	return LoadFixture("./fixtures")(name)
}

func TestGetNumberOfResourcesForOneResource(t *testing.T) {
	script := loadFixture("get_number_of_resources_test/one_resource.yaml")
	if result := GetNumberOfResources(script); result != 1 {
		t.Errorf("The result was incorrect, got: %v, want: %v.", result, 1)
	}
}

func TestGetNumberOfResourcesForMultipleResources(t *testing.T) {
	script := loadFixture("get_number_of_resources_test/two_resources.yaml")
	if result := GetNumberOfResources(script); result != 2 {
		t.Errorf("The result was incorrect, got: %v, want: %v.", result, 2)
	}
}
