package jesus

import "time"

type UProfile struct {
	Id          int64
	Phone       string
	Balance     int64
	UpdatedAt   time.Time
	CreatedAT   time.Time
	Withdrawals []Withdrawal
	Deposits    []Deposit
}
type Withdrawal struct {
	Id         int64
	Amount     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UProfileId int64
}

type Deposit struct {
	Id         int64
	Amount     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UProfileId int64
}


func (u *UProfile)PrepareDeposit(s *Inbox){
	d:=Deposit{}
	d.Amount=s.DepositedAmount()
	_=append(u.Deposits, d)
}

