package scrape

import (
	"encoding/base64"

	"RouteHub.Service.Dashboard/ent/schema/types"
	"github.com/RouteHub-Link/DomainUtils/validator"
	"github.com/gocolly/colly/v2"
)

// Source from : Chrome104.0.5112.79
const internalUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.79 Safari/537.36"

type CollyClient struct {
	UserAgent           string
	MaxBodySizeInBytes  int
	MaxImageSizeInBytes int
}

type CollyClientOptions func(*CollyClient)

func WithUserAgent(ua string) CollyClientOptions {
	return func(c *CollyClient) {
		c.UserAgent = ua
	}
}

func WithMaxBodySizeInBytes(size int) CollyClientOptions {
	return func(c *CollyClient) {
		c.MaxBodySizeInBytes = size
	}
}

func WithMaxImageSizeInBytes(size int) CollyClientOptions {
	return func(c *CollyClient) {
		c.MaxImageSizeInBytes = size
	}
}

func NewCollyClient(options ...CollyClientOptions) *CollyClient {
	c := &CollyClient{
		UserAgent:           internalUserAgent,
		MaxBodySizeInBytes:  10485760, // 10MB
		MaxImageSizeInBytes: 26214400, // 25MB
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (cc CollyClient) VisitScrapeOG(url string) (meta *types.MetaDescription, err error) {
	var _validator = validator.DefaultValidator()
	isValid, err := _validator.ValidateURL(url)

	if !isValid {
		return nil, err
	}

	images := make(map[string]string)

	c := colly.NewCollector(
		colly.UserAgent(cc.UserAgent),
		colly.IgnoreRobotsTxt(),
		colly.MaxDepth(2),
		colly.Async(true),
		colly.MaxBodySize(cc.MaxBodySizeInBytes),
		colly.CheckHead(),
	)

	c.OnHTML("head", func(e *colly.HTMLElement) {
		meta = &types.MetaDescription{
			Title:         e.ChildAttr("meta[property='og:title']", "content"),
			FavIcon:       e.ChildAttr("link[rel='shortcut icon']", "href"),
			Description:   e.ChildAttr("meta[property='og:description']", "content"),
			Locale:        e.ChildAttr("meta[property='og:locale']", "content"),
			OGTitle:       e.ChildAttr("meta[property='og:title']", "content"),
			OGDescription: e.ChildAttr("meta[property='og:description']", "content"),
			OGURL:         e.ChildAttr("meta[property='og:url']", "content"),
			OGSiteName:    e.ChildAttr("meta[property='og:site_name']", "content"),
			OGLocale:      e.ChildAttr("meta[property='og:locale']", "content"),
			OGBigImage:    e.ChildAttr("meta[property='og:image']", "content"),
			OGBigWidth:    e.ChildAttr("meta[property='og:image:width']", "content"),
			OGBigHeight:   e.ChildAttr("meta[property='og:image:height']", "content"),
			OGSmallImage:  e.ChildAttr("meta[property='og:image:secure_url']", "content"),
			OGSmallWidth:  e.ChildAttr("meta[property='og:image:secure_url:width']", "content"),
			OGSmallHeight: e.ChildAttr("meta[property='og:image:secure_url:height']", "content"),
			OGCard:        e.ChildAttr("meta[property='twitter:card']", "content"),
			OGType:        e.ChildAttr("meta[property='twitter:type']", "content"),
			OGCreator:     e.ChildAttr("meta[property='twitter:creator']", "content"),
		}
	})

	// download favicon
	c.OnResponse(func(r *colly.Response) {
		// download favicon & images
		if r.Headers.Get("Content-Type") == "image/x-icon" ||
			r.Headers.Get("Content-Type") == "image/png" ||
			r.Headers.Get("Content-Type") == "image/jpeg" ||
			r.Headers.Get("Content-Type") == "image/jpg" ||
			r.Headers.Get("Content-Type") == "image/gif" {
			// check file size
			if len(r.Body) < cc.MaxImageSizeInBytes {
				fileName := r.FileName()
				images[fileName] = base64.StdEncoding.EncodeToString(r.Body)
			}
		}
	})

	err = c.Visit(url)
	c.Wait()

	if err != nil {
		return nil, err
	}

	transformMetaImages(meta, cc, images)

	return
}

func transformMetaImages(meta *types.MetaDescription, cc CollyClient, images map[string]string) {
	if meta.FavIcon != "" {
		if len(meta.FavIcon) < cc.MaxImageSizeInBytes {
			meta.FavIcon = images[meta.FavIcon]
		}
	}

	if meta.OGBigImage != "" {
		if len(meta.OGBigImage) < cc.MaxImageSizeInBytes {
			meta.OGBigImage = images[meta.OGBigImage]
		}
	}

	if meta.OGSmallImage != "" {
		if len(meta.OGSmallImage) < cc.MaxImageSizeInBytes {
			meta.OGSmallImage = images[meta.OGSmallImage]
		}
	}
}
