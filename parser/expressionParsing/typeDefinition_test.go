package expressionParsing_test

import (
	. "github.com/apoydence/gof/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeDefinition", func() {
	Context("Parse and GenerateGo", func() {
		It("Should return a primitive type", func() {
			code := "int32"
			t, err, _ := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("int32"))
		})
		It("Should return a simple function type", func() {
			code := "func a int32 -> int32"
			t, err, _ := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("func (a int32) int32"))
		})
		It("Should return a function type with a function as an argument", func() {
			code := "func a func x int32 -> int32 -> int32"
			t, err, _ := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("func (a func (x int32) int32) int32"))
		})
		It("Should return a curried function type", func() {
			code := "func a int32 -> b int32 -> int32"
			t, err, _ := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("func (a int32) func (b int32) int32"))
		})
		It("Should return a curried function type with argument as a function", func() {
			code := "func a func x int32 -> int32 -> b int32 -> int32"
			t, err, _ := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("func (a func (x int32) int32) func (b int32) int32"))
		})
		It("Should return a curried function type with argument as a function and the stuff after", func() {
			code := "func a func x int32 -> int32 -> b int32 -> int32 -> a b"
			t, err, rest := ParseTypeDef(code)
			Expect(err).To(BeNil())
			Expect(t.GenerateGo()).To(Equal("func (a func (x int32) int32) func (b int32) int32"))
			Expect(rest).To(Equal("-> a b"))
		})
		It("Should return an error", func() {
			code := "func (a func (x int32) int32) func (b B) int32"
			_, err, _ := ParseTypeDef(code)
			Expect(err).ToNot(BeNil())
		})
	})
})
