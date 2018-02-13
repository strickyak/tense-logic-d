package D

type Story []bool

type Prop interface {
	Eval(Story) bool
}

type Rains struct{}
type Tomorrow struct {
	A Prop
}
type Eventually struct {
	A Prop
}
type Forevermore struct {
	A Prop
}
type Not struct {
	A Prop
}
type And struct {
	A, B Prop
}
type Or struct {
	A, B Prop
}
type Implies struct {
	A, B Prop
}

func (o Rains) Eval(s Story) bool {
	return s[0]
}
func (o Tomorrow) Eval(s Story) bool {
	if len(s) > 1 {
		return o.A.Eval(s[1:])
	}
	// Repeat final day.
	return o.A.Eval(s)
}
func (o Eventually) Eval(s Story) bool {
	for i, _ := range s {
		if o.A.Eval(s[i:]) == true {
			return true
		}
	}
	return false
}
func (o Forevermore) Eval(s Story) bool {
	for i, _ := range s {
		if o.A.Eval(s[i:]) == false {
			return false
		}
	}
	return true
}
func (o Not) Eval(s Story) bool {
	return !o.A.Eval(s)
}
func (o And) Eval(s Story) bool {
	return o.A.Eval(s) && o.B.Eval(s)
}
func (o Or) Eval(s Story) bool {
	return o.A.Eval(s) || o.B.Eval(s)
}
func (o Implies) Eval(s Story) bool {
	return !o.A.Eval(s) || o.B.Eval(s)
}
