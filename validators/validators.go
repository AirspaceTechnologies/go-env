package validators

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type RangeOpts struct {
	MinExclusive bool
	MaxExclusive bool
}

func RangeWithOpts[T constraints.Ordered](min, max T, opts RangeOpts) Func[T] {
	return func(v T) error {
		if err := satisfyCompareOp(compareOpGreater, v, min, !opts.MinExclusive); err != nil {
			return err
		}

		return satisfyCompareOp(compareOpLess, v, max, !opts.MaxExclusive)
	}
}

func Range[T constraints.Ordered](min, max T) Func[T] {
	return RangeWithOpts(min, max, RangeOpts{})
}

func Greater[T constraints.Ordered](than T, orEqual bool) Func[T] {
	return func(v T) error {
		return satisfyCompareOp[T](compareOpGreater, v, than, orEqual)
	}
}

func Less[T constraints.Ordered](than T, orEqual bool) Func[T] {
	return func(v T) error {
		return satisfyCompareOp[T](compareOpLess, v, than, orEqual)
	}
}

func In[T comparable](allowed ...T) Func[T] {
	return func(v T) error {
		for _, allowedVal := range allowed {
			if v == allowedVal {
				return nil
			}
		}

		return fmt.Errorf("%v is not in allowed set", v)
	}
}

func NotIn[T comparable](disallowed ...T) Func[T] {
	return func(v T) error {
		for _, disallowedVal := range disallowed {
			if v == disallowedVal {
				return fmt.Errorf("%v is a non allowed value", v)
			}
		}

		return nil
	}
}
