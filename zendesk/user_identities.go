package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
)

// UserIdentity is struct for user_field payload
type UserIdentity struct {
	ID                 int64  `json:"id,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	DeliverableState   string `json:"deliverable_state,omitempty"`
	Primary            bool   `json:"primary,omitempty"`
	IDType             string `json:"type,omitempty"`
	UndeliverableCount int64  `json:"undeliverable_count,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	URL                string `json:"url,omitempty"`
	UserID             int64  `json:"user_id,omitempty"`
	Value              string `json:"value,omitempty"`
	Verified           bool   `json:"verified,omitempty"`
}

type UserIdentityListOptions struct {
	PageOptions
}

type UserIdentitiesAPI interface {
	GetEndUserIdentites(ctx context.Context, opts *UserIdentityListOptions) ([]UserIdentity, Page, error)
}

// GetUserIdentites fetch trigger list
//
// https://developer.zendesk.com/api-reference/ticketing/users/user_identities/#list-identities
func (z *Client) GetEndUserIdentites(ctx context.Context, userID int64) ([]UserIdentity, error) {

	var result struct {
		UserIdentities []UserIdentity `json:"identities"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/end_users/%d/identities.json", userID))
	if err != nil {
		return []UserIdentity{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return []UserIdentity{}, err
	}
	return result.UserIdentities, nil
}

func (z *Client) CreateEndUserIdentity(ctx context.Context, userID int64, userIdentity UserIdentity) (UserIdentity, error) {
	var data, result struct {
		UserIdentity UserIdentity `json:"identity"`
	}
	data.UserIdentity = userIdentity

	body, err := z.post(ctx, fmt.Sprintf("/end_users/%d/identities.json", userID), data)
	if err != nil {
		return UserIdentity{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return UserIdentity{}, err
	}
	return result.UserIdentity, nil
}

func (z *Client) MakeEndUserIdentityPrimary(ctx context.Context, userID int64, userIdentityID int64) ([]UserIdentity, error) {
	var result struct {
		UserIdentities []UserIdentity `json:"identities"`
	}
	_, err := z.put(ctx, fmt.Sprintf("/end_users/%d/identities/%d/make_primary", userID, userIdentityID), result)

	if err != nil {
		return []UserIdentity{}, err
	}

	return result.UserIdentities, nil
}

func (z *Client) VerifyEndUserIdentity(ctx context.Context, userID int64, userIdentityID int64) (UserIdentity, error) {
	var result struct {
		UserIdentity UserIdentity `json:"identity"`
	}
	_, err := z.put(ctx, fmt.Sprintf("/users/%d/identities/%d/verify", userID, userIdentityID), result)

	if err != nil {
		return UserIdentity{}, err
	}

	return result.UserIdentity, nil
}

func (z *Client) DeleteEndUserIdentity(ctx context.Context, userID int64, userIdentityID int64) error {
	err := z.delete(ctx, fmt.Sprintf("/end_users/%d/identities/%d", userID, userIdentityID))

	if err != nil {
		return err
	}

	return nil
}
