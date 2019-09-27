package main

import (
	"fmt"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
//////////////////////////    Common part    //////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

const (
	// DEFAULT_PAGE_LIMIT is the default value for the pagination's limit.
	DEFAULT_PAGE_LIMIT = 10
	// MAX_PAGE_LIMIT is the defmaxault value for the pagination's limit.
	MAX_PAGE_LIMIT = 100
)

// Pagination defines the pagination information.
type Pagination struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// Response defines the common information for the response payload.
type Response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

// NewResponseErr returns a new response with ok value false.
func NewResponseErr(err error) Response {
	return Response{
		OK:    false,
		Error: err.Error(),
	}
}

///////////////////////////////////////////////////////////////////////////////
//////////////////////////    ListFriends    //////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// ListFriendsRequest defines the request payload
// for list friends endpoint.
type ListFriendsRequest struct {
	TargetID string `json:"target_id"`

	Pagination
}

// Validate checks if the given ListFriendsRequest is valid.
func (lfr *ListFriendsRequest) Validate() error {
	if len(lfr.TargetID) == 0 {
		return fmt.Errorf("target_id is required")
	}

	if lfr.Limit <= 0 {
		lfr.Limit = DEFAULT_PAGE_LIMIT
	}

	if lfr.Limit > MAX_PAGE_LIMIT {
		lfr.Limit = MAX_PAGE_LIMIT
	}

	return nil
}

// ListFriendsResponse defines the response payload
// for list friends endpoint.
type ListFriendsResponse struct {
	Friends []*Friends `json:"friends"`

	Pagination

	Response
}

// NewListFriendsResponseOK returns a new ListFriendsResponse with ok value true.
func NewListFriendsResponseOK(fs []*Friends, p Pagination) *ListFriendsResponse {
	return &ListFriendsResponse{
		Friends:    fs,
		Pagination: p,

		Response: Response{
			OK: true,
		},
	}
}

// NewListFriendsResponseErr returns a new ListFriendsResponse with ok value false.
func NewListFriendsResponseErr(err error) *ListFriendsResponse {
	return &ListFriendsResponse{
		Response: NewResponseErr(err),
	}
}

///////////////////////////////////////////////////////////////////////////////
//////////////////////////    GetProfiles    //////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GetProfilesResponse defines the response payload
// for get profiles endpoint.
type GetProfilesResponse struct {
	Profiles []*Profile `json:"profiles"`

	Response
}

// NewGetProfilesResponseOK returns a new GetProfilesResponse with ok value true.
func NewGetProfilesResponseOK(ps []*Profile) *GetProfilesResponse {
	return &GetProfilesResponse{
		Profiles: ps,

		Response: Response{
			OK: true,
		},
	}
}

// NewGetProfilesResponseErr returns a new GetProfilesResponse with ok value false.
func NewGetProfilesResponseErr(err error) *GetProfilesResponse {
	return &GetProfilesResponse{
		Response: NewResponseErr(err),
	}
}

///////////////////////////////////////////////////////////////////////////////
//////////////////////////    AskFriends     //////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// AskFriendsRequest defines the request payload
// when playerA ask playerB to be friends
// for the ask friends endpoint.
type AskFriendsRequest struct {
	TargetID string `json:"target_id"`
	Message  string `json:"message,omitempty"`
}

// Validate checks if the given AskFriendsRequest is valid.
func (afr *AskFriendsRequest) Validate() error {
	if len(afr.TargetID) == 0 {
		return fmt.Errorf("target_id is required")
	}

	return nil
}

// AskFriendsResponse defines the response payload
// for the ask friends endpoint.
type AskFriendsResponse struct {
	Created time.Time `json:"created"`
	Message string    `json:"message,omitempty"`

	Response
}

// NewAskFriendsResponseOK returns a new AskFriendsResponse with ok value true.
func NewAskFriendsResponseOK(a *AskFriends) *AskFriendsResponse {
	return &AskFriendsResponse{
		Created: a.CreatedAt,
		Message: a.Message,

		Response: Response{
			OK: true,
		},
	}
}

// NewAskFriendsResponseErr returns a new AskFriendsResponse with ok value false.
func NewAskFriendsResponseErr(err error) *AskFriendsResponse {
	return &AskFriendsResponse{
		Response: NewResponseErr(err),
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////// ListAskFriendsResponse  /////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// ListAskFriendsResponse defines the response payloads
// for the list ask friends endpoint.
type ListAskFriendsResponse struct {
	Requests []*AskFriends `json:"requests"`

	Response
}

// NewListAskFriendsResponseOK returns a new ListAskFriendsResponse with ok value true.
func NewListAskFriendsResponseOK(as []*AskFriends) *ListAskFriendsResponse {
	return &ListAskFriendsResponse{
		Requests: as,

		Response: Response{
			OK: true,
		},
	}
}

// NewListAskFriendsResponseErr returns a new ListAskFriendsResponse with ok value false.
func NewListAskFriendsResponseErr(err error) *ListAskFriendsResponse {
	return &ListAskFriendsResponse{
		Response: NewResponseErr(err),
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////// TreatAskFriendsRequest  /////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// TreatAskFriendsAction defines the action for an ask friends request.
type TreatAskFriendsAction string

const (
	ActionUnkown TreatAskFriendsAction = ""
	ActionAccept TreatAskFriendsAction = "a"
	ActionRefuse TreatAskFriendsAction = "r"
	ActionIgnore TreatAskFriendsAction = "i"
)

// TreatAskFriendsRequest defines the request payload
// when playerB accepts/refuses playerA's ask friends action
// for the treat friends endpoint.
type TreatAskFriendsRequest struct {
	ActionID string                `json:"action_id"`
	Action   TreatAskFriendsAction `json:"action"`
}

// Validate checks if the given AcceptFriendsRequest is valid.
func (afr *TreatAskFriendsRequest) Validate() error {
	if len(afr.ActionID) == 0 {
		return fmt.Errorf("action id is required")
	}

	if afr.Action == ActionUnkown {
		afr.Action = ActionIgnore // TODO: we can also return error in this case.
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////      ReportRequest      /////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// CreateReportRequest defines the request payload
// for the create report endpoint.
type CreateReportRequest struct {
	TargetID string `json:"target_id"`
	Motive   string `json:"motive"`
}

// Validate checks if the given ReportRequest is valid.
func (crr *CreateReportRequest) Validate() error {
	if len(crr.TargetID) == 0 {
		return fmt.Errorf("target id is required")
	}

	if len(crr.Motive) == 0 {
		return fmt.Errorf("motive is required")
	}

	return nil
}
