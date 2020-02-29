package server

import (
	"testing"
)

var (
	errMsgExpectedInsteadOfResult = "Expected %v instead of %v"
)

func TestGetRunningContextTypeLambda(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeLambda)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType for RunningContextTypeLambda - Error: %v", err)
	}

	if runningContextType != RunningContextTypeLambda {
		t.Fatalf(errMsgExpectedInsteadOfResult, RunningContextTypeLambda, runningContextType)
	}
}

func TestGetRunningContextTypeStandalone(t *testing.T) {
	runningContextType, err := GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		t.Fatalf("Error failed to GetRunningContextType for RunningContextTypeStandalone- Error: %v", err)
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf(errMsgExpectedInsteadOfResult, RunningContextTypeStandalone, runningContextType)
	}
}

func TestGetRunningContextTypeRandom(t *testing.T) {
	runningContextType, err := GetRunningContextType("RunAway")
	if err == nil {
		t.Fatal("Error expect when providing invalid RunningContext")
	}

	if runningContextType != RunningContextTypeStandalone {
		t.Fatalf(errMsgExpectedInsteadOfResult, RunningContextTypeStandalone, runningContextType)
	}
}
