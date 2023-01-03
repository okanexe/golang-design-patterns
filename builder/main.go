package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type HTTPBuilder interface {
	AddHeader(key, value string) HTTPBuilder
	Body(r io.Reader) HTTPBuilder
	Method(method string) HTTPBuilder
	Close(close bool) HTTPBuilder
	Build() (*http.Request, error)
}

type builder struct {
	headers map[string][]string
	url     string
	method  string
	body    io.Reader
	close   bool
	ctx     context.Context
}

func newBuilder(url string) *builder {
	return &builder{
		headers: map[string][]string{},
		url:     url,
		body:    nil,
		method:  http.MethodGet,
		ctx:     context.Background(),
		close:   false,
	}
}

func (b builder) AddHeader(key, value string) HTTPBuilder {
	values, ok := b.headers[key]
	if !ok {
		values = make([]string, 0, 10)
	}
	b.headers[key] = append(values, value)
	return b
}

func (b builder) Body(r io.Reader) HTTPBuilder {
	b.body = r
	return b
}

func (b builder) Close(close bool) HTTPBuilder {
	b.close = close
	return b
}

func (b builder) Method(method string) HTTPBuilder {
	b.method = method
	return b
}

func (b builder) Build() (*http.Request, error) {
	r, err := http.NewRequestWithContext(b.ctx, b.method, b.url, b.body)
	if err != nil {
		return nil, err
	}

	for key, values := range b.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	r.Close = b.close
	return r, nil
}

type director struct {
	builder *builder
}

func newDirector(b *builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b builder) {
	d.builder = &b
}

func (d *director) buildRequest() (*http.Request, error) {
	return d.builder.AddHeader("User-Agent", "Golang patterns").Build()
}

func main() {
	b := newBuilder("https://example.com/")
	director := newDirector(b)
	r, _ := director.buildRequest()
	fmt.Println(r)
}
