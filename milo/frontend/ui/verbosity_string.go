// Code generated by "stringer -type=Verbosity"; DO NOT EDIT.

package ui

import "strconv"

const _Verbosity_name = "NormalHiddenInteresting"

var _Verbosity_index = [...]uint8{0, 6, 12, 23}

func (i Verbosity) String() string {
	if i < 0 || i >= Verbosity(len(_Verbosity_index)-1) {
		return "Verbosity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Verbosity_name[_Verbosity_index[i]:_Verbosity_index[i+1]]
}