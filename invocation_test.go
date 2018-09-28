package resolvers

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Invocation", func() {
	Context("With Arguments", func() {
		data := Invocation{
			Resolve: "exaple.resolver",
			Context: ContextData{
				Arguments: json.RawMessage(`{ "foo": "bar" }`),
			},
		}

		It("should be root", func() {
			Expect(data.isRoot()).To(BeTrue())
		})

		It("should detect data", func() {
			Expect(data.payload()).To(Equal(json.RawMessage(`{ "foo": "bar" }`)))
		})
	})

	Context("With Source", func() {
		data := Invocation{
			Resolve: "exaple.resolver",
			Context: ContextData{
				Source: json.RawMessage(`{ "bar": "foo" }`),
			},
		}

		It("should be root", func() {
			Expect(data.isRoot()).To(BeFalse())
		})

		It("should detect data", func() {
			Expect(data.payload()).To(Equal(json.RawMessage(`{ "bar": "foo" }`)))
		})
	})
})
