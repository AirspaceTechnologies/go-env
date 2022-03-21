package validators

type Func[T any] func(T) error

func (vf Func[T]) Wrap(f Func[T]) Func[T] {
	if vf == nil {
		return f
	}

	return func(v T) error {
		if err := vf(v); err != nil {
			return err
		}

		return f(v)
	}
}
