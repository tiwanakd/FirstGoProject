package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"projecta/movies"
	"projecta/moviestui"
)

func main() {
	var movie movies.Movie
	var command string

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
			fmt.Print("Enter Rating: ")
			rating, err := strconv.ParseFloat(inputDetails(), 64)
			if err != nil {
				fmt.Println("Invalid Rating")
				continue
			}

			moviesByRating, err := movie.GetMoviesByRating(rating)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Printf("Movies with rating greater than %.1f\n", rating)
			movie.PrintAll(moviesByRating)
		case "N":
			fmt.Print("Enter Movie Name: ")
			movieName := inputDetails()
			serchMovie, err := movie.SearchMoviebyName(movieName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(serchMovie)
		case "A":
			fmt.Println("Enter New Movie Details")
			newMovie, err := newMovieInfo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = newMovie.AddMovie()
			if err != nil {
				fmt.Printf("Unable to add movie, received error %v\n", err)
				continue
			}
			fmt.Println("New Movie added!")
			fmt.Println(newMovie)
		case "D":
			fmt.Print("Enter the Moive name to Delete: ")
			movieName := inputDetails()

			if !moviestui.ConfimDelete(movieName) {
				continue
			}

			if err := movie.DeleteMovie(movieName); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s deleted!\n", movieName)

		case "U":
			fmt.Print("Movie to Update: ")
			movieName := inputDetails()

			field := moviestui.UpdateFieldOption()

			updatedMovie, err := updateMovieInfo(field)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = updatedMovie.UpdateMovie(movieName, field)
			if err != nil {
				fmt.Printf("Unable to update movie, error %v\n", err)
				continue
			}
			fmt.Printf("%s updated!\n", movieName)

		case "E":
			fmt.Println("Exiting...")
			break loop
		default:
			fmt.Println("Invalid Command")
		}
	}
}
