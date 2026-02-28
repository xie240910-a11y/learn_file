package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FocusModel = (*customFocusModel)(nil)

type (
	// FocusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFocusModel.
	FocusModel interface {
		focusModel
		withSession(session sqlx.Session) FocusModel
	}

	customFocusModel struct {
		*defaultFocusModel
	}
)

// NewFocusModel returns a model for the database table.
func NewFocusModel(conn sqlx.SqlConn) FocusModel {
	return &customFocusModel{
		defaultFocusModel: newFocusModel(conn),
	}
}

func (m *customFocusModel) withSession(session sqlx.Session) FocusModel {
	return NewFocusModel(sqlx.NewSqlConnFromSession(session))
}
