package smells_detector

import "testing"

func TestInCompleteTasks(t *testing.T) {
	script := `# FIXME: Inconsistent naming in services.
# prefix with some strings FIXME: Meow
# prefix with some strings TODO: add something
apiVersion: v1
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local
	
# FIXME: Meow Meow Meow
# Normal Comment
# TODO: Add more configmaps`
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForTODO(t *testing.T) {
	script := `# prefix with some strings TODO: add something
apiVersion: v1
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local
# TODO: Add more configmaps`
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForFIXME(t *testing.T) {
	script := `# FIXME: Inconsistent naming in services.
# prefix with some strings FIXME: Meow
apiVersion: v1
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local`
	if result, err := inCompleteTasks(script); result != true || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, true)
	}
}

func TestInCompleteTaskForNoSmells(t *testing.T) {
	script := `apiVersion: v1
kind: ConfigMap
metadata:
	name: gateway-staging-config
	namespace: staging
data:
	auth-service: staging.auth-service.svc.cluster.local
	PAYMENT_SERVICE: staging.payment-service.svc.cluster.local`
	if result, err := inCompleteTasks(script); result != false || err != nil {
		t.Errorf("Detection result was incorrect, got: %v, want: %v.", result, false)
	}
}
