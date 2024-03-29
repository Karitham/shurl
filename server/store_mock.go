// Code generated by go-mockgen 1.1.2; DO NOT EDIT.

package server

import "sync"

// MockStore is a mock implementation of the Store interface (from the
// package github.com/Karitham/shurl/server) used for unit testing.
type MockStore struct {
	// GetFunc is an instance of a mock function object controlling the
	// behavior of the method Get.
	GetFunc *StoreGetFunc
	// SetFunc is an instance of a mock function object controlling the
	// behavior of the method Set.
	SetFunc *StoreSetFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		GetFunc: &StoreGetFunc{
			defaultHook: func([]byte) ([]byte, error) {
				return nil, nil
			},
		},
		SetFunc: &StoreSetFunc{
			defaultHook: func([]byte, []byte) error {
				return nil
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i Store) *MockStore {
	return &MockStore{
		GetFunc: &StoreGetFunc{
			defaultHook: i.Get,
		},
		SetFunc: &StoreSetFunc{
			defaultHook: i.Set,
		},
	}
}

// StoreGetFunc describes the behavior when the Get method of the parent
// MockStore instance is invoked.
type StoreGetFunc struct {
	defaultHook func([]byte) ([]byte, error)
	hooks       []func([]byte) ([]byte, error)
	history     []StoreGetFuncCall
	mutex       sync.Mutex
}

// Get delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Get(v0 []byte) ([]byte, error) {
	r0, r1 := m.GetFunc.nextHook()(v0)
	m.GetFunc.appendCall(StoreGetFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Get method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreGetFunc) SetDefaultHook(hook func([]byte) ([]byte, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Get method of the parent MockStore instance invokes the hook at the front
// of the queue and discards it. After the queue is empty, the default hook
// function is invoked for any future action.
func (f *StoreGetFunc) PushHook(hook func([]byte) ([]byte, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreGetFunc) SetDefaultReturn(r0 []byte, r1 error) {
	f.SetDefaultHook(func([]byte) ([]byte, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreGetFunc) PushReturn(r0 []byte, r1 error) {
	f.PushHook(func([]byte) ([]byte, error) {
		return r0, r1
	})
}

func (f *StoreGetFunc) nextHook() func([]byte) ([]byte, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetFunc) appendCall(r0 StoreGetFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetFuncCall objects describing the
// invocations of this function.
func (f *StoreGetFunc) History() []StoreGetFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetFuncCall is an object that describes an invocation of method Get
// on an instance of MockStore.
type StoreGetFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 []byte
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []byte
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreSetFunc describes the behavior when the Set method of the parent
// MockStore instance is invoked.
type StoreSetFunc struct {
	defaultHook func([]byte, []byte) error
	hooks       []func([]byte, []byte) error
	history     []StoreSetFuncCall
	mutex       sync.Mutex
}

// Set delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Set(v0 []byte, v1 []byte) error {
	r0 := m.SetFunc.nextHook()(v0, v1)
	m.SetFunc.appendCall(StoreSetFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Set method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreSetFunc) SetDefaultHook(hook func([]byte, []byte) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Set method of the parent MockStore instance invokes the hook at the front
// of the queue and discards it. After the queue is empty, the default hook
// function is invoked for any future action.
func (f *StoreSetFunc) PushHook(hook func([]byte, []byte) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreSetFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func([]byte, []byte) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreSetFunc) PushReturn(r0 error) {
	f.PushHook(func([]byte, []byte) error {
		return r0
	})
}

func (f *StoreSetFunc) nextHook() func([]byte, []byte) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreSetFunc) appendCall(r0 StoreSetFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreSetFuncCall objects describing the
// invocations of this function.
func (f *StoreSetFunc) History() []StoreSetFuncCall {
	f.mutex.Lock()
	history := make([]StoreSetFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreSetFuncCall is an object that describes an invocation of method Set
// on an instance of MockStore.
type StoreSetFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 []byte
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 []byte
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreSetFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreSetFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
