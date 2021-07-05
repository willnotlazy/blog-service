package convert

import "strconv"

type Strto string

func (s Strto) 	String() string {
	return string(s)
}

func (s Strto) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s Strto) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s Strto) UInt32() (uint32, error){
	v, err := s.Int()
	return uint32(v), err
}

func (s Strto) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}


