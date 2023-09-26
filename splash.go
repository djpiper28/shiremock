package shiremock

import (
	"fmt"
)

const PROGRAM_NAME = "ShireMock"
const COPYRIGHT = "Public Domain (Unlicense)"
const AUTHORS = "Danny Piper <djpiper28@gmail.com>"

// / Prints the program name, copyright and authors, used in startup.
func PrintSplashScreen() {
	fmt.Printf("%s\nLicense: %s\nBy: %s", PROGRAM_NAME, COPYRIGHT, AUTHORS)
}
