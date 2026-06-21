package audit

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
)

// Log writes an audit entry asynchronously so it never blocks the handler.
func Log(ctx context.Context, db *sqlx.DB, tenantID, schoolID, actorID int64, action, entity string, entityID int64, meta interface{}) {
	var metaJSON []byte
	if meta != nil {
		metaJSON, _ = json.Marshal(meta)
	}
	go func() {
		_, err := db.Exec(`
			INSERT INTO audit_logs (tenant_id, school_id, actor_id, action, entity, entity_id, meta)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			tenantID, schoolID, actorID, action, entity, entityID, string(metaJSON))
		if err != nil {
			log.Printf("[AUDIT ERROR] %v", err)
		}
	}()
}
