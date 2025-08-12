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

			size, err := path_size.GetSize(path, false, false)

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

	// flag.Parse()

	// if flag.NArg() < 1 {
	// 	fmt.Fprintln(os.Stderr, "Error: path is required")
	// 	os.Exit(1)
	// }

	// path := flag.Arg(0)

	// size, err := getSize(path, flagRecursive, flagAll)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error:", err)
	// 	os.Exit(0)
	// }

	// var sizeStr string

	// if flagHumanReadable {
	// 	sizeStr = humanize.Bytes(uint64(size))
	// } else {
	// 	sizeStr = fmt.Sprint(size)
	// }

	// fmt.Println(sizeStr, path)
}
