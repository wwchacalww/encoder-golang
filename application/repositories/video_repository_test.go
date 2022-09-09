package repositories_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/wwchacalww/encoder-golang/application/repositories"
	"github.com/wwchacalww/encoder-golang/domain"
	"github.com/wwchacalww/encoder-golang/framework/database"
)

func TestVideRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.Filepath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)

}
