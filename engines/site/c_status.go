package site

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) getAdminSiteStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	status := web.H{"os": p.runtimeStatus()}
	var err error
	if status["db"], err = p.dbStatus(); err != nil {
		return nil, err
	}
	if status["redis"], err = p.redisStatus(); err != nil {
		return nil, err
	}
	return status, nil
}

func (p *Engine) dbStatus() (interface{}, error) {
	type Status struct {
		Version string
	}
	var sts Status
	if err := p.Db.Raw("SELECT VERSION() AS version").Scan(&sts).Error; err != nil {
		return nil, err
	}
	return web.H{
		"version": sts.Version,
	}, nil
}

func (p *Engine) redisStatus() (interface{}, error) {
	c := p.Redis.Get()
	defer c.Close()

	sts, err := redis.String(c.Do("INFO"))
	if err != nil {
		return nil, err
	}
	return string(sts), nil
}

func (p *Engine) runtimeStatus() interface{} {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return web.H{
		"GO":           runtime.Version(),
		"OS":           runtime.GOOS,
		"ARCH":         runtime.GOARCH,
		"CPUS":         runtime.NumCPU(),
		"MEMORY USAGE": fmt.Sprintf("%d/%d MB", mem.Alloc/(1024*1024), mem.Sys/(1024*1024)),
		"LAST GC":      time.Unix(0, int64(mem.LastGC)),
	}
}
