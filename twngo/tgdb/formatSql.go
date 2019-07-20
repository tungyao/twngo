package tgdb

type operation interface {

}

type _SQL struct {
	
}

func ConvertMapString(m map[string]string)string  {
	sl :=""
	for k,v :=range m{
		sl += k+"='"+v+"'"+","
	}
	return sl[:len(sl)-1]
}
func ConvertArrayStrign(arr []string) string {
	sl :=""
	for _,v :=range arr{
		sl += "'"+v + "',"
	}
	return "("+sl[:len(sl)-1]+")"
}

