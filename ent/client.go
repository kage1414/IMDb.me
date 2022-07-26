// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"imdb-db/ent/migrate"

	"imdb-db/ent/akas"
	"imdb-db/ent/crew"
	"imdb-db/ent/episode"
	"imdb-db/ent/name"
	"imdb-db/ent/principals"
	"imdb-db/ent/ratings"
	"imdb-db/ent/title"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Akas is the client for interacting with the Akas builders.
	Akas *AkasClient
	// Crew is the client for interacting with the Crew builders.
	Crew *CrewClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// Name is the client for interacting with the Name builders.
	Name *NameClient
	// Principals is the client for interacting with the Principals builders.
	Principals *PrincipalsClient
	// Ratings is the client for interacting with the Ratings builders.
	Ratings *RatingsClient
	// Title is the client for interacting with the Title builders.
	Title *TitleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Akas = NewAkasClient(c.config)
	c.Crew = NewCrewClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.Name = NewNameClient(c.config)
	c.Principals = NewPrincipalsClient(c.config)
	c.Ratings = NewRatingsClient(c.config)
	c.Title = NewTitleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Akas:       NewAkasClient(cfg),
		Crew:       NewCrewClient(cfg),
		Episode:    NewEpisodeClient(cfg),
		Name:       NewNameClient(cfg),
		Principals: NewPrincipalsClient(cfg),
		Ratings:    NewRatingsClient(cfg),
		Title:      NewTitleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Akas:       NewAkasClient(cfg),
		Crew:       NewCrewClient(cfg),
		Episode:    NewEpisodeClient(cfg),
		Name:       NewNameClient(cfg),
		Principals: NewPrincipalsClient(cfg),
		Ratings:    NewRatingsClient(cfg),
		Title:      NewTitleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Akas.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Akas.Use(hooks...)
	c.Crew.Use(hooks...)
	c.Episode.Use(hooks...)
	c.Name.Use(hooks...)
	c.Principals.Use(hooks...)
	c.Ratings.Use(hooks...)
	c.Title.Use(hooks...)
}

// AkasClient is a client for the Akas schema.
type AkasClient struct {
	config
}

