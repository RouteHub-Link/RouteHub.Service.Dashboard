package customize

import "github.com/labstack/echo/v4"

func (h *Handlers) ConfigureRoutes(group *echo.Group) {
	group.GET("/customize", h.CustomizeGetHandler)
	group.GET("/customize/meta", h.MetaPageGet)
	group.GET("/customize/partial/meta", h.MetaPartialGet)
	group.POST("/customize/partial/meta", h.MetaPartialPost)

	group.GET("/customize/navbar", h.NavbarPageGet)
	group.GET("/customize/partial/navbar", h.NavbarPartialGet)
	group.POST("/customize/partial/navbar", h.NavbarPartialPost)

	// Two button for new navbar item and new navbar button
	group.GET("/customize/navbar/new/selection", h.NavbarNewSelectionGet)

	group.GET("/customize/navbar/item/:itemID/new", h.NavbarItemAddFormGet)
	group.POST("/customize/navbar/item/:itemID/new", h.NavbarItemAddFormPost)

	group.GET("/customize/navbar/item/:itemID/edit", h.NavbarItemEditFormGet)
	group.POST("/customize/navbar/item/:itemID/edit", h.NavbarItemEditFormPost)

	group.GET("/customize/navbar/item/:itemID/delete", h.NavbarItemDeleteFormGet)
	group.POST("/customize/navbar/item/:itemID/delete", h.NavbarItemDeletePost)

	group.GET("/customize/navbar/button/new", h.NavbarButtonEditFormGet)
	group.POST("/customize/navbar/button/new", h.NavbarButtonEditFormPost)

	group.GET("/customize/navbar/button/:itemID/edit", h.NavbarButtonEditFormGet)
	group.POST("/customize/navbar/button/:itemID/edit", h.NavbarButtonEditFormPost)

	group.GET("/customize/navbar/button/:itemID/delete", h.NavbarButtonDeleteFormGet)
	group.POST("/customize/navbar/button/:itemID/delete", h.NavbarButtonDeleteFormPost)

	group.GET("/customize/partial/navbar/tree", h.NavbarItemsTreeGet)
	group.GET("/customize/partial/navbar/shadow", h.NavbarItemsShadowGet)

	group.GET("/customize/footer", h.FooterPageGet)
	group.GET("/customize/partial/footer", h.FooterPartialGet)
	group.POST("/customize/partial/footer", h.FooterPartialPost)

	group.GET("/customize/footer/social-media-container", h.FooterSocialMediaContainerGet)
	group.GET("/customize/footer/social-media-links", h.FooterSocialMediaLinksGet)
	group.POST("/customize/footer/social-media-links", h.FooterSocialMediaLinksPost)
	group.GET("/customize/partial/footer/shadow", h.FooterShadow)

	group.GET("/customize/footer/branding-html", h.FooterBrandingHTMLGet)
	group.POST("/customize/footer/branding-html", h.FooterBrandingHTMLPost)
}
