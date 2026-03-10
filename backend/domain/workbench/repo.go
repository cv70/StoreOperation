package workbench

import (
	"backend/datasource/dbdao"
)

type Repository interface {
	ListStatuses() (map[string]string, error)
	UpsertStatus(cardID string, status string) error
	AppendEvent(cardID string, fromState string, toState string, reason string) error
	ListEvents(cardID string) ([]ActionEvent, error)
}

type noopRepo struct{}

func (noopRepo) ListStatuses() (map[string]string, error) {
	return map[string]string{}, nil
}

func (noopRepo) UpsertStatus(_ string, _ string) error {
	return nil
}

func (noopRepo) AppendEvent(_ string, _ string, _ string, _ string) error {
	return nil
}

func (noopRepo) ListEvents(_ string) ([]ActionEvent, error) {
	return []ActionEvent{}, nil
}

type dbRepo struct {
	db *dbdao.DB
}

func NewRepository(db *dbdao.DB) (Repository, error) {
	if db == nil {
		return noopRepo{}, nil
	}
	if err := db.DB().AutoMigrate(&dbdao.WorkbenchActionCard{}, &dbdao.WorkbenchActionEvent{}); err != nil {
		return nil, err
	}
	return &dbRepo{db: db}, nil
}

func (r *dbRepo) ListStatuses() (map[string]string, error) {
	return r.db.ListWorkbenchActionStatuses()
}

func (r *dbRepo) UpsertStatus(cardID string, status string) error {
	return r.db.UpsertWorkbenchActionStatus(cardID, status)
}

func (r *dbRepo) AppendEvent(cardID string, fromState string, toState string, reason string) error {
	return r.db.CreateWorkbenchActionEvent(cardID, fromState, toState, reason)
}

func (r *dbRepo) ListEvents(cardID string) ([]ActionEvent, error) {
	rows, err := r.db.ListWorkbenchActionEvents(cardID)
	if err != nil {
		return nil, err
	}
	out := make([]ActionEvent, 0, len(rows))
	for _, row := range rows {
		out = append(out, ActionEvent{
			CardID:    row.CardID,
			FromState: row.FromState,
			ToState:   row.ToState,
			Reason:    row.Reason,
			CreatedAt: row.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return out, nil
}
