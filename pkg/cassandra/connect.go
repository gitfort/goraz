package cassandra

import (
	"github.com/gocql/gocql"
	"time"
)

func Connect(db, user, pass string, hosts ...string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: pass,
	}
	cluster.Keyspace = db
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Minute
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}
