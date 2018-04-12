// Code generated by "stringer -type=ItemType -output=gen_ItemType.go"; DO NOT EDIT.

package itemtype

import "strconv"

const _ItemType_name = "ItemErrorItemEOFItemTextItemCodeFenceItemCodeItemTmplBlockStartItemJsonBlockStartItemBlockEndItemCommEndItemArgItemQuoteArgItemArgCommentItemOption"

var _ItemType_index = [...]uint8{0, 9, 16, 24, 37, 45, 63, 81, 93, 104, 111, 123, 137, 147}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
