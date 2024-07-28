package main

import (
	"bufio"
	"fmt"
	"os"
	"projecta/movies"
	"strings"
)

func main() {
	var movie movies.Movie
	const fileName = "marvel_movies.csv"

	var command string

loop:
	for {

		fmt.Print("Enter Command: ")
		fmt.Scan(&command)

		switch command {
		case "A":
			allMovies, err := movie.GetAllMovies(fileName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			movie.PrintAll(allMovies)
		case "R":
			var rating float64
			fmt.Print("Enter Rating: ")
			fmt.Scan(&rating)
			moviesByRating, err := movie.GetMoviesByRating(fileName, rating)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			movie.PrintAll(moviesByRating)
		case "N":
			fmt.Print("Enter Movie Name: ")
			movieName := inputMovieName()
			serchMovie, err := movie.SearchMoviebyName(fileName, movieName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			serchMovie.PrintMovieDetails()

		case "E":
			break loop
		default:
			fmt.Println("Invalid Command")
		}

	}
}

func inputMovieName() string {

	reader := bufio.NewReader(os.Stdin)
	moveName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	moveName = strings.TrimSuffix(moveName, "\n")
	return moveName

}
