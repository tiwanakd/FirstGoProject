package movies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
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

// Use text/tabwriter package for alignment
// Delcaring a global varible to use the writer provided by tabwriter
var writer = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent|tabwriter.Debug)

// Print all the movies in Tabular format
func (mv Movie) PrintAll(allmovies *[]Movie) {
	fmt.Fprintln(writer, "Name\tDirector\tMainProtagonist\tMainAntagonist\tReleaseDate\tRating")
	fmt.Fprintln(writer, "----------------------\t-----------------\t-----------------\t-------------------\t-----------\t------")
	for _, movie := range *allmovies {
		fmt.Fprintln(writer, movie)
	}
	writer.Flush()
}

// Print Single Movie
func (mv Movie) Print() {
	fmt.Fprintln(writer, "Name\tDirector\tMainProtagonist\tMainAntagonist\tReleaseDate\tRating")
	fmt.Fprintln(writer, "----------------------\t-----------------\t-----------------\t-------------------\t-----------\t------")
	fmt.Fprintln(writer, mv)
	writer.Flush()
}

// Implement the Stringer Interface
func (mv Movie) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%v\t%.1f", mv.MovieName, mv.Director, mv.MainProtagonist, mv.MainAntagonist, mv.ReleaseDate.Format("2006-Jan-02"), mv.Rating)
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

	//make a movies slice which has the length of csvData or number or rows
	movies := make([]Movie, len(csvData))

	for index, movie := range csvData {
		mv.MovieName = movie[0]
		mv.Director = movie[1]
		mv.MainProtagonist = movie[2]
		mv.MainAntagonist = movie[3]
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
func readMoviesByName(movieName string) ([][]string, error) {

	if len(movieName) < 4 {
		return nil, errors.New("movie name must be more that 3 letters")
	}

	//open the csv file to read
	moviesFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read movies csv file %w", err)
	}

	defer moviesFile.Close()

	//move to the begennning of the file
	moviesFile.Seek(0, 0)
	//create a NewReader using csv std lib which takes io.Reader
	//*os.File from Open satisfies io.Reader interface
	csvReader := csv.NewReader(moviesFile)

	//create emtpy slice to hold movie movie
	var moviesSlice [][]string

	//loop over the csv file and read one movie at a time using Read method on reader
	for {
		movie, err := csvReader.Read()
		//if the file is reaches end of file return as movie with given names does not exist
		if err == io.EOF {
			break
		}
		//return if there are additional errors
		if err != nil {
			return nil, err
		}

		// use containts check the all the movie name contacts the give name
		// append to moviesSlice if true
		if strings.Contains(strings.ToLower(movie[0]), strings.ToLower(movieName)) {
			moviesSlice = append(moviesSlice, movie)
		}

	}

	//if the movie slice is empty return error as the movie with given name was not found
	if len(moviesSlice) < 1 {
		return nil, fmt.Errorf("movie with name %q not found", movieName)
	}

	return moviesSlice, nil
}

// Function to get the Movie by given name
func (mv Movie) SearchMoviesbyName(movieName string) (*[]Movie, error) {
	movies, err := readMoviesByName(movieName)
	if err != nil {
		return nil, err
	}

	//make a movies slice which has the length of csvData or number or rows
	moviesByName := make([]Movie, len(movies))

	for index, movie := range movies {
		mv.MovieName = movie[0]
		mv.Director = movie[1]
		mv.MainProtagonist = movie[2]
		mv.MainAntagonist = movie[3]
		mv.ReleaseDate, _ = GetDate(movie[4])
		mv.Rating, _ = strconv.ParseFloat(movie[5], 64) //convert to float.

		moviesByName[index] = mv
	}

	return &moviesByName, nil
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
exact duplicate names are not allowed.
create a function to check if movie has a duplicate name
*/
func dupicateMovieName(movieName string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			return false
		}
		if err != nil {
			return false
		}

		if strings.EqualFold(record[0], movieName) {
			return true
		}
	}
}

/*
This func add a new Movie field to the csv file.
Created this as standalone function without receiver argument
*/
func (mv Movie) AddMovie() error {

	if dupicateMovieName(mv.MovieName) {
		return fmt.Errorf("exact duplicate Movie names not allowed, %q already exists", mv.MovieName)
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
		return fmt.Errorf("error writing csv: %w", err)
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

	//get all movies
	allMovies, err := readAllMovies()
	if err != nil {
		return err
	}

	//set a traking var to check if the provided movie name exists
	found := false

	//loop over allMoives slice
	//on the same index as movieIndex update the passed in filed
	for _, movieField := range allMovies {
		if strings.EqualFold(movieField[0], movieName) {
			found = true
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
		} else {
			found = false
		}
	}

	if !found {
		return fmt.Errorf("movie with name %q not found", movieName)
	}

	//use writeAll function to write back the update slice to csv file
	return writeAll(allMovies)
}
