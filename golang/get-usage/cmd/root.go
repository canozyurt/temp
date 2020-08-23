package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "get-usage",
	Short: "Get pods with their resource consumption highlighted",
	Long: `Fetches both pod list and top list to colorize pods based on their resource consumption. >%50 Yellow >%80 Red"`,
	Run: func(cmd *cobra.Command, args []string) {

		MyList := make(map[string]*PodAttributes)
		kubectl := newClientSetFromConfigFile()
		pods := kubectl.getPods()
		err := parsePods(pods, &MyList)
		if err != nil {
			panic(err.Error())
		}

		metrics := kubectl.topPods()
		err = parseMetrics(metrics, &MyList)
		if err != nil {
			panic(err.Error())
		}

		for key := range MyList {
			ColorPrint(MyList[key])
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
