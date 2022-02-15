package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return fmt.Println(emoji.Sprint("Hello :world_map:!"))
}

