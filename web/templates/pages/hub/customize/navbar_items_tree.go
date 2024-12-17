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

	root.AdditionalTemplate = navRootFunctions(hubSlug)

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
			node.AdditionalTemplate = navbarButtonFunctions(node, hubSlug)
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

// Common error type for better error handling
type NavError struct {
	Path string
	Msg  string
}

func (e *NavError) Error() string {
	return fmt.Sprintf("navigation error at %s: %s", e.Path, e.Msg)
}

// Common interface for index traversal
type IndexedItem interface {
	GetButtons() *[]types.NavbarButton
	GetItems() *[]types.NavbarItem
}

// Implement the interface for NavbarDescription
func GetButtons(nav *types.NavbarDescription) *[]types.NavbarButton {
	return nav.EndButtons
}

func GetItems(nav *types.NavbarDescription) *[]types.NavbarItem {
	return nav.StartItems
}

// Generic traversal function
func traverse(nav *types.NavbarDescription, indexPath string) (interface{}, int, error) {
	const MaxDepth = 100

	// Validate path prefix
	if !strings.HasPrefix(indexPath, "start-") && !strings.HasPrefix(indexPath, "end-") {
		return nil, -1, &NavError{indexPath, "path must start with 'start-' or 'end-'"}
	}

	isButton := strings.HasPrefix(indexPath, "end-")
	prefix := "start-"
	if isButton {
		prefix = "end-"
	}

	indices := strings.Split(strings.TrimPrefix(indexPath, prefix), "-")

	if len(indices) > MaxDepth {
		return nil, -1, &NavError{indexPath, fmt.Sprintf("path too deep: maximum depth is %d", MaxDepth)}
	}

	if isButton {
		if nav.EndButtons == nil {
			return nil, -1, &NavError{indexPath, "no end buttons found"}
		}

		// For buttons, we only support single-level indexing
		if len(indices) != 1 {
			return nil, -1, &NavError{indexPath, "buttons only support single-level indexing"}
		}

		index, err := strconv.Atoi(indices[0])
		if err != nil {
			return nil, -1, &NavError{indexPath, "invalid button index"}
		}

		if index >= len(*nav.EndButtons) {
			return nil, -1, &NavError{indexPath, "button index out of bounds"}
		}

		return nav.EndButtons, index, nil
	}

	// Handle NavbarItems
	if nav.StartItems == nil {
		return nil, -1, &NavError{indexPath, "no start items found"}
	}

	currentItems := nav.StartItems
	var currentItem *types.NavbarItem

	// Traverse until the second-to-last index for items
	for i := 0; i < len(indices)-1; i++ {
		index, err := strconv.Atoi(indices[i])
		if err != nil {
			return nil, -1, &NavError{indexPath, fmt.Sprintf("invalid index at position %d", i)}
		}

		if index >= len(*currentItems) {
			return nil, -1, &NavError{indexPath, fmt.Sprintf("index out of bounds at position %d", i)}
		}

		currentItem = &(*currentItems)[index]
		if currentItem.Dropdown == nil {
			return nil, -1, &NavError{indexPath, fmt.Sprintf("no dropdown found at depth %d", i)}
		}
		currentItems = currentItem.Dropdown
	}

	lastIndex, err := strconv.Atoi(indices[len(indices)-1])
	if err != nil {
		return nil, -1, &NavError{indexPath, "invalid final index"}
	}

	if lastIndex >= len(*currentItems) {
		return nil, -1, &NavError{indexPath, "final index out of bounds"}
	}

	return currentItems, lastIndex, nil
}

// Functions for NavbarItems
func GetItemByIndex(nav *types.NavbarDescription, indexPath string) (*types.NavbarItem, error) {
	items, index, err := traverse(nav, indexPath)
	if err != nil {
		return nil, err
	}
	itemsSlice := items.(*[]types.NavbarItem)
	return &(*itemsSlice)[index], nil
}

func ReplaceItemByIndex(nav *types.NavbarDescription, indexPath string, newItem types.NavbarItem) error {
	items, index, err := traverse(nav, indexPath)
	if err != nil {
		return err
	}
	itemsSlice := items.(*[]types.NavbarItem)
	(*itemsSlice)[index] = newItem
	return nil
}

func UpdateItemByIndex(nav *types.NavbarDescription, indexPath string, updater func(*types.NavbarItem)) error {
	items, index, err := traverse(nav, indexPath)
	if err != nil {
		return err
	}
	itemsSlice := items.(*[]types.NavbarItem)
	updater(&(*itemsSlice)[index])
	return nil
}

// Functions for NavbarButtons
func GetButtonByIndex(nav *types.NavbarDescription, indexPath string) (*types.NavbarButton, error) {
	buttons, index, err := traverse(nav, indexPath)
	if err != nil {
		return nil, err
	}
	buttonsSlice := buttons.(*[]types.NavbarButton)
	return &(*buttonsSlice)[index], nil
}

func ReplaceButtonByIndex(nav *types.NavbarDescription, indexPath string, newButton types.NavbarButton) error {
	buttons, index, err := traverse(nav, indexPath)
	if err != nil {
		return err
	}
	buttonsSlice := buttons.(*[]types.NavbarButton)
	(*buttonsSlice)[index] = newButton
	return nil
}

func UpdateButtonByIndex(nav *types.NavbarDescription, indexPath string, updater func(*types.NavbarButton)) error {
	buttons, index, err := traverse(nav, indexPath)
	if err != nil {
		return err
	}
	buttonsSlice := buttons.(*[]types.NavbarButton)
	updater(&(*buttonsSlice)[index])
	return nil
}

func AppendButton(nav *types.NavbarDescription, button types.NavbarButton) {
	if nav.EndButtons == nil {
		nav.EndButtons = &[]types.NavbarButton{}
	}
	*nav.EndButtons = append(*nav.EndButtons, button)
}

// Convenience wrapper functions for clarity
func DeleteItemByIndex(nav *types.NavbarDescription, indexPath string) error {
	if !strings.HasPrefix(indexPath, "start-") {
		return &NavError{indexPath, "item path must start with 'start-'"}
	}
	return DeleteByIndex(nav, indexPath)
}

func DeleteButtonByIndex(nav *types.NavbarDescription, indexPath string) error {
	if !strings.HasPrefix(indexPath, "end-") {
		return &NavError{indexPath, "button path must start with 'end-'"}
	}
	return DeleteByIndex(nav, indexPath)
}

func DeleteByIndex(nav *types.NavbarDescription, indexPath string) error {
	items, index, err := traverse(nav, indexPath)
	if err != nil {
		return err
	}

	// Handle deletion based on type
	switch slice := items.(type) {
	case *[]types.NavbarItem:
		*slice = append((*slice)[:index], (*slice)[index+1:]...)
	case *[]types.NavbarButton:
		*slice = append((*slice)[:index], (*slice)[index+1:]...)
	default:
		return &NavError{indexPath, "unknown type for deletion"}
	}

	return nil
}

/*
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

*/
