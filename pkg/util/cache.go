package util

type Cache[T comparable, K any] interface {
	Get(T) (K, error)
	MustGet(T) K
}

type cache[T comparable, K any] struct {
	data map[T]K
	fn   func(T) (K, error)
}

func NewCache[T comparable, K any](fn func(T) (K, error)) Cache[T, K] {
	return &cache[T, K]{
		data: make(map[T]K),
		fn:   fn,
	}
}

func (c *cache[T, K]) Get(key T) (K, error) {
	v, ok := c.data[key]
	if ok {
		return v, nil
	}
	v, err := c.fn(key)
	if err != nil {
		return v, err
	}
	c.data[key] = v
	return v, nil
}

func (c *cache[T, K]) MustGet(key T) K {
	v, ok := c.data[key]
	if ok {
		return v
	}
	v, err := c.fn(key)
	if err != nil {
		panic(err)
	}
	c.data[key] = v
	return v
}
