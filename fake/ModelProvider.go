// This file was generated by counterfeiter
package fake

import (
	"sync"

	"github.com/phogolabs/oak/model"
)

type ModelProvider struct {
	TablesStub        func(schema string) ([]string, error)
	tablesMutex       sync.RWMutex
	tablesArgsForCall []struct {
		schema string
	}
	tablesReturns struct {
		result1 []string
		result2 error
	}
	SchemaStub        func(schema string, tables ...string) (*model.Schema, error)
	schemaMutex       sync.RWMutex
	schemaArgsForCall []struct {
		schema string
		tables []string
	}
	schemaReturns struct {
		result1 *model.Schema
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ModelProvider) Tables(schema string) ([]string, error) {
	fake.tablesMutex.Lock()
	fake.tablesArgsForCall = append(fake.tablesArgsForCall, struct {
		schema string
	}{schema})
	fake.recordInvocation("Tables", []interface{}{schema})
	fake.tablesMutex.Unlock()
	if fake.TablesStub != nil {
		return fake.TablesStub(schema)
	}
	return fake.tablesReturns.result1, fake.tablesReturns.result2
}

func (fake *ModelProvider) TablesCallCount() int {
	fake.tablesMutex.RLock()
	defer fake.tablesMutex.RUnlock()
	return len(fake.tablesArgsForCall)
}

func (fake *ModelProvider) TablesArgsForCall(i int) string {
	fake.tablesMutex.RLock()
	defer fake.tablesMutex.RUnlock()
	return fake.tablesArgsForCall[i].schema
}

func (fake *ModelProvider) TablesReturns(result1 []string, result2 error) {
	fake.TablesStub = nil
	fake.tablesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *ModelProvider) Schema(schema string, tables ...string) (*model.Schema, error) {
	fake.schemaMutex.Lock()
	fake.schemaArgsForCall = append(fake.schemaArgsForCall, struct {
		schema string
		tables []string
	}{schema, tables})
	fake.recordInvocation("Schema", []interface{}{schema, tables})
	fake.schemaMutex.Unlock()
	if fake.SchemaStub != nil {
		return fake.SchemaStub(schema, tables...)
	}
	return fake.schemaReturns.result1, fake.schemaReturns.result2
}

func (fake *ModelProvider) SchemaCallCount() int {
	fake.schemaMutex.RLock()
	defer fake.schemaMutex.RUnlock()
	return len(fake.schemaArgsForCall)
}

func (fake *ModelProvider) SchemaArgsForCall(i int) (string, []string) {
	fake.schemaMutex.RLock()
	defer fake.schemaMutex.RUnlock()
	return fake.schemaArgsForCall[i].schema, fake.schemaArgsForCall[i].tables
}

func (fake *ModelProvider) SchemaReturns(result1 *model.Schema, result2 error) {
	fake.SchemaStub = nil
	fake.schemaReturns = struct {
		result1 *model.Schema
		result2 error
	}{result1, result2}
}

func (fake *ModelProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tablesMutex.RLock()
	defer fake.tablesMutex.RUnlock()
	fake.schemaMutex.RLock()
	defer fake.schemaMutex.RUnlock()
	return fake.invocations
}

func (fake *ModelProvider) recordInvocation(key string, args []interface{}) {
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

var _ model.Provider = new(ModelProvider)