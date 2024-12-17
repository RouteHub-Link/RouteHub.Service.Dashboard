package partial

import "github.com/a-h/templ"

type TreeIcon int

const (
	TreeIconFolder TreeIcon = iota
	TreeIconFile
	TreeIconMenu
	TreeIconArrow
	TreeIconBullet
	TreeIconLink
	TreeIconButton
)

type TreeNode struct {
	Value              string
	IsDir              bool
	Icon               TreeIcon
	ID                 string
	Children           []*TreeNode
	AdditionalTemplate templ.Component
}

func DemoTreeNode() *TreeNode {
	return &TreeNode{
		Value: "assets",
		IsDir: true,
		Icon:  TreeIconFolder,
		Children: []*TreeNode{
			{
				Value: "css",
				IsDir: true,
				Icon:  TreeIconFolder,
				Children: []*TreeNode{
					{
						Value: "main",
						IsDir: true,
						Icon:  TreeIconFolder,
						Children: []*TreeNode{
							{Value: "main.css", IsDir: false, Icon: TreeIconLink},
							{Value: "docs.css", IsDir: false, Icon: TreeIconLink},
							{Value: "README.txt", IsDir: false, Icon: TreeIconLink},
						},
					},
					{
						Value: "tailwind",
						IsDir: true,
						Children: []*TreeNode{
							{Value: "input.css", IsDir: false, Icon: TreeIconLink},
						},
					},
					{Value: ".gitignore", IsDir: false, Icon: TreeIconLink},
				},
			},
			{
				Value: "img",
				IsDir: true,
				Children: []*TreeNode{
					{Value: "hero.jpg", IsDir: false, Icon: TreeIconLink},
					{Value: "tailwind.png", IsDir: false, Icon: TreeIconLink},
					{Value: "untitled.png", IsDir: false, Icon: TreeIconLink},
				},
			},
			{
				Value: "js",
				IsDir: true,
				Children: []*TreeNode{
					{Value: "preline.jpg", IsDir: false, Icon: TreeIconLink},
				},
			},
		},
	}
}
