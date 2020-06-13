package sqlinq

import "testing"

func TestSelectQuery(t *testing.T) {
	builder := NewCouchbaseBuilder()
	Select("aa", "112").
		From("bucket").
		ToSql(builder)

}
