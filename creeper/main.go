package main

import "github.com/wspl/creeper"

func main() {
	c := creeper.Open("/Users/jpfaria/Projects/Go/src/github.com/jpfaria/golang-tests/creeper/test.crs")
	c.Array("news").Each(func(c *creeper.Creeper) {
		println("title: ", c.String("title"))
		println("site: ", c.String("site"))
		println("link: ", c.String("link"))
		println("===")
	})
}