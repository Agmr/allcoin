package allcoin

import (
	"fmt"
	"testing"
)

func TestNewFromAPI(t *testing.T) {

	coins, err := NewFromAPI()

	if err != nil {
		t.Fatalf("Error occured while parsing coins: %v\n", err)
	}

	fmt.Println(len(coins))
}
