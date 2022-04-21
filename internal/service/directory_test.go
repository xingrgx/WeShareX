package service

import (
	"context"
	"testing"
)

var ctx context.Context

func Test_deleteDirRecursively(t *testing.T) {
	deleteDirRecursively(ctx, "1vlyjjj8is0cjfl6lnhpm6sr00sb6kik", 1)
}
