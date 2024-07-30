package main

import (
	"errors"
	"fmt"
	"projecta/movies"
	"strconv"
	"strings"
)

/*
	Function that pompts the user to enter the new movie data
	Return a moview type that will be used in AddNewMovie function
*/

func newMovieInfo() (movies.Movie, error) {

	//prompt the user for the required record info
	//use inputDetails() function to get values from Stdin
	//check if the provided input is empty
	//if so return empty movie type and a custom error

	fmt.Print("Movie Name: ")
	movieName := inputDetails()
	if movieName == "" {
		return movies.Movie{}, errors.New("error: no movie name entered")
	}

	//chekc if movei name is duplicate
	if isDuplicate(movieName) {
		return movies.Movie{}, errors.New("error: duplicate movie name")
	}

	fmt.Print("Director: ")
	director := inputDetails()
	if director == "" {
		return movies.Movie{}, errors.New("error: no director name entered")
	}

	fmt.Print("Main Protagonist: ")
	mainProtagonist := inputDetails()
	if mainProtagonist == "" {
		return movies.Movie{}, errors.New("error: no Main Protagonist name entered")
	}

	fmt.Print("Main Antagonist: ")
	mainAntagonist := inputDetails()
	if mainAntagonist == "" {
		return movies.Movie{}, errors.New("error: no Main Antagonist name entered")
	}

	fmt.Print("Release Date (2024-01-01): ")
	releaseDate := inputDetails()
	if releaseDate == "" {
		return movies.Movie{}, errors.New("error: invalid date")
	}

	releaseDateConverted, err := movies.GetDate(releaseDate)
	if err != nil {
		return movies.Movie{}, errors.New("error: invalid date entered, date should follow the following format: 2024-01-01")
	}

	fmt.Print("Rating: ")
	rating := inputDetails()
	if rating == "" {
		return movies.Movie{}, errors.New("error: no Rating entered")
	}

	ratingtoFloat64, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		return movies.Movie{}, fmt.Errorf("error: invaid rating")
	}

	if ratingtoFloat64 > 10 {
		return movies.Movie{}, errors.New("error: rating cannot be more that 10")
	}

	return movies.Movie{
		MovieName:       movieName,
		Director:        director,
		MainProtagonist: mainProtagonist,
		MainAntagonist:  mainAntagonist,
		ReleaseDate:     releaseDateConverted,
		Rating:          ratingtoFloat64,
	}, nil
}

/*
check if the movie name provided is a dupliate
use SearchMoviebyName to check if a field with given moive name exits
compare only the movie names using strings.EqualFold
*/
func isDuplicate(movieName string) bool {
	var movie movies.Movie
	serchMovie, _ := movie.SearchMoviebyName(fileName, movieName)
	return strings.EqualFold(serchMovie.MovieName, movieName)
}

// func emptyValue(value string) (movies.Movie, error) {
// 	if value == "" {
// 		return movies.Movie{}, fmt.Errorf("error: no %s name entered", value)
// 	}
// 	return
// }
