package main

import "fmt"

type DataSource interface {
	WriteData(s string)
	ReadData() string
}

type dataSourceDecorator struct {
	wrappee DataSource
	content string
}

func DataSourceDecorator(s DataSource) *dataSourceDecorator {
	return &dataSourceDecorator{
		wrappee: s,
	}
}

func (d *dataSourceDecorator) WriteData(s string) {
	d.content = s
}

func (s *dataSourceDecorator) ReadData() string {
	return s.content
}

type EncryptionDecorator struct {
	wrappee DataSource
	content string
	// other fields to specific encryption
}

func (e *EncryptionDecorator) WriteData(s string) {
	wrappedContent := e.wrappee.ReadData() + " wrapped by encryption decorator with content: " + s
	e.content = wrappedContent
}

func (e *EncryptionDecorator) ReadData() string {
	return e.content
}

type CompressionDecorator struct {
	wrappee DataSource
	content string
	// other fields to specific compression
}

func (c *CompressionDecorator) WriteData(s string) {
	wrappedContent := c.wrappee.ReadData() + " wrapped by compression decorator with content: " + s
	c.content = wrappedContent
}

func (c *CompressionDecorator) ReadData() string {
	return c.content
}

func main() {
	data := dataSourceDecorator{}
	encrypt := EncryptionDecorator{
		wrappee: &data,
		content: "encrpyt successed.",
	}
	compressioned := CompressionDecorator{
		wrappee: &encrypt,
	}

	compressioned.WriteData("data compressed")
	d := compressioned.ReadData()
	fmt.Println(d)
}
