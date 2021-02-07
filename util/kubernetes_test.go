package util

import "testing"

type Node map[interface{}]interface{}

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
	manifest := map[interface{}]interface{}{
		"name":  "liveness",
		"image": "k8s.grc.io/busybox",
		"livenessProbe": map[interface{}]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
		"readinessProbe": map[interface{}]interface{}{
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
	manifest := map[interface{}]interface{}{
		"name":  "liveness",
		"image": "k8s.grc.io/busybox",
		"livenessProbe": map[interface{}]interface{}{
			"exec": map[interface{}]interface{}{
				"command": []interface{}{"cat", "/tmp/healthy"},
			},
		},
		"readinessProbe": map[interface{}]interface{}{
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
