package campaign

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(id, name, content string, contacts []Contact) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("Name is required")
	}

	if content == "" {
		return nil, errors.New("Content is required")
	}

	if contacts[0].Email == "" {
		return nil, errors.New("Email is required")
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
