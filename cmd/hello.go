package cmd

import (
	"strings"

	"github.com/derphilipp/gocast/cast"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search []",
	Short: "Search for elements",
}

var searchPodcastCmd = &cobra.Command{
	Use:   "podcast [no options!]",
	Short: "Search for podcast information",
	Long: `Search a podcast on the Apple podcast directory
and display relevant informations for further search. For example:

gocast search podcast DE Triumvirat`,

	Run: func(cmd *cobra.Command, args []string) {
		searchterm := strings.Join(args[1:], " ")
		countrycode := args[0]
		cast.SearchPodcast(searchterm, countrycode)
	},
}

var globalrankCmd = &cobra.Command{
	Use:   "globalrank COUNTRYCODE ID",
	Short: "Display rank of a podcasts a genre in a given country",
	Long: `Search a podcast on the Apple podcast charts.
Print rank if found, else return errorcode to os. For example:

gocast globalrank DE 1220156551`,
	Run: func(cmd *cobra.Command, args []string) {
		podcastID := strings.Join(args[1:], " ")
		countrycode := args[0]
		cast.PrintPodcastGlobalRank(podcastID, countrycode)
	},
}
var genrerankCmd = &cobra.Command{
	Use:   "genrerank COUNTRYCODE GENREID ID",
	Short: "Display rank of a podcasts in a given country within its genre",
	Long: `Search a podcast on the Apple podcast charts if a give genre.
Print rank if found, else return errorcode to os. For example:

gocast genrerank DE 1487 1220156551`,
	Run: func(cmd *cobra.Command, args []string) {
		podcastID := strings.Join(args[2:], " ")
		genreID := args[1]
		countrycode := args[0]
		cast.PrintPodcastGenreRank(podcastID, genreID, countrycode)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(globalrankCmd)
	rootCmd.AddCommand(genrerankCmd)
	searchCmd.AddCommand(searchPodcastCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
