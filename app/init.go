package app

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	db "github.com/revel/revel/modules/db/app"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func InitDB() {
	db.Init()

	qstr := `

-- ----------------------------
--  Sequence structure for users_id_seq
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" INCREMENT 1 START 3 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

-- ----------------------------
--  Table structure for users
-- ----------------------------
CREATE TABLE "public"."users" (
	"id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
	"firstname" varchar(255) NOT NULL COLLATE "default",
	"lastname" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of users
-- ----------------------------
BEGIN;
INSERT INTO "public"."users" VALUES ('1', 'John', 'Doe');
INSERT INTO "public"."users" VALUES ('2', 'Jane', 'Doe');
COMMIT;


-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq" RESTART 3 OWNED BY "users"."id";

-- ----------------------------
--  Primary key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
 `

	_, err := db.Db.Exec(qstr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database rebuilded successfully")
}
