package jesus

import "time"

type Inbox struct {
	ID                int       `gorm:"column:ID;primary_key:yes"`
	Text              string    `sql:"type:text" gorm:"column:Text"`
	ReceivingDateTime time.Time `gorm:"column:ReceivingDateTime"`
	UpdatedInDB       time.Time `gorm:"column:UpdatedInDB"`
	SenderNumber      string    `sql:"type:varchar(20);not null;DEFAULT:''" gorm:"column:SenderNumber"`
	Coding            string    `sql:"type:varchar(255);DEFAULT:'Default_No_Compression'" gorm:"column:Coding"`
	UDH               string    `sql:"type:text;not null" gorm:"column:UDH"`
	SMSCNumber        string    `sql:"not null;type:varchar(20);DEFAULT:''" gorm:"column:SMSCNumber"`
	Class             int       `sql:"not null;DEFAULT:'-1'" gorm:"column:Class"`
	TextDecoded       string    `sql:"type:text;not null;DEFAULT:''" gorm:"column:TextDecoded"`
	RecipientID       string    `sql:"type:text;not null" gorm:"column:RecipientID"`
	Processed         bool      `sql:"not null;DEFAULT:'false'" gorm:"column:Processed"`
}

func (in Inbox) DepositedAmount() int64 {
	// calculate the ammount received from text
	return 0
}
