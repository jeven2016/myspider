package pkg

import (
	"github.com/traefik/yaegi/interp"
	"regexp"
	"testing"
)

const src = `package foo
func Bar(s string) string { return s + "-Foo" }`

func TestStringCode(t *testing.T) {
	i := interp.New(interp.Options{})

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	println(r)
}

func TestLastPageLink(t *testing.T) {
	compile, err := regexp.Compile("http.*/\\d+_(\\d+)/?$")
	if err != nil {
		panic(err)
	}
	match := compile.FindStringSubmatch("http://m.xinbanzhu.net/sort/7_782/")

	println(match[1])

}

func TestPageLinkPrefix(t *testing.T) {
	compile, err := regexp.Compile("\\s*(http.*/\\d+_)\\d+/?$")
	if err != nil {
		panic(err)
	}
	match := compile.FindStringSubmatch("http://m.xinbanzhu.net/sort/7_782/")

	println(match[1])

}
