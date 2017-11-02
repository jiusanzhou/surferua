package surferua

import (
	"testing"
	"fmt"
)

func TestNewEngineInfo(t *testing.T) {
	maps := []interface{}{
		map[interface{}]interface{}{
			"name": "Blink",
			"version": map[interface{}]interface{}{
				"major": []int{534, 603},
				"minor": []int64{0, 21},
				"patch": []int{0, 10},
			},
		},
		map[interface{}]interface{}{
			"name": "Gecko",
			"version": map[interface{}]interface{}{
				"major": []int{534, 603},
				"minor": []int64{0, 21},
				"patch": []int{0, 10},
			},
		},
	}

	for _, m := range maps {
		ei := NewEngineInfo(m)
		e := ei.Random()
		fmt.Println(e.String())
	}
}
