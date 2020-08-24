package cmd

import (
	"github.com/fatih/color"
	v1 "k8s.io/api/core/v1"
)

func ColorPrint (MyList *map[string]*PodAttributes, podList *v1.PodList) {

	var c *color.Color
	for _, pod := range podList.Items {

		if (*MyList)[pod.Name].cpuLimit == 0 || (*MyList)[pod.Name].memoryLimit == 0 {
			c = color.New(color.FgHiMagenta)
		} else if (*MyList)[pod.Name].cpuRatio > 80.0 || (*MyList)[pod.Name].memoryRatio > 80.0 {
			c = color.New(color.FgHiRed)
		} else if (*MyList)[pod.Name].cpuRatio > 50.0 || (*MyList)[pod.Name].memoryRatio > 50.0 {
			c = color.New(color.FgHiYellow)
		} else {
			c = color.New(color.FgHiWhite)
		}
		_, err := c.Println((*MyList)[pod.Name])
		if err != nil {
			panic(err.Error())
		}
	}
}