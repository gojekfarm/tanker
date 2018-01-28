package shippers

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/postgres"
	"github.com/sudhanshuraheja/tanker/pkg/postgresmock"
)

var shipperDatastoreTestContext *appcontext.AppContext

func NewTestDatastoreContext() *appcontext.AppContext {
	if shipperDatastoreTestContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		shipperDatastoreTestContext = appcontext.NewAppContext(conf, log)
	}
	return shipperDatastoreTestContext
}

func NewTestPostgresDatastore(ctx *appcontext.AppContext) *sqlx.DB {
	// Not required unless you want to hit an actual DB
	return postgres.NewPostgres(ctx.GetLogger(), ctx.GetConfig().Database().ConnectionURL(), 10)
}

func TestShippersDatastoreAdd(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^INSERT INTO shippers (.+) RETURNING id$"
	mockRows := sqlmock.NewRows([]string{"id", "access_key", "name", "machine_name", "created_at", "updated_at"}).AddRow(10, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", "test", "machine.test", 1517161676, 1517161676)
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	id, _, err := shipperDatastore.Add("test", "machine.test")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), id)
}

func TestShippersDatastoreDelete(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^DELETE FROM shippers"
	mockRows := sqlmock.NewResult(0, 0)
	mock.ExpectExec(mockQuery).WillReturnResult(mockRows)

	err := shipperDatastore.Delete(10)
	assert.Nil(t, err)
}

func TestShippersDatastoreView(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^SELECT \\* FROM shippers WHERE (.+)$"
	mockRows := sqlmock.NewRows([]string{"id", "access_key", "name", "machine_name", "created_at", "updated_at"}).AddRow(10, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", "test1", "machine.test1", 1517161676, 1517161676)
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	shipper, err := shipperDatastore.View(10)
	assert.Nil(t, err)
	assert.Equal(t, "test1", shipper.Name)
	assert.Equal(t, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", shipper.AccessKey)
}

func TestShippersDatastoreViewAll(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^SELECT \\* FROM shippers LIMIT 0, 100$"
	mockRows := sqlmock.NewRows([]string{"id", "access_key", "name", "machine_name", "created_at", "updated_at"}).AddRow(10, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", "test1", "machine.test1", 1517161676, 1517161676).AddRow(11, "8b0047c1-9e6a-46fb-9664-75ac60c3596b", "test2", "machine.test2", 1517161676, 1517161676).AddRow(12, "8b0047c1-9e6a-46fb-9664-75ac60c3596c", "test3", "machine.test3", 1517161676, 1517161676)
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	shippers, err := shipperDatastore.ViewAll()
	assert.Nil(t, err)
	assert.Equal(t, "test1", shippers[0].Name)
	assert.Equal(t, "8b0047c1-9e6a-46fb-9664-75ac60c3596b", shippers[1].AccessKey)
}
