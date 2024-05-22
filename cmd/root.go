package cmd

import (
	"blog-golang/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "notion",
	Short: "A brief description of your command",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("NOTION_TOKEN")
		dbId := viper.GetString("NOTION_DB_ID")

		if token == "" {
			log.Error("NOTION_TOKEN is not set")
			os.Exit(1)
		}
		if dbId == "" {
			log.Error("NOTION_DB_ID is not set")
			os.Exit(1)

		}
		wr := services.NewNotionWrapper(token, dbId)

		writer := services.NewWriter(postPath)
		writer.DeleteFilesInFolder()
		posts := wr.GetPostsFromNotionDB()

		for post := range posts {
			writer.WriteToFile(posts[post].String(), posts[post].TitleToFilename())
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var postPath = ""

func init() {
	//rootCmd.AddCommand(notionCmd)
	rootCmd.Flags().StringVarP(&postPath, "path", "p", "", "path to the folder where the posts will be written")
	rootCmd.MarkFlagsOneRequired("path")

	viper.AutomaticEnv()
}
