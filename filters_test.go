package jesus_test

import (
	. "github.com/gernest/jesus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filters", func() {
	Describe("date filter", func() {
		var date = "15/12/14"
		failureCase := []string{
			"12/24/20", "2/06/20",
		}
		It("ticks", func() {
			a, b := Tarehe(date)
			Expect(a).ShouldNot(BeEmpty())
			Expect(a).Should(Equal("Tarehe"))
			Expect(b).Should(Equal(date))
		})
		It("bumps", func() {
			for _, v := range failureCase {
				_, b := Tarehe(v)
				Expect(b).Should(BeEmpty())
			}
		})
	})
})
