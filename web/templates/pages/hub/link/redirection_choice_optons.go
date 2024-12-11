package link

import (
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
)

func RedirectionChoiceOptions(selected enums.RedirectionChoice) partial.SelectOptions {

	optionsArray := enums.RedirectionChoiceTimed.HumanizeAll()
	selectOptionArray := make([]partial.SelectOption, len(optionsArray))
	for i, option := range optionsArray {
		selectOptionArray[i] = partial.SelectOption{
			Value: option,
			Label: option,
		}
		if option == selected.Humanize() {
			selectOptionArray[i].Selected = true
		}
	}

	options := partial.NewSelectOptions(selectOptionArray)

	return options
}
