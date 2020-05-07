package cng

type IteratorFunc func(v interface{}) error

type iterator struct {
	data  []interface{}
	index int
}

func newIterator() *iterator {
	return &iterator{
		data:  nil,
		index: 0,
	}
}

func (i *iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *iterator) Next() interface{} {
	defer func() {
		i.index++
	}()
	if i.index < len(i.data) {
		return i.data[i.index]
	}

	return nil
}

func (i *iterator) Reset() {
	i.index = 0
}

func (i *iterator) Add(v interface{}) {
	i.data = append(i.data, v)
}

func (i *iterator) Size() int {
	return len(i.data)
}

func (i *iterator) Iterator(f IteratorFunc) error {
	i.Reset()
	for i.HasNext() {
		if err := f(i.Next()); err != nil {
			return err
		}
	}
	return nil
}
