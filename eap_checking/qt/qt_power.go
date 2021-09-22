package main

import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	a := qt.NewQApplication(len(os.Args), os.Args) //ch: widgets.NewQ..
	hello := widgets.QPushButton("Hello world!", 0)
	a.setMainWidget(&hello)
	hello.Show()
	a.Exec()
}
