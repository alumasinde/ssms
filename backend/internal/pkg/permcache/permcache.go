// Package permcache provides a lightweight in-process permission cache
// so that RequirePermission does not hit the DB on every single request.
//
// FIX: previous version queried `WHERE rp.role = ?` which does not exist
// in the schema. The correct query joins through the roles table.
package permcache

import (
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

const ttl = 5 * time.Minute

type entry struct {
	perms     map[string]struct{}
	expiresAt time.Time
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]*entry // key: role name
	db    *sqlx.DB
}

func New(db *sqlx.DB) *Cache {
	c := &Cache{store: make(map[string]*entry), db: db}
	go c.sweep()
	return c
}

// Has returns true when the role has the named permission.
func (c *Cache) Has(role, permission string) bool {
	c.mu.RLock()
	e, ok := c.store[role]
	c.mu.RUnlock()

	if ok && time.Now().Before(e.expiresAt) {
		_, has := e.perms[permission]
		return has
	}

	// Cache miss or expired — load from DB
	perms, err := c.load(role)
	if err != nil {
		return false
	}
	c.mu.Lock()
	c.store[role] = &entry{perms: perms, expiresAt: time.Now().Add(ttl)}
	c.mu.Unlock()

	_, has := perms[permission]
	return has
}

// Invalidate removes all cached entries for a role (call after permission changes).
func (c *Cache) Invalidate(role string) {
	c.mu.Lock()
	delete(c.store, role)
	c.mu.Unlock()
}

func (c *Cache) load(role string) (map[string]struct{}, error) {
	var names []string
	err := c.db.Select(&names, `
		SELECT p.name
		FROM permissions p
		JOIN role_permissions rp ON rp.permission_id = p.id
		JOIN roles r             ON r.id             = rp.role_id
		WHERE r.code     = ?
		  AND r.is_active = 1
	`, role)
	if err != nil {
		return nil, err
	}
	m := make(map[string]struct{}, len(names))
	for _, n := range names {
		m[n] = struct{}{}
	}
	return m, nil
}

// sweep removes expired entries every minute to keep memory bounded.
func (c *Cache) sweep() {
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for range t.C {
		now := time.Now()
		c.mu.Lock()
		for k, e := range c.store {
			if now.After(e.expiresAt) {
				delete(c.store, k)
			}
		}
		c.mu.Unlock()
	}
}
