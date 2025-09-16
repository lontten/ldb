package ldb

func (w *WhereBuilder) And(wb *WhereBuilder, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	if wb == nil {
		return w
	}
	if wb.Invalid() {
		return w
	}
	if wb.err != nil {
		w.err = wb.err
		return w
	}
	w.andWheres = append(w.andWheres, *wb)
	return w
}

func (w *WhereBuilder) Or(wb *WhereBuilder, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	if wb == nil {
		return w
	}
	if wb.Invalid() {
		return w
	}
	if wb.err != nil {
		w.err = wb.err
		return w
	}
	w.wheres = append(w.wheres, *wb)
	return w
}

func (w *WhereBuilder) Not(condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	if w.Invalid() {
		return w
	}
	w.not = true
	return w
}
