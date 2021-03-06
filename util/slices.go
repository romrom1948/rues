// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package util

import (
	"errors"
)

// returns the first match in a slice or an error if none can be found
func IsAt(str string, slice []string) (int, error) {
	for i, s := range slice {
		if s == str {
			return i, nil
		}
	}

	str += " not in slice"
	return 0, errors.New(str)
}

func IsIn(str string, slice []string) bool {
	_, err := IsAt(str, slice);
	
	if err != nil { return false } else { return true }
}
