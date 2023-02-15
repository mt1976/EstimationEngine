package das

func ID(in string) string {
	return "'" + in + "'"
}

func ISNULL(in string) string {
	return " datalength(" + in + ") = 0"
}
