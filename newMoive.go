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
	movieName, err := getValue()
	if err != nil {
		return movies.Movie{}, err
	}

	//chekc if movie name is duplicate
	if isDuplicate(movieName) {
		return movies.Movie{}, errors.New("error: duplicate movie name")
	}

	fmt.Print("Director: ")
	director, err := getValue()
	if err != nil {
		return movies.Movie{}, err
	}

	fmt.Print("Main Protagonist: ")
	mainProtagonist, err := getValue()
	if err != nil {
		return movies.Movie{}, err
	}

	fmt.Print("Main Antagonist: ")
	mainAntagonist, err := getValue()
	if err != nil {
		return movies.Movie{}, err
	}

	fmt.Print("Release Date (2024-01-01): ")
	releaseDate, err := getValue()
	if err != nil {
		return movies.Movie{}, err
	}

	releaseDateConverted, err := movies.GetDate(releaseDate)
	if err != nil {
		return movies.Movie{}, errors.New("error: invalid date entered, date should follow the following format: 2024-01-01")
	}

	fmt.Print("Rating: ")
	rating, err := getValue()
	if err != nil {
		return movies.Movie{}, err
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
check if the movie name provided is a duplicate
use SearchMoviebyName to check if a field with given moive name exits
compare only the movie names using strings.EqualFold
*/
func isDuplicate(movieName string) bool {
	var movie movies.Movie
	serchMovie, _ := movie.SearchMoviebyName(movieName)
	return strings.EqualFold(serchMovie.MovieName, movieName)
}

func getValue() (string, error) {
	value := inputDetails()
	if value == "" {
		return "", fmt.Errorf("error: cannot accept empty values")
	}

	return value, nil
}
