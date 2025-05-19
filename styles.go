package runner

import (
	"github.com/morikuni/aec"
)

type Style int

const (
    StyleTitle Style = iota
    StyleMessage
    StyleSuccess
    StyleError
)

var Styles = map[Style]aec.ANSI{
    StyleTitle: aec.CyanF,
    StyleMessage: aec.Color8BitF(aec.NewRGB8Bit(132, 132, 132)),
    StyleSuccess: aec.GreenF,
    StyleError: aec.RedF,
}

func ApplyStyle(style Style, s string) string {
    return Styles[style].Apply(s)
}