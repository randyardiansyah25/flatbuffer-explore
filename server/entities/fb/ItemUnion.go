// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import "strconv"

type ItemUnion byte

const (
	ItemUnionNONE        ItemUnion = 0
	ItemUnionArchiveItem ItemUnion = 1
	ItemUnionHistoryItem ItemUnion = 2
)

var EnumNamesItemUnion = map[ItemUnion]string{
	ItemUnionNONE:        "NONE",
	ItemUnionArchiveItem: "ArchiveItem",
	ItemUnionHistoryItem: "HistoryItem",
}

var EnumValuesItemUnion = map[string]ItemUnion{
	"NONE":        ItemUnionNONE,
	"ArchiveItem": ItemUnionArchiveItem,
	"HistoryItem": ItemUnionHistoryItem,
}

func (v ItemUnion) String() string {
	if s, ok := EnumNamesItemUnion[v]; ok {
		return s
	}
	return "ItemUnion(" + strconv.FormatInt(int64(v), 10) + ")"
}
