package database

import (
	"time"

	"github.com/kerezsiz42/kinko/internal/utils"
)

var RecordFields = []string{
	"Key",
	"Name",
	"Public",
	"Private",
	"Info",
	"UpdatedAt",
	"CreatedAt",
}

type Record struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Public    string `json:"public"`
	Private   string `json:"private"`
	Info      string `json:"info"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func RecordFromCSV(r []string) Record {
	return Record{
		Key:       r[0],
		Name:      r[1],
		Public:    r[2],
		Private:   r[3],
		Info:      r[4],
		UpdatedAt: r[5],
		CreatedAt: r[6],
	}
}

func NewRecord(
	name string,
	public string,
	private string,
	info string,
) Record {
	currentTime := time.Now().Format(time.RFC3339Nano)

	return Record{
		Key:       utils.GenerateString(utils.Alphanum, 16),
		Name:      name,
		Public:    public,
		Private:   private,
		Info:      info,
		UpdatedAt: currentTime,
		CreatedAt: currentTime,
	}
}

func (r *Record) Modify(name, public, private, info *string) {
	if name != nil {
		r.Name = *name
	}

	if public != nil {
		r.Public = *public
	}

	if private != nil {
		r.Private = *private
	}

	if info != nil {
		r.Info = *info
	}

	r.UpdatedAt = time.Now().Format(time.RFC3339Nano)
}

func (r *Record) ToCSV() []string {
	return []string{
		r.Key,
		r.Name,
		r.Public,
		r.Private,
		r.Info,
		r.UpdatedAt,
		r.CreatedAt,
	}
}
