package main

import (
	"fmt"
	"time"
)

// Store defines a fake db store.
type Store struct {
	sessions   map[string]*Session
	profiles   map[string]*Profile
	friends    []*Friends
	askFriends map[string]*AskFriends
}

// NewStore returns a default store.
func NewStore() *Store {
	return &Store{
		sessions:   make(map[string]*Session),
		profiles:   make(map[string]*Profile),
		friends:    []*Friends{},
		askFriends: make(map[string]*AskFriends),
	}
}

// GetSession gets session from db.
func (s *Store) GetSession(sessionID string) (*Session, error) {
	if len(sessionID) == 0 {
		return nil, fmt.Errorf("empty session ID")
	}

	session, found := s.sessions[sessionID]
	if !found {
		return nil, fmt.Errorf("session is not found in db")
	}

	return session, nil
}

// GetProfiles gets profiles from db.
func (s *Store) GetProfiles(profileIDs []string) ([]*Profile, error) {
	ids := RemoveDuplicate(profileIDs)

	if len(ids) == 0 {
		return nil, fmt.Errorf("empty profile ID")
	}

	res := make([]*Profile, 0, len(ids))
	for _, id := range ids {
		profile, found := s.profiles[id]
		if found {
			res = append(res, profile)
		}
	}

	return res, nil
}

// ListFriends fetches all friends of the given profile.
func (s *Store) ListFriends(targetProfileID string) ([]*Friends, error) {
	var res []*Friends

	for _, f := range s.friends {
		if f.MyID == targetProfileID {
			res = append(res, f)
		}
	}

	return res, nil
}

// InsertSession inserts the given session into db.
func (s *Store) InsertSession(session *Session) error {
	_, existed := s.sessions[session.ID]
	if existed {
		return fmt.Errorf("session already existed")
	}

	s.sessions[session.ID] = session

	return nil
}

// InsertProfile inserts the given profile into db.
func (s *Store) InsertProfile(profile *Profile) error {
	_, existed := s.profiles[profile.ID]
	if existed {
		return fmt.Errorf("profile already existed")
	}

	s.profiles[profile.ID] = profile
	return nil
}

// InserFriends inserts the given friends relations into db.
func (s *Store) InserFriends(fs []*Friends) error {
	s.friends = append(s.friends, fs...)
	return nil
}

// AskFriendsAlreadyExisted checks if the given ask friends existed in db.
func (s *Store) AskFriendsAlreadyExisted(a *AskFriends) (bool, error) {
	now := time.Now()
	for _, storedAction := range s.askFriends {
		if storedAction.Asker == a.Asker &&
			storedAction.Receiver == a.Receiver {

			// ask friends existed, but already treated.
			if storedAction.ReceiverTreatedAt != nil &&
				storedAction.ReceiverTreatedAt.Before(now) {
				continue
			}
			return true, nil
		}
	}

	return false, nil
}

// InserAskFriends inserts the given ask friends request into db.
func (s *Store) InserAskFriends(a *AskFriends) error {
	a.ID = GenerateID()
	s.askFriends[a.ID] = a

	return nil
}

// GetAskFriends fetch a AskFriends by ID from db.
func (s *Store) GetAskFriends(askFriendsID string) (*AskFriends, error) {
	if len(askFriendsID) == 0 {
		return nil, fmt.Errorf("empty ask friends ID")
	}

	if res, stored := s.askFriends[askFriendsID]; stored {
		return res, nil
	}

	return nil, fmt.Errorf("ask friends is not found")
}

// UpdateAskFriends updated an existing ask friends in db.
func (s *Store) UpdateAskFriends(a *AskFriends) error {
	_, stored := s.askFriends[a.ID]
	if !stored {
		return fmt.Errorf("ask friends is not found")
	}

	s.askFriends[a.ID] = a
	return nil
}

// GetMyAskFriends returns all my ask friends requests.
// TODO: pagination
func (s *Store) GetMyAskFriends(myProfileID string) ([]*AskFriends, error) {
	if len(myProfileID) == 0 {
		return nil, fmt.Errorf("empty profile ID")
	}

	var res []*AskFriends
	for _, a := range s.askFriends {
		if a.Receiver == myProfileID {
			res = append(res, a)
		}
	}

	return res, nil
}

// InsertReport inserts a report into db.
func (s *Store) InsertReport(report *Report) error {
	return nil
}
