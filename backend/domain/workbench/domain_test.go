package workbench

import "testing"

type mockRepo struct {
	statuses map[string]string
	updated  map[string]string
	events   map[string][]ActionEvent
}

func (m *mockRepo) ListStatuses() (map[string]string, error) {
	if m.statuses == nil {
		return map[string]string{}, nil
	}
	return m.statuses, nil
}

func (m *mockRepo) UpsertStatus(cardID string, status string) error {
	if m.updated == nil {
		m.updated = map[string]string{}
	}
	m.updated[cardID] = status
	return nil
}

func (m *mockRepo) AppendEvent(cardID string, fromState string, toState string, reason string) error {
	if m.events == nil {
		m.events = map[string][]ActionEvent{}
	}
	m.events[cardID] = append(m.events[cardID], ActionEvent{
		CardID:    cardID,
		FromState: fromState,
		ToState:   toState,
		Reason:    reason,
	})
	return nil
}

func (m *mockRepo) ListEvents(cardID string) ([]ActionEvent, error) {
	return m.events[cardID], nil
}

func TestTransitionActionSuccess(t *testing.T) {
	store := NewStore(&mockRepo{})

	overview := store.Overview("store_manager")
	if len(overview.Opportunities) == 0 {
		t.Fatalf("expected opportunities")
	}

	cardID := overview.Opportunities[0].ID
	if overview.Opportunities[0].Status != "new" {
		t.Fatalf("expected initial status new, got %s", overview.Opportunities[0].Status)
	}

	updated, err := store.Transition(cardID, "accepted", "")
	if err != nil {
		t.Fatalf("expected transition success, got error: %v", err)
	}
	if updated.Status != "accepted" {
		t.Fatalf("expected accepted, got %s", updated.Status)
	}

	overview = store.Overview("store_manager")
	var found *Card
	for _, c := range overview.Opportunities {
		if c.ID == cardID {
			copy := c
			found = &copy
			break
		}
	}
	if found == nil {
		t.Fatalf("expected card in overview")
	}
	if found.Status != "accepted" {
		t.Fatalf("expected accepted in overview, got %s", found.Status)
	}
}

func TestTransitionActionUnknownID(t *testing.T) {
	store := NewStore(&mockRepo{})
	_, err := store.Transition("not-found", "accepted", "")
	if err == nil {
		t.Fatalf("expected error for unknown id")
	}
}

func TestTransitionActionPersistsStatus(t *testing.T) {
	repo := &mockRepo{}
	store := NewStore(repo)

	overview := store.Overview("store_manager")
	cardID := overview.Opportunities[0].ID

	_, err := store.Transition(cardID, "accepted", "")
	if err != nil {
		t.Fatalf("expected transition success, got: %v", err)
	}

	if repo.updated[cardID] != "accepted" {
		t.Fatalf("expected persisted status accepted, got %q", repo.updated[cardID])
	}
}

func TestOverviewLoadsPersistedStatuses(t *testing.T) {
	repo := &mockRepo{
		statuses: map[string]string{
			"opp-1": "in_progress",
		},
	}
	store := NewStore(repo)
	overview := store.Overview("store_manager")

	if overview.Opportunities[0].Status != "in_progress" {
		t.Fatalf("expected status from repo, got %s", overview.Opportunities[0].Status)
	}
}

func TestTransitionActionPersistsEvent(t *testing.T) {
	repo := &mockRepo{}
	store := NewStore(repo)
	overview := store.Overview("store_manager")
	cardID := overview.Opportunities[0].ID

	_, err := store.Transition(cardID, "accepted", "operator_accepted")
	if err != nil {
		t.Fatalf("expected transition success, got: %v", err)
	}

	events := repo.events[cardID]
	if len(events) != 1 {
		t.Fatalf("expected one event, got %d", len(events))
	}
	if events[0].FromState != "new" || events[0].ToState != "accepted" {
		t.Fatalf("unexpected event transition %s -> %s", events[0].FromState, events[0].ToState)
	}
}
