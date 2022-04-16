package strtocamel_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/strtocamel"
)

func dotest(str, exp string) {
	fmt.Println("input:", str)
	var ans = ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases 1 ", func() {
		dotest("", "")
	})

	It("should handle basic cases 2", func() {
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
	})

	It("should handle basic cases 3", func() {
		dotest("the-Stealth-Warrior", "theStealthWarrior")
	})
	It("should handle basic cases 3", func() {
		dotest("You_have_chosen_to_translate_this_kata_For_your_convenience_we_have_provided_the_existing_test_cases_used_for_the_language_that_you_have_already_completed_as_well_as_all_of_the_other_related_fields",
			"YouHaveChosenToTranslateThisKataForYourConvenienceWeHaveProvidedTheExistingTestCasesUsedForTheLanguageThatYouHaveAlreadyCompletedAsWellAsAllOfTheOtherRelatedFields")
	})
	It("", func() {
		dotest("down_Street_right_down_desert_Blue_north", "downStreetRightDownDesertBlueNorth")
	})
})
