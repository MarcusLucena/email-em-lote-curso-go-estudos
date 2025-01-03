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
