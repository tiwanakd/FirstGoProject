package movies

import (
	"encoding/csv"
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

func (m Movie) PrintAll() {
	allMoves, err := m.getAllMovies("marvel_movies.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for i, movie := range *allMoves {
		if i == 0 {
			continue //skipping the first row for column names
		}

		movie.PrintMovieDetails()
	}

}

// Creae a method to print all the reuiqed details for movie type
func (mv Movie) PrintMovieDetails() {

	fmt.Printf("%s:\n", mv.MovieName)
	fmt.Printf("\tDirector: %s\n", mv.Director)
	fmt.Printf("\tMain Protagonist: %s\n", mv.MainProtagonist)
	fmt.Printf("\tMain Antagonist: %s\n", mv.MainAntagonist)
	fmt.Printf("\tRelease Date: %v\n", mv.ReleaseDate.Format("2006-Jan-02"))
	fmt.Printf("\tRating: %.2f\n", mv.Rating)
	fmt.Println("-------------------------------------------------------")
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
func (mv Movie) getAllMovies(filepath string) (*[]Movie, error) {

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
			fmt.Fprintln(os.Stderr, "error: invalid rating")
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
