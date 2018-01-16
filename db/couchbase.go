package db

import (
	"fmt"

	"github.com/couchbase/gocb"
)

// ConnectCouchBase : ConnectCouchBase to CouchBase Server
func ConnectCouchBase(host, user, pass, path string) *gocb.Bucket {
	cluster, err := gocb.Connect("couchbase://" + host)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: pass,
	})

	if err != nil {
		fmt.Println("[CouchBase] [ERROR] CouchBase Connection Error: ", err.Error())
		panic(err)
	}

	bucket, err := cluster.OpenBucket(path, "")

	if err != nil {
		fmt.Printf("[CouchBase] [ERROR] CouchBase Bucket Error: %s\n", err.Error())
		panic(err)
	}

	fmt.Printf("[CouchBase] [INFO] Connection Established: couchbase://%s\n", host)

	return bucket
}
