package campaign

import (
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	contacts = []Contact{
		{Email: "teste@uorak.com"},
	}
	id          = xid.New().String()
	campaign, _ = NewCampaign(id, "Test Campaign", "Test Content", contacts)
)

func TestNewCampaign(t *testing.T) {
	assert.NotNil(t, campaign.ID, "expected campaign id to be created")
	assert.Equal(t, "Test Campaign", campaign.Name, "expected Name to be 'Test Campaign'")
	assert.Equal(t, "Test Content", campaign.Content, "expected Content to be 'Test Content'")
	assert.Len(t, campaign.Contacts, 1, "expected Contacts to contain one contact")
	assert.Equal(t, "teste@uorak.com", campaign.Contacts[0].Email, "expected Contacts to contain one contact with email 'teste@uorak.com'")
	assert.False(t, campaign.CreatedOn.IsZero(), "expected CreatedOn to be set")
}

func TestNewCampaignEmptyName(t *testing.T) {
	_, err := NewCampaign(id, "", "Test Content", contacts)
	assert.NotNil(t, err, "expected error when creating campaign with empty name")
	assert.Equal(t, "Name is required", err.Error(), "expected error message to be 'Name is required'")
}

func TestNewContentEmptyContent(t *testing.T) {
	_, err := NewCampaign(id, "Test Campaign", "", contacts)
	assert.NotNil(t, err, "expected error when creating campaign with empty content")
	assert.Equal(t, "Content is required", err.Error(), "expected error message to be 'Content is required'")
}

func TestNewContactEmptyEmail(t *testing.T) {
	_, err := NewCampaign(id, "Test Campaign", "Test Content", []Contact{{Email: ""}})
	assert.NotNil(t, err, "expected error when creating campaign with empty email")
	assert.Equal(t, "Email is required", err.Error(), "expected error message to be 'Email is required'")
}
