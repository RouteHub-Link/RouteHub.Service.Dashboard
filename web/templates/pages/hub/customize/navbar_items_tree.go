package customize

import (
	"fmt"
	"strconv"
	"strings"

	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
)

func NavbarToTree(navbar *types.NavbarDescription, hubSlug string) *partial.TreeNode {

	root := &partial.TreeNode{
		Value: navbar.BrandName,
		IsDir: true,
		ID:    "nav-root",
		Icon:  partial.TreeIconFolder,
	}

	root.AdditionalTemplate = navbarFunctions(root, true, false, false, hubSlug)

	// Process StartItems
	if navbar.StartItems != nil {
		for i, item := range *navbar.StartItems {
			id := fmt.Sprintf("start-%d", i)
			node := convertNavbarItemToTree(item, id, hubSlug)
			root.Children = append(root.Children, node)
		}
	}

	// Process EndButtons
	if navbar.EndButtons != nil {
		for i, button := range *navbar.EndButtons {
			id := fmt.Sprintf("end-%d", i)
			node := convertNavbarButtonToTree(button, id)
			node.AdditionalTemplate = navbarFunctions(node, false, false, true, hubSlug)
			root.Children = append(root.Children, node)
		}
	}

	return root
}

func convertNavbarItemToTree(item types.NavbarItem, baseID string, hubSlug string) *partial.TreeNode {
	hasDropdown := item.Dropdown != nil && len(*item.Dropdown) > 0

	node := &partial.TreeNode{
		Value: item.Text,
		IsDir: hasDropdown,
		ID:    baseID,
		Icon:  getNavItemIcon(hasDropdown),
	}

	node.AdditionalTemplate = navbarFunctions(node, true, true, true, hubSlug)

	// Process dropdown items if they exist
	if hasDropdown {
		for i, dropdownItem := range *item.Dropdown {
			childID := fmt.Sprintf("%s-%d", baseID, i)
			childNode := convertNavbarItemToTree(dropdownItem, childID, hubSlug)
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}

func convertNavbarButtonToTree(button types.NavbarButton, baseID string) *partial.TreeNode {
	return &partial.TreeNode{
		Value: button.Text,
		IsDir: false,
		ID:    baseID,
		Icon:  partial.TreeIconButton,
	}
}

func getNavItemIcon(hasDropdown bool) partial.TreeIcon {
	if hasDropdown {
		return partial.TreeIconFolder
	}
	return partial.TreeIconLink
}
func (h Handlers) traverseToIndex(nav *types.NavbarDescription, indexPath string) (*[]types.NavbarItem, int, error) {
	const MaxDepth = 100

	if !strings.HasPrefix(indexPath, "start-") {
		return nil, -1, fmt.Errorf("invalid index path: must start with 'start-'")
	}

	indices := strings.Split(strings.TrimPrefix(indexPath, "start-"), "-")

	if len(indices) > MaxDepth {
		return nil, -1, fmt.Errorf("index path too deep: maximum depth is %d", MaxDepth)
	}

	if nav.StartItems == nil {
		return nil, -1, fmt.Errorf("no start items found")
	}

	currentItems := nav.StartItems
	var currentItem *types.NavbarItem

	// Traverse until the second-to-last index
	for i := 0; i < len(indices)-1; i++ {
		index, err := strconv.Atoi(indices[i])
		if err != nil {
			return nil, -1, fmt.Errorf("invalid index at position %d: %s", i, indices[i])
		}

		if index >= len(*currentItems) {
			return nil, -1, fmt.Errorf("index out of bounds at position %d", i)
		}

		currentItem = &(*currentItems)[index]
		if currentItem.Dropdown == nil {
			return nil, -1, fmt.Errorf("no dropdown found at depth %d", i)
		}
		currentItems = currentItem.Dropdown
	}

	// Get the final index
	lastIndex, err := strconv.Atoi(indices[len(indices)-1])
	if err != nil {
		return nil, -1, fmt.Errorf("invalid final index: %s", indices[len(indices)-1])
	}

	if lastIndex >= len(*currentItems) {
		return nil, -1, fmt.Errorf("final index out of bounds")
	}

	return currentItems, lastIndex, nil
}

// GetItemByIndex returns the item at the specified index path
func (h Handlers) GetItemByIndex(nav *types.NavbarDescription, indexPath string) (*types.NavbarItem, error) {
	items, index, err := h.traverseToIndex(nav, indexPath)
	if err != nil {
		return nil, err
	}
	return &(*items)[index], nil
}

// ReplaceItemByIndex replaces the item at the specified index path
func (h Handlers) ReplaceItemByIndex(nav *types.NavbarDescription, indexPath string, newItem types.NavbarItem) error {
	items, index, err := h.traverseToIndex(nav, indexPath)
	if err != nil {
		return err
	}
	(*items)[index] = newItem
	return nil
}

// UpdateItemByIndex updates the item at the specified index path using a modifier function
func (h Handlers) UpdateItemByIndex(nav *types.NavbarDescription, indexPath string, updater func(*types.NavbarItem)) error {
	items, index, err := h.traverseToIndex(nav, indexPath)
	if err != nil {
		return err
	}
	updater(&(*items)[index])
	return nil
}
