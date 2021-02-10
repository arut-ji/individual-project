package smells_detector

import (
	"github.com/arut-ji/individual-project/linter/smells_detector/avoid_comments"
	"github.com/arut-ji/individual-project/linter/smells_detector/duplicated_liveness_and_readiness"
	"github.com/arut-ji/individual-project/linter/smells_detector/improper_alignment"
	"github.com/arut-ji/individual-project/linter/smells_detector/incomplete_tasks"
	"github.com/arut-ji/individual-project/linter/smells_detector/long_statement"
	"github.com/arut-ji/individual-project/linter/smells_detector/missing_readiness_probes"
)

type (
	getNumberOfInstancesFn func(string) (int, error)
	Detectors              map[string]getNumberOfInstancesFn
	DetectionResult        map[string]int
)

func newDetector() Detectors {
	return Detectors{
		"avoid-comments":                         avoid_comments.GetNumberOfInstances,
		"duplicate-liveness-and-readiness-probe": duplicated_liveness_and_readiness.GetNumberOfInstances,
		"improper-alignment":                     improper_alignment.GetNumberOfInstances,
		"in-complete-tasks":                      incomplete_tasks.GetNumberOfInstances,
		"long-statement":                         long_statement.GetNumberOfInstances,
		"missing-readiness-probe":                missing_readiness_probes.GetNumberOfInstances,
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
