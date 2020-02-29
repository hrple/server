package server

import (
	"testing"
)

func TestGetRunningContextTypeLambda(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeLambda)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType - Error: %v", err)
	}

	if runningContextType != RunningContextTypeLambda {
		t.Fatalf("Error failed to GetRunningContextType - Error: %v", "Expected: "+RunningContextTypeLambda)
	}
}

func TestGetRunningContextTypeStandalone(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType - Error: %v", err)
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf("Error failed to GetRunningContextType - Error: %v", "Expected: "+RunningContextTypeStandalone)
	}
}
