package gomono

import (
	"log"

	"github.com/golang/freetype/truetype"
	. "golang.org/x/image/font/gofont/gomono"
)

var Font, err = truetype.Parse(TTF)

func init() {
	if err != nil {
		log.Fatalln("goregular", err)
	}
}
