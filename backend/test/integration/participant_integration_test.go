package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoResource *dockertest.Resource
	mongoPool     *dockertest.Pool
	mongoClient   *mongo.Client
	mongoDB       *mongo.Database
)

func setupMongoDockerTest(t *testing.T) {
	var err error
	mongoPool, err = dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	mongoResource, err = mongoPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "6.0",
		Env:        []string{},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	if err = mongoPool.Retry(func() error {
		var err error
		mongoURI := fmt.Sprintf("mongodb://localhost:%s", mongoResource.GetPort("27017/tcp"))
		mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
		if err != nil {
			return err
		}
		return mongoClient.Ping(context.Background(), nil)
	}); err != nil {
		t.Fatalf("Could not connect to docker mongo: %s", err)
	}

	mongoDB = mongoClient.Database("decidr_test_db")
}

func teardownMongoDockerTest(t *testing.T) {
	if mongoClient != nil {
		_ = mongoClient.Disconnect(context.Background())
	}
	if mongoPool != nil && mongoResource != nil {
		_ = mongoPool.Purge(mongoResource)
	}
}

func insertParticipants(t *testing.T, participants []interface{}) {
	ctx := context.Background()
	collection := mongoDB.Collection("participants")
	_, err := collection.InsertMany(ctx, participants)
	require.NoError(t, err)
}

func TestIntegration_GetParticipants(t *testing.T) {
	setupMongoDockerTest(t)
	defer teardownMongoDockerTest(t)

	ctx := context.Background()

	participants := []interface{}{
		entities.Participant{ID: "1", Name: "Juliette", IsNominated: true, IsEliminated: false},
		entities.Participant{ID: "2", Name: "Gil", IsNominated: false, IsEliminated: false},
		entities.Participant{ID: "3", Name: "Camilla", IsNominated: true, IsEliminated: false},
		entities.Participant{ID: "4", Name: "Vitoria", IsNominated: false, IsEliminated: true},
	}
	insertParticipants(t, participants)

	repo := repositories.NewParticipantRepository(mongoDB)
	svc := services.NewParticipantService(repo)

	t.Run("should return only nominated participants", func(t *testing.T) {
		nominated := true
		result, err := svc.GetParticipants(ctx, &nominated)
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.ElementsMatch(t, []string{"Juliette", "Camilla"}, []string{result[0].Name, result[1].Name})
	})

	t.Run("should return non nominated participants", func(t *testing.T) {
		nominated := false
		result, err := svc.GetParticipants(ctx, &nominated)
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.ElementsMatch(t, []string{"Gil", "Vitoria"}, []string{result[0].Name, result[1].Name})
	})

	t.Run("should return all participants without filter", func(t *testing.T) {
		result, err := svc.GetParticipants(ctx, nil)
		require.NoError(t, err)
		assert.Len(t, result, 4)
	})
}
