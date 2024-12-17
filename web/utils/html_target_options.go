package utils

import "RouteHub.Service.Dashboard/web/templates/pages/partial"

func ElementTargetOptions(selected string) partial.SelectOptions {

	if selected == "" {
		selected = "_blank"
	}

	options := partial.NewSelectOptions([]partial.SelectOption{
		{Value: "_self", Label: "Same Window"},
		{Value: "_blank", Label: "New Window"},
		{Value: "_parent", Label: "Parent Frame"},
		{Value: "_top", Label: "Top Frame"},
	})

	options.Select(selected)

	return options
}
