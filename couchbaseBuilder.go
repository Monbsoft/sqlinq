package sqlinq

import (
	"bytes"
)

// CouchbaseBuilder represents couchbase builder
type CouchbaseBuilder struct {
	sql *bytes.Buffer
	params []interface{}
	err    error
}

// NewCouchbaseBuilder initializes couchbase builder
func NewCouchbaseBuilder() *CouchbaseBuilder {
	return &CouchbaseBuilder{sql: &bytes.Buffer{}}
}

// Visit visits select statement
func (b *CouchbaseBuilder) Visit(s *SelectStatement) {
	if len(s.selectClause) == 0 {
		
	}

}

// Build builds query and parameters
func (b *CouchbaseBuilder) Build() (query string, params []interface{}, err error) {
	return "", nil, nil
}
