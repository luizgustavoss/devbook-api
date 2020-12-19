package persistence

import (
	"database/sql"
	"devbook/src/models"
)

// persists publication data
type PublicationRepository struct {
	db *sql.DB
}

// CreatePublication creates a new publication
func (repository PublicationRepository) CreatePublication(publication models.Publication) (uint64, error) {

	stmt, error := repository.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer stmt.Close()

	insert, error := stmt.Exec(publication.Title, publication.Content, publication.AuthorId)
	if error != nil {
		return 0, error
	}

	id, error := insert.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(id), nil;
}

// GetPublicationById gets a specific publication by its id
func (repository PublicationRepository) GetPublicationById(id uint64) (models.Publication, error) {

	var publication models.Publication

	resultSet, error := repository.db.Query(
		`select p.id, p.title, p.content, p.author_id, u.nick, p.likes, p.created_at 
				from publications p 
				join users u on (p.author_id = u.id) 
				where p.id = ?`, id)

	if error != nil {
		return publication, error
	}
	defer resultSet.Close()

	if resultSet.Next() {
		if error = resultSet.Scan(
			&publication.ID, &publication.Title,
			&publication.Content, &publication.AuthorId,
			&publication.AuthorNick, &publication.Likes,
			&publication.CreatedAt,); error != nil {
			return publication, error
		}
	}
	return publication, nil
}

// GetPublicationsForUserId gets publications of a user and the users he/she follows
func (repository PublicationRepository) GetPublicationsForUserId(userId uint64) ([]models.Publication, error) {

	resultSet, error := repository.db.Query(
		`select p.id, p.title, p.content, p.author_id, u.nick, p.likes, p.created_at 
				from publications p 
				join users u on (p.author_id = u.id) 
				where u.id = ? or u.id in (select f.user_id from followers f where f.follower_id = ?) order by 1 desc`,
		userId, userId)
	if error != nil {
		return nil, error
	}
	defer resultSet.Close()

	var publications []models.Publication

	for resultSet.Next() {
		var publication models.Publication
		if error = resultSet.Scan(
			&publication.ID, &publication.Title,
			&publication.Content, &publication.AuthorId,
			&publication.AuthorNick, &publication.Likes,
			&publication.CreatedAt,); error != nil {
			return nil, error
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

// UpdatePublication updates a publication
func (repository PublicationRepository) UpdatePublication(id uint64, publication models.Publication) error {

	stmt, error := repository.db.Prepare(
		"update publications set title = ?, content = ? where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(publication.Title, publication.Content, id); error != nil  {
		return error
	}
	return nil
}

// DeletePublication deletes a publication
func (repository PublicationRepository) DeletePublication(id uint64) error {

	stmt, error := repository.db.Prepare(
		"delete from publications where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(id); error != nil {
		return error
	}
	return nil
}

// GetUserPublicationById get all publications of a specific user
func (repository PublicationRepository) GetUserPublicationById(userId uint64) ([]models.Publication, error) {

	resultSet, error := repository.db.Query(
		`select p.id, p.title, p.content, p.author_id, u.nick, p.likes, p.created_at 
				from publications p 
				join users u on (p.author_id = u.id) 
				where u.id = ? order by 1 desc`,
			userId)
	if error != nil {
		return nil, error
	}
	defer resultSet.Close()

	var publications []models.Publication

	for resultSet.Next() {
		var publication models.Publication
		if error = resultSet.Scan(
			&publication.ID, &publication.Title,
			&publication.Content, &publication.AuthorId,
			&publication.AuthorNick, &publication.Likes,
			&publication.CreatedAt,); error != nil {
			return nil, error
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

// RegisterPublicationLike registers a like in publication
func (repository PublicationRepository) RegisterPublicationLike(id uint64) error {
	stmt, error := repository.db.Prepare(
		"update publications p set p.likes = p.likes + 1 where id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(id); error != nil  {
		return error
	}
	return nil
}

// RegisterPublicationUnlike registers an unlike in publication
func (repository PublicationRepository) RegisterPublicationUnlike(id uint64) error {
	stmt, error := repository.db.Prepare(
		"update publications p set p.likes = p.likes - 1 where p.likes > 0 and id = ? ")
	if error != nil {
		return error
	}
	defer stmt.Close()

	if _, error = stmt.Exec(id); error != nil  {
		return error
	}
	return nil
}




// factory
func NewPublicationRepository(db *sql.DB) *PublicationRepository {
	return &PublicationRepository{db}
}
