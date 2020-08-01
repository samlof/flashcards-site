// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"flashcards-backend/ent/migrate"

	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CardLog is the client for interacting with the CardLog builders.
	CardLog *CardLogClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Word is the client for interacting with the Word builders.
	Word *WordClient
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
	c.CardLog = NewCardLogClient(c.config)
	c.User = NewUserClient(c.config)
	c.Word = NewWordClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		CardLog: NewCardLogClient(cfg),
		User:    NewUserClient(cfg),
		Word:    NewWordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:  cfg,
		CardLog: NewCardLogClient(cfg),
		User:    NewUserClient(cfg),
		Word:    NewWordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CardLog.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.CardLog.Use(hooks...)
	c.User.Use(hooks...)
	c.Word.Use(hooks...)
}

// CardLogClient is a client for the CardLog schema.
type CardLogClient struct {
	config
}

// NewCardLogClient returns a client for the CardLog from the given config.
func NewCardLogClient(c config) *CardLogClient {
	return &CardLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cardlog.Hooks(f(g(h())))`.
func (c *CardLogClient) Use(hooks ...Hook) {
	c.hooks.CardLog = append(c.hooks.CardLog, hooks...)
}

// Create returns a create builder for CardLog.
func (c *CardLogClient) Create() *CardLogCreate {
	mutation := newCardLogMutation(c.config, OpCreate)
	return &CardLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for CardLog.
func (c *CardLogClient) Update() *CardLogUpdate {
	mutation := newCardLogMutation(c.config, OpUpdate)
	return &CardLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardLogClient) UpdateOne(cl *CardLog) *CardLogUpdateOne {
	mutation := newCardLogMutation(c.config, OpUpdateOne, withCardLog(cl))
	return &CardLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardLogClient) UpdateOneID(id int) *CardLogUpdateOne {
	mutation := newCardLogMutation(c.config, OpUpdateOne, withCardLogID(id))
	return &CardLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CardLog.
func (c *CardLogClient) Delete() *CardLogDelete {
	mutation := newCardLogMutation(c.config, OpDelete)
	return &CardLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CardLogClient) DeleteOne(cl *CardLog) *CardLogDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CardLogClient) DeleteOneID(id int) *CardLogDeleteOne {
	builder := c.Delete().Where(cardlog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardLogDeleteOne{builder}
}

// Create returns a query builder for CardLog.
func (c *CardLogClient) Query() *CardLogQuery {
	return &CardLogQuery{config: c.config}
}

// Get returns a CardLog entity by its id.
func (c *CardLogClient) Get(ctx context.Context, id int) (*CardLog, error) {
	return c.Query().Where(cardlog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardLogClient) GetX(ctx context.Context, id int) *CardLog {
	cl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return cl
}

// QueryUser queries the user edge of a CardLog.
func (c *CardLogClient) QueryUser(cl *CardLog) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cardlog.Table, cardlog.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cardlog.UserTable, cardlog.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCard queries the card edge of a CardLog.
func (c *CardLogClient) QueryCard(cl *CardLog) *WordQuery {
	query := &WordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cardlog.Table, cardlog.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cardlog.CardTable, cardlog.CardColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CardLogClient) Hooks() []Hook {
	return c.hooks.CardLog
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryCardLogs queries the cardLogs edge of a User.
func (c *UserClient) QueryCardLogs(u *User) *CardLogQuery {
	query := &CardLogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(cardlog.Table, cardlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CardLogsTable, user.CardLogsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// WordClient is a client for the Word schema.
type WordClient struct {
	config
}

// NewWordClient returns a client for the Word from the given config.
func NewWordClient(c config) *WordClient {
	return &WordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `word.Hooks(f(g(h())))`.
func (c *WordClient) Use(hooks ...Hook) {
	c.hooks.Word = append(c.hooks.Word, hooks...)
}

// Create returns a create builder for Word.
func (c *WordClient) Create() *WordCreate {
	mutation := newWordMutation(c.config, OpCreate)
	return &WordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Word.
func (c *WordClient) Update() *WordUpdate {
	mutation := newWordMutation(c.config, OpUpdate)
	return &WordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WordClient) UpdateOne(w *Word) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWord(w))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WordClient) UpdateOneID(id int) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWordID(id))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Word.
func (c *WordClient) Delete() *WordDelete {
	mutation := newWordMutation(c.config, OpDelete)
	return &WordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WordClient) DeleteOne(w *Word) *WordDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WordClient) DeleteOneID(id int) *WordDeleteOne {
	builder := c.Delete().Where(word.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WordDeleteOne{builder}
}

// Create returns a query builder for Word.
func (c *WordClient) Query() *WordQuery {
	return &WordQuery{config: c.config}
}

// Get returns a Word entity by its id.
func (c *WordClient) Get(ctx context.Context, id int) (*Word, error) {
	return c.Query().Where(word.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WordClient) GetX(ctx context.Context, id int) *Word {
	w, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return w
}

// QueryCardLogs queries the cardLogs edge of a Word.
func (c *WordClient) QueryCardLogs(w *Word) *CardLogQuery {
	query := &CardLogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(cardlog.Table, cardlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, word.CardLogsTable, word.CardLogsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WordClient) Hooks() []Hook {
	return c.hooks.Word
}
