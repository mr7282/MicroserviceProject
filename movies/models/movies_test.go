// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMovies(t *testing.T) {
	t.Parallel()

	query := Movies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMoviesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMoviesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Movies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMoviesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MovieSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMoviesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MovieExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Movie exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MovieExists to return true, but got false.")
	}
}

func testMoviesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	movieFound, err := FindMovie(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if movieFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMoviesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Movies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMoviesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Movies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMoviesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	movieOne := &Movie{}
	movieTwo := &Movie{}
	if err = randomize.Struct(seed, movieOne, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}
	if err = randomize.Struct(seed, movieTwo, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = movieOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = movieTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Movies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMoviesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	movieOne := &Movie{}
	movieTwo := &Movie{}
	if err = randomize.Struct(seed, movieOne, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}
	if err = randomize.Struct(seed, movieTwo, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = movieOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = movieTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func movieBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func movieAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Movie) error {
	*o = Movie{}
	return nil
}

func testMoviesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Movie{}
	o := &Movie{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, movieDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Movie object: %s", err)
	}

	AddMovieHook(boil.BeforeInsertHook, movieBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	movieBeforeInsertHooks = []MovieHook{}

	AddMovieHook(boil.AfterInsertHook, movieAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	movieAfterInsertHooks = []MovieHook{}

	AddMovieHook(boil.AfterSelectHook, movieAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	movieAfterSelectHooks = []MovieHook{}

	AddMovieHook(boil.BeforeUpdateHook, movieBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	movieBeforeUpdateHooks = []MovieHook{}

	AddMovieHook(boil.AfterUpdateHook, movieAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	movieAfterUpdateHooks = []MovieHook{}

	AddMovieHook(boil.BeforeDeleteHook, movieBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	movieBeforeDeleteHooks = []MovieHook{}

	AddMovieHook(boil.AfterDeleteHook, movieAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	movieAfterDeleteHooks = []MovieHook{}

	AddMovieHook(boil.BeforeUpsertHook, movieBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	movieBeforeUpsertHooks = []MovieHook{}

	AddMovieHook(boil.AfterUpsertHook, movieAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	movieAfterUpsertHooks = []MovieHook{}
}

func testMoviesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMoviesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(movieColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMoviesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMoviesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MovieSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMoviesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Movies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	movieDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `ReleaseMovie`: `int`, `Genere`: `varchar`, `Country`: `varchar`, `Duration`: `varchar`, `PosterURL`: `text`, `MovieURL`: `text`, `AddedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_            = bytes.MinRead
)

func testMoviesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(moviePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(movieAllColumns) == len(moviePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, movieDBTypes, true, moviePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMoviesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(movieAllColumns) == len(moviePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Movie{}
	if err = randomize.Struct(seed, o, movieDBTypes, true, movieColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, movieDBTypes, true, moviePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(movieAllColumns, moviePrimaryKeyColumns) {
		fields = movieAllColumns
	} else {
		fields = strmangle.SetComplement(
			movieAllColumns,
			moviePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MovieSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMoviesUpsert(t *testing.T) {
	t.Parallel()

	if len(movieAllColumns) == len(moviePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMovieUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Movie{}
	if err = randomize.Struct(seed, &o, movieDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Movie: %s", err)
	}

	count, err := Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, movieDBTypes, false, moviePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Movie struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Movie: %s", err)
	}

	count, err = Movies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
