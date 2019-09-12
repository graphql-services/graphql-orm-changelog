package src

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm-changelog/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	return resolver
}

func CloudEventsReceive(db *gen.DB) func(event cloudevents.Event) error {
	// do something with event.Context and event.Data (via event.DataAs(foo)
	// fmt.Println("new event")
	return func(event cloudevents.Event) (err error) {
		log.Println("new event")
		ormEvent := &events.Event{}
		err = event.DataAs(ormEvent)
		if err != nil {
			return
		}

		tx := db.Query().Begin()

		err = storeEvent(tx, ormEvent)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Commit().Error
		return
	}
}

func storeEvent(tx *gorm.DB, event *events.Event) error {
	log := &gen.Changelog{
		Entity:      event.Entity,
		EntityID:    event.EntityID,
		Date:        time.Now(),
		Type:        gen.ChangelogType(event.Type),
		PrincipalID: event.PrincipalID,
	}

	changes := []*gen.ChangelogChange{}
	for _, ch := range event.Changes {
		oldValue := fmt.Sprintf("%v", ch.OldValue)
		newValue := fmt.Sprintf("%v", ch.NewValue)
		changes = append(changes, &gen.ChangelogChange{
			Column:   ch.Name,
			OldValue: &oldValue,
			NewValue: &newValue,
			Log:      log,
		})
	}

	return tx.Create(log).Error
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