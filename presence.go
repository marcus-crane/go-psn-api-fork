package psn

import (
	"context"
	"fmt"
)

const presenceApi = "https://m.np.playstation.net/api/userProfile/v1/internal/users/me/basicPresences"

type Presence struct {
	Availability        string              `json:"availability"`
	PrimaryPlatformInfo PrimaryPlatformInfo `json:"primaryPlatformInfo"`
	GameTitleInfoList   []GameTitleInfo     `json:"gameTitleInfoList"`
}

type PrimaryPlatformInfo struct {
	OnlineStatus   string `json:"online"`
	Platform       string `json:"platform"`
	LastOnlineDate string `json:"lastOnlineDate"`
}

type GameTitleInfo struct {
	TitleID        string `json:"npTitleId"`
	Name           string `json:"titleName"`
	Platform       string `json:"format"`
	LaunchPlatform string `json:"launchPlatform"`
	TitleImage     string `json:"npTitleIconUrl"`
}

type PresenceResponse struct {
	BasicPresence Presence `json:"basicPresence"`
}

// GetPresenceRequest retrieves the current users presence info
func (p *psn) GetPresenceRequest(ctx context.Context) (presence *Presence, err error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)

	presenceResponse := &PresenceResponse{}
	err = p.get(
		ctx,
		presenceApi,
		h,
		&presenceResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &presenceResponse.BasicPresence, nil
}
