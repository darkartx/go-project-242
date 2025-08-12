package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	path_size "github.com/darkartx/go-project-242"
	cli "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    "all",
				Value:   false,
				Usage:   "include hidden files and directories",
				Aliases: []string{"a"},
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Value:   false,
				Usage:   "include hidden files and directories",
				Aliases: []string{"r"},
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "path",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")

			if len(path) == 0 {
				return errors.New("path requred")
			}

			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")

			size, err := path_size.GetSize(path, recursive, all)

			if err != nil {
				return err
			}

			fmt.Println(path_size.FormatSize(size, cmd.Bool("human")), path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
