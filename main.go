package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"projecta/movies"
	"projecta/moviestui"

	"github.com/fatih/color"
)

func main() {
	var movie movies.Movie
	var command string
	errorRed := color.New(color.FgRed)

loop:
	for {
		command = moviestui.IntialCommand()

		switch strings.ToUpper(command) {
		case "L":
			allMovies, err := movie.GetAllMovies()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			movie.PrintAll(allMovies)
		case "R":
			rating, err := strconv.ParseFloat(moviestui.InputDetails("Enter Rating: "), 64)
			if err != nil {
				errorRed.Fprintln(os.Stderr, "Invalid Rating")
				continue
			}

			moviesByRating, err := movie.GetMoviesByRating(rating)
			if err != nil {
				errorRed.Fprintln(os.Stderr, err)
				continue
			}
			color.Green("Movies with rating greater than %.1f\n", rating)
			movie.PrintAll(moviesByRating)
		case "N":
			movieName := moviestui.InputDetails("Enter Movie Name: ")
			searchMovies, err := movie.SearchMoviesbyName(movieName)
			if err != nil {
				errorRed.Fprintln(os.Stderr, err)
				continue
			}
			movie.PrintAll(searchMovies)

		case "A":
			fmt.Println("Enter New Movie Details")
			newMovie, err := moviestui.NewMovieInfo()
			if err != nil {
				errorRed.Fprintln(os.Stderr, err)
				continue
			}
			err = newMovie.AddMovie()
			if err != nil {
				errorRed.Fprintf(os.Stderr, "Unable to add movie, received error: %v\n", err)
				continue
			}
			color.Green("New Movie added!")
			newMovie.Print()
		case "D":
			movieName := moviestui.InputDetails("Enter the Moive name to Delete: ")

			if !moviestui.ConfimDelete(movieName) {
				continue
			}

			if err := movie.DeleteMovie(movieName); err != nil {
				errorRed.Fprintf(os.Stderr, "unable to delete movie, receveid error: %v\n", err)
				continue
			}
			color.Green("%q deleted!\n", movieName)

		case "U":
			movieName := moviestui.InputDetails("Movie to Update: ")
			field := moviestui.UpdateFieldOption()
			updatedMovie, err := moviestui.UpdateMovieInfo(field)
			if err != nil {
				errorRed.Fprintln(os.Stderr, err)
				continue
			}

			err = updatedMovie.UpdateMovie(movieName, field)
			if err != nil {
				errorRed.Fprintf(os.Stderr, "Unable to update movie, error: %v\n", err)
				continue
			}
			color.Green("%q updated!\n", movieName)

		case "E":
			color.Yellow("Exiting...\n")
			//fmt.Println("Exiting...")
			break loop
		default:
			fmt.Println("Invalid Command")
		}
	}
}
