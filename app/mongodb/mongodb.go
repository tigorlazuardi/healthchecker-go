package mongodb

import "github.com/tigorlazuardi/healthchecker/pkg"

type MongoDBHealthChecker struct {
	client pkg.Doer
}
