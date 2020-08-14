package main

import (
"fmt"
"bufio"
"os"

	"github.com/tebeka/selenium"
	"github.com/xiote/interparkcontroller"
)

func main() {
	Example()
}

func Example() {

	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	selenium.SetDebug(true)

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	c := interparkcontroller.NewController2(wd, interparkcontroller.LoginInfo{"xiote12", "gkswlsdn78#"})
	if err := c.Login(); err != nil {
		panic(err)
	}

fmt.Print("What is your name? ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        return
    
    }

}
