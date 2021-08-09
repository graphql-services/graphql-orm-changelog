package main

import (
	"github.com/graphql-services/graphql-orm-changelog/gen"
	"github.com/graphql-services/graphql-orm-changelog/src"
	handler "github.com/jakubknejzlik/cloudevents-lambda-handler"
)

func main() {
	db := gen.NewDBFromEnvVars()
	h := handler.NewCloudEventsLambdaHandler(src.CloudEventsReceive(db))
	h.Start()
}
