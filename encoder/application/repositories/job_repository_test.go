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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_pat", "pending", video)
	require.NotNil(t, job)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job2, err := repoJob.Find(job.ID)
	require.NotNil(t, job2)
	require.NotEmpty(t, job2.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, job2.ID)
	require.Equal(t, job2.VideioID, video.ID)

}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_pat", "pending", video)
	require.NotNil(t, job)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"
	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotNil(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
