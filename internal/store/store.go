package store

import "public-podcast-episode-room/internal/models"

type Store struct {
	Records []models.Record
}

func New(records []models.Record) *Store {
	return &Store{Records: records}
}

func (s *Store) List() []models.Record {
	return s.Records
}

func (s *Store) Find(id string) *models.Record {
	for i := range s.Records {
		if s.Records[i].ID == id {
			return &s.Records[i]
		}
	}
	return nil
}

func (s *Store) Create(title string, owner string, note string) models.Record {
	item := models.Record{
		ID:       "r" + string(rune(len(s.Records)+49)),
		Title:    title,
		Owner:    owner,
		Status:   "新建",
		Priority: "中",
		Note:     note,
	}
	s.Records = append([]models.Record{item}, s.Records...)
	return item
}
