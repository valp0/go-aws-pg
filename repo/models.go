package repo

// Schema of the user data table
type User struct {
	ID           string `json:"user_id"`
	Name         string `json:"user_name"`
	ThumbnailUrl string `json:"user_thumbnail,omitempty"`
}

// Schema of the favorites table
type Favorite struct {
	ID          string `json:"id"`
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Title       string `json:"title"`
}

// Schema of the user_favs table
type UserFav struct {
	UserId string `json:"user_id"`
	FavId  string `json:"fav_id"`
}
