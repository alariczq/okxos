package client

import (
	"strconv"
)

type Integer int64

func (n Integer) Int() int64 {
	return int64(n)
}

func (n Integer) String() string {
	return strconv.FormatInt(int64(n), 10)
}

func (n *Integer) UnmarshalJSON(b []byte) error {
	num := string(b)
	if num[0] == '"' {
		num = num[1 : len(num)-1]
	}

	i, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return err
	}

	*n = Integer(i)

	return nil
}
