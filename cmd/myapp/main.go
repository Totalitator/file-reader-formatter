package main

import (
	"frf/internal"
	pkg "frf/pkg/formatter"
)

func main() {
	for {
		file := pkg.CreateFile()
		if file == nil {
			continue
		}

		pkg.FillFile(file)

		data := pkg.Reader(file)

		method := internal.ChooseMethod()

		formatData := pkg.GetFormatter(method, data)

		newFormat := pkg.Formating(formatData)

		pkg.SaveFormat(file, newFormat)

		if !internal.Cycle() {
			break
		}
	}
}
