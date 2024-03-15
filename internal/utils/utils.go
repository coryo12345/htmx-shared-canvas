package utils

import (
	"errors"
	"image/color"
	"strconv"
)

func HexColor(hex string) (color.RGBA, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return color.RGBA{}, errors.New("hex strings must be in the form #xxxxxx")
	}
	values, _ := strconv.ParseUint(string(hex[1:]), 16, 32)
	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}, nil
}
