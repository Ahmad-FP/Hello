package mixed

type row struct{ Schema, Table, Kind string }
type tableInfo struct{ Name, Kind string }
type schemaInfo struct {
	Name   string
	Tables []tableInfo
}
func group(rows []row) []schemaInfo{
	var out []schemaInfo
	inside := map[string] int{}
	for  _, r :=  range rows {
		if _, seen := inside[r.Schema]; !seen {
			out = append(out, schemaInfo{Name: r.Schema})
			i := len(out)-1
			inside[r.Schema] = i
			out[i].Tables = append(out[i].Tables, tableInfo{r.Table, r.Kind})
		}else{
			i := len(out)-1
			out[i].Tables = append(out[i].Tables, tableInfo{r.Table, r.Kind})
		}
	}
	return out
}

func Truncate(s string, max int) string{
	if len([]rune(s)) <= max {
		return s
	} else{
		return string([]rune(s)[:max]) + "\u2026"
	}
}