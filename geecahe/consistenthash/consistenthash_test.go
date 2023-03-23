package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		//fmt.Println(uint32(i))
		return uint32(i)
	})
	hash.Add("6", "4", "2")
	//hash.Add("6")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	hash.Add("8")
	//
	//// 27 should now map to 8.
	testCases["27"] = "8"
	//
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
