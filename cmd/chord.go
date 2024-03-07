/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"fretz/music"
	"github.com/spf13/cobra"
	"strings"
)

// chordCmd represents the chord command
var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: command,
}

func command(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		err := cmd.Help()
		if err != nil {
			return
		}
		return
	}

	note := music.NewNoteFromString(strings.ToUpper(args[0]))
	chordName := args[1]

	switch chordName {
	case "major":
		fmt.Println(note.MajorChord())
	case "minor":
		fmt.Println(note.MinorChord())
	case "diminished":
		fmt.Println(note.DiminishedChord())
	case "augmented":
		fmt.Println(note.AugmentedChord())
	case "seventh":
		fmt.Println(note.SeventhChord())
	case "major-seventh":
		fmt.Println(note.MajorSeventhChord())
	case "minor-seventh":
		fmt.Println(note.MinorSeventhChord())
	case "minor-major-seventh":
		fmt.Println(note.MinorMajorSeventhChord())
	case "diminished-seventh":
		fmt.Println(note.DiminishedSeventhChord())
	case "half-diminished-seventh":
		fmt.Println(note.HalfDiminishedSeventhChord())
	case "augmented-seventh":
		fmt.Println(note.AugmentedSeventhChord())
	case "augmented-major-seventh":
		fmt.Println(note.AugmentedMajorSeventhChord())
	case "ninth":
		fmt.Println(note.NinthChord())
	case "major-ninth":
		fmt.Println(note.MajorNinthChord())
	case "minor-ninth":
		fmt.Println(note.MinorNinthChord())

	default:
		fmt.Println("Unknown chord")
	}

}

func init() {
	rootCmd.AddCommand(chordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
