//http://www.jb51.net/article/127805.htm

package main

import (
	"./server/router"
	"google.golang.org/appengine"
)

func main() {
	router.Init() // init router
	appengine.Main()
}
