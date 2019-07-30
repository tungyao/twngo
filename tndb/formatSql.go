package tndb

type operation interface {
}

type _SQL struct {
}

func ConvertMapString(m map[string]string) string {
	sl := ""
	for k, v := range m {
		sl += k + "='" + v + "'" + ","
	}
	return sl[:len(sl)-1]
}
func ConvertArrayString(arr []string) string {
	sl := ""
	for _, v := range arr {
		sl += "'" + v + "',"
	}
	return "(" + sl[:len(sl)-1] + ")"
}

func keyForInsertOrUpdate(k interface{}) string {
	sl := ""
	switch k.(type) {
	case map[string]string:
		sl = " set " + ConvertMapString(k.(map[string]string))
	case []string:
		sl = " values " + ConvertArrayString(k.([]string))
	default:
		sl = ""
	}
	return sl
}
