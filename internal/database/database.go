package database

import (
	"time"

	"github.com/kerezsiz42/kinko/internal/utils"
)

const RECORD_LEN = 7

type Record struct {
	Key       string
	Name      string
	Public    string
	Private   string
	Info      string
	UpdatedAt string
	CreatedAt string
}

func NewRecord(
	name string,
	public string,
	private string,
	info string,
) *Record {
	currentTime := time.Now().Format(time.RFC3339)

	return &Record{
		Key:       utils.GenerateID(16),
		Name:      name,
		Public:    public,
		Private:   private,
		Info:      info,
		UpdatedAt: currentTime,
		CreatedAt: currentTime,
	}
}

func (r *Record) UpdateName(name string) {
	r.Name = name
	r.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (r *Record) UpdatePublic(public string) {
	r.Public = public
	r.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (r *Record) UpdatePrivate(private string) {
	r.Private = private
	r.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (r *Record) UpdateInfo(info string) {
	r.Info = info
	r.UpdatedAt = time.Now().Format(time.RFC3339)
}

type Database interface {
	Load() ([]Record, error)
	Persist([]Record) error
}
