package server

import (
	"errors"
	"strings"
)

//Convert to context object?
const (
	RunningContextTypeLambda     = "LAMBDA"
	RunningContextTypeStandalone = "STANDALONE"
)

// GetRunningContextType allows for the retrieval of a valid run time environment i.e. Standalone aka docker or Lambda
func GetRunningContextType(contextType string) (string, error) {
	contextType = strings.ToUpper(contextType)

	if contextType != RunningContextTypeLambda && contextType != RunningContextTypeStandalone {
		return RunningContextTypeStandalone, errors.New("WARNING: Invalid GetRunningContextType specified, using STANDALONE")
	}

	if contextType == RunningContextTypeLambda {
		contextType = RunningContextTypeLambda
	} else {
		contextType = RunningContextTypeStandalone
	}

	return contextType, nil
}
