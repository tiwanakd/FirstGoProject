package moviestui

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"projecta/movies"

	"github.com/charmbracelet/huh"
)

// propmt user to add a new movie
// reutrn a movies.Movie type with new information
func NewMovieInfo() (movies.Movie, error) {

	movieName, err := getInputData("MovieName")
	if err != nil {
		return movies.Movie{}, err
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
func UpdateMovieInfo(field string) (movies.Movie, error) {
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

// get the input from the termial
func InputDetails(title string) string {

	var input string
	form := huh.NewInput().
		Title(title).
		Value(&input)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return input
}

// Function that get the input from the user as per the given prompt
// Returns a Moive type
func getInputData(prompt string) (movies.Movie, error) {

	//create a nil movie type
	var movie movies.Movie

	//prompt the user for the required field
	value := InputDetails(prompt)
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
