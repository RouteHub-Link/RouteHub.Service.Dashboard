package partial

type SelectOption struct {
	Label    string `json:"label,omitempty"`
	Value    string `json:"value,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}

type SelectOptions []SelectOption

func NewSelectOptions(options []SelectOption) SelectOptions {
	return options
}

func (sos SelectOptions) Select(value string) SelectOptions {
	selected := false
	for i, so := range sos {
		if so.Value == value {
			sos[i].Selected = true
			selected = true
		}
	}

	if !selected {
		sos[len(sos)-1].Selected = true
	}

	return sos

}

func (sos SelectOptions) SelectedAny() bool {
	for _, so := range sos {
		if so.Selected {
			return true
		}
	}

	return false
}
