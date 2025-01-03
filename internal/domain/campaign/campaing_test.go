package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	contacts := []Contact{
		{Email: "teste@uorak.com"},
	}
	campaign := NewCampaign("1", "Test Campaign", "Test Content", contacts)

	assert.Equal(t, "1", campaign.ID, "expected ID to be '1'")
	assert.Equal(t, "Test Campaign", campaign.Name, "expected Name to be 'Test Campaign'")
	assert.Equal(t, "Test Content", campaign.Content, "expected Content to be 'Test Content'")
	assert.Len(t, campaign.Contacts, 1, "expected Contacts to contain one contact")
	assert.Equal(t, "teste@uorak.com", campaign.Contacts[0].Email, "expected Contacts to contain one contact with email 'teste@uorak.com'")
	assert.False(t, campaign.CreatedOn.IsZero(), "expected CreatedOn to be set")
}
