package validators

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

const (
	compareOpGreater compareOp = iota
	compareOpLess
)

type compareOp int

func (op compareOp) String() string {
	switch op {
	case compareOpLess:
		return "<"
	case compareOpGreater:
		return ">"
	}

	return strconv.FormatInt(int64(op), 10)
}

func satisfyCompareOp[T constraints.Ordered](op compareOp, v, than T, orEqual bool) error {
	switch op {
	case compareOpGreater:
		if v > than {
			return nil
		}
	case compareOpLess:
		if v < than {
			return nil
		}
	}

	if orEqual && v == than {
		return nil
	}

	var conditionString string
	if orEqual {
		conditionString = fmt.Sprintf("%v%v", op, "=")
	} else {
		conditionString = op.String()
	}

	return fmt.Errorf("the following condition failed: %v %v %v", v, conditionString, than)
}
