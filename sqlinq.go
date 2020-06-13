package sqlinq

// Represents a sql statement (SELECT / INSERT / UPDATE and other )
type SqlStatement interface {
	ToSql(b Builder) (query string, params []interface{}, err error)
}

// Builder represents a query builder
type Builder interface {
	Build() (query string, params []interface{}, err error)
	Visit(s *SelectStatement)
}
