package src

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/gofrs/uuid"
	"github.com/graphql-services/graphql-orm-changelog/gen"
	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

func CloudEventsReceive(db *gen.DB) func(event cloudevents.Event) error {
	return func(event cloudevents.Event) (err error) {
		ormEvent := &events.Event{}
		err = event.DataAs(ormEvent)
		if err != nil {
			return
		}

		tx := db.Query().Begin()

		err = StoreEvent(tx, ormEvent)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Commit().Error
		return
	}
}

func StoreEvent(tx *gorm.DB, event *events.Event) (err error) {
	log := &gen.Changelog{
		ID:          event.ID,
		Entity:      event.Entity,
		EntityID:    event.EntityID,
		Date:        time.Now(),
		Type:        gen.ChangelogType(event.Type),
		PrincipalID: event.PrincipalID,
	}

	err = tx.Create(log).Error
	if err != nil {
		return
	}

	changes := []*gen.ChangelogChange{}
	for _, ch := range event.Changes {
		var oldValue string
		var newValue string
		err = ch.OldValueAs(&oldValue)
		if err != nil {
			oldValue = ch.OldValue
		}
		err = ch.NewValueAs(&newValue)
		if err != nil {
			newValue = ch.NewValue
		}
		change := &gen.ChangelogChange{
			ID:       uuid.Must(uuid.NewV4()).String(),
			Column:   ch.Name,
			OldValue: &oldValue,
			NewValue: &newValue,
			LogID:    &event.ID,
		}
		if err := tx.Create(change).Error; err != nil {
			return err
		}
		changes = append(changes, change)
	}
	log.Changes = changes

	return nil
}

func StartCloudEventsServer(db *gen.DB) error {
	portString := os.Getenv("CLOUDEVENTS_PORT")
	if portString == "" {
		portString = "8080"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return err
	}

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithStructuredEncoding(),
	)
	if err != nil {
		return err
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		return err
	}
	log.Printf("cloudevents listening on http://localhost:%d", port)

	log.Fatal(c.StartReceiver(context.Background(), CloudEventsReceive(db)))

	return nil
}
