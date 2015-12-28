package ansible

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/jackspirou/tfs/state"
)

// Inventory parses a terraform statefile
func Inventory(statefile string) error {

	raw, err := ioutil.ReadFile(statefile)
	if err != nil {
		return err
	}

	s, err := state.ReadState(bytes.NewReader(raw))
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", s)

	return nil
}
