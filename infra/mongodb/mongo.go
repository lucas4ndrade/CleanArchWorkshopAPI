package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo represents the interface of functions we can use in Mongo
type Mongo interface {
	Close()
	GetConnection() *mongo.Client
	Ping() error
}

// Client represents the connection of Mongo
type Client struct {
	conn *mongo.Client
}

// New Connect and returns the Mongo Client that implements the Mongo interface.
func New(uri string) (m Mongo, err error) {
	c := &Client{}
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(100).
		SetReadPreference(readpref.SecondaryPreferred())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c.conn, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}

	m = c
	return
}

// Close will close the connection.
func (c *Client) Close() {
	if c.conn != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = c.conn.Disconnect(ctx)
	}
}

// Ping will show if mongo is connected.
func (c *Client) Ping() (err error) {
	if c.conn != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = c.conn.Ping(ctx, nil)
	}
	return
}

// GetConnection returns the instance of mongo client.
func (c *Client) GetConnection() *mongo.Client {
	return c.conn
}
