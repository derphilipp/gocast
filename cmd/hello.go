/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
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
		SearchPodcast(searchterm, countrycode)
	},
}

var globalrankCmd = &cobra.Command{
	Use:   "globalrank COUNTRYCODE ID",
	Short: "Display rank of podcast of all podcasts in given country",
	Run: func(cmd *cobra.Command, args []string) {
		podcastID := strings.Join(args[1:], " ")
		countrycode := args[0]
		printPodcastGlobalRank(podcastID, countrycode)
	},
}
var genrerankCmd = &cobra.Command{
	Use:   "genrerank COUNTRYCODE GENREID ID",
	Short: "Display rank of podcast of all podcasts in given country",
	Run: func(cmd *cobra.Command, args []string) {
		podcastID := strings.Join(args[2:], " ")
		genreID := args[1]
		countrycode := args[0]
		printPodcastGenreRank(podcastID, genreID, countrycode)
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
