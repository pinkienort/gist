package cmd

import (
	"fmt"
	"os"

	"github.com/b4b4r07/gist/cli"
	"github.com/b4b4r07/gist/cli/gist"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete gist files",
	Long:  "Delete gist files on the remote",
	RunE:  delete,
}

func delete(cmd *cobra.Command, args []string) (err error) {
	screen, err := cli.NewScreen()
	if err != nil {
		return
	}

	items, err := screen.Select()
	if err != nil {
		return
	}

	// lines = lines.Uniq()
	// if len(lines) > 0 {
	// 	cli.NewCache().Clear()
	// }

	client, err := gist.NewClient(cli.Conf.Gist.Token)
	if err != nil {
		return
	}

	for _, item := range items {
		err := client.Delete(item.ID)
		if err != nil {
			continue
		}
		// remove from local
		path, err := cli.GetPath(item.ID)
		if err != nil {
			continue
		}
		_ = os.Remove(path)
		fmt.Printf("Deleted %s\n", item.ID)
	}

	return nil
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}