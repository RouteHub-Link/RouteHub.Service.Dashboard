package partial

import (
	"strconv"

	"RouteHub.Service.Dashboard/ent/schema/enums"
)

func StatusStateOptions(selected enums.StatusState) SelectOptions {
	optionsArray := enums.StatusInactive.HumanizeWithValue()
	selectOptionArray := make([]SelectOption, len(optionsArray))
	for i, option := range optionsArray {
		valueAsString := strconv.Itoa(option.Value)
		selectOptionArray[i] = SelectOption{
			Value: valueAsString,
			Label: option.Label,
		}
		if option.Label == selected.Humanize() {
			selectOptionArray[i].Selected = true
		}
	}

	options := NewSelectOptions(selectOptionArray)

	return options
}
