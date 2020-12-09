package smells_detector

import "testing"

func TestHasCommentsForNoSmell(t *testing.T) {
	script := `apiVersion: v1
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local`
	if result, err := hasComments(script); result != false || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestHasCommentsForContainingSmell(t *testing.T) {
	script := `
apiVersion: v1 # ConfigMap comment
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local`
	if result, err := hasComments(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}
