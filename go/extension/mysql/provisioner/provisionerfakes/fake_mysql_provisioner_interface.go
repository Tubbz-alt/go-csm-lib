// This file was generated by counterfeiter
package provisionerfakes

import (
	"database/sql"
	"sync"

	"github.com/hpcloud/sidecar-extensions/go/extensions/mysql/config"
	"github.com/hpcloud/sidecar-extensions/go/extensions/mysql/provisioner"
)

type FakeMysqlProvisionerInterface struct {
	ConnectStub        func(config.MysqlDriverConfig) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		arg1 config.MysqlDriverConfig
	}
	connectReturns struct {
		result1 error
	}
	IsDatabaseCreatedStub        func(string) (bool, error)
	isDatabaseCreatedMutex       sync.RWMutex
	isDatabaseCreatedArgsForCall []struct {
		arg1 string
	}
	isDatabaseCreatedReturns struct {
		result1 bool
		result2 error
	}
	IsUserCreatedStub        func(string) (bool, error)
	isUserCreatedMutex       sync.RWMutex
	isUserCreatedArgsForCall []struct {
		arg1 string
	}
	isUserCreatedReturns struct {
		result1 bool
		result2 error
	}
	CreateDatabaseStub        func(string) error
	createDatabaseMutex       sync.RWMutex
	createDatabaseArgsForCall []struct {
		arg1 string
	}
	createDatabaseReturns struct {
		result1 error
	}
	DeleteDatabaseStub        func(string) error
	deleteDatabaseMutex       sync.RWMutex
	deleteDatabaseArgsForCall []struct {
		arg1 string
	}
	deleteDatabaseReturns struct {
		result1 error
	}
	QueryStub        func(string, ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	CreateUserStub        func(string, string, string) error
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 error
	}
	DeleteUserStub        func(string) error
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
	}
	deleteUserReturns struct {
		result1 error
	}
}

func (fake *FakeMysqlProvisionerInterface) Connect(arg1 config.MysqlDriverConfig) error {
	fake.connectMutex.Lock()
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		arg1 config.MysqlDriverConfig
	}{arg1})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(arg1)
	} else {
		return fake.connectReturns.result1
	}
}

func (fake *FakeMysqlProvisionerInterface) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) ConnectArgsForCall(i int) config.MysqlDriverConfig {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return fake.connectArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) ConnectReturns(result1 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMysqlProvisionerInterface) IsDatabaseCreated(arg1 string) (bool, error) {
	fake.isDatabaseCreatedMutex.Lock()
	fake.isDatabaseCreatedArgsForCall = append(fake.isDatabaseCreatedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.isDatabaseCreatedMutex.Unlock()
	if fake.IsDatabaseCreatedStub != nil {
		return fake.IsDatabaseCreatedStub(arg1)
	} else {
		return fake.isDatabaseCreatedReturns.result1, fake.isDatabaseCreatedReturns.result2
	}
}

func (fake *FakeMysqlProvisionerInterface) IsDatabaseCreatedCallCount() int {
	fake.isDatabaseCreatedMutex.RLock()
	defer fake.isDatabaseCreatedMutex.RUnlock()
	return len(fake.isDatabaseCreatedArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) IsDatabaseCreatedArgsForCall(i int) string {
	fake.isDatabaseCreatedMutex.RLock()
	defer fake.isDatabaseCreatedMutex.RUnlock()
	return fake.isDatabaseCreatedArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) IsDatabaseCreatedReturns(result1 bool, result2 error) {
	fake.IsDatabaseCreatedStub = nil
	fake.isDatabaseCreatedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMysqlProvisionerInterface) IsUserCreated(arg1 string) (bool, error) {
	fake.isUserCreatedMutex.Lock()
	fake.isUserCreatedArgsForCall = append(fake.isUserCreatedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.isUserCreatedMutex.Unlock()
	if fake.IsUserCreatedStub != nil {
		return fake.IsUserCreatedStub(arg1)
	} else {
		return fake.isUserCreatedReturns.result1, fake.isUserCreatedReturns.result2
	}
}

func (fake *FakeMysqlProvisionerInterface) IsUserCreatedCallCount() int {
	fake.isUserCreatedMutex.RLock()
	defer fake.isUserCreatedMutex.RUnlock()
	return len(fake.isUserCreatedArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) IsUserCreatedArgsForCall(i int) string {
	fake.isUserCreatedMutex.RLock()
	defer fake.isUserCreatedMutex.RUnlock()
	return fake.isUserCreatedArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) IsUserCreatedReturns(result1 bool, result2 error) {
	fake.IsUserCreatedStub = nil
	fake.isUserCreatedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMysqlProvisionerInterface) CreateDatabase(arg1 string) error {
	fake.createDatabaseMutex.Lock()
	fake.createDatabaseArgsForCall = append(fake.createDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.createDatabaseMutex.Unlock()
	if fake.CreateDatabaseStub != nil {
		return fake.CreateDatabaseStub(arg1)
	} else {
		return fake.createDatabaseReturns.result1
	}
}

func (fake *FakeMysqlProvisionerInterface) CreateDatabaseCallCount() int {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return len(fake.createDatabaseArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) CreateDatabaseArgsForCall(i int) string {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return fake.createDatabaseArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) CreateDatabaseReturns(result1 error) {
	fake.CreateDatabaseStub = nil
	fake.createDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMysqlProvisionerInterface) DeleteDatabase(arg1 string) error {
	fake.deleteDatabaseMutex.Lock()
	fake.deleteDatabaseArgsForCall = append(fake.deleteDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.deleteDatabaseMutex.Unlock()
	if fake.DeleteDatabaseStub != nil {
		return fake.DeleteDatabaseStub(arg1)
	} else {
		return fake.deleteDatabaseReturns.result1
	}
}

func (fake *FakeMysqlProvisionerInterface) DeleteDatabaseCallCount() int {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return len(fake.deleteDatabaseArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) DeleteDatabaseArgsForCall(i int) string {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return fake.deleteDatabaseArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) DeleteDatabaseReturns(result1 error) {
	fake.DeleteDatabaseStub = nil
	fake.deleteDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMysqlProvisionerInterface) Query(arg1 string, arg2 ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(arg1, arg2...)
	} else {
		return fake.queryReturns.result1, fake.queryReturns.result2
	}
}

func (fake *FakeMysqlProvisionerInterface) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].arg1, fake.queryArgsForCall[i].arg2
}

func (fake *FakeMysqlProvisionerInterface) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeMysqlProvisionerInterface) CreateUser(arg1 string, arg2 string, arg3 string) error {
	fake.createUserMutex.Lock()
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	} else {
		return fake.createUserReturns.result1
	}
}

func (fake *FakeMysqlProvisionerInterface) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return fake.createUserArgsForCall[i].arg1, fake.createUserArgsForCall[i].arg2, fake.createUserArgsForCall[i].arg3
}

func (fake *FakeMysqlProvisionerInterface) CreateUserReturns(result1 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMysqlProvisionerInterface) DeleteUser(arg1 string) error {
	fake.deleteUserMutex.Lock()
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1)
	} else {
		return fake.deleteUserReturns.result1
	}
}

func (fake *FakeMysqlProvisionerInterface) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeMysqlProvisionerInterface) DeleteUserArgsForCall(i int) string {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return fake.deleteUserArgsForCall[i].arg1
}

func (fake *FakeMysqlProvisionerInterface) DeleteUserReturns(result1 error) {
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 error
	}{result1}
}

var _ provisioner.MysqlProvisionerInterface = new(FakeMysqlProvisionerInterface)
