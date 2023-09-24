package services

import (
	"log"
)

const (
	ServiceTypeElasticSearch = "elasticsearch"
	ServiceTypeMemcached     = "memcached"
	ServiceTypePostgreSQL    = "pgsql"
	ServiceTypeRedis         = "redis"
)

func ServiceTypeValues() []string {
	return []string{
		ServiceTypeElasticSearch,
		ServiceTypeMemcached,
		ServiceTypePostgreSQL,
		ServiceTypeRedis,
	}
}

const (
	ServiceClassCacher   = "cacher"
	ServiceClassDatabase = "database"
	ServiceClassSearch   = "search"
)

// Map with ServiceManager objects for each supported PaaS service.
var managers = map[string]ServiceManager{
	ElasticSearch.ServiceType(): ElasticSearch,
	Memcached.ServiceType():     Memcached,
}

func GetServiceManager(serviceType string) ServiceManager {
	if v, ok := managers[serviceType]; ok {
		return v
	}

	log.Printf("[ERROR] Can't found parameters manager for service type %s", serviceType)
	return nil
}
