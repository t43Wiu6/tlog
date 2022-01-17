package log

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
)

var DEBUG bool

func Fatal(msg interface{}){
	pterm.Println(pterm.Red(fmt.Sprintf("[x] %v",msg)))
	os.Exit(0)
}

func Fatalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	pterm.Println(pterm.Red(fmt.Sprintf("[x] %v",msg)))
	os.Exit(0)
}

func Error(msg interface{}){
	pterm.Println(pterm.Red(fmt.Sprintf("[-] %v" ,msg)))
}

func Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	pterm.Println(pterm.Red(fmt.Sprintf("[-] %v" ,msg)))
}

func Warn(msg interface{}){
	pterm.Println(pterm.Yellow(fmt.Sprintf("[*] %v" ,msg)))
}

func Warnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	pterm.Println(pterm.Yellow(fmt.Sprintf("[*] %v" ,msg)))
}

func Info(msg interface{}) {
	pterm.Println(pterm.Blue(fmt.Sprintf("[+] %v", msg)))
}

func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	pterm.Println(pterm.Blue(fmt.Sprintf("[+] %v", msg)))
}

func Debug(msg interface{}){
	if DEBUG{
		pterm.Println(pterm.Green(fmt.Sprintf("[+] %v", msg)))
	}
}

func Debugf(format string, a ...interface{}) {
	if DEBUG{
		msg := fmt.Sprintf(format, a...)
		pterm.Println(pterm.Green(fmt.Sprintf("[+] %v", msg)))
	}
}
