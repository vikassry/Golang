package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	var interested_people []string
	for key, value := range data {
		if contains(value, interest) {
			interested_people = append(interested_people, key)
		}
	}
	return interested_people
}

func contains(src []string, desired_item string) bool {
	for _, item := range src {
		if item == desired_item {
			return true
		}
	}
	return false
}
