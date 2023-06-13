package main

import "fmt"

// Printer 简单工厂要返回的接口类型
type Printer interface {
	Print(name string) string
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("Non"))
}
