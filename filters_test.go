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
	Describe("Imethibitishwa filter", func() {
		var node = "Imethibitishwa"
		var failureNodes = []string{
			"imethibitishwa", "Imethibitishwa ", " Imethibitishwa",
		}
		It("ticks", func() {
			a, b := Imethibitishwa(node)
			Expect(a).ShouldNot(BeEmpty())
			Expect(a).Should(Equal("Thibitisha"))
			Expect(b).Should(Equal(node))
		})
		It("bumps", func() {
			for _, v := range failureNodes {
				_, b := Imethibitishwa(v)
				Expect(b).Should(BeEmpty())
			}
		})
	})
	Describe("Kiasi", func() {
		var nodes = []string{
			"Tsh200", "Tsh2,500", "Tsh2,555", "Tsh2,500,500", "Tsh2,000,000,000",
			"Ksh200", "Ksh2,500", "Ksh2,555", "Ksh2,500,500", "Ksh2,000,000,000",
		}
		var failureNodes = []string{
			"Tsh2", "Tsh20", "Tsh2,55", "Tsh2500,500", "Tsh2,000,000,00",
			"Ksh200O", "Ksh2,500wap", "Ksh2555", "Ksh2,500500", "Ksh2,000,000000",
		}

		It("Ticks", func() {
			for _, node := range nodes {
				a, b := Kiasi(node)
				Expect(a).ShouldNot(BeEmpty())
				Expect(a).Should(Equal("Kiasi"))
				Expect(b).Should(Equal(node))
			}
		})
		It("bumps", func() {
			for _, node := range failureNodes {
				_, b := Kiasi(node)
				Expect(b).Should(BeEmpty())
			}
		})
	})
})
