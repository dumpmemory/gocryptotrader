// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
)

// Datahistoryjobresult is an object representing the database table.
type Datahistoryjobresult struct {
	ID                string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	JobID             string      `boil:"job_id" json:"job_id" toml:"job_id" yaml:"job_id"`
	Result            null.String `boil:"result" json:"result,omitempty" toml:"result" yaml:"result,omitempty"`
	Status            float64     `boil:"status" json:"status" toml:"status" yaml:"status"`
	IntervalStartTime string      `boil:"interval_start_time" json:"interval_start_time" toml:"interval_start_time" yaml:"interval_start_time"`
	IntervalEndTime   string      `boil:"interval_end_time" json:"interval_end_time" toml:"interval_end_time" yaml:"interval_end_time"`
	RunTime           string      `boil:"run_time" json:"run_time" toml:"run_time" yaml:"run_time"`

	R *datahistoryjobresultR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L datahistoryjobresultL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DatahistoryjobresultColumns = struct {
	ID                string
	JobID             string
	Result            string
	Status            string
	IntervalStartTime string
	IntervalEndTime   string
	RunTime           string
}{
	ID:                "id",
	JobID:             "job_id",
	Result:            "result",
	Status:            "status",
	IntervalStartTime: "interval_start_time",
	IntervalEndTime:   "interval_end_time",
	RunTime:           "run_time",
}

// Generated where

var DatahistoryjobresultWhere = struct {
	ID                whereHelperstring
	JobID             whereHelperstring
	Result            whereHelpernull_String
	Status            whereHelperfloat64
	IntervalStartTime whereHelperstring
	IntervalEndTime   whereHelperstring
	RunTime           whereHelperstring
}{
	ID:                whereHelperstring{field: "\"datahistoryjobresult\".\"id\""},
	JobID:             whereHelperstring{field: "\"datahistoryjobresult\".\"job_id\""},
	Result:            whereHelpernull_String{field: "\"datahistoryjobresult\".\"result\""},
	Status:            whereHelperfloat64{field: "\"datahistoryjobresult\".\"status\""},
	IntervalStartTime: whereHelperstring{field: "\"datahistoryjobresult\".\"interval_start_time\""},
	IntervalEndTime:   whereHelperstring{field: "\"datahistoryjobresult\".\"interval_end_time\""},
	RunTime:           whereHelperstring{field: "\"datahistoryjobresult\".\"run_time\""},
}

// DatahistoryjobresultRels is where relationship names are stored.
var DatahistoryjobresultRels = struct {
	Job string
}{
	Job: "Job",
}

// datahistoryjobresultR is where relationships are stored.
type datahistoryjobresultR struct {
	Job *Datahistoryjob
}

// NewStruct creates a new relationship struct
func (*datahistoryjobresultR) NewStruct() *datahistoryjobresultR {
	return &datahistoryjobresultR{}
}

// datahistoryjobresultL is where Load methods for each relationship are stored.
type datahistoryjobresultL struct{}

var (
	datahistoryjobresultAllColumns            = []string{"id", "job_id", "result", "status", "interval_start_time", "interval_end_time", "run_time"}
	datahistoryjobresultColumnsWithoutDefault = []string{"id", "job_id", "result", "status", "interval_start_time", "interval_end_time"}
	datahistoryjobresultColumnsWithDefault    = []string{"run_time"}
	datahistoryjobresultPrimaryKeyColumns     = []string{"id"}
)

type (
	// DatahistoryjobresultSlice is an alias for a slice of pointers to Datahistoryjobresult.
	// This should generally be used opposed to []Datahistoryjobresult.
	DatahistoryjobresultSlice []*Datahistoryjobresult
	// DatahistoryjobresultHook is the signature for custom Datahistoryjobresult hook methods
	DatahistoryjobresultHook func(context.Context, boil.ContextExecutor, *Datahistoryjobresult) error

	datahistoryjobresultQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	datahistoryjobresultType                 = reflect.TypeOf(&Datahistoryjobresult{})
	datahistoryjobresultMapping              = queries.MakeStructMapping(datahistoryjobresultType)
	datahistoryjobresultPrimaryKeyMapping, _ = queries.BindMapping(datahistoryjobresultType, datahistoryjobresultMapping, datahistoryjobresultPrimaryKeyColumns)
	datahistoryjobresultInsertCacheMut       sync.RWMutex
	datahistoryjobresultInsertCache          = make(map[string]insertCache)
	datahistoryjobresultUpdateCacheMut       sync.RWMutex
	datahistoryjobresultUpdateCache          = make(map[string]updateCache)
	datahistoryjobresultUpsertCacheMut       sync.RWMutex
	datahistoryjobresultUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var datahistoryjobresultBeforeInsertHooks []DatahistoryjobresultHook
var datahistoryjobresultBeforeUpdateHooks []DatahistoryjobresultHook
var datahistoryjobresultBeforeDeleteHooks []DatahistoryjobresultHook
var datahistoryjobresultBeforeUpsertHooks []DatahistoryjobresultHook

var datahistoryjobresultAfterInsertHooks []DatahistoryjobresultHook
var datahistoryjobresultAfterSelectHooks []DatahistoryjobresultHook
var datahistoryjobresultAfterUpdateHooks []DatahistoryjobresultHook
var datahistoryjobresultAfterDeleteHooks []DatahistoryjobresultHook
var datahistoryjobresultAfterUpsertHooks []DatahistoryjobresultHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Datahistoryjobresult) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Datahistoryjobresult) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Datahistoryjobresult) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Datahistoryjobresult) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Datahistoryjobresult) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Datahistoryjobresult) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Datahistoryjobresult) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Datahistoryjobresult) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Datahistoryjobresult) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range datahistoryjobresultAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDatahistoryjobresultHook registers your hook function for all future operations.
func AddDatahistoryjobresultHook(hookPoint boil.HookPoint, datahistoryjobresultHook DatahistoryjobresultHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		datahistoryjobresultBeforeInsertHooks = append(datahistoryjobresultBeforeInsertHooks, datahistoryjobresultHook)
	case boil.BeforeUpdateHook:
		datahistoryjobresultBeforeUpdateHooks = append(datahistoryjobresultBeforeUpdateHooks, datahistoryjobresultHook)
	case boil.BeforeDeleteHook:
		datahistoryjobresultBeforeDeleteHooks = append(datahistoryjobresultBeforeDeleteHooks, datahistoryjobresultHook)
	case boil.BeforeUpsertHook:
		datahistoryjobresultBeforeUpsertHooks = append(datahistoryjobresultBeforeUpsertHooks, datahistoryjobresultHook)
	case boil.AfterInsertHook:
		datahistoryjobresultAfterInsertHooks = append(datahistoryjobresultAfterInsertHooks, datahistoryjobresultHook)
	case boil.AfterSelectHook:
		datahistoryjobresultAfterSelectHooks = append(datahistoryjobresultAfterSelectHooks, datahistoryjobresultHook)
	case boil.AfterUpdateHook:
		datahistoryjobresultAfterUpdateHooks = append(datahistoryjobresultAfterUpdateHooks, datahistoryjobresultHook)
	case boil.AfterDeleteHook:
		datahistoryjobresultAfterDeleteHooks = append(datahistoryjobresultAfterDeleteHooks, datahistoryjobresultHook)
	case boil.AfterUpsertHook:
		datahistoryjobresultAfterUpsertHooks = append(datahistoryjobresultAfterUpsertHooks, datahistoryjobresultHook)
	}
}

