package database

import (
	"errors"

	"github.com/thernande/app-pedidos-fondos-backend/pkg/appLogs"
	"gorm.io/gorm"
)

type Db struct {
	Instance *gorm.DB
	Err      error
	Log      appLogs.Logger
}

func NewDb(Instance *gorm.DB) *Db {
	looger := appLogs.Logger{}
	looger.Init()
	return &Db{
		Instance: Instance,
		Log:      looger,
	}
}

func (m *Db) GetDataByQuery(query string, data interface{}) {
	if m.Instance == nil {
		m.Err = errors.New("Instance connection is not initialized")
		m.Log.ErrorLogPrint(m.Err.Error())
		return
	}
	result := m.Instance.Raw(query).Scan(data)
	if result.Error != nil {
		m.Err = result.Error
		m.Log.ErrorLogPrint(result.Error.Error())
		return
	}
	m.Err = nil
}

func (m *Db) GetDataByQueryWithConditions(query string, conditions []interface{}, data interface{}) {
	if m.Instance == nil {
		m.Err = errors.New("instance connection is not initialized")
		m.Log.ErrorLogPrint(m.Err.Error())
		return
	}
	// fmt.Printf("%#v \n %s \n", data, query)
	if res := m.Instance.Raw(query, conditions...).Scan(data); res.Error != nil {
		m.Log.ErrorLogPrint(res.Error.Error())
		m.Err = res.Error
		return
	}
	m.Err = nil
}

func (m *Db) ExecByQuery(query string) error {
	if m.Instance == nil {
		m.Err = errors.New("instance connection is not initialized")
		m.Log.ErrorLogPrint(m.Err.Error())
		return m.Err
	}
	if res := m.Instance.Exec(query); res.Error != nil {
		m.Log.ErrorLogPrint(res.Error.Error())
		m.Err = res.Error
		return res.Error
	}
	m.Err = nil
	return nil
}

func (m *Db) InitTransaction() {
	if m.Instance == nil {
		m.Err = errors.New("instance connection is not initialized")
		m.Log.ErrorLogPrint(m.Err.Error())
		return
	}
	res := m.Instance.Begin()
	if res.Error != nil {
		m.Log.ErrorLogPrint(res.Error.Error())
		m.Err = res.Error
		return
	}
	m.Instance = res
	m.Err = nil
}

func (m *Db) CommitTransaction() {
	if m.Instance == nil {
		m.Err = errors.New("instance connection is not initialized")
		m.Log.ErrorLogPrint(m.Err.Error())
		return
	}
	if res := m.Instance.Commit(); res.Error != nil {
		m.Log.ErrorLogPrint(res.Error.Error())
		m.Err = res.Error
		m.Instance.Rollback()
		return
	}
	m.Err = nil
}
