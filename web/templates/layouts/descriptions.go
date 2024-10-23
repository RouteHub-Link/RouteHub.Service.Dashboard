package layouts

import "github.com/a-h/templ"

type PageDescription struct {
	AdditionalHead   *templ.Component
	MainContent      templ.Component
	AdditionalFooter *templ.Component
}

func (pd PageDescription) GetAdditionalFooter() templ.Component {
	return *pd.AdditionalFooter
}

func (pd PageDescription) GetAdditionalHead() templ.Component {
	return *pd.AdditionalHead
}
