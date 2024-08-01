package main

import (
	"bufio"
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
		fmt.Println("R - Enter Rating")
		fmt.Println("N - Serach Movie by Name")
		fmt.Println("A - Add New Movie")
		fmt.Println("D - Delete a Movie")
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
			rating, _ := strconv.ParseFloat(inputDetails(), 64)

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
			fmt.Printf("%s deleted!", movieName)
			fmt.Println()
		case "E":
			fmt.Println("Exiting...")
			break loop
		default:
			fmt.Println("Invalid Command")
		}

	}
}

func inputDetails() string {
	reader := bufio.NewReader(os.Stdin)
	info, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	info = strings.TrimSuffix(info, "\n")
	return info
}
