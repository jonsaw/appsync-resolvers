package resolvers

import (
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	type arguments struct {
		Bar string `json:"bar"`
	}
	type identity struct {
		Bar string `json:"bar"`
	}
	type response struct {
		Foo string
	}
	r := New()
	r.Add("example.resolver", func(arg arguments) (response, error) { return response{"bar"}, nil })
	r.Add("example.resolver.with.error", func(arg arguments) (response, error) { return response{"bar"}, errors.New("Has Error") })
	r.Add("example.resolver.with.identity", func(arg arguments, ident identity) (response, error) { return response{"foo"}, nil })
	r.Add("example.resolver.with.identity.with.error", func(arg arguments, ident identity) (response, error) { return response{"foo"}, errors.New("Has Error") })
	Context("Matching invocation", func() {
		res, err := r.Handle(Invocation{
			Resolve: "example.resolver",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
			},
		})

		It("Should not error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Should have data", func() {
			Expect(res.(response).Foo).To(Equal("bar"))
		})
	})

	Context("Matching invocation with identity", func() {
		identityMessage := json.RawMessage(`{"bar":"foo"}`)
		res, err := r.Handle(Invocation{
			Resolve: "example.resolver.with.identity",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
				Identity:  &identityMessage,
			},
		})

		It("Should not error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Should have data", func() {
			Expect(res.(response).Foo).To(Equal("foo"))
		})
	})

	Context("Matching invocation with error", func() {
		_, err := r.Handle(Invocation{
			Resolve: "example.resolver.with.error",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
			},
		})

		It("Should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Matching invocation with identity and error", func() {
		identityMessage := json.RawMessage(`{"bar:foo"}`)
		_, err := r.Handle(Invocation{
			Resolve: "example.resolver.with.identity.with.error",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
				Identity:  &identityMessage,
			},
		})

		It("Should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Matching invocation with invalid payload", func() {
		_, err := r.Handle(Invocation{
			Resolve: "example.resolver.with.error",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar:foo"}`),
			},
		})

		It("Should error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Not matching invocation", func() {
		res, err := r.Handle(Invocation{
			Resolve: "example.resolver.not.found",
			Context: ContextData{
				Arguments: json.RawMessage(`{"bar":"foo"}`),
			},
		})

		It("Should error", func() {
			Expect(err).To(HaveOccurred())
		})

		It("Should have no data", func() {
			Expect(res).To(BeNil())
		})
	})
})
