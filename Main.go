package main

import (
	"Day8-Http-Context-Signal/DemoSignal"
	"fmt"
)

func main() {
	fmt.Println("----------------------- HTTP Clients -----------------------------")
	//ExampleHTTP.TestHttp()
	//ExampleHTTP.TestHttpPost()
	fmt.Println("----------------------- CONTEXT -----------------------------")
	//DemoContext.DemoContext()
	//DemoContext.TestContextBackground()
	//DemoContext.DemoValueContext()
	//DemoContext.DemoDeadlinesContext()
	fmt.Println("----------------------- SIGNALS -----------------------------")
	DemoSignal.TestSignal()
}
