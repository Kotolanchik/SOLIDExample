package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Fax interface {
	Fax()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Fax
}

type SimplePrinter struct {
}

type AdvancedPrinter struct {
}

func (sp SimplePrinter) Print() {
	fmt.Println("SimplePrinter: Print")
}

func (ap AdvancedPrinter) Print() {
	fmt.Println("AdvancedPrinter: Print")
}

type SimpleScanner struct {
}

type AdvancedScanner struct {
}

func (ss SimpleScanner) Scan() {
	fmt.Println("SimpleScanner: Scan")
}

func (as AdvancedScanner) Scan() {
	fmt.Println("AdvancedScanner: Scan")
}

type SimpleFax struct {
}

type AdvancedFax struct {
}

func (sf SimpleFax) Fax() {
	fmt.Println("SimpleFax: Fax")
}

func (af AdvancedFax) Fax() {
	fmt.Println("AdvancedFax: Fax")
}

func main() {
	simplePrinter := SimplePrinter{}
	advancedPrinter := AdvancedPrinter{}
	simpleScanner := SimpleScanner{}
	advancedScanner := AdvancedScanner{}
	simpleFax := SimpleFax{}
	advancedFax := AdvancedFax{}

	simplePrinter.Print()
	advancedPrinter.Print()
	simpleScanner.Scan()
	advancedScanner.Scan()
	simpleFax.Fax()
	advancedFax.Fax()

}
