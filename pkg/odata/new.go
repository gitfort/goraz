package odata

func Select(fields ...string) Query {
	return make(Query).Select(fields...)
}

func AndFilter(field string, operate Operate, value interface{}) Query {
	return make(Query).AndFilter(field, operate, value)
}

func OrFilter(field string, operate Operate, value interface{}) Query {
	return make(Query).OrFilter(field, operate, value)
}
