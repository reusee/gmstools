package slicetable

func (t *Table[T]) Update(slice *[]T) {
	t.ptr.Store(slice)
}
