package usecase

import (
	"context"

	"github.com/maousmud/memory_share/domain/model"
	"github.com/maousmud/memory_share/domain/repository"
)

// MovieUseCase : Movie における UseCase のインターフェース
type MovieUseCase interface {
	GetAll(context.Context) ([]*model.Movie, error)
}

type movieUseCase struct {
	movieRepository repository.MovieRepository
}

// NewMovieUseCase : Movie データに関する UseCase を生成
func NewBookUseCase(mr repository.MovieRepository) MovieUseCase {
	return &movieUseCase{
		movieRepository: mr,
	}
}

// GetAll : Movie データを全件取得するためのユースケース
//  -> 本システムではあまりユースケース層の恩恵を受けれないが、もう少し大きなシステムになってくると、
//    「ドメインモデルの調節者」としての役割が見えてくる
func (mu movieUseCase) GetAll(ctx context.Context) (movie []*model.Movie, err error) {
	// Persistence（Repository）を呼出
	movie, err = mu.movieRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
