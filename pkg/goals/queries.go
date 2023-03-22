package goals

import (
	"context"
	"database/sql"

	"github.com/gocopper/copper/csql"
)

var ErrRecordNotFound = sql.ErrNoRows

func NewQueries(querier csql.Querier) *Queries {
	return &Queries{
		querier: querier,
	}
}

type Queries struct {
	querier csql.Querier
}

func (q *Queries) CreateGoal(ctx context.Context, goal *Goal) error {
	const query = `
	INSERT INTO goals (name, total_needed, current_savings, monthly_contribution, start_date, end_date, months_to_goal)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := q.querier.Exec(ctx, query,
		goal.Name,
		goal.TotalNeeded,
		goal.CurrentSavings,
		goal.MonthlyContribution,
		goal.StartDate,
		goal.EndDate,
		goal.MonthsToGoal,
	)

	return err
}

/*
Here are some example queries that use Querier to unmarshal results into Go strcuts

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	const query = "SELECT * FROM posts ORDER BY created_at DESC"

	var (
	    posts []Post
	    err = q.querier.Select(ctx, &posts, query)
    )

	return posts, err
}

func (q *Queries) GetPostByID(ctx context.Context, id string) (*Post, error) {
	const query = "SELECT * from posts where id=?"

	var (
	    post Post
	    err = q.querier.Get(ctx, &post, query, id)
    )

	return &post, err
}

func (q *Queries) SavePost(ctx context.Context, post *Post) error {
	const query = `
	INSERT INTO posts (id, title, url, poster)
	VALUES (?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET title=?, url=?`

	_, err := q.querier.Exec(ctx, query,
		post.ID,
		post.Title,
		post.URL,
		post.Poster,
		post.Title,
		post.URL,
	)

	return err
}
*/
