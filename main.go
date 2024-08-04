package main

import (
	"fmt"
	"os"
	"projecta/movies"
	"strconv"
	"strings"
)

func main() {
	var movie movies.Movie
	var command string

loop:
	for {
		fmt.Println("Choose from the following: ")
		fmt.Println("L - List all Movies")
		fmt.Println("R - Search By Rating")
		fmt.Println("N - Serach Movie by Name")
		fmt.Println("A - Add New Movie")
		fmt.Println("D - Delete a Movie")
		fmt.Println("U - Update a Movie Field")
		fmt.Println("E - Exit")

		fmt.Print("Enter Command: ")
		fmt.Scan(&command)

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
				return
			}
			fmt.Printf("Movies with rating greater than %.1f\n", rating)
			movie.PrintAll(moviesByRating)
		case "N":
			fmt.Print("Enter Movie Name: ")
			movieName := inputDetails()
			serchMovie, err := movie.SearchMoviebyName(movieName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
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
				return
			}
			fmt.Println("New Movie added!")
			fmt.Println(newMovie)
		case "D":
			fmt.Print("Enter the Moive name to Delete: ")
			movieName := inputDetails()
			if err := movie.DeleteMovie(movieName); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s deleted!\n", movieName)

		case "U":
			fmt.Print("Movie to Update: ")
			movieName := inputDetails()

			fmt.Print("Field to Update (name|dir|pro|ant|rdate|rating): ")
			field := inputDetails()

			updatedMovie, err := updateMovieInfo(field)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = updatedMovie.UpdateMovie(movieName, field)
			if err != nil {
				fmt.Printf("Unable to update movie, error %v\n", err)
				return
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
