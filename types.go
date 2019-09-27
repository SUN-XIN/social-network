package main

import (
	"fmt"
	"time"
)

// ProfileSex defines the sex of a profile.
type ProfileSex int

// ProfileStatus defines the status of a profile.
type ProfileStatus string

const (
	DEFAULT_SESSION_EXPIRED_DURATION = 10 * time.Minute
	SEPARATOR                        = "_"

	SexUnknown ProfileSex = 0
	SexMale    ProfileSex = 1
	SexFemale  ProfileSex = 2

	StatusUnknown ProfileStatus = ""
	StatusNormal  ProfileStatus = "n"
	StatusSpecial ProfileStatus = "s"
)

// Profile define the profile information of an player.
type Profile struct {
	// Required fields
	ID        string        `json:"id"` // primary key
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Status    ProfileStatus `json:"status"`

	// Optional fileds
	Age int
	Sex ProfileSex

	// TODO:
	// it is possible to add other fileds such like
	// general info: 	descriptipion
	// geo info: 		city, country ...
	// media info: 		photo, small format photo, large format photo ...
}

// NewProfile returns a new profile.
func NewProfile(name string) *Profile {
	now := time.Now()
	return &Profile{
		ID:        GenerateID(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// SetStatus sets the given profile's status with the given value.
func (p *Profile) SetStatus(s ProfileStatus) *Profile {
	p.Status = s
	return p
}

// Friends defines the friends relation.
type Friends struct {
	ID        string    `json:"id"` // primary key
	MyID      string    `json:"my_id"`
	FriendID  string    `json:"friend_id"`
	CreatedAt time.Time `json:"created_at"`
}

// NewFriends returns a new friend by the given asker profile ID and
// the given acceptor profile ID.
func NewFriends(myID, friendID string) []*Friends {
	now := time.Now()
	f1 := &Friends{
		MyID:      myID,
		FriendID:  friendID,
		CreatedAt: now,
	}
	f1.SetID()

	f2 := &Friends{
		MyID:      friendID,
		FriendID:  myID,
		CreatedAt: now,
	}
	f2.SetID()

	return []*Friends{f1, f2}
}

// SetID sets the ID of a friends by composite IDs.
func (f *Friends) SetID() {
	f.ID = fmt.Sprintf("%s%s%s", f.MyID, SEPARATOR, f.FriendID)
}

// Session defines the information of a session.
type Session struct {
	ID        string // primary key
	ProfileID string
	CreatedAt time.Time
	ExpiredAt time.Time
}

// Validate checks if the given session is valid.
func (s *Session) Validate() error {
	if s.ExpiredAt.Before(time.Now()) {
		return fmt.Errorf("Session is already expired")
	}

	return nil
}

// NewSession returns a new session.
func NewSession(pID string) *Session {
	return &Session{
		ID:        GenerateID(),
		ProfileID: pID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(DEFAULT_SESSION_EXPIRED_DURATION),
	}
}

// AskFriends defines the information when
// playerA ask playerB to be friends.
type AskFriends struct {
	// Required fields.
	ID        string    `json:"id"`       // primary key
	Asker     string    `json:"asker"`    // profile id of the initiator.
	Receiver  string    `json:"receiver"` // profile id of the influencer.
	CreatedAt time.Time `json:"created_at"`

	// Optional fields.

	Message        string     `json:"message,omitempty"` // playerA can ask playerB to be friends with a message.
	OwnerDeletedAt *time.Time `json:"-"`

	ReceiverReadAt    *time.Time `json:"-"`
	ReceiverTreatedAt *time.Time `json:"-"`

	Accepted bool `json:"-"`
}

// NewAskFriendsByAsker returns a new AskFriends.
func NewAskFriendsByAsker(asker, receiver, msg string) *AskFriends {
	return &AskFriends{
		Asker:     asker,
		Receiver:  receiver,
		CreatedAt: time.Now(),

		Message: msg,
	}
}

// AcceptByReceiver updates the given AskFriends with a positive response.
func (af *AskFriends) AcceptByReceiver() {
	now := time.Now()
	af.ReceiverTreatedAt = &now
	af.Accepted = true
}

// RefuseByReceiver updates the given AskFriends with a negative response.
func (af *AskFriends) RefuseByReceiver() {
	now := time.Now()
	af.ReceiverTreatedAt = &now
	af.Accepted = false
}

// Report defines the report information.
type Report struct {
	// Required fields.
	ID        string
	Reporter  string
	Target    string
	Motive    string
	CreatedAt time.Time

	// Optional fields.
	DeletedAt time.Time
	HandledAt time.Time
}

// NewReport returns a new report.
func NewReport(reporter, target, motive string) *Report {
	return &Report{
		Reporter:  reporter,
		Target:    target,
		Motive:    motive,
		CreatedAt: time.Now(),
	}
}
