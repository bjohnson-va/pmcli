package cssdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

//SocialPost represents a cs social post (doesn't matter what service it came from)
type SocialPost struct {
	AccountGroupID      string `json:"agid"`
	IsError             bool   `json:"isError"`
	DeletionStatus      string `json:"deletionStatus"`
	Permalink           string `json:"permalink"`
	PostCreatedDateTime string `json:"postCreatedDateTime"`
	PostText            string `json:"postText"`
	PostedDateTime      string `json:"postedDateTime"`
	ProfileURL          string `json:"profileUrl"`
	ProfileImageURL     string `json:"profileImageUrl"`
	ScheduledDateTime   string `json:"scheduledDateTime"`
	SocialPostID        string `json:"socialPostId"`
	Ssid                string `json:"ssid"`
	Status              string `json:"status"`
	ImageURL            string `json:"imageUrl"`
	Name                string `json:"name"`
	Username            string `json:"username"`
	ParentSocialPostID  string
}

// socialPostFromResponse converts an http response from core services to a list of social services
func socialPostFromResponse(r *http.Response) (*SocialPost, error) {
	defer r.Body.Close()
	type Response struct {
		SocialPost *SocialPost `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to SocialPost: " + err.Error()
		return nil, errors.New(reason)
	}

	return res.SocialPost, nil
}

type listSocialPostResponse struct {
	NextQueryString string                `json:"nextQueryString"`
	Data            []*listSocialPostData `json:"data"`
}
type listSocialPostData struct {
	Services           []*SocialPost `json:"services"`
	ParentSocialPostID string        `json:"socialPostId"`
}

func socialPostsFromResponse(r *http.Response) ([]*SocialPost, string, error) {
	defer r.Body.Close()

	res := listSocialPostResponse{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to SocialPost: " + err.Error()
		return nil, "", errors.New(reason)
	}
	var posts []*SocialPost
	for _, item := range res.Data {
		for _, post := range item.Services {
			post.ParentSocialPostID = item.ParentSocialPostID
		}
		posts = append(posts, item.Services...)
	}

	cursor, err := parseCursorFromNextQueryString(res.NextQueryString)
	if err != nil {
		return nil, "", errors.New("Failed to convert cursor")
	}

	return posts, cursor, nil
}

func parseCursorFromNextQueryString(qs string) (string, error) {
	m, err := url.ParseQuery(qs)
	cursors, ok := m["cursor"]
	if err != nil {
		return "", errors.New("Error parsing cursor")
	}
	if !ok {
		return "", nil
	}
	return cursors[0], nil
}
