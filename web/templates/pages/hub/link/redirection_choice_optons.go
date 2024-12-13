package link

import (
	"strconv"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
)

func RedirectionChoiceOptions(selected enums.RedirectionChoice) partial.SelectOptions {
	optionsArray := enums.RedirectionChoiceTimed.HumanizeWithValue()
	selectOptionArray := make([]partial.SelectOption, len(optionsArray))
	for i, option := range optionsArray {
		valueAsString := strconv.Itoa(option.Value)
		selectOptionArray[i] = partial.SelectOption{
			Value: valueAsString,
			Label: option.Label,
		}
		if option.Label == selected.Humanize() {
			selectOptionArray[i].Selected = true
		}
	}

	options := partial.NewSelectOptions(selectOptionArray)

	return options
}
