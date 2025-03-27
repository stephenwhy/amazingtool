package find

import "github.com/stephenwhy/amazingtool/convert"

func LookupNames(ids string, nameMap map[int64]string) []string {
	names := make([]string, 0)
	idList := convert.StrToInt64Slice(ids, ",")
	if len(idList) == 0 {
		return names
	}
	for _, id := range idList {
		if name, ok := nameMap[id]; ok {
			names = append(names, name)
		}
	}
	return names
}
