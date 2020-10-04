package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	amount = 10
	return &Dollar{amount: amount}
}

func (d *Dollar) Times(multiplier int) {

}
