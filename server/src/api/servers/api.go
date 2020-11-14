package servers

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/openmultiplayer/server-index/server/src/queryer"
	"github.com/openmultiplayer/server-index/server/src/serverdb"
)

type service struct {
	ctx     context.Context
	storer  serverdb.Storer
	queryer queryer.Queryer
}

func New(ctx context.Context, storer serverdb.Storer, queryer queryer.Queryer) *chi.Mux {
	rtr := chi.NewRouter()
	svc := service{ctx, storer, queryer}

	rtr.Get("/{address}", svc.get)

	rtr.Get("/", svc.list)

	return rtr
}