// NewAkasClient returns a client for the Akas from the given config.
func NewAkasClient(c config) *AkasClient {
	return &AkasClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `akas.Hooks(f(g(h())))`.
func (c *AkasClient) Use(hooks ...Hook) {
	c.hooks.Akas = append(c.hooks.Akas, hooks...)
}

// Create returns a builder for creating a Akas entity.
func (c *AkasClient) Create() *AkasCreate {
	mutation := newAkasMutation(c.config, OpCreate)
	return &AkasCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Akas entities.
func (c *AkasClient) CreateBulk(builders ...*AkasCreate) *AkasCreateBulk {
	return &AkasCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Akas.
func (c *AkasClient) Update() *AkasUpdate {
	mutation := newAkasMutation(c.config, OpUpdate)
	return &AkasUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AkasClient) UpdateOne(a *Akas) *AkasUpdateOne {
	mutation := newAkasMutation(c.config, OpUpdateOne, withAkas(a))
	return &AkasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AkasClient) UpdateOneID(id int) *AkasUpdateOne {
	mutation := newAkasMutation(c.config, OpUpdateOne, withAkasID(id))
	return &AkasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Akas.
func (c *AkasClient) Delete() *AkasDelete {
	mutation := newAkasMutation(c.config, OpDelete)
	return &AkasDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AkasClient) DeleteOne(a *Akas) *AkasDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AkasClient) DeleteOneID(id int) *AkasDeleteOne {
	builder := c.Delete().Where(akas.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AkasDeleteOne{builder}
}

// Query returns a query builder for Akas.
func (c *AkasClient) Query() *AkasQuery {
	return &AkasQuery{
		config: c.config,
	}
}

// Get returns a Akas entity by its id.
func (c *AkasClient) Get(ctx context.Context, id int) (*Akas, error) {
	return c.Query().Where(akas.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AkasClient) GetX(ctx context.Context, id int) *Akas {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AkasClient) Hooks() []Hook {
	return c.hooks.Akas
}

// CrewClient is a client for the Crew schema.
type CrewClient struct {
	config
}

// NewCrewClient returns a client for the Crew from the given config.
func NewCrewClient(c config) *CrewClient {
	return &CrewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `crew.Hooks(f(g(h())))`.
func (c *CrewClient) Use(hooks ...Hook) {
	c.hooks.Crew = append(c.hooks.Crew, hooks...)
}

// Create returns a builder for creating a Crew entity.
func (c *CrewClient) Create() *CrewCreate {
	mutation := newCrewMutation(c.config, OpCreate)
	return &CrewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Crew entities.
func (c *CrewClient) CreateBulk(builders ...*CrewCreate) *CrewCreateBulk {
	return &CrewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Crew.
func (c *CrewClient) Update() *CrewUpdate {
	mutation := newCrewMutation(c.config, OpUpdate)
	return &CrewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CrewClient) UpdateOne(cr *Crew) *CrewUpdateOne {
	mutation := newCrewMutation(c.config, OpUpdateOne, withCrew(cr))
	return &CrewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CrewClient) UpdateOneID(id int) *CrewUpdateOne {
	mutation := newCrewMutation(c.config, OpUpdateOne, withCrewID(id))
	return &CrewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Crew.
func (c *CrewClient) Delete() *CrewDelete {
	mutation := newCrewMutation(c.config, OpDelete)
	return &CrewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CrewClient) DeleteOne(cr *Crew) *CrewDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CrewClient) DeleteOneID(id int) *CrewDeleteOne {
	builder := c.Delete().Where(crew.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CrewDeleteOne{builder}
}

// Query returns a query builder for Crew.
func (c *CrewClient) Query() *CrewQuery {
	return &CrewQuery{
		config: c.config,
	}
}

// Get returns a Crew entity by its id.
func (c *CrewClient) Get(ctx context.Context, id int) (*Crew, error) {
	return c.Query().Where(crew.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CrewClient) GetX(ctx context.Context, id int) *Crew {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CrewClient) Hooks() []Hook {
	return c.hooks.Crew
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Create returns a builder for creating a Episode entity.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id int) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EpisodeClient) DeleteOneID(id int) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id int) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id int) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// NameClient is a client for the Name schema.
type NameClient struct {
	config
}

// NewNameClient returns a client for the Name from the given config.
func NewNameClient(c config) *NameClient {
	return &NameClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `name.Hooks(f(g(h())))`.
func (c *NameClient) Use(hooks ...Hook) {
	c.hooks.Name = append(c.hooks.Name, hooks...)
}

// Create returns a builder for creating a Name entity.
func (c *NameClient) Create() *NameCreate {
	mutation := newNameMutation(c.config, OpCreate)
	return &NameCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Name entities.
func (c *NameClient) CreateBulk(builders ...*NameCreate) *NameCreateBulk {
	return &NameCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Name.
func (c *NameClient) Update() *NameUpdate {
	mutation := newNameMutation(c.config, OpUpdate)
	return &NameUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NameClient) UpdateOne(n *Name) *NameUpdateOne {
	mutation := newNameMutation(c.config, OpUpdateOne, withName(n))
	return &NameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NameClient) UpdateOneID(id int) *NameUpdateOne {
	mutation := newNameMutation(c.config, OpUpdateOne, withNameID(id))
	return &NameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Name.
func (c *NameClient) Delete() *NameDelete {
	mutation := newNameMutation(c.config, OpDelete)
	return &NameDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NameClient) DeleteOne(n *Name) *NameDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NameClient) DeleteOneID(id int) *NameDeleteOne {
	builder := c.Delete().Where(name.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NameDeleteOne{builder}
}

// Query returns a query builder for Name.
func (c *NameClient) Query() *NameQuery {
	return &NameQuery{
		config: c.config,
	}
}

// Get returns a Name entity by its id.
func (c *NameClient) Get(ctx context.Context, id int) (*Name, error) {
	return c.Query().Where(name.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NameClient) GetX(ctx context.Context, id int) *Name {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NameClient) Hooks() []Hook {
	return c.hooks.Name
}

// PrincipalsClient is a client for the Principals schema.
type PrincipalsClient struct {
	config
}

// NewPrincipalsClient returns a client for the Principals from the given config.
func NewPrincipalsClient(c config) *PrincipalsClient {
	return &PrincipalsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `principals.Hooks(f(g(h())))`.
func (c *PrincipalsClient) Use(hooks ...Hook) {
	c.hooks.Principals = append(c.hooks.Principals, hooks...)
}

// Create returns a builder for creating a Principals entity.
func (c *PrincipalsClient) Create() *PrincipalsCreate {
	mutation := newPrincipalsMutation(c.config, OpCreate)
	return &PrincipalsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Principals entities.
func (c *PrincipalsClient) CreateBulk(builders ...*PrincipalsCreate) *PrincipalsCreateBulk {
	return &PrincipalsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Principals.
func (c *PrincipalsClient) Update() *PrincipalsUpdate {
	mutation := newPrincipalsMutation(c.config, OpUpdate)
	return &PrincipalsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrincipalsClient) UpdateOne(pr *Principals) *PrincipalsUpdateOne {
	mutation := newPrincipalsMutation(c.config, OpUpdateOne, withPrincipals(pr))
	return &PrincipalsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrincipalsClient) UpdateOneID(id int) *PrincipalsUpdateOne {
	mutation := newPrincipalsMutation(c.config, OpUpdateOne, withPrincipalsID(id))
	return &PrincipalsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Principals.
func (c *PrincipalsClient) Delete() *PrincipalsDelete {
	mutation := newPrincipalsMutation(c.config, OpDelete)
	return &PrincipalsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PrincipalsClient) DeleteOne(pr *Principals) *PrincipalsDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PrincipalsClient) DeleteOneID(id int) *PrincipalsDeleteOne {
	builder := c.Delete().Where(principals.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrincipalsDeleteOne{builder}
}

// Query returns a query builder for Principals.
func (c *PrincipalsClient) Query() *PrincipalsQuery {
	return &PrincipalsQuery{
		config: c.config,
	}
}

// Get returns a Principals entity by its id.
func (c *PrincipalsClient) Get(ctx context.Context, id int) (*Principals, error) {
	return c.Query().Where(principals.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrincipalsClient) GetX(ctx context.Context, id int) *Principals {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PrincipalsClient) Hooks() []Hook {
	return c.hooks.Principals
}

// RatingsClient is a client for the Ratings schema.
type RatingsClient struct {
	config
}

// NewRatingsClient returns a client for the Ratings from the given config.
func NewRatingsClient(c config) *RatingsClient {
	return &RatingsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ratings.Hooks(f(g(h())))`.
func (c *RatingsClient) Use(hooks ...Hook) {
	c.hooks.Ratings = append(c.hooks.Ratings, hooks...)
}

// Create returns a builder for creating a Ratings entity.
func (c *RatingsClient) Create() *RatingsCreate {
	mutation := newRatingsMutation(c.config, OpCreate)
	return &RatingsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ratings entities.
func (c *RatingsClient) CreateBulk(builders ...*RatingsCreate) *RatingsCreateBulk {
	return &RatingsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ratings.
func (c *RatingsClient) Update() *RatingsUpdate {
	mutation := newRatingsMutation(c.config, OpUpdate)
	return &RatingsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RatingsClient) UpdateOne(r *Ratings) *RatingsUpdateOne {
	mutation := newRatingsMutation(c.config, OpUpdateOne, withRatings(r))
	return &RatingsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RatingsClient) UpdateOneID(id int) *RatingsUpdateOne {
	mutation := newRatingsMutation(c.config, OpUpdateOne, withRatingsID(id))
	return &RatingsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ratings.
func (c *RatingsClient) Delete() *RatingsDelete {
	mutation := newRatingsMutation(c.config, OpDelete)
	return &RatingsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RatingsClient) DeleteOne(r *Ratings) *RatingsDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *RatingsClient) DeleteOneID(id int) *RatingsDeleteOne {
	builder := c.Delete().Where(ratings.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RatingsDeleteOne{builder}
}

// Query returns a query builder for Ratings.
func (c *RatingsClient) Query() *RatingsQuery {
	return &RatingsQuery{
		config: c.config,
	}
}

// Get returns a Ratings entity by its id.
func (c *RatingsClient) Get(ctx context.Context, id int) (*Ratings, error) {
	return c.Query().Where(ratings.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RatingsClient) GetX(ctx context.Context, id int) *Ratings {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RatingsClient) Hooks() []Hook {
	return c.hooks.Ratings
}

// TitleClient is a client for the Title schema.
type TitleClient struct {
	config
}

// NewTitleClient returns a client for the Title from the given config.
func NewTitleClient(c config) *TitleClient {
	return &TitleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `title.Hooks(f(g(h())))`.
func (c *TitleClient) Use(hooks ...Hook) {
	c.hooks.Title = append(c.hooks.Title, hooks...)
}

// Create returns a builder for creating a Title entity.
func (c *TitleClient) Create() *TitleCreate {
	mutation := newTitleMutation(c.config, OpCreate)
	return &TitleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Title entities.
func (c *TitleClient) CreateBulk(builders ...*TitleCreate) *TitleCreateBulk {
	return &TitleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Title.
func (c *TitleClient) Update() *TitleUpdate {
	mutation := newTitleMutation(c.config, OpUpdate)
	return &TitleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TitleClient) UpdateOne(t *Title) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitle(t))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TitleClient) UpdateOneID(id int) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitleID(id))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Title.
func (c *TitleClient) Delete() *TitleDelete {
	mutation := newTitleMutation(c.config, OpDelete)
	return &TitleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TitleClient) DeleteOne(t *Title) *TitleDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TitleClient) DeleteOneID(id int) *TitleDeleteOne {
	builder := c.Delete().Where(title.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TitleDeleteOne{builder}
}

// Query returns a query builder for Title.
func (c *TitleClient) Query() *TitleQuery {
	return &TitleQuery{
		config: c.config,
	}
}

// Get returns a Title entity by its id.
func (c *TitleClient) Get(ctx context.Context, id int) (*Title, error) {
	return c.Query().Where(title.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TitleClient) GetX(ctx context.Context, id int) *Title {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TitleClient) Hooks() []Hook {
	return c.hooks.Title
}
