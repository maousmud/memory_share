package repository

import (
	"context"

	"github.com/maousmud/memory_share/domain/model"
)

// MovieRepository : Movie における Repository のインターフェース
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type MovieRepository interface {
	GetAll(context.Context) ([]*model.Movie, error)
}
