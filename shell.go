package shell

import (
	"fmt"

	c1 "github.com/GillainHsi/Core_1"
	c2 "github.com/GillainHsi/Core_2"
)

func Show() string {
	fmt.Println("shell::Show() v2.0 ")

	return fmt.Sprintf("c1 version: %s, c2 version: %s", c1.Version(), c2.Version())
}
