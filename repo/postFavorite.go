package repo

import (
	"log"

	_ "github.com/lib/pq"
)

// Will post a favorite to favorites table and add the relation between fav_id and user_id in the user_favs table
func PostFavorite(id string, fav Favorite) ([]Favorite, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Checking video was not present in the favorites table already.
	if alreadyInFavs(db, fav.ID) {
		log.Println("Video with id", fav.ID, "was already present in the favorites table.")
	} else {
		if err := insertToFavs(db, fav); err != nil {
			return nil, err
		}
	}

	// Checking a video was not already linked as favorite to the same user.
	if alreadyInUserFavs(db, id, fav.ID) {
		log.Println("Video with id", fav.ID, "was already present in the user_favs table.")
	} else {
		if err := insertToUserFavs(db, id, fav.ID); err != nil {
			return nil, err
		}
	}

	favId := fav.ID
	fav = Favorite{}
	fav, err = getFav(db, favId)
	if err != nil {
		return nil, err
	}

	return []Favorite{fav}, nil
}
