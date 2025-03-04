package router

type QueryParam struct {
	Name     string
	Regex    *string
	Required bool
}

func (p QueryParam) String() string {
	return p.Name
}
