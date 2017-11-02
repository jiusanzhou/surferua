package surferua

import (
	"fmt"
	"testing"
)

func TestNewVersionInfo(t *testing.T) {
	maps := []interface{}{
		map[interface{}]interface{}{
			"major": []int{534, 603},
			"minor": []int{0, 21},
			"patch": []int{0, 10},
		},
		map[interface{}]interface{}{
			"major": []int{534, 603},
			"minor": []int64{0, 21},
			"patch": []int{0, 10},
		},
		map[interface{}]interface{}{
			"major": []int{534, 603},
			"minor": []int{0, 21},
			"patch": 0,
		},
	}

	for _, m := range maps {
		vi := NewVersionInfo(m)
		s := vi.Random()
		fmt.Println(s.String())
	}
}
