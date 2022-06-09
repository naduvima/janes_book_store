package main

import (
	goji "goji.io"
)


func main() {
	//init all http handlers called here.
	mux := goji.NewMux()
	initHandlers(mux)
}

