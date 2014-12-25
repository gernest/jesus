package jesus_test

import (
	. "github.com/gernest/jesus"
	"github.com/jinzhu/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This does not work ontravis so I have to skip it till when I'm sure
// How to sort it out
var _ = PDescribe("Sms", func() {
	var (
		DB  *gorm.DB
		err error
	)
	BeforeEach(func() {
		DB, err = ConnectLocalDB("jesus_test")
		Expect(err).ShouldNot(HaveOccurred())
	})
	AfterEach(func() {
		DB.Close()
	})
	Describe("It should have a compatible schema with the smsd", func() {
		var receivedSMS *Inbox
		BeforeEach(func() {
			receivedSMS = &Inbox{
				Text:        "yoyo",
				UDH:         "Yoyo ma",
				RecipientID: "yoyo man",
				Coding:      "Default_No_Compression",
			}
		})
		AfterEach(func() {
			DB.Delete(receivedSMS)
		})

		It("saves to db", func() {
			err = DB.Create(receivedSMS).Error
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
