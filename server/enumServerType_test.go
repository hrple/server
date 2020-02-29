package server

import (
	"testing"
)

func TestGetRunningContextTypeLambda(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeLambda)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType for RunningContextTypeLambda - Error: %v", err)
	}

	if runningContextType != RunningContextTypeLambda {
		t.Fatalf("Error expected %v as default instead of %v", RunningContextTypeLambda, runningContextType)
	}
}

func TestGetRunningContextTypeStandalone(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType for RunningContextTypeStandalone- Error: %v", err)
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf("Error expected %v as default instead of %v", RunningContextTypeStandalone, runningContextType)
	}
}

func TestGetRunningContextTypeRandom(t *testing.T) {
	runningContextType, err := GetRunningContextType("RunAway")
	if err == nil {
		t.Fatal("Error expect when providing invalid RunningContext")
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf("Error expected %v as default instead of %v", RunningContextTypeStandalone, runningContextType)
	}
}
