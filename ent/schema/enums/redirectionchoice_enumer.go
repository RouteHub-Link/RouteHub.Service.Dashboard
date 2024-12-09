// Code generated by "enumer -type=RedirectionChoice -trimprefix=RedirectionChoice -transform=snake-upper -json -sql -values -gqlgen"; DO NOT EDIT.

package enums

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const _RedirectionChoiceName = "TIMEDNOT_AUTODIRECT_HTTPCONFIRMCUSTOM"

var _RedirectionChoiceIndex = [...]uint8{0, 5, 13, 24, 31, 37}

const _RedirectionChoiceLowerName = "timednot_autodirect_httpconfirmcustom"

func (i RedirectionChoice) String() string {
	if i < 0 || i >= RedirectionChoice(len(_RedirectionChoiceIndex)-1) {
		return fmt.Sprintf("RedirectionChoice(%d)", i)
	}
	return _RedirectionChoiceName[_RedirectionChoiceIndex[i]:_RedirectionChoiceIndex[i+1]]
}

func (RedirectionChoice) Values() []string {
	return RedirectionChoiceStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RedirectionChoiceNoOp() {
	var x [1]struct{}
	_ = x[RedirectionChoiceTimed-(0)]
	_ = x[RedirectionChoiceNotAuto-(1)]
	_ = x[RedirectionChoiceDirectHTTP-(2)]
	_ = x[RedirectionChoiceConfirm-(3)]
	_ = x[RedirectionChoiceCustom-(4)]
}

var _RedirectionChoiceValues = []RedirectionChoice{RedirectionChoiceTimed, RedirectionChoiceNotAuto, RedirectionChoiceDirectHTTP, RedirectionChoiceConfirm, RedirectionChoiceCustom}

var _RedirectionChoiceNameToValueMap = map[string]RedirectionChoice{
	_RedirectionChoiceName[0:5]:        RedirectionChoiceTimed,
	_RedirectionChoiceLowerName[0:5]:   RedirectionChoiceTimed,
	_RedirectionChoiceName[5:13]:       RedirectionChoiceNotAuto,
	_RedirectionChoiceLowerName[5:13]:  RedirectionChoiceNotAuto,
	_RedirectionChoiceName[13:24]:      RedirectionChoiceDirectHTTP,
	_RedirectionChoiceLowerName[13:24]: RedirectionChoiceDirectHTTP,
	_RedirectionChoiceName[24:31]:      RedirectionChoiceConfirm,
	_RedirectionChoiceLowerName[24:31]: RedirectionChoiceConfirm,
	_RedirectionChoiceName[31:37]:      RedirectionChoiceCustom,
	_RedirectionChoiceLowerName[31:37]: RedirectionChoiceCustom,
}

var _RedirectionChoiceNames = []string{
	_RedirectionChoiceName[0:5],
	_RedirectionChoiceName[5:13],
	_RedirectionChoiceName[13:24],
	_RedirectionChoiceName[24:31],
	_RedirectionChoiceName[31:37],
}

// RedirectionChoiceString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RedirectionChoiceString(s string) (RedirectionChoice, error) {
	if val, ok := _RedirectionChoiceNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RedirectionChoiceNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RedirectionChoice values", s)
}

// RedirectionChoiceValues returns all values of the enum
func RedirectionChoiceValues() []RedirectionChoice {
	return _RedirectionChoiceValues
}

// RedirectionChoiceStrings returns a slice of all String values of the enum
func RedirectionChoiceStrings() []string {
	strs := make([]string, len(_RedirectionChoiceNames))
	copy(strs, _RedirectionChoiceNames)
	return strs
}

// IsARedirectionChoice returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RedirectionChoice) IsARedirectionChoice() bool {
	for _, v := range _RedirectionChoiceValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RedirectionChoice
func (i RedirectionChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RedirectionChoice
func (i *RedirectionChoice) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RedirectionChoice should be a string, got %s", data)
	}

	var err error
	*i, err = RedirectionChoiceString(s)
	return err
}

func (i RedirectionChoice) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *RedirectionChoice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of RedirectionChoice: %[1]T(%[1]v)", value)
	}

	val, err := RedirectionChoiceString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for RedirectionChoice
func (i RedirectionChoice) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(i.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for RedirectionChoice
func (i *RedirectionChoice) UnmarshalGQL(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("RedirectionChoice should be a string, got %T", value)
	}

	var err error
	*i, err = RedirectionChoiceString(str)
	return err
}