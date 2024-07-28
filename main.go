package main

import (
	"projecta/movies"
)

func main() {
	var movie movies.Movie
	// allMovies, _ := movies.GetAllMovies()

	// fmt.Println(len(*allMovies))

	// for _, m := range *allMovies {
	// 	fmt.Println(m.ReleaseDate.Date())
	// }

	//movies.PrintAllMovies()

	// mv, err := movie.SearchMoviebyName("marvel_movies.csv", "iron Man")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// mv.PrintMovieDetails()

	//movie.PrintAll()

	byRating, _ := movie.GetMoviesByRating("marvel_movies.csv", 8)
	for _, mv := range *byRating {
		mv.PrintMovieDetails()
	}
}
