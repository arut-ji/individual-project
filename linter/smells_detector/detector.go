package smells_detector

import (
	"github.com/arut-ji/individual-project/linter/smells_detector/avoid_comments"
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
		"longStatement":         long_statement.Scan,
		"inCompleteTasks":       incomplete_tasks.Scan,
		"avoidComments":         avoid_comments.Scan,
		"improperAlignment":     improper_alignment.Scan,
		"missingReadinessProbe": missing_readiness_probes.Scan,
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
