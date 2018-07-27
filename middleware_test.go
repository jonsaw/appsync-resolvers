package resolvers

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func sequence(ch chan string, seq ...string) error {
	for i, str := range seq {
		if msg := <-ch; msg != str {
			return fmt.Errorf("[%d] expected %s, got %s", i, str, msg)
		}
	}
	return nil
}

var _ = Describe("Middleware", func() {
	type arguments struct {
		Bar string `json:"bar"`
	}
	type response struct {
		Foo string
	}
	type graphQLError struct {
	}

	Context("With no hijacking", func() {
		ch := make(chan string, 10)
		r := New()
		r.Add("example.resolver", func(arg arguments) (response, error) {
			ch <- "handler"
			return response{"bar"}, nil
		})
		r.Use(func(h Handler) Handler {
			m := func(in invocation) (interface{}, error) {
				ch <- "before 1"
				out, err := h.Serve(in)
				ch <- "after 1"
				return out, err
			}
			return HandlerFunc(m)
		})
		r.Use(func(h Handler) Handler {
			m := func(in invocation) (interface{}, error) {
				ch <- "before 2"
				out, err := h.Serve(in)
				ch <- "after 2"
				return out, err
			}
			return HandlerFunc(m)
		})
		res, err := r.Handle(invocation{
			Resolve: "example.resolver",
			Context: context{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
			},
		})

		It("Should not error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Should have data", func() {
			Expect(res.(response).Foo).To(Equal("bar"))
		})

		It("Should be in sequence", func() {
			Expect(
				sequence(ch,
					"before 1",
					"before 2",
					"handler",
					"after 2",
					"after 1",
				)).ToNot(HaveOccurred())
		})
	})
})
