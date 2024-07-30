package main

import (
	"bufio"
	"fmt"
	"os"
	"projecta/movies"
	"strconv"
	"strings"
)

const fileName = "marvel_movies.csv"

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
		fmt.Println("E - Exit")

		fmt.Print("Enter Command: ")
		fmt.Scan(&command)

		switch strings.ToUpper(command) {
		case "L":
			allMovies, err := movie.GetAllMovies(fileName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			movie.PrintAll(allMovies)
		case "R":
			fmt.Print("Enter Rating: ")
			rating, _ := strconv.ParseFloat(inputDetails(), 64)

			moviesByRating, err := movie.GetMoviesByRating(fileName, rating)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Printf("Movies with rating greater than %.1f\n", rating)
			movie.PrintAll(moviesByRating)
		case "N":
			fmt.Print("Enter Movie Name: ")
			movieName := inputDetails()
			serchMovie, err := movie.SearchMoviebyName(fileName, movieName)
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

			err = movies.AddNewMovie(fileName, newMovie)
			if err != nil {
				fmt.Printf("Unable to add movie, received error %v\n", err)
				return

			}
			fmt.Println("New Movie added!")
			fmt.Println(newMovie)

		case "E":
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
