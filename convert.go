package ldb

type ConvertCtx struct {
	convertFuncs ConvertFuncMap
	valBox       ConvertValBoxMap
}

type Convert struct {
	name        string
	val         any
	convertFunc ConvertFunc
}

type ConvertFunc func(o any) any
type ConvertFuncMap map[string]ConvertFunc
type ConvertValBoxMap map[string]any

func (c ConvertCtx) Init() ConvertCtx {
	c.convertFuncs = ConvertFuncMap{}
	c.valBox = ConvertValBoxMap{}
	return c
}
func (c *ConvertCtx) Add(v Convert) {
	name := v.name
	c.valBox[name] = v.val
	c.convertFuncs[name] = v.convertFunc
}

func (c ConvertCtx) Get(name string) (any, ConvertFunc) {
	vb, ok := c.valBox[name]
	if !ok {
		return nil, nil
	}
	f := c.convertFuncs[name]
	return vb, f
}

func ConvertRegister[T any](name string, f func(v T) any) Convert {
	var t = new(T)
	return Convert{
		name: name,
		val:  t,
		convertFunc: func(val any) any {
			return f(*val.(*T))
		},
	}
}
