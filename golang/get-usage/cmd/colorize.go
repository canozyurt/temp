package cmd

import (
	"github.com/fatih/color"
)

func ColorPrint (p *PodAttributes) {

	var c *color.Color

	if p.cpuLimit == 0 || p.memoryLimit == 0 {
		c = color.New(color.FgHiMagenta)
	} else if p.cpuRatio > 80.0 || p.memoryRatio > 80.0 {
		c = color.New(color.FgHiRed)
	} else if p.cpuRatio > 50.0 || p.memoryRatio > 50.0{
		c = color.New(color.FgHiYellow)
	} else {
		c = color.New(color.FgHiWhite)
	}
	_, err := c.Println(p)
	if err != nil {
		panic(err.Error())
	}

}