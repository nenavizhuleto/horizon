package main

import (
	"github.com/nenavizhuleto/horizon/examples/scenario/components/adder"
	"github.com/nenavizhuleto/horizon/examples/scenario/components/printer"
	"github.com/nenavizhuleto/horizon/scenario"
)

func main() {

	adder := adder.New()
	printer := printer.New()

	script := scenario.NewScript("test")

	script.NewComponent("adder", adder)
	script.NewComponent("printer", printer)

	script.Connect("adder", 0, "printer", 0)
	script.Connect("printer", 0, "adder", 0)
	script.Connect("printer", 0, "adder", 1)

	script.Run()
}
