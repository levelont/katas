package kata_test
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "codewarrior/kata"
)

func dotest(n, p int, exp int) {
	var ans = Diagonal(n, p)
	Expect(ans).To(Equal(exp))
}

func testPyramid(n, p int, exp [][]int) {
	Expect(Pyramid(n,p)).To(Equal(exp))
}

var _ = Describe("Test Pyramid", func() {

	It("generate pascal triangle", func() {
		testPyramid(5,3, [][]int{  {1, 1, 1, 1},
			{1, 2, 3, 4},
			{1, 3, 6, 10},
		})
		testPyramid(7,7, [][]int{  {1, 1, 1, 1, 1, 1, 1, 1},
		})
		testPyramid(7,6, [][]int{  {1, 1, 1, 1, 1, 1, 1},
			{1, 2, 3, 4, 5, 6, 7},
		})
	})
})

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(7,0, 8)
		dotest(7,1, 28)
		dotest(7,2, 56)
		dotest(20,3, 5985)
		dotest(20,3, 5985)
		dotest(20,4, 20349)
		dotest(20,5, 54264)
		dotest(20,15, 20349)
		dotest(100,0, 101)
		dotest(100,10, 158940114100040)
	})
})

