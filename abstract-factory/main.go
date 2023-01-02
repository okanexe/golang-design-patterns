package main

import (
	"fmt"
	"reflect"
)

type file struct {
	name    string
	content string
}

type FileSystem interface {
	createFile(string)
	findFile(string) file
}

type ntfs struct {
	files map[string]file
}

func (ntfs ntfs) createFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ntfs ntfs) findFile(key string) file {
	if _, ok := ntfs.files[key]; ok {
		return ntfs.files[key]
	}
	return file{}
}

type ext4 struct {
	files map[string]file
}

func (ext ext4) createFile(path string) {
	file := file{content: "EXT4 file", name: path}
	ext.files[path] = file
	fmt.Println("EXT4")
}

func (ext ext4) findFile(key string) file {
	if _, ok := ext.files[key]; ok {
		return ext.files[key]
	}
	return file{}
}

func FileSystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

type Database interface {
	getData(string) string
	putData(string, string)
}

type mongoDB struct {
	database map[string]string
}

func (m mongoDB) getData(key string) string {
	val, ok := m.database[key]
	if !ok {
		return ""
	}
	fmt.Println("MongoDB")
	return val
}

func (m mongoDB) putData(key string, val string) {
	if _, ok := m.database[key]; !ok {
		m.database[key] = ""
	}
	m.database[key] = val
}

type postgresql struct {
	database map[string]string
}

func (p postgresql) getData(key string) string {
	val, ok := p.database[key]
	if !ok {
		return ""
	}
	fmt.Println("PostgreSQL")
	return val
}

func (p postgresql) putData(key string, val string) {
	if _, ok := p.database[key]; !ok {
		p.database[key] = ""
	}
	p.database[key] = val
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return postgresql{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

type Factory func(string) interface{}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "filesystem":
		return FileSystemFactory
	case "database":
		return DatabaseFactory
	default:
		return nil
	}
}

func setupConstructors(env string) (Database, FileSystem) {
	fs := AbstractFactory("filesystem")
	db := AbstractFactory("database")
	return db(env).(Database), fs(env).(FileSystem)
}

func main() {
	envprod := "production"
	envdev := "development"

	db1, fs1 := setupConstructors(envprod)
	db2, fs2 := setupConstructors(envdev)

	db1.putData("test", "this is mongodb")
	fmt.Println(db1.getData("test"))

	db2.putData("test", "this is postgresql")
	fmt.Println(db2.getData("test"))

	fs1.createFile("../example/testntfs.txt")
	fmt.Println(fs1.findFile("../example/testntfs.txt"))

	fs2.createFile("../example/testext4.txt")
	fmt.Println(fs2.findFile("../example/testext4.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(fs1).Name())
	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(fs2).Name())
	fmt.Println(reflect.TypeOf(&fs2).Elem())
}
