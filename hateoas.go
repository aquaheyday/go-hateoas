package hateoas

import (
	"encoding/json"
)

// Link represents a hypermedia link with an href and optional rel
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel,omitempty"`
}

// Links is a map of relation name to Link
type Links map[string]Link

// Resource wraps any payload data with HATEOAS-compliant _links
type Resource struct {
	Data  interface{} `json:"data"`
	Links Links       `json:"_links"`
}

// ResourceBuilder builds a HATEOAS Resource with chained methods
type ResourceBuilder struct {
	data  interface{}
	links Links
}

// Wrap initializes a ResourceBuilder for the given data
func Wrap(data interface{}) *ResourceBuilder {
	return &ResourceBuilder{
		data:  data,
		links: make(Links),
	}
}

// Link adds a new hypermedia link to the builder
func (rb *ResourceBuilder) Link(rel string, href string) *ResourceBuilder {
	rb.links[rel] = Link{
		Href: href,
		Rel:  rel,
	}
	return rb
}

// Build constructs the final Resource object
func (rb *ResourceBuilder) Build() Resource {
	return Resource{
		Data:  rb.data,
		Links: rb.links,
	}
}

// JSON serializes the built Resource into JSON bytes
func (rb *ResourceBuilder) JSON() ([]byte, error) {
	return json.Marshal(rb.Build())
}
