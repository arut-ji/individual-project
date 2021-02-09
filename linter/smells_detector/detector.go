package smells_detector

import (
	"github.com/arut-ji/individual-project/linter/smells_detector/avoid_comments"
	"github.com/arut-ji/individual-project/linter/smells_detector/incomplete_tasks"
)

type (
	getNumberOfInstancesFn func(string) (int, error)
	Detectors              map[string]getNumberOfInstancesFn
	DetectionResult        map[string]int
)

func newDetector() Detectors {
	return Detectors{
		"avoid-comments": avoid_comments.GetNumberOfInstances,
		//"duplicate-liveness-and-readiness-probe": duplicated_liveness_and_readiness.Scan,
		//"improper-alignment":                     improper_alignment.Scan,
		"in-complete-tasks": incomplete_tasks.GetNumberOfInstances,
		//"long-statement":                         long_statement.Scan,
		//"missing-readiness-probe":                missing_readiness_probes.Scan,
	}
}

func Detect(script string) (DetectionResult, error) {
	detectors := newDetector()
	results := make(DetectionResult)
	for key, fn := range detectors {
		result, err := fn(script)
		if err != nil {
			return nil, err
		}
		results[key] = result
	}
	return results, nil
}