// One returns a single datahistoryjobresult record from the query.
func (q datahistoryjobresultQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Datahistoryjobresult, error) {
	o := &Datahistoryjobresult{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: failed to execute a one query for datahistoryjobresult")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Datahistoryjobresult records from the query.
func (q datahistoryjobresultQuery) All(ctx context.Context, exec boil.ContextExecutor) (DatahistoryjobresultSlice, error) {
	var o []*Datahistoryjobresult

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlite3: failed to assign all query results to Datahistoryjobresult slice")
	}

	if len(datahistoryjobresultAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Datahistoryjobresult records in the query.
func (q datahistoryjobresultQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to count datahistoryjobresult rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q datahistoryjobresultQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: failed to check if datahistoryjobresult exists")
	}

	return count > 0, nil
}

// Job pointed to by the foreign key.
func (o *Datahistoryjobresult) Job(mods ...qm.QueryMod) datahistoryjobQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.JobID),
	}

	queryMods = append(queryMods, mods...)

	query := Datahistoryjobs(queryMods...)
	queries.SetFrom(query.Query, "\"datahistoryjob\"")

	return query
}

// LoadJob allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (datahistoryjobresultL) LoadJob(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDatahistoryjobresult interface{}, mods queries.Applicator) error {
	var slice []*Datahistoryjobresult
	var object *Datahistoryjobresult

	if singular {
		object = maybeDatahistoryjobresult.(*Datahistoryjobresult)
	} else {
		slice = *maybeDatahistoryjobresult.(*[]*Datahistoryjobresult)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &datahistoryjobresultR{}
		}
		args = append(args, object.JobID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &datahistoryjobresultR{}
			}

			for _, a := range args {
				if a == obj.JobID {
					continue Outer
				}
			}

			args = append(args, obj.JobID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`datahistoryjob`), qm.WhereIn(`datahistoryjob.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Datahistoryjob")
	}

	var resultSlice []*Datahistoryjob
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Datahistoryjob")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for datahistoryjob")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for datahistoryjob")
	}

	if len(datahistoryjobresultAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Job = foreign
		if foreign.R == nil {
			foreign.R = &datahistoryjobR{}
		}
		foreign.R.JobDatahistoryjobresults = append(foreign.R.JobDatahistoryjobresults, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.JobID == foreign.ID {
				local.R.Job = foreign
				if foreign.R == nil {
					foreign.R = &datahistoryjobR{}
				}
				foreign.R.JobDatahistoryjobresults = append(foreign.R.JobDatahistoryjobresults, local)
				break
			}
		}
	}

	return nil
}

// SetJob of the datahistoryjobresult to the related item.
// Sets o.R.Job to related.
// Adds o to related.R.JobDatahistoryjobresults.
func (o *Datahistoryjobresult) SetJob(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Datahistoryjob) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"datahistoryjobresult\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"job_id"}),
		strmangle.WhereClause("\"", "\"", 0, datahistoryjobresultPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.JobID = related.ID
	if o.R == nil {
		o.R = &datahistoryjobresultR{
			Job: related,
		}
	} else {
		o.R.Job = related
	}

	if related.R == nil {
		related.R = &datahistoryjobR{
			JobDatahistoryjobresults: DatahistoryjobresultSlice{o},
		}
	} else {
		related.R.JobDatahistoryjobresults = append(related.R.JobDatahistoryjobresults, o)
	}

	return nil
}

// Datahistoryjobresults retrieves all the records using an executor.
func Datahistoryjobresults(mods ...qm.QueryMod) datahistoryjobresultQuery {
	mods = append(mods, qm.From("\"datahistoryjobresult\""))
	return datahistoryjobresultQuery{NewQuery(mods...)}
}

// FindDatahistoryjobresult retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDatahistoryjobresult(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Datahistoryjobresult, error) {
	datahistoryjobresultObj := &Datahistoryjobresult{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"datahistoryjobresult\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, datahistoryjobresultObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: unable to select from datahistoryjobresult")
	}

	return datahistoryjobresultObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Datahistoryjobresult) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlite3: no datahistoryjobresult provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(datahistoryjobresultColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	datahistoryjobresultInsertCacheMut.RLock()
	cache, cached := datahistoryjobresultInsertCache[key]
	datahistoryjobresultInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			datahistoryjobresultAllColumns,
			datahistoryjobresultColumnsWithDefault,
			datahistoryjobresultColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(datahistoryjobresultType, datahistoryjobresultMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(datahistoryjobresultType, datahistoryjobresultMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"datahistoryjobresult\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"datahistoryjobresult\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"datahistoryjobresult\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, datahistoryjobresultPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to insert into datahistoryjobresult")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to populate default values for datahistoryjobresult")
	}

CacheNoHooks:
	if !cached {
		datahistoryjobresultInsertCacheMut.Lock()
		datahistoryjobresultInsertCache[key] = cache
		datahistoryjobresultInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Datahistoryjobresult.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Datahistoryjobresult) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	datahistoryjobresultUpdateCacheMut.RLock()
	cache, cached := datahistoryjobresultUpdateCache[key]
	datahistoryjobresultUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			datahistoryjobresultAllColumns,
			datahistoryjobresultPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("sqlite3: unable to update datahistoryjobresult, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"datahistoryjobresult\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, datahistoryjobresultPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(datahistoryjobresultType, datahistoryjobresultMapping, append(wl, datahistoryjobresultPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update datahistoryjobresult row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by update for datahistoryjobresult")
	}

	if !cached {
		datahistoryjobresultUpdateCacheMut.Lock()
		datahistoryjobresultUpdateCache[key] = cache
		datahistoryjobresultUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q datahistoryjobresultQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all for datahistoryjobresult")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected for datahistoryjobresult")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DatahistoryjobresultSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlite3: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datahistoryjobresultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"datahistoryjobresult\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, datahistoryjobresultPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all in datahistoryjobresult slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected all in update all datahistoryjobresult")
	}
	return rowsAff, nil
}

// Delete deletes a single Datahistoryjobresult record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Datahistoryjobresult) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlite3: no Datahistoryjobresult provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), datahistoryjobresultPrimaryKeyMapping)
	sql := "DELETE FROM \"datahistoryjobresult\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete from datahistoryjobresult")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by delete for datahistoryjobresult")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q datahistoryjobresultQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlite3: no datahistoryjobresultQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from datahistoryjobresult")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for datahistoryjobresult")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DatahistoryjobresultSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(datahistoryjobresultBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datahistoryjobresultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"datahistoryjobresult\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, datahistoryjobresultPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from datahistoryjobresult slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for datahistoryjobresult")
	}

	if len(datahistoryjobresultAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Datahistoryjobresult) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDatahistoryjobresult(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DatahistoryjobresultSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DatahistoryjobresultSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), datahistoryjobresultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"datahistoryjobresult\".* FROM \"datahistoryjobresult\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, datahistoryjobresultPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to reload all in DatahistoryjobresultSlice")
	}

	*o = slice

	return nil
}

// DatahistoryjobresultExists checks if the Datahistoryjobresult row exists.
func DatahistoryjobresultExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"datahistoryjobresult\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: unable to check if datahistoryjobresult exists")
	}

	return exists, nil
}