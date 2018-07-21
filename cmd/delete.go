package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// ゴミ箱に入っている、指定した一つのファイルを削除する。
func del(file string) bool {
	fmt.Printf("target: %s\n", file)
	fmt.Println("'yes' or 'no'")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Text() != "yes" && scanner.Text() != "no" {
		fmt.Println("'yes' or 'no'")
		scanner.Scan()
	}

	if scanner.Text() == "yes" {
		return true
	} else {
		return false
	}
}

func createDeleteCmd(trashPath string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a file in the trash",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]

			if del(file) {
				if err := os.RemoveAll(trashPath + "/" + file); err != nil {
					log.Fatalln(err)
				}
			}
		},
	}

	return cmd
}
