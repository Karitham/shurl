// Code generated by go-mockgen 1.3.3; DO NOT EDIT.

package server

import "sync"

// MockCoder is a mock implementation of the Coder interface (from the
// package github.com/Karitham/shurl/server) used for unit testing.
type MockCoder struct {
	// DecodeStringFunc is an instance of a mock function object controlling
	// the behavior of the method DecodeString.
	DecodeStringFunc *CoderDecodeStringFunc
	// EncodeToStringFunc is an instance of a mock function object
	// controlling the behavior of the method EncodeToString.
	EncodeToStringFunc *CoderEncodeToStringFunc
}

// NewMockCoder creates a new mock of the Coder interface. All methods
// return zero values for all results, unless overwritten.
func NewMockCoder() *MockCoder {
	return &MockCoder{
		DecodeStringFunc: &CoderDecodeStringFunc{
			defaultHook: func(string) (r0 []byte, r1 error) {
				return
			},
		},
		EncodeToStringFunc: &CoderEncodeToStringFunc{
			defaultHook: func([]byte) (r0 string) {
				return
			},
		},
	}
}

// NewStrictMockCoder creates a new mock of the Coder interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockCoder() *MockCoder {
	return &MockCoder{
		DecodeStringFunc: &CoderDecodeStringFunc{
			defaultHook: func(string) ([]byte, error) {
				panic("unexpected invocation of MockCoder.DecodeString")
			},
		},
		EncodeToStringFunc: &CoderEncodeToStringFunc{
			defaultHook: func([]byte) string {
				panic("unexpected invocation of MockCoder.EncodeToString")
			},
		},
	}
}

// NewMockCoderFrom creates a new mock of the MockCoder interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockCoderFrom(i Coder) *MockCoder {
	return &MockCoder{
		DecodeStringFunc: &CoderDecodeStringFunc{
			defaultHook: i.DecodeString,
		},
		EncodeToStringFunc: &CoderEncodeToStringFunc{
			defaultHook: i.EncodeToString,
		},
	}
}

// CoderDecodeStringFunc describes the behavior when the DecodeString method
// of the parent MockCoder instance is invoked.
type CoderDecodeStringFunc struct {
	defaultHook func(string) ([]byte, error)
	hooks       []func(string) ([]byte, error)
	history     []CoderDecodeStringFuncCall
	mutex       sync.Mutex
}

// DecodeString delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockCoder) DecodeString(v0 string) ([]byte, error) {
	r0, r1 := m.DecodeStringFunc.nextHook()(v0)
	m.DecodeStringFunc.appendCall(CoderDecodeStringFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the DecodeString method
// of the parent MockCoder instance is invoked and the hook queue is empty.
func (f *CoderDecodeStringFunc) SetDefaultHook(hook func(string) ([]byte, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DecodeString method of the parent MockCoder instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *CoderDecodeStringFunc) PushHook(hook func(string) ([]byte, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *CoderDecodeStringFunc) SetDefaultReturn(r0 []byte, r1 error) {
	f.SetDefaultHook(func(string) ([]byte, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *CoderDecodeStringFunc) PushReturn(r0 []byte, r1 error) {
	f.PushHook(func(string) ([]byte, error) {
		return r0, r1
	})
}

func (f *CoderDecodeStringFunc) nextHook() func(string) ([]byte, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CoderDecodeStringFunc) appendCall(r0 CoderDecodeStringFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CoderDecodeStringFuncCall objects
// describing the invocations of this function.
func (f *CoderDecodeStringFunc) History() []CoderDecodeStringFuncCall {
	f.mutex.Lock()
	history := make([]CoderDecodeStringFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CoderDecodeStringFuncCall is an object that describes an invocation of
// method DecodeString on an instance of MockCoder.
type CoderDecodeStringFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []byte
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c CoderDecodeStringFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CoderDecodeStringFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// CoderEncodeToStringFunc describes the behavior when the EncodeToString
// method of the parent MockCoder instance is invoked.
type CoderEncodeToStringFunc struct {
	defaultHook func([]byte) string
	hooks       []func([]byte) string
	history     []CoderEncodeToStringFuncCall
	mutex       sync.Mutex
}

// EncodeToString delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockCoder) EncodeToString(v0 []byte) string {
	r0 := m.EncodeToStringFunc.nextHook()(v0)
	m.EncodeToStringFunc.appendCall(CoderEncodeToStringFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the EncodeToString
// method of the parent MockCoder instance is invoked and the hook queue is
// empty.
func (f *CoderEncodeToStringFunc) SetDefaultHook(hook func([]byte) string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// EncodeToString method of the parent MockCoder instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *CoderEncodeToStringFunc) PushHook(hook func([]byte) string) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *CoderEncodeToStringFunc) SetDefaultReturn(r0 string) {
	f.SetDefaultHook(func([]byte) string {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *CoderEncodeToStringFunc) PushReturn(r0 string) {
	f.PushHook(func([]byte) string {
		return r0
	})
}

func (f *CoderEncodeToStringFunc) nextHook() func([]byte) string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CoderEncodeToStringFunc) appendCall(r0 CoderEncodeToStringFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CoderEncodeToStringFuncCall objects
// describing the invocations of this function.
func (f *CoderEncodeToStringFunc) History() []CoderEncodeToStringFuncCall {
	f.mutex.Lock()
	history := make([]CoderEncodeToStringFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CoderEncodeToStringFuncCall is an object that describes an invocation of
// method EncodeToString on an instance of MockCoder.
type CoderEncodeToStringFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 []byte
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c CoderEncodeToStringFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CoderEncodeToStringFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
