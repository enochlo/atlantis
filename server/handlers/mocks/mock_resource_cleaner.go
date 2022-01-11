// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/handlers (interfaces: ResourceCleaner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockResourceCleaner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockResourceCleaner(options ...pegomock.Option) *MockResourceCleaner {
	mock := &MockResourceCleaner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockResourceCleaner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockResourceCleaner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockResourceCleaner) CleanUp(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockResourceCleaner().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("CleanUp", params, []reflect.Type{})
}

func (mock *MockResourceCleaner) VerifyWasCalledOnce() *VerifierMockResourceCleaner {
	return &VerifierMockResourceCleaner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockResourceCleaner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockResourceCleaner {
	return &VerifierMockResourceCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockResourceCleaner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockResourceCleaner {
	return &VerifierMockResourceCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockResourceCleaner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockResourceCleaner {
	return &VerifierMockResourceCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockResourceCleaner struct {
	mock                   *MockResourceCleaner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockResourceCleaner) CleanUp(_param0 string) *MockResourceCleaner_CleanUp_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CleanUp", params, verifier.timeout)
	return &MockResourceCleaner_CleanUp_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockResourceCleaner_CleanUp_OngoingVerification struct {
	mock              *MockResourceCleaner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockResourceCleaner_CleanUp_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockResourceCleaner_CleanUp_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
