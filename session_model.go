package xorm

// Statement 数据库操作
// go-component/orm 定制
func (session *Session) Statement() *Statement {
	return &session.statement
}
