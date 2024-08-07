package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"projecta/movies"
)

// propmt user to add a new movie
// reutrn a movies.Movie type with new information
func newMovieInfo() (movies.Movie, error) {

	movieName, err := getInputData("MovieName")
	if err != nil {
		return movies.Movie{}, err
	}
	//chekc if movie name is duplicate
	if isDuplicate(movieName.MovieName) {
		return movies.Movie{}, errors.New("error: duplicate movie name")
	}

	moivieDirector, err := getInputData("Director")
	if err != nil {
		return movies.Movie{}, err
	}

	movieMainProtagonist, err := getInputData("MainProtagonist")
	if err != nil {
		return movies.Movie{}, err
	}

	movieMainAntagonist, err := getInputData("MainAntagonist")
	if err != nil {
		return movies.Movie{}, err
	}

	movieReleaseDate, err := getInputData("ReleaseDate")
	if err != nil {
		return movies.Movie{}, err
	}

	movieRating, err := getInputData("Rating")
	if err != nil {
		return movies.Movie{}, err
	}

	return movies.Movie{
		MovieName:       movieName.MovieName,
		Director:        moivieDirector.Director,
		MainProtagonist: movieMainProtagonist.MainProtagonist,
		MainAntagonist:  movieMainAntagonist.MainAntagonist,
		ReleaseDate:     movieReleaseDate.ReleaseDate,
		Rating:          movieRating.Rating,
	}, nil
}

// prompt user to update a movie field
// user will proivde the moviename and field that needs to be updated
func updateMovieInfo(field string) (movies.Movie, error) {
	var updatedMovie movies.Movie

	switch strings.ToLower(field) {
	case "name":
		movie, err := getInputData("MovieName")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MovieName = movie.MovieName

	case "director":
		movie, err := getInputData("Director")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.Director = movie.Director

	case "protagonist":
		movie, err := getInputData("MainProtagonist")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MainProtagonist = movie.MainProtagonist

	case "antagonist":
		movie, err := getInputData("MainAntagonist")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MainAntagonist = movie.MainAntagonist

	case "releasedate":
		movie, err := getInputData("ReleaseDate")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.ReleaseDate = movie.ReleaseDate

	case "rating":
		movie, err := getInputData("Rating")
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.Rating = movie.Rating
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
func getInputData(prompt string) (movies.Movie, error) {

	//create a nil movie type
	var movie movies.Movie

	//prompt the user for the required field
	fmt.Printf("%s: ", prompt)
	value := inputDetails()
	if value == "" {
		return movies.Movie{}, errors.New("error: cannot accept empty values")
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
			return movies.Movie{}, errors.New("error: invalid date entered, date should follow the following format: 2024-01-01")
		}
		movie.ReleaseDate = releaseDateConverted
	case "Rating":
		//provided string values needs to be converted to float
		ratingtoFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return movies.Movie{}, errors.New("error: invaid rating")
		}

		if ratingtoFloat64 > 10 {
			return movies.Movie{}, errors.New("error: rating cannot be more that 10")
		}

		movie.Rating = ratingtoFloat64
	}

	return movie, nil
}
