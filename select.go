package sqlinq

type SelectStatement struct {
	selectClause []string
	fromClause   []string
	limit        int
}

// Initializes Select statement with Select clause.
func Select(columns ...string) *SelectStatement {
	s := &SelectStatement{}
	for _, c := range columns {
		s.selectClause = append(s.selectClause, c)
	}
	return s
}

// Sets From clause.
func (s *SelectStatement) From(data string) *SelectStatement {
	return s
}

// Sets Limit clause
func (s *SelectStatement) Limit(limit int) *SelectStatement {
	s.limit = limit
	return s
}

func (s *SelectStatement) ToSql(b Builder) (query string, params []interface{}, err error) {
	b.Visit(s)
	return b.Build()
}
