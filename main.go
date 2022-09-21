package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func getFirstWord(line string) string {
	r := regexp.MustCompile(`\w*`)

	return r.FindString(line)
}

func main() {
	list, err := exec.Command("gh", "gist", "list", "-L", "50").CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for index, line := range strings.Split(string(list), "\n") {
		id := getFirstWord(line)
		fmt.Println(index, id)

		// my idea was execute another command here like "gh gist edit <id> make it secret"
		// but github doesn't support this feature anymore :'(
		// in the end of the day this is just a useless code but I will keep in my repository for future reference
	}
}
