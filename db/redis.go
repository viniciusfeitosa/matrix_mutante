package db

import (
	"log"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// Pool is the interface to pool of redis
type Pool interface {
	Get() redigo.Conn
}

// DB is the struc with db configuration
type DB struct {
	Enable          bool
	MaxIdle         int
	MaxActive       int
	IdleTimeoutSecs int
	Address         string
	Auth            string
	DB              string
	Pool            *redigo.Pool
}

// NewDBPool return a new instance of the redis pool
func (db *DB) NewDBPool() *redigo.Pool {
	if db.Enable {
		pool := &redigo.Pool{
			MaxIdle:     db.MaxIdle,
			MaxActive:   db.MaxActive,
			IdleTimeout: time.Second * time.Duration(db.IdleTimeoutSecs),
			Dial: func() (redigo.Conn, error) {
				c, err := redigo.Dial("tcp", db.Address)
				if err != nil {
					return nil, err
				}
				// if _, err = c.Do("AUTH", db.Auth); err != nil {
				// 	c.Close()
				// 	return nil, err
				// }
				if _, err = c.Do("SELECT", db.DB); err != nil {
					c.Close()
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redigo.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}

		c := pool.Get() // Test connection during init
		if _, err := c.Do("PING"); err != nil {
			log.Fatal("Cannot connect to Redis: ", err)
		}
		return pool
	}
	return nil
}

// GetValue is a method resposible to get a value from the DB using a key
func (db *DB) GetValue(key interface{}) (string, error) {
	if db.Enable {
		conn := db.Pool.Get()
		defer conn.Close()
		value, err := redigo.String(conn.Do("GET", key))
		return value, err
	}
	return "", nil
}

// SetValue is a method resposible to set a value on the DB appling a key as reference
func (db *DB) SetValue(key interface{}, value interface{}) error {
	if db.Enable {
		conn := db.Pool.Get()
		defer conn.Close()
		_, err := redigo.String(conn.Do("SET", key, value))
		return err
	}
	return nil
}

// DelValue is a method resposible to delete a value from the DB by a key
func (db *DB) DelValue(key interface{}) error {
	if db.Enable {
		conn := db.Pool.Get()
		defer conn.Close()
		_, err := redigo.String(conn.Do("DEL", key))
		return err
	}
	return nil
}

// EnqueueValue is the responsible to add a value on the queue
func (db *DB) EnqueueValue(queue string, uuid uint32) error {
	if db.Enable {
		conn := db.Pool.Get()
		defer conn.Close()
		_, err := conn.Do("RPUSH", queue, uuid)
		return err
	}
	return nil
}

// PopQueue is the responsible to get a value from the queue
func (db *DB) PopQueue(queue string, id int) (uint32, string, error) {
	if db.Enable {
		conn := db.Pool.Get()
		defer conn.Close()

		var channel string
		var uuid uint32
		var values string
		if reply, err := redigo.Values(conn.Do("BLPOP", queue, 30+id)); err == nil {

			if _, err := redigo.Scan(reply, &channel, &uuid); err != nil {
				db.EnqueueValue(queue, uuid)
				return 0, "", err
			}

			values, err = redigo.String(conn.Do("GET", uuid))
			if err != nil {
				db.EnqueueValue(queue, uuid)
				return 0, "", err
			}

		} else if err != redigo.ErrNil {
			log.Fatal(err)
		}
		return uuid, values, nil
	}
	return 0, "", nil
}
