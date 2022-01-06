package persistence

// repository という名前にしたいが domain 配下の repository とパッケージ名が被ってしまうため persistence で代替

import (
	"context"
	"time"

	"github.com/maousmud/memory_share/domain/model"
	"github.com/maousmud/memory_share/domain/repository"
)

type moviePersistence struct{}

// NewMoviePersistence : Movie データに関する Persistence を生成
func NewMoviePersistence() repository.MovieRepository {
	return &moviePersistence{}
}

// GetAll : DB から Movie データを全件取得（MovieRepository インターフェースの GetAll() を実装したもの）
//  -> 本来は DB からデータを取得するが、簡略化のために省略（モックデータを返却）
func (mp moviePersistence) GetAll(context.Context) ([]*model.Movie, error) {
	movie := model.Movie{}
	movie.Id = 1
	movie.Title = "Title"
	movie.Movie_file = []byte{75, 76}
	movie.Register_username = "Register_username"
	movie.Remarks = "Remarks"
	movie.Memory_shubetsu = "Memory_shubetsu"
	movie.Created_date = time.Now().Add(-24 * time.Hour)
	movie.Update_date = time.Now().Add(-24 * time.Hour)
	movie.Del_flg = true

	return []*model.Movie{&movie}, nil
}
