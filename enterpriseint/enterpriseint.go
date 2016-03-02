package enterpriseint

import (
	"math"
	"sync/atomic"
)

type EnterpriseInt int64

const Zero = EnterpriseInt(0)

func (i EnterpriseInt) Abs() EnterpriseInt {
	if i.IsNegative() {
		return 0 - i
	}
	return i
}

func (i EnterpriseInt) IsPositive() bool {
	return i > 0
}

func (i EnterpriseInt) IsNegative() bool {
	return i < 0
}

func (i EnterpriseInt) IsEven() bool {
	return i&1 == 0
}

func (i EnterpriseInt) IsOdd() bool {
	return !i.IsEven()
}

func (i EnterpriseInt) IsPrime() bool {
	if i == Zero {
		return false
	}
	a := float64(i.Abs())
	if a == 2 {
		return true
	}
	if i.IsEven() {
		return false
	}

	sqrt := int64(math.Sqrt(a))
	for j := int64(2); j <= sqrt+1; j++ {
		if math.Remainder(a, float64(j)) == 0 {
			return false
		}
	}
	return true
}

func (i *EnterpriseInt) AtomicIncrement(delta int64) EnterpriseInt {
	return EnterpriseInt(atomic.AddInt64((*int64)(i), delta))
}

func (i EnterpriseInt) ForEach(f func(EnterpriseInt) error) error {
	var err error
	for k := Zero; k < i; k++ {
		if err = f(k); err != nil {
			return err
		}
	}
	return nil
}

func (i EnterpriseInt) Read(p []byte) (int, error) {
	for k := 0; k < len(p); k++ {
		p[k] = byte(i)
	}
	return len(p), nil
}
