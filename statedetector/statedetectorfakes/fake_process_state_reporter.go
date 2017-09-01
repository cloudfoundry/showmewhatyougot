package statedetectorfakes

import (
	"sync"

	"code.cloudfoundry.org/showmewhatyougot/statedetector"
)

type FakeProcessStateReporter struct {
	RunStub        func(pidList []int, processesList []string) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		pidList       []int
		processesList []string
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessStateReporter) Run(pidList []int, processesList []string) error {
	var pidListCopy []int
	if pidList != nil {
		pidListCopy = make([]int, len(pidList))
		copy(pidListCopy, pidList)
	}
	var processesListCopy []string
	if processesList != nil {
		processesListCopy = make([]string, len(processesList))
		copy(processesListCopy, processesList)
	}
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		pidList       []int
		processesList []string
	}{pidListCopy, processesListCopy})
	fake.recordInvocation("Run", []interface{}{pidListCopy, processesListCopy})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(pidList, processesList)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeProcessStateReporter) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeProcessStateReporter) RunArgsForCall(i int) ([]int, []string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].pidList, fake.runArgsForCall[i].processesList
}

func (fake *FakeProcessStateReporter) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcessStateReporter) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcessStateReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessStateReporter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ statedetector.DataCollector = new(FakeProcessStateReporter)
