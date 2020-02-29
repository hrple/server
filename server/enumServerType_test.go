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
		t.Fatalf(ErrMsgExpectedInsteadOfResult, RunningContextTypeLambda, runningContextType)
	}
}

func TestGetRunningContextTypeStandalone(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType for RunningContextTypeStandalone- Error: %v", err)
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf(ErrMsgExpectedInsteadOfResult, RunningContextTypeStandalone, runningContextType)
	}
}

func TestGetRunningContextTypeRandom(t *testing.T) {
	runningContextType, err := GetRunningContextType("RunAway")
	if err == nil {
		t.Fatal("Error expect when providing invalid RunningContext")
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf(ErrMsgExpectedInsteadOfResult, RunningContextTypeStandalone, runningContextType)
	}
}
