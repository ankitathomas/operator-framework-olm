// Code generated by counterfeiter. DO NOT EDIT.
package solver

import (
	"sync"
	"time"

	"github.com/go-air/gini/inter"
	"github.com/go-air/gini/z"
)

type FakeS struct {
	AddStub        func(z.Lit)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 z.Lit
	}
	AssumeStub        func(...z.Lit)
	assumeMutex       sync.RWMutex
	assumeArgsForCall []struct {
		arg1 []z.Lit
	}
	GoSolveStub        func() inter.Solve
	goSolveMutex       sync.RWMutex
	goSolveArgsForCall []struct {
	}
	goSolveReturns struct {
		result1 inter.Solve
	}
	goSolveReturnsOnCall map[int]struct {
		result1 inter.Solve
	}
	LitStub        func() z.Lit
	litMutex       sync.RWMutex
	litArgsForCall []struct {
	}
	litReturns struct {
		result1 z.Lit
	}
	litReturnsOnCall map[int]struct {
		result1 z.Lit
	}
	MaxVarStub        func() z.Var
	maxVarMutex       sync.RWMutex
	maxVarArgsForCall []struct {
	}
	maxVarReturns struct {
		result1 z.Var
	}
	maxVarReturnsOnCall map[int]struct {
		result1 z.Var
	}
	ReasonsStub        func([]z.Lit, z.Lit) []z.Lit
	reasonsMutex       sync.RWMutex
	reasonsArgsForCall []struct {
		arg1 []z.Lit
		arg2 z.Lit
	}
	reasonsReturns struct {
		result1 []z.Lit
	}
	reasonsReturnsOnCall map[int]struct {
		result1 []z.Lit
	}
	SCopyStub        func() inter.S
	sCopyMutex       sync.RWMutex
	sCopyArgsForCall []struct {
	}
	sCopyReturns struct {
		result1 inter.S
	}
	sCopyReturnsOnCall map[int]struct {
		result1 inter.S
	}
	SolveStub        func() int
	solveMutex       sync.RWMutex
	solveArgsForCall []struct {
	}
	solveReturns struct {
		result1 int
	}
	solveReturnsOnCall map[int]struct {
		result1 int
	}
	TestStub        func([]z.Lit) (int, []z.Lit)
	testMutex       sync.RWMutex
	testArgsForCall []struct {
		arg1 []z.Lit
	}
	testReturns struct {
		result1 int
		result2 []z.Lit
	}
	testReturnsOnCall map[int]struct {
		result1 int
		result2 []z.Lit
	}
	TryStub        func(time.Duration) int
	tryMutex       sync.RWMutex
	tryArgsForCall []struct {
		arg1 time.Duration
	}
	tryReturns struct {
		result1 int
	}
	tryReturnsOnCall map[int]struct {
		result1 int
	}
	UntestStub        func() int
	untestMutex       sync.RWMutex
	untestArgsForCall []struct {
	}
	untestReturns struct {
		result1 int
	}
	untestReturnsOnCall map[int]struct {
		result1 int
	}
	ValueStub        func(z.Lit) bool
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
		arg1 z.Lit
	}
	valueReturns struct {
		result1 bool
	}
	valueReturnsOnCall map[int]struct {
		result1 bool
	}
	WhyStub        func([]z.Lit) []z.Lit
	whyMutex       sync.RWMutex
	whyArgsForCall []struct {
		arg1 []z.Lit
	}
	whyReturns struct {
		result1 []z.Lit
	}
	whyReturnsOnCall map[int]struct {
		result1 []z.Lit
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeS) Add(arg1 z.Lit) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 z.Lit
	}{arg1})
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(arg1)
	}
}

func (fake *FakeS) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeS) AddCalls(stub func(z.Lit)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeS) AddArgsForCall(i int) z.Lit {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) Assume(arg1 ...z.Lit) {
	fake.assumeMutex.Lock()
	fake.assumeArgsForCall = append(fake.assumeArgsForCall, struct {
		arg1 []z.Lit
	}{arg1})
	fake.recordInvocation("Assume", []interface{}{arg1})
	fake.assumeMutex.Unlock()
	if fake.AssumeStub != nil {
		fake.AssumeStub(arg1...)
	}
}

func (fake *FakeS) AssumeCallCount() int {
	fake.assumeMutex.RLock()
	defer fake.assumeMutex.RUnlock()
	return len(fake.assumeArgsForCall)
}

func (fake *FakeS) AssumeCalls(stub func(...z.Lit)) {
	fake.assumeMutex.Lock()
	defer fake.assumeMutex.Unlock()
	fake.AssumeStub = stub
}

