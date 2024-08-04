package movies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
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

// CSV file that is to be used
const filePath = "marvel_movies.csv"

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
func (mv Movie) GetAllMovies() (*[]Movie, error) {

	csvData, err := readAllMovies()
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

/*
locat function to read the csv and return all the fields
*/
func readAllMovies() ([][]string, error) {

	moviesFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}
	defer moviesFile.Close()
	csvReader := csv.NewReader(moviesFile)
	return csvReader.ReadAll()
}

/*
local function to read a single given movie and
return a single field
*/
func readMovie(movieName string) ([]string, int, error) {

	//open the csv file to read
	moviesFile, err := os.Open(filePath)
	if err != nil {
		return nil, -1, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	//move to the begennning of the file
	moviesFile.Seek(0, 0)
	//create a NewReader using csv std lib which takes io.Reader
	//*os.File from Open satisfies io.Reader interface
	csvReader := csv.NewReader(moviesFile)

	//create emtpy slice to hold movie movie
	var moviemovie []string

	//set up an index to track where the movie is located in file
	var movieIndex int
	//set up count which will incremented everytime a record is read
	//count will be set to the movieIndex if the movie is found
	count := 0

	//tracking if the movie for the given name is found
	found := false

	//loop over the csv file and read one movie at a time using Read method on reader
	for {
		movie, err := csvReader.Read()
		//if the file is reaches end of file return as movie with given names does not exist
		if err == io.EOF {
			found = false
			break
		}
		//return if there are additional errors
		if err != nil {
			return nil, -1, err
		}

		//if the first reocrd that is the movie name is equal append to moviemovie
		if strings.EqualFold(movie[0], movieName) {
			moviemovie = append(moviemovie, movie...)
			found = true // Movie was found
			movieIndex = count
			break
		}

		count++ //increment the count
	}

	if !found {
		return nil, -1, fmt.Errorf("movie with name %s not found", movieName)
	}

	return moviemovie, movieIndex, nil

}

// Function to get the Movie by given name
func (mv Movie) SearchMoviebyName(movieName string) (Movie, error) {
	movieRecord, _, err := readMovie(movieName)
	if err != nil {
		return Movie{}, err
	}

	mv.MovieName = movieRecord[0]
	mv.Director = movieRecord[1]
	mv.MainProtagonist = movieRecord[2]
	mv.MainAntagonist = movieRecord[3]
	mv.ReleaseDate, _ = GetDate(movieRecord[4])
	mv.Rating, _ = strconv.ParseFloat(movieRecord[5], 64)

	return mv, nil
}

/*
Function to get the Movies as per the given rating
Return slice of movies which have more rating than the provided rating.
*/

func (mv Movie) GetMoviesByRating(rating float64) (*[]Movie, error) {

	//If the rating provided is 0 or more that 10 return
	if rating == 0 || rating > 10 {
		return nil, errors.New("rating provided is not valid")
	}

	moviesFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	moviesFile.Seek(0, 0)
	csvReader := csv.NewReader(moviesFile)

	//since we need more that one movies to be returned a slice of Movie type is used
	var moviesByRating []Movie

	for {
		movie, err := csvReader.Read()
		if err == io.EOF {
			// break once the end of file reached
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
		//since Read gives []string, convert it to float64 to compare with given float64 rating
		if r, _ := strconv.ParseFloat(movie[5], 64); r > rating {
			mv.MovieName = movie[0]
			mv.Director = movie[1]
			mv.MainProtagonist = movie[2]
			mv.MainAntagonist = movie[3]
			mv.ReleaseDate, _ = GetDate(movie[4])
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

func (mv Movie) AddMovie() error {

	//duplicate movie names are not allowed
	//user SearchMovie func to check the movie with given name exists
	serchMovie, _ := mv.SearchMoviebyName(mv.MovieName)
	if strings.EqualFold(serchMovie.MovieName, mv.MovieName) {
		return fmt.Errorf("movie with name %s already exists", mv.MovieName)
	}

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
		mv.MovieName,
		mv.Director,
		mv.MainProtagonist,
		mv.MainAntagonist,
		mv.ReleaseDate.Format("2006-01-02"),
		fmt.Sprintf("%.1f", mv.Rating),
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

/*
Function that writes multiple rows to CSV File
*/
func writeAll(allmovies [][]string) error {
	moviesFile, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("unable to open the file, err %w", err)
	}

	// Truncate the file and move to the beginning
	moviesFile.Truncate(0)
	moviesFile.Seek(0, 0)
	csvWriter := csv.NewWriter(moviesFile)

	//wrtie the enite slice to file using csvWrite
	csvWriter.WriteAll(allmovies)

	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("erro writing csv: %w", err)
	}

	return nil
}

/*
Method to delete the movie with provided name
*/
func (mv Movie) DeleteMovie(movieName string) error {
	if movieName == "" {
		return errors.New("no movie name provided to delete")
	}

	// read all the movie data into memory
	allMovies, err := readAllMovies()
	if err != nil {
		return err
	}

	var deleteIndex int //will use to get the index of row to delete

	found := false
	for index, movie := range allMovies {
		if movie[0] == movieName {
			// once the movie name is found set the current index to the delete index
			deleteIndex = index
			found = true
			break
		}
	}

	//return if the movie name is not found
	if !found {
		return fmt.Errorf("no movie with name: %s found", movieName)
	}

	//using deleteindex delete the the row of that index using slices.Delete
	//assisng to new a slice, that still points to same underlying array
	newMoviesSlice := slices.Delete(allMovies, deleteIndex, deleteIndex+1)

	//use the wrtieAll fucntion to write the new slice to cav file
	return writeAll(newMoviesSlice)
}

func (mv Movie) UpdateMovie(movieName, field string) error {

	//get the required index from readMovie function
	_, movieIndex, err := readMovie(movieName)
	if err != nil {
		return err
	}

	//get all movies
	allMovies, err := readAllMovies()
	if err != nil {
		return err
	}

	//loop over allMoives slice
	//on the same index as movieIndex update the passed in filed
	for index, movieField := range allMovies {
		if index == movieIndex {
			//use swtich to check which field is to be updated
			switch strings.ToLower(field) {
			case "name":
				movieField[0] = mv.MovieName
			case "director", "dir":
				movieField[1] = mv.Director
			case "protagonist", "pro":
				movieField[2] = mv.MainProtagonist
			case "antagonist", "ant":
				movieField[3] = mv.MainAntagonist
			case "releasedate", "rdate":
				movieField[4] = mv.ReleaseDate.Format("2006-01-02")
			case "rating":
				movieField[5] = fmt.Sprintf("%.1f", mv.Rating)
			default:
				return fmt.Errorf("invalid Input name %s", field)
			}
		}
	}

	//use writeAll function to write back the update slice to csv file
	return writeAll(allMovies)
}
