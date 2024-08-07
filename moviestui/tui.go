package moviestui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func IntialCommand() string {
	var command string

	form := huh.NewSelect[string]().
		Title("Select Your Option").
		Options(
			huh.NewOption("List all Movies", "L"),
			huh.NewOption("Search By Rating", "R"),
			huh.NewOption("Serach Movie by Name", "N"),
			huh.NewOption("Add New Movie", "A"),
			huh.NewOption("Delete Movie", "D"),
			huh.NewOption("Update Movie", "U"),
			huh.NewOption("Exit", "E"),
		).
		Value(&command)
	form.WithTheme(huh.ThemeCharm())

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return command
}

func UpdateFieldOption() string {
	var field string

	form := huh.NewSelect[string]().
		Title("Select the Moive Field to Update.").
		Options(
			huh.NewOption("Name", "name"),
			huh.NewOption("Director", "director"),
			huh.NewOption("Main Protagonist", "protagonist"),
			huh.NewOption("Main Antagonist", "antagonist"),
			huh.NewOption("Release Date", "releasedate"),
			huh.NewOption("Rating", "rating"),
		).Value(&field)
	form.WithTheme(huh.ThemeCharm())

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return field
}

func ConfimDelete(movieName string) bool {
	var confirm bool
	form := huh.NewConfirm().
		TitleFunc(func() string {
			return fmt.Sprintf("Are you sure you want to delete %q movie\n", movieName)
		}, &movieName).
		Affirmative("Yes").
		Negative("No").
		Value(&confirm).WithTheme(huh.ThemeBase16())

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return confirm
}
