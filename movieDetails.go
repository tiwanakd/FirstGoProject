package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"projecta/movies"
	"strconv"
	"strings"
)

// propmt user to add a new movie
// reutrn a movies.Movie type with new information
func newMovieInfo() (movies.Movie, error) {

	movieName := getInputData("MovieName").MovieName
	//chekc if movie name is duplicate
	if isDuplicate(movieName) {
		return movies.Movie{}, errors.New("error: duplicate movie name")
	}

	director := getInputData("Director").Director
	mainProtagonist := getInputData("MainProtagonist").MainProtagonist
	mainAntagonist := getInputData("MainAntagonist").MainAntagonist
	releaseDate := getInputData("ReleaseDate").ReleaseDate
	rating := getInputData("Rating").Rating

	return movies.Movie{
		MovieName:       movieName,
		Director:        director,
		MainProtagonist: mainProtagonist,
		MainAntagonist:  mainAntagonist,
		ReleaseDate:     releaseDate,
		Rating:          rating,
	}, nil
}

// prompt user to update a movie field
// user will proivde the moviename and field that needs to be updated
func updateMovieInfo(field string) (movies.Movie, error) {
	var updatedMovie movies.Movie

	switch strings.ToLower(field) {
	case "name":
		movieName := getInputData("MovieName").MovieName
		updatedMovie.MovieName = movieName

	case "director", "dir":
		director := getInputData("Director").Director
		updatedMovie.Director = director

	case "protagonist", "pro":
		mainProtagonist := getInputData("MainProtagonist").MainProtagonist
		updatedMovie.MainProtagonist = mainProtagonist

	case "antagonist", "ant":
		mainAntagonist := getInputData("MainAntagonist").MainAntagonist
		updatedMovie.MainAntagonist = mainAntagonist

	case "releasedate", "rdate":
		releaseDate := getInputData("ReleaseDate").ReleaseDate
		updatedMovie.ReleaseDate = releaseDate

	case "rating":
		rating := getInputData("Rating").Rating
		updatedMovie.Rating = rating
	default:
		return movies.Movie{}, fmt.Errorf("invalid Field name")
	}

	return updatedMovie, nil
}

/*
check if the movie name provided is a duplicate
use SearchMoviebyName to check if a field with given moive name exits
compare only the movie names using strings.EqualFold
*/
func isDuplicate(movieName string) bool {
	var movie movies.Movie
	serchMovie, _ := movie.SearchMoviebyName(movieName)
	return strings.EqualFold(serchMovie.MovieName, movieName)
}

// get the input from the termial
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

// Function that get the input from the user as per the given prompt
// Returns a Moive type
func getInputData(prompt string) movies.Movie {

	//create a nil movie type
	var movie movies.Movie

	//prompt the user for the required field
	fmt.Printf("%s: ", prompt)
	value := inputDetails()
	if value == "" {
		log.Fatalln("error: cannot accept empty values")
	}

	//assign the values inputted by the user to their corresponding property
	switch prompt {
	case "MovieName":
		movie.MovieName = value
	case "Director":
		movie.Director = value
	case "MainProtagonist":
		movie.MainProtagonist = value
	case "MainAntagonist":
		movie.MainAntagonist = value
	case "ReleaseDate":
		//date needs to be converted to time.Time type
		//user GetDate Function provided by movies package
		releaseDateConverted, err := movies.GetDate(value)
		if err != nil {
			log.Fatalln("error: invalid date entered, date should follow the following format: 2024-01-01")
		}
		movie.ReleaseDate = releaseDateConverted
	case "Rating":
		//provided string values needs to be converted to float
		ratingtoFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatalln("error: invaid rating")
		}

		if ratingtoFloat64 > 10 {
			log.Fatalln("error: rating cannot be more that 10")
		}

		movie.Rating = ratingtoFloat64
	}

	return movie
}
