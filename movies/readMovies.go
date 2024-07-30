package movies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Movie struct {
	MovieName       string
	Director        string
	MainProtagonist string
	MainAntagonist  string
	ReleaseDate     time.Time
	Rating          float64
}

func (mv Movie) PrintAll(allmovies *[]Movie) {
	for _, movie := range *allmovies {
		fmt.Println(movie)
	}
}

// Implement the Stringer Interface
func (mv Movie) String() string {

	movieName := fmt.Sprintf("%s:\n", mv.MovieName)
	director := fmt.Sprintf("\tDirector: %s\n", mv.Director)
	mainProtagonist := fmt.Sprintf("\tMain Protagonist: %s\n", mv.MainProtagonist)
	mainAntagonist := fmt.Sprintf("\tMain Antagonist: %s\n", mv.MainAntagonist)
	releaseDate := fmt.Sprintf("\tRelease Date: %v\n", mv.ReleaseDate.Format("2006-Jan-02"))
	rating := fmt.Sprintf("\tRating: %.2f\n", mv.Rating)
	lineBreak := fmt.Sprintln("-------------------------------------------------------")

	return movieName + director + mainProtagonist + mainAntagonist + releaseDate + rating + lineBreak
}

// Since date is being provided as string
// Fucntion to return a Time struct for given date
func getDate(dateString string) time.Time {
	const shortForm = "2006-01-02"
	t, _ := time.Parse(shortForm, dateString)
	return t
}

// Method to get all the movies from the csv file
// Using ReadAll() will reuturn a string slices of slice
// Arrange them into the Movie Strutct
// Make a slice of Movie structs and return a pointer to it.
func (mv Movie) GetAllMovies(filepath string) (*[]Movie, error) {

	moviesFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	csvReader := csv.NewReader(moviesFile)

	csvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	csvData = csvData[1:] //skipping the first row for column names

	//make a movies slice which has the length of slice returned by readMoviesFile
	movies := make([]Movie, len(csvData))

	for index, movie := range csvData {
		mv.MovieName = movie[0]
		mv.Director = movie[1]
		mv.MainProtagonist = movie[2]
		mv.MainProtagonist = movie[3]
		mv.ReleaseDate = getDate(movie[4])
		mv.Rating, _ = strconv.ParseFloat(movie[5], 64) //convert to float.

		//Since the movies slice is already made with len(csvData)
		//It should hold 0 val for Movie on each index till len(csvData)
		//Append cannot be used here as this start appening to the end of slice
		//With zero values still being precedded.
		movies[index] = mv
	}

	return &movies, nil
}

// Function to get the Movie by given name
func (mv Movie) SearchMoviebyName(file, movieName string) (Movie, error) {
	moviesFile, err := os.Open(file)
	if err != nil {
		return Movie{}, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	csvReader := csv.NewReader(moviesFile)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			return Movie{}, fmt.Errorf("no movies by given name found")
		}
		if err != nil {
			log.Fatal(err)
		}

		if strings.EqualFold(record[0], movieName) {
			mv.MovieName = record[0]
			mv.Director = record[1]
			mv.MainProtagonist = record[2]
			mv.MainAntagonist = record[3]
			mv.ReleaseDate = getDate(record[4])
			mv.Rating, _ = strconv.ParseFloat(record[5], 64)

			return mv, nil
		}
	}

}

func (mv Movie) GetMoviesByRating(file string, rating float64) (*[]Movie, error) {

	if rating == 0 {
		return nil, errors.New("rating provided is not a valid number")
	}

	moviesFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	csvReader := csv.NewReader(moviesFile)

	var moviesByRating []Movie

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			fmt.Fprintln(os.Stderr, "error: invalid rating or too high!")
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if r, _ := strconv.ParseFloat(record[5], 64); r > rating {
			mv.MovieName = record[0]
			mv.Director = record[1]
			mv.MainProtagonist = record[2]
			mv.MainAntagonist = record[3]
			mv.ReleaseDate = getDate(record[4])
			mv.Rating = r

			moviesByRating = append(moviesByRating, mv)
		}

	}
	return &moviesByRating, nil
}
