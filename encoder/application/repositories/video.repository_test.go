package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	v, err := repo.Insert(video)
	require.NotNil(t, v)
	require.Nil(t, err)

	v2, err := repo.Find(video.ID)
	require.NotNil(t, v2)
	require.Nil(t, err)
	require.Equal(t, v.ID, v2.ID)
}
