package movies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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
	rating := fmt.Sprintf("\tRating: %.1f\n", mv.Rating)
	lineBreak := fmt.Sprintln("-------------------------------------------------------")

	return movieName + director + mainProtagonist + mainAntagonist + releaseDate + rating + lineBreak
}

// Since date is being provided as string
// Fucntion to return a Time struct for given date
func GetDate(dateString string) (time.Time, error) {
	const shortForm = "2006-01-02"
	t, err := time.Parse(shortForm, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
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
	csvData = csvData[1:] //skipping the first row that has column names

	//make a movies slice which has the length of csvData or number or rows
	movies := make([]Movie, len(csvData))

	for index, movie := range csvData {
		mv.MovieName = movie[0]
		mv.Director = movie[1]
		mv.MainProtagonist = movie[2]
		mv.MainProtagonist = movie[3]
		mv.ReleaseDate, _ = GetDate(movie[4])
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

	//open the csv file to read
	moviesFile, err := os.Open(file)
	if err != nil {
		return Movie{}, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	//create a NewReader using csv std lib which takes io.Reader
	//*os.File from Open satisfies io.Reader interface
	csvReader := csv.NewReader(moviesFile)

	//loop over the csv file and read one record at a time using Read method on reader
	for {
		record, err := csvReader.Read()
		//if the file is reaches end of file return as movie with given names does not exist
		if err == io.EOF {
			return Movie{}, fmt.Errorf("no movies by given name found")
		}
		//return if there are additional errors
		if err != nil {
			return Movie{}, err
		}
		//check if the provided name matches the first record which is the moive name
		//assign values of each record to its correspoding fields from Movie type
		//after assigning the values return as we only need a signle Movie field
		if strings.EqualFold(record[0], movieName) {
			mv.MovieName = record[0]
			mv.Director = record[1]
			mv.MainProtagonist = record[2]
			mv.MainAntagonist = record[3]
			mv.ReleaseDate, _ = GetDate(record[4])
			mv.Rating, _ = strconv.ParseFloat(record[5], 64)

			return mv, nil
		}
	}

}

/*
Function to get the Movies as per the given rating
Return slice of movies which have more rating than the provided rating.
*/

func (mv Movie) GetMoviesByRating(file string, rating float64) (*[]Movie, error) {

	//If the rating provided is 0 or more that 10 return
	if rating == 0 || rating > 10 {
		return nil, errors.New("rating provided is not valid")
	}

	moviesFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	csvReader := csv.NewReader(moviesFile)

	//since we need more that one records to be returned a slice of Movie type is used
	var moviesByRating []Movie

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			// entire file needs to be looped over unlike SearchMoviebyName func as
			// that func returned as soon as match was recevied
			// cannot return as we need to keep final return out of the loop
			// returning here would leave that code unreachable - using break instead
			break

		}
		if err != nil {
			return nil, err
		}

		//since Read gives []string, convert it to float64 to compare with given float64 rating
		if r, _ := strconv.ParseFloat(record[5], 64); r > rating {
			mv.MovieName = record[0]
			mv.Director = record[1]
			mv.MainProtagonist = record[2]
			mv.MainAntagonist = record[3]
			mv.ReleaseDate, _ = GetDate(record[4])
			mv.Rating = r

			//append each match to the slice
			moviesByRating = append(moviesByRating, mv)
		}

	}
	return &moviesByRating, nil
}

/*
This func add a new Movie field to the csv file.
Created this as standalone function without receiver argument
*/

func AddNewMovie(filePath string, newMovie Movie) error {

	//use OpenFile method as Open only allows to open file as Read only.
	// both flags os.O_APPEND|os.O_WRONLY are required to the end of file
	moviesFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer moviesFile.Close()

	csvWriter := csv.NewWriter(moviesFile)

	//create a field of slice of string
	//this will be passed to Write method o csvWriter
	moviefield := []string{
		newMovie.MovieName,
		newMovie.Director,
		newMovie.MainProtagonist,
		newMovie.MainAntagonist,
		newMovie.ReleaseDate.Format("2006-01-02"),
		fmt.Sprintf("%.1f", newMovie.Rating),
	}

	// write the field to the buffer
	if err := csvWriter.Write(moviefield); err != nil {
		return err
	}

	// Write any buffered data to the underlying writer (standard output).
	csvWriter.Flush()

	// use Error() method on csv writer to check if the write to file was successful
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
