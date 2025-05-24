package runner

import "context"

type Resolver func(ctx context.Context, task *Task) error