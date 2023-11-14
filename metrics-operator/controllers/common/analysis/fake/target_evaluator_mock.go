// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1beta1"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/controllers/common/analysis/types"
	"sync"
)

// ITargetEvaluatorMock is a mock implementation of analysis.ITargetEvaluator.
//
//	func TestSomethingThatUsesITargetEvaluator(t *testing.T) {
//
//		// make and configure a mocked analysis.ITargetEvaluator
//		mockedITargetEvaluator := &ITargetEvaluatorMock{
//			EvaluateFunc: func(val float64, target *metricsapi.Target) types.TargetResult {
//				panic("mock out the Evaluate method")
//			},
//		}
//
//		// use mockedITargetEvaluator in code that requires analysis.ITargetEvaluator
//		// and then make assertions.
//
//	}
type ITargetEvaluatorMock struct {
	// EvaluateFunc mocks the Evaluate method.
	EvaluateFunc func(val float64, target *metricsapi.Target) types.TargetResult

	// calls tracks calls to the methods.
	calls struct {
		// Evaluate holds details about calls to the Evaluate method.
		Evaluate []struct {
			// Val is the val argument value.
			Val float64
			// Target is the target argument value.
			Target *metricsapi.Target
		}
	}
	lockEvaluate sync.RWMutex
}

// Evaluate calls EvaluateFunc.
func (mock *ITargetEvaluatorMock) Evaluate(val float64, target *metricsapi.Target) types.TargetResult {
	if mock.EvaluateFunc == nil {
		panic("ITargetEvaluatorMock.EvaluateFunc: method is nil but ITargetEvaluator.Evaluate was just called")
	}
	callInfo := struct {
		Val    float64
		Target *metricsapi.Target
	}{
		Val:    val,
		Target: target,
	}
	mock.lockEvaluate.Lock()
	mock.calls.Evaluate = append(mock.calls.Evaluate, callInfo)
	mock.lockEvaluate.Unlock()
	return mock.EvaluateFunc(val, target)
}

// EvaluateCalls gets all the calls that were made to Evaluate.
// Check the length with:
//
//	len(mockedITargetEvaluator.EvaluateCalls())
func (mock *ITargetEvaluatorMock) EvaluateCalls() []struct {
	Val    float64
	Target *metricsapi.Target
} {
	var calls []struct {
		Val    float64
		Target *metricsapi.Target
	}
	mock.lockEvaluate.RLock()
	calls = mock.calls.Evaluate
	mock.lockEvaluate.RUnlock()
	return calls
}
