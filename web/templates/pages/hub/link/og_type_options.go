package link

import (
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
)

func OGTypeOptions(md *types.MetaDescription) partial.SelectOptions {
	options := partial.NewSelectOptions(
		[]partial.SelectOption{
			{Label: "Website", Value: "website"},
			{Label: "Article", Value: "article"},
			{Label: "Book", Value: "book"},
			{Label: "Profile", Value: "profile"},
			{Label: "Video", Value: "video"},
			{Label: "Music", Value: "music"},
			{Label: "Movie", Value: "movie"},
			{Label: "TV Show", Value: "tv_show"},
			{Label: "Game", Value: "game"},
			{Label: "Place", Value: "place"},
			{Label: "Product", Value: "product"},
			{Label: "Music", Value: "music"},
			{Label: "Song", Value: "music.song"},
			{Label: "Album", Value: "music.album"},
			{Label: "Playlist", Value: "music.playlist"},
			{Label: "Radio Station", Value: "music.radio_station"},
			{Label: "Fitness", Value: "fitness"},
			{Label: "Courses", Value: "course"},
			{Label: "Other", Value: "other"},
		},
	)

	if md != nil {
		options.Select(md.OGType)
	}

	return options
}
