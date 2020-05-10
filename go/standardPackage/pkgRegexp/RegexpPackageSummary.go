package pkgRegexp

import (
	"fmt"
	"regexp"
)

func RegexpPackageSummary() {
	// very simple regx
	testMatcherList := []string{"AAA.BBB", "A.C", ""}
	for _, s := range testMatcherList {
		fmt.Println(regexp.MatchString("A+.C", s))
	}
	//custom regx
	re := regexp.MustCompile(`\d`)
	fmt.Println(re.MatchString("2021"))

	testCustomMatcherList := []string{"abcdef_", "あ", "nil", "a_-cdgo"}
	customRe := regexp.MustCompile(`\w+go$`)
	for _, s := range testCustomMatcherList {
		fmt.Println(customRe.MatchString(s))
	}
	//split
	testStr := "12 go 3 python"
	splitRe := regexp.MustCompile(`\s+`)
	fmt.Println(splitRe.Split(testStr, -1))
}
