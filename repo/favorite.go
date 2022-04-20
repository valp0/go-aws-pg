package repo

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Helper function that gets all favorites of a user, given its id.
func (r repository) getFavorites(id string) ([]Favorite, error) {
	favs := []Favorite{}
	sqlStatement := `SELECT f.* FROM user_favs uf JOIN favorites f ON uf.fav_id = f.id WHERE uf.user_id = $1;`

	rows, _ := r.db.Query(sqlStatement, id)
	defer rows.Close()

	for rows.Next() {
		var fav Favorite
		err := rows.Scan(&fav.ID, &fav.Description, &fav.Thumbnail, &fav.Title)
		if err != nil {
			return nil, fmt.Errorf("an error occured while getting favorites of user with id %s, %v", id, err)
		}

		favs = append(favs, fav)
	}

	return favs, nil
}

// An auxiliary repo function to insert a video safely into the favorites table.
func (r repository) insertToFavs(fav Favorite) error {
	insertStatement := `INSERT INTO Favorites (id, description, thumbnail, title) VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(insertStatement, fav.ID, fav.Description, fav.Thumbnail, fav.Title)
	if err != nil {
		return fmt.Errorf("couldn't insert favorite with id %s, %v", fav.ID, err)
	}

	log.Printf("Favorite with id %s was inserted to favorites table successfully.\n", fav.ID)
	return nil
}

// An auxiliary repo function to insert a user-favorite relation safely into the user_favs table.
func (r repository) insertToUserFavs(userId, favId string) error {
	insertStatement := `INSERT INTO user_favs (user_id, fav_id) VALUES ($1, $2);`

	_, err := r.db.Exec(insertStatement, userId, favId)
	if err != nil {
		return fmt.Errorf("couldn't insert favorite with id %s to user with id %s, %v", favId, userId, err)
	}

	log.Printf("Favorite with id %s was inserted to user_favs table successfully.\n", favId)
	return nil
}

// An auxiliary repo function to get a specific video from the favorites table, given its id.
func (r repository) getFav(id string) (Favorite, error) {
	var fav Favorite
	selectStatement := `SELECT * FROM Favorites WHERE id = $1;`

	err := r.db.QueryRow(selectStatement, id).Scan(&fav.ID, &fav.Description, &fav.Thumbnail, &fav.Title)
	if err != nil {
		return Favorite{}, fmt.Errorf("video with id %s could not be fetched, %s", id, err.Error())
	}

	return fav, nil
}

// Checks a favorite is not already present in the table to avoid duplicity.
func (r repository) alreadyInFavs(id string) bool {
	sqlStatement := `SELECT * FROM Favorites WHERE id = $1;`
	var fav Favorite

	err := r.db.QueryRow(sqlStatement, id).Scan(&fav.ID, &fav.Description, &fav.Thumbnail, &fav.Title)
	return err == nil
}

// Checks a user favorite is not already present in the table to avoid duplicity.
func (r repository) alreadyInUserFavs(userId, favId string) bool {
	sqlStatement := `SELECT * FROM user_favs WHERE user_id = $1 AND fav_id = $2;`
	var userFav UserFav

	err := r.db.QueryRow(sqlStatement, userId, favId).Scan(&userFav.UserId, &userFav.FavId)
	return err == nil
}

// Will delete every row in the user_favs table containing id as user_id.
func (r repository) deleteFavsFromUser(id string) error {
	sqlStatement := `DELETE FROM user_favs WHERE user_id = $1;`
	if _, err := r.db.Exec(sqlStatement, id); err != nil {
		return err
	}

	return nil
}

// Deletes a user_favs row containing the provided user id and favorite id.
func (r repository) deleteUserFav(userId, favId string) error {
	deleteStatement := `DELETE FROM user_favs WHERE user_id = $1 AND fav_id = $2;`

	result, err := r.db.Exec(deleteStatement, userId, favId)
	if err != nil {
		return fmt.Errorf("video with id %s  for user with id %s could not be removed from user_favs table, %s", favId, userId, err.Error())
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows were affected, probably user with id %s is not related to video with id %s", userId, favId)
	}

	r.cleanFavs()
	return nil
}

// Deletes any row in the favorites table that isn't related to a user in the user_favs table.
func (r repository) cleanFavs() {
	favs, err := r.getUnmatchedFavs()
	if err != nil {
		log.Println(err)
		return
	}

	for _, fav := range favs {
		if err := r.deleteFromFavorites(fav.ID); err != nil {
			log.Println(err)
			return
		}
	}

}

// Will return any video that is present in the favorites table but not in the user_favs table.
func (r repository) getUnmatchedFavs() ([]Favorite, error) {
	sqlStatement := `SELECT f.* FROM favorites f LEFT JOIN user_favs uf ON f.id = uf.fav_id WHERE uf.user_id is NULL;`
	favs := []Favorite{}

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var fav Favorite
		if err := rows.Scan(&fav.ID, &fav.Description, &fav.Thumbnail, &fav.Title); err != nil {
			return nil, err
		}

		favs = append(favs, fav)
	}

	return favs, nil
}

// Deletes a row from Favorites table.
func (r repository) deleteFromFavorites(id string) error {
	sqlStatement := `DELETE FROM Favorites WHERE id = $1;`

	_, err := r.db.Exec(sqlStatement, id)
	return err
}
