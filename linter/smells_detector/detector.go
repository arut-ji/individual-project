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
	scan            func(string) (bool, error)
	Detectors       map[string]scan
	DetectionResult map[string]bool
)

func newDetector() Detectors {
	return Detectors{
		"avoid-comments":                         avoid_comments.Scan,
		"duplicate-liveness-and-readiness-probe": duplicated_liveness_and_readiness.Scan,
		"improper-alignment":                     improper_alignment.Scan,
		"in-complete-tasks":                      incomplete_tasks.Scan,
		"long-statement":                         long_statement.Scan,
		"missing-readiness-probe":                missing_readiness_probes.Scan,
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