func (fake *FakeS) AssumeArgsForCall(i int) []z.Lit {
	fake.assumeMutex.RLock()
	defer fake.assumeMutex.RUnlock()
	argsForCall := fake.assumeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) GoSolve() inter.Solve {
	fake.goSolveMutex.Lock()
	ret, specificReturn := fake.goSolveReturnsOnCall[len(fake.goSolveArgsForCall)]
	fake.goSolveArgsForCall = append(fake.goSolveArgsForCall, struct {
	}{})
	fake.recordInvocation("GoSolve", []interface{}{})
	fake.goSolveMutex.Unlock()
	if fake.GoSolveStub != nil {
		return fake.GoSolveStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.goSolveReturns
	return fakeReturns.result1
}

func (fake *FakeS) GoSolveCallCount() int {
	fake.goSolveMutex.RLock()
	defer fake.goSolveMutex.RUnlock()
	return len(fake.goSolveArgsForCall)
}

func (fake *FakeS) GoSolveCalls(stub func() inter.Solve) {
	fake.goSolveMutex.Lock()
	defer fake.goSolveMutex.Unlock()
	fake.GoSolveStub = stub
}

func (fake *FakeS) GoSolveReturns(result1 inter.Solve) {
	fake.goSolveMutex.Lock()
	defer fake.goSolveMutex.Unlock()
	fake.GoSolveStub = nil
	fake.goSolveReturns = struct {
		result1 inter.Solve
	}{result1}
}

func (fake *FakeS) GoSolveReturnsOnCall(i int, result1 inter.Solve) {
	fake.goSolveMutex.Lock()
	defer fake.goSolveMutex.Unlock()
	fake.GoSolveStub = nil
	if fake.goSolveReturnsOnCall == nil {
		fake.goSolveReturnsOnCall = make(map[int]struct {
			result1 inter.Solve
		})
	}
	fake.goSolveReturnsOnCall[i] = struct {
		result1 inter.Solve
	}{result1}
}

func (fake *FakeS) Lit() z.Lit {
	fake.litMutex.Lock()
	ret, specificReturn := fake.litReturnsOnCall[len(fake.litArgsForCall)]
	fake.litArgsForCall = append(fake.litArgsForCall, struct {
	}{})
	fake.recordInvocation("Lit", []interface{}{})
	fake.litMutex.Unlock()
	if fake.LitStub != nil {
		return fake.LitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.litReturns
	return fakeReturns.result1
}

func (fake *FakeS) LitCallCount() int {
	fake.litMutex.RLock()
	defer fake.litMutex.RUnlock()
	return len(fake.litArgsForCall)
}

func (fake *FakeS) LitCalls(stub func() z.Lit) {
	fake.litMutex.Lock()
	defer fake.litMutex.Unlock()
	fake.LitStub = stub
}

func (fake *FakeS) LitReturns(result1 z.Lit) {
	fake.litMutex.Lock()
	defer fake.litMutex.Unlock()
	fake.LitStub = nil
	fake.litReturns = struct {
		result1 z.Lit
	}{result1}
}

func (fake *FakeS) LitReturnsOnCall(i int, result1 z.Lit) {
	fake.litMutex.Lock()
	defer fake.litMutex.Unlock()
	fake.LitStub = nil
	if fake.litReturnsOnCall == nil {
		fake.litReturnsOnCall = make(map[int]struct {
			result1 z.Lit
		})
	}
	fake.litReturnsOnCall[i] = struct {
		result1 z.Lit
	}{result1}
}

func (fake *FakeS) MaxVar() z.Var {
	fake.maxVarMutex.Lock()
	ret, specificReturn := fake.maxVarReturnsOnCall[len(fake.maxVarArgsForCall)]
	fake.maxVarArgsForCall = append(fake.maxVarArgsForCall, struct {
	}{})
	fake.recordInvocation("MaxVar", []interface{}{})
	fake.maxVarMutex.Unlock()
	if fake.MaxVarStub != nil {
		return fake.MaxVarStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.maxVarReturns
	return fakeReturns.result1
}

func (fake *FakeS) MaxVarCallCount() int {
	fake.maxVarMutex.RLock()
	defer fake.maxVarMutex.RUnlock()
	return len(fake.maxVarArgsForCall)
}

func (fake *FakeS) MaxVarCalls(stub func() z.Var) {
	fake.maxVarMutex.Lock()
	defer fake.maxVarMutex.Unlock()
	fake.MaxVarStub = stub
}

func (fake *FakeS) MaxVarReturns(result1 z.Var) {
	fake.maxVarMutex.Lock()
	defer fake.maxVarMutex.Unlock()
	fake.MaxVarStub = nil
	fake.maxVarReturns = struct {
		result1 z.Var
	}{result1}
}

func (fake *FakeS) MaxVarReturnsOnCall(i int, result1 z.Var) {
	fake.maxVarMutex.Lock()
	defer fake.maxVarMutex.Unlock()
	fake.MaxVarStub = nil
	if fake.maxVarReturnsOnCall == nil {
		fake.maxVarReturnsOnCall = make(map[int]struct {
			result1 z.Var
		})
	}
	fake.maxVarReturnsOnCall[i] = struct {
		result1 z.Var
	}{result1}
}

func (fake *FakeS) Reasons(arg1 []z.Lit, arg2 z.Lit) []z.Lit {
	var arg1Copy []z.Lit
	if arg1 != nil {
		arg1Copy = make([]z.Lit, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.reasonsMutex.Lock()
	ret, specificReturn := fake.reasonsReturnsOnCall[len(fake.reasonsArgsForCall)]
	fake.reasonsArgsForCall = append(fake.reasonsArgsForCall, struct {
		arg1 []z.Lit
		arg2 z.Lit
	}{arg1Copy, arg2})
	fake.recordInvocation("Reasons", []interface{}{arg1Copy, arg2})
	fake.reasonsMutex.Unlock()
	if fake.ReasonsStub != nil {
		return fake.ReasonsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.reasonsReturns
	return fakeReturns.result1
}

func (fake *FakeS) ReasonsCallCount() int {
	fake.reasonsMutex.RLock()
	defer fake.reasonsMutex.RUnlock()
	return len(fake.reasonsArgsForCall)
}

func (fake *FakeS) ReasonsCalls(stub func([]z.Lit, z.Lit) []z.Lit) {
	fake.reasonsMutex.Lock()
	defer fake.reasonsMutex.Unlock()
	fake.ReasonsStub = stub
}

func (fake *FakeS) ReasonsArgsForCall(i int) ([]z.Lit, z.Lit) {
	fake.reasonsMutex.RLock()
	defer fake.reasonsMutex.RUnlock()
	argsForCall := fake.reasonsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeS) ReasonsReturns(result1 []z.Lit) {
	fake.reasonsMutex.Lock()
	defer fake.reasonsMutex.Unlock()
	fake.ReasonsStub = nil
	fake.reasonsReturns = struct {
		result1 []z.Lit
	}{result1}
}

func (fake *FakeS) ReasonsReturnsOnCall(i int, result1 []z.Lit) {
	fake.reasonsMutex.Lock()
	defer fake.reasonsMutex.Unlock()
	fake.ReasonsStub = nil
	if fake.reasonsReturnsOnCall == nil {
		fake.reasonsReturnsOnCall = make(map[int]struct {
			result1 []z.Lit
		})
	}
	fake.reasonsReturnsOnCall[i] = struct {
		result1 []z.Lit
	}{result1}
}

func (fake *FakeS) SCopy() inter.S {
	fake.sCopyMutex.Lock()
	ret, specificReturn := fake.sCopyReturnsOnCall[len(fake.sCopyArgsForCall)]
	fake.sCopyArgsForCall = append(fake.sCopyArgsForCall, struct {
	}{})
	fake.recordInvocation("SCopy", []interface{}{})
	fake.sCopyMutex.Unlock()
	if fake.SCopyStub != nil {
		return fake.SCopyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sCopyReturns
	return fakeReturns.result1
}

func (fake *FakeS) SCopyCallCount() int {
	fake.sCopyMutex.RLock()
	defer fake.sCopyMutex.RUnlock()
	return len(fake.sCopyArgsForCall)
}

func (fake *FakeS) SCopyCalls(stub func() inter.S) {
	fake.sCopyMutex.Lock()
	defer fake.sCopyMutex.Unlock()
	fake.SCopyStub = stub
}

func (fake *FakeS) SCopyReturns(result1 inter.S) {
	fake.sCopyMutex.Lock()
	defer fake.sCopyMutex.Unlock()
	fake.SCopyStub = nil
	fake.sCopyReturns = struct {
		result1 inter.S
	}{result1}
}

func (fake *FakeS) SCopyReturnsOnCall(i int, result1 inter.S) {
	fake.sCopyMutex.Lock()
	defer fake.sCopyMutex.Unlock()
	fake.SCopyStub = nil
	if fake.sCopyReturnsOnCall == nil {
		fake.sCopyReturnsOnCall = make(map[int]struct {
			result1 inter.S
		})
	}
	fake.sCopyReturnsOnCall[i] = struct {
		result1 inter.S
	}{result1}
}

func (fake *FakeS) Solve() int {
	fake.solveMutex.Lock()
	ret, specificReturn := fake.solveReturnsOnCall[len(fake.solveArgsForCall)]
	fake.solveArgsForCall = append(fake.solveArgsForCall, struct {
	}{})
	fake.recordInvocation("Solve", []interface{}{})
	fake.solveMutex.Unlock()
	if fake.SolveStub != nil {
		return fake.SolveStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.solveReturns
	return fakeReturns.result1
}

func (fake *FakeS) SolveCallCount() int {
	fake.solveMutex.RLock()
	defer fake.solveMutex.RUnlock()
	return len(fake.solveArgsForCall)
}

func (fake *FakeS) SolveCalls(stub func() int) {
	fake.solveMutex.Lock()
	defer fake.solveMutex.Unlock()
	fake.SolveStub = stub
}

func (fake *FakeS) SolveReturns(result1 int) {
	fake.solveMutex.Lock()
	defer fake.solveMutex.Unlock()
	fake.SolveStub = nil
	fake.solveReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) SolveReturnsOnCall(i int, result1 int) {
	fake.solveMutex.Lock()
	defer fake.solveMutex.Unlock()
	fake.SolveStub = nil
	if fake.solveReturnsOnCall == nil {
		fake.solveReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.solveReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) Test(arg1 []z.Lit) (int, []z.Lit) {
	var arg1Copy []z.Lit
	if arg1 != nil {
		arg1Copy = make([]z.Lit, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.testMutex.Lock()
	ret, specificReturn := fake.testReturnsOnCall[len(fake.testArgsForCall)]
	fake.testArgsForCall = append(fake.testArgsForCall, struct {
		arg1 []z.Lit
	}{arg1Copy})
	fake.recordInvocation("Test", []interface{}{arg1Copy})
	fake.testMutex.Unlock()
	if fake.TestStub != nil {
		return fake.TestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.testReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeS) TestCallCount() int {
	fake.testMutex.RLock()
	defer fake.testMutex.RUnlock()
	return len(fake.testArgsForCall)
}

func (fake *FakeS) TestCalls(stub func([]z.Lit) (int, []z.Lit)) {
	fake.testMutex.Lock()
	defer fake.testMutex.Unlock()
	fake.TestStub = stub
}

func (fake *FakeS) TestArgsForCall(i int) []z.Lit {
	fake.testMutex.RLock()
	defer fake.testMutex.RUnlock()
	argsForCall := fake.testArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) TestReturns(result1 int, result2 []z.Lit) {
	fake.testMutex.Lock()
	defer fake.testMutex.Unlock()
	fake.TestStub = nil
	fake.testReturns = struct {
		result1 int
		result2 []z.Lit
	}{result1, result2}
}

func (fake *FakeS) TestReturnsOnCall(i int, result1 int, result2 []z.Lit) {
	fake.testMutex.Lock()
	defer fake.testMutex.Unlock()
	fake.TestStub = nil
	if fake.testReturnsOnCall == nil {
		fake.testReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []z.Lit
		})
	}
	fake.testReturnsOnCall[i] = struct {
		result1 int
		result2 []z.Lit
	}{result1, result2}
}

func (fake *FakeS) Try(arg1 time.Duration) int {
	fake.tryMutex.Lock()
	ret, specificReturn := fake.tryReturnsOnCall[len(fake.tryArgsForCall)]
	fake.tryArgsForCall = append(fake.tryArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("Try", []interface{}{arg1})
	fake.tryMutex.Unlock()
	if fake.TryStub != nil {
		return fake.TryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tryReturns
	return fakeReturns.result1
}

func (fake *FakeS) TryCallCount() int {
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	return len(fake.tryArgsForCall)
}

func (fake *FakeS) TryCalls(stub func(time.Duration) int) {
	fake.tryMutex.Lock()
	defer fake.tryMutex.Unlock()
	fake.TryStub = stub
}

func (fake *FakeS) TryArgsForCall(i int) time.Duration {
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	argsForCall := fake.tryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) TryReturns(result1 int) {
	fake.tryMutex.Lock()
	defer fake.tryMutex.Unlock()
	fake.TryStub = nil
	fake.tryReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) TryReturnsOnCall(i int, result1 int) {
	fake.tryMutex.Lock()
	defer fake.tryMutex.Unlock()
	fake.TryStub = nil
	if fake.tryReturnsOnCall == nil {
		fake.tryReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.tryReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) Untest() int {
	fake.untestMutex.Lock()
	ret, specificReturn := fake.untestReturnsOnCall[len(fake.untestArgsForCall)]
	fake.untestArgsForCall = append(fake.untestArgsForCall, struct {
	}{})
	fake.recordInvocation("Untest", []interface{}{})
	fake.untestMutex.Unlock()
	if fake.UntestStub != nil {
		return fake.UntestStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.untestReturns
	return fakeReturns.result1
}

func (fake *FakeS) UntestCallCount() int {
	fake.untestMutex.RLock()
	defer fake.untestMutex.RUnlock()
	return len(fake.untestArgsForCall)
}

func (fake *FakeS) UntestCalls(stub func() int) {
	fake.untestMutex.Lock()
	defer fake.untestMutex.Unlock()
	fake.UntestStub = stub
}

func (fake *FakeS) UntestReturns(result1 int) {
	fake.untestMutex.Lock()
	defer fake.untestMutex.Unlock()
	fake.UntestStub = nil
	fake.untestReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) UntestReturnsOnCall(i int, result1 int) {
	fake.untestMutex.Lock()
	defer fake.untestMutex.Unlock()
	fake.UntestStub = nil
	if fake.untestReturnsOnCall == nil {
		fake.untestReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.untestReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeS) Value(arg1 z.Lit) bool {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
		arg1 z.Lit
	}{arg1})
	fake.recordInvocation("Value", []interface{}{arg1})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1
}

func (fake *FakeS) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeS) ValueCalls(stub func(z.Lit) bool) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeS) ValueArgsForCall(i int) z.Lit {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) ValueReturns(result1 bool) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeS) ValueReturnsOnCall(i int, result1 bool) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeS) Why(arg1 []z.Lit) []z.Lit {
	var arg1Copy []z.Lit
	if arg1 != nil {
		arg1Copy = make([]z.Lit, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.whyMutex.Lock()
	ret, specificReturn := fake.whyReturnsOnCall[len(fake.whyArgsForCall)]
	fake.whyArgsForCall = append(fake.whyArgsForCall, struct {
		arg1 []z.Lit
	}{arg1Copy})
	fake.recordInvocation("Why", []interface{}{arg1Copy})
	fake.whyMutex.Unlock()
	if fake.WhyStub != nil {
		return fake.WhyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.whyReturns
	return fakeReturns.result1
}

func (fake *FakeS) WhyCallCount() int {
	fake.whyMutex.RLock()
	defer fake.whyMutex.RUnlock()
	return len(fake.whyArgsForCall)
}

func (fake *FakeS) WhyCalls(stub func([]z.Lit) []z.Lit) {
	fake.whyMutex.Lock()
	defer fake.whyMutex.Unlock()
	fake.WhyStub = stub
}

func (fake *FakeS) WhyArgsForCall(i int) []z.Lit {
	fake.whyMutex.RLock()
	defer fake.whyMutex.RUnlock()
	argsForCall := fake.whyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeS) WhyReturns(result1 []z.Lit) {
	fake.whyMutex.Lock()
	defer fake.whyMutex.Unlock()
	fake.WhyStub = nil
	fake.whyReturns = struct {
		result1 []z.Lit
	}{result1}
}

func (fake *FakeS) WhyReturnsOnCall(i int, result1 []z.Lit) {
	fake.whyMutex.Lock()
	defer fake.whyMutex.Unlock()
	fake.WhyStub = nil
	if fake.whyReturnsOnCall == nil {
		fake.whyReturnsOnCall = make(map[int]struct {
			result1 []z.Lit
		})
	}
	fake.whyReturnsOnCall[i] = struct {
		result1 []z.Lit
	}{result1}
}

func (fake *FakeS) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.assumeMutex.RLock()
	defer fake.assumeMutex.RUnlock()
	fake.goSolveMutex.RLock()
	defer fake.goSolveMutex.RUnlock()
	fake.litMutex.RLock()
	defer fake.litMutex.RUnlock()
	fake.maxVarMutex.RLock()
	defer fake.maxVarMutex.RUnlock()
	fake.reasonsMutex.RLock()
	defer fake.reasonsMutex.RUnlock()
	fake.sCopyMutex.RLock()
	defer fake.sCopyMutex.RUnlock()
	fake.solveMutex.RLock()
	defer fake.solveMutex.RUnlock()
	fake.testMutex.RLock()
	defer fake.testMutex.RUnlock()
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	fake.untestMutex.RLock()
	defer fake.untestMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	fake.whyMutex.RLock()
	defer fake.whyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeS) recordInvocation(key string, args []interface{}) {
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

var _ inter.S = new(FakeS)
