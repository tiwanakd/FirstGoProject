package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

func updateMovieInfo(field string) (movies.Movie, error) {
	var updatedMovie movies.Movie

	switch strings.ToLower(field) {
	case "name":
		fmt.Print("Enter the new Movie Name: ")
		newName, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MovieName = newName

	case "director", "dir":
		fmt.Print("Enter the new Dirctor Name: ")
		newDirector, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.Director = newDirector

	case "protagonist", "pro":
		fmt.Print("Enter the new Main Protagonist: ")
		newProtagonist, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MainProtagonist = newProtagonist

	case "antagonist", "ant":
		fmt.Print("Enter the new Main Antagonist: ")
		newAntagonist, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}
		updatedMovie.MainAntagonist = newAntagonist

	case "releasedate", "rdate":
		fmt.Print("Enter the new Relase Date: ")
		newDate, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}

		releaseDateConverted, err := movies.GetDate(newDate)
		if err != nil {
			return movies.Movie{}, errors.New("error: invalid date entered, date should follow the following format: 2024-01-01")
		}

		updatedMovie.ReleaseDate = releaseDateConverted

	case "rating":
		fmt.Print("Enter the new Rating: ")
		newRating, err := getValue()
		if err != nil {
			return movies.Movie{}, err
		}
		ratingtoFloat64, err := strconv.ParseFloat(newRating, 64)
		if err != nil {
			return movies.Movie{}, fmt.Errorf("error: invaid rating")
		}

		if ratingtoFloat64 > 10 {
			return movies.Movie{}, errors.New("error: rating cannot be more that 10")
		}

		updatedMovie.Rating = ratingtoFloat64
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

func getValue() (string, error) {
	value := inputDetails()
	if value == "" {
		return "", fmt.Errorf("error: cannot accept empty values")
	}

	return value, nil
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
