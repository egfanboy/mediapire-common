package router

import "fmt"

type QueryParam struct {
	Name     string
	Regex    *string
	Required bool
}

func buildQuery(qs []QueryParam) (result []string) {
	for _, q := range qs {
		result = append(result, q.Name)

		if q.Regex != nil {
			result = append(result, fmt.Sprintf("{%s:%s}", q.Name, *q.Regex))
		} else {
			result = append(result, fmt.Sprintf("{%s}", q.Name))
		}
	}

	return
}
