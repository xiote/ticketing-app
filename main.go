package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xiote/ticketing-app/interpark"

	"github.com/tebeka/selenium"
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

	c := interpark.NewController2(wd, interpark.LoginInfo{"xiote12", "gkswlsdn78#"})
	if err := c.Login(); err != nil {
		panic(err)
	}

	fmt.Print("Press ENTER or type command to continue")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return

	}

}
