package jesus_test

import (
	. "github.com/gernest/jesus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processor", func() {
	Describe("Pipeline", func() {
		var smsp *TextPipeline
		Describe("Mpesa", func() {
			mpesa := `BJ86VC805 Imethibitishwa Umepokea Tsh3,000 kutok kwa MAMA MIA
			Tarehe 15/12/14 saa1:13pm salio lako la M-PESA ni Tsh25,000`
			mpesaMatch := map[string]string{
				"Thibitisha": "Imethibitishwa",
				"Umepokea":   "Umepokea",
				"Tarehe":     "15/12/14",
				"Kiasi":      "Tsh3,000",
			}
			BeforeEach(func() {
				smsp = NewTextPipeline(mpesa)
			})
			It("process m-pesa", func() {
				Expect(smsp.Process()).Should(Succeed())
			})
			It("should report to output map", func() {
				_ = smsp.Process()
				c := smsp.CheckOutput()
				for k, v := range c {
					Expect(mpesaMatch[k]).Should(Equal(v))
				}
			})
		})

	})
})
