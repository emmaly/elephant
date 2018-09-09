package log

import (
	"fmt"

	"github.com/dennwc/dom"
)

func init() {
	input := dom.GetDocument().NewInput("text")
	body := dom.GetDocument().GetElementsByTagName("body")[0]
	body.AppendChild(input)
	input.SetValue("input")
	input.OnChange(func(evt dom.Event) {
		fmt.Printf("%v\n", evt)
	})
}
