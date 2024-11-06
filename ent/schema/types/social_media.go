package types

type SocialMedia struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"`
}

type SocialMedias struct {
	SocialMedias *[]SocialMedia `json:"social_medias"`
}
