// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGuitarStrings(t *testing.T) {
	t.Parallel()

	query := GuitarStrings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGuitarStringsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
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

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarStringsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GuitarStrings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarStringsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuitarStringSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarStringsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GuitarStringExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if GuitarString exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GuitarStringExists to return true, but got false.")
	}
}

func testGuitarStringsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	guitarStringFound, err := FindGuitarString(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if guitarStringFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGuitarStringsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GuitarStrings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGuitarStringsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GuitarStrings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGuitarStringsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guitarStringOne := &GuitarString{}
	guitarStringTwo := &GuitarString{}
	if err = randomize.Struct(seed, guitarStringOne, guitarStringDBTypes, false, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}
	if err = randomize.Struct(seed, guitarStringTwo, guitarStringDBTypes, false, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guitarStringOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guitarStringTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GuitarStrings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGuitarStringsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	guitarStringOne := &GuitarString{}
	guitarStringTwo := &GuitarString{}
	if err = randomize.Struct(seed, guitarStringOne, guitarStringDBTypes, false, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}
	if err = randomize.Struct(seed, guitarStringTwo, guitarStringDBTypes, false, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guitarStringOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guitarStringTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func guitarStringBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func guitarStringAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GuitarString) error {
	*o = GuitarString{}
	return nil
}

func testGuitarStringsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &GuitarString{}
	o := &GuitarString{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, guitarStringDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GuitarString object: %s", err)
	}

	AddGuitarStringHook(boil.BeforeInsertHook, guitarStringBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	guitarStringBeforeInsertHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.AfterInsertHook, guitarStringAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	guitarStringAfterInsertHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.AfterSelectHook, guitarStringAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	guitarStringAfterSelectHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.BeforeUpdateHook, guitarStringBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	guitarStringBeforeUpdateHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.AfterUpdateHook, guitarStringAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	guitarStringAfterUpdateHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.BeforeDeleteHook, guitarStringBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	guitarStringBeforeDeleteHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.AfterDeleteHook, guitarStringAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	guitarStringAfterDeleteHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.BeforeUpsertHook, guitarStringBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	guitarStringBeforeUpsertHooks = []GuitarStringHook{}

	AddGuitarStringHook(boil.AfterUpsertHook, guitarStringAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	guitarStringAfterUpsertHooks = []GuitarStringHook{}
}

func testGuitarStringsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuitarStringsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(guitarStringColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuitarStringToManyStringStringExchangeLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GuitarString
	var b, c StringExchangeLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, stringExchangeLogDBTypes, false, stringExchangeLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stringExchangeLogDBTypes, false, stringExchangeLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.StringID = a.ID
	c.StringID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.StringStringExchangeLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.StringID == b.StringID {
			bFound = true
		}
		if v.StringID == c.StringID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GuitarStringSlice{&a}
	if err = a.L.LoadStringStringExchangeLogs(ctx, tx, false, (*[]*GuitarString)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StringStringExchangeLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.StringStringExchangeLogs = nil
	if err = a.L.LoadStringStringExchangeLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StringStringExchangeLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGuitarStringToManyAddOpStringStringExchangeLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GuitarString
	var b, c, d, e StringExchangeLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guitarStringDBTypes, false, strmangle.SetComplement(guitarStringPrimaryKeyColumns, guitarStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StringExchangeLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stringExchangeLogDBTypes, false, strmangle.SetComplement(stringExchangeLogPrimaryKeyColumns, stringExchangeLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*StringExchangeLog{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStringStringExchangeLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.StringID {
			t.Error("foreign key was wrong value", a.ID, first.StringID)
		}
		if a.ID != second.StringID {
			t.Error("foreign key was wrong value", a.ID, second.StringID)
		}

		if first.R.String != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.String != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.StringStringExchangeLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.StringStringExchangeLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.StringStringExchangeLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGuitarStringsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
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

func testGuitarStringsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuitarStringSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGuitarStringsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GuitarStrings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	guitarStringDBTypes = map[string]string{`ID`: `varchar`, `Name`: `varchar`, `Description`: `varchar`, `Maker`: `varchar`, `ThinGauge`: `tinyint`, `ThickGauge`: `tinyint`, `URL`: `varchar`, `IsDeleted`: `tinyint`, `Version`: `int`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                   = bytes.MinRead
)

func testGuitarStringsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(guitarStringPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(guitarStringAllColumns) == len(guitarStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGuitarStringsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(guitarStringAllColumns) == len(guitarStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GuitarString{}
	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guitarStringDBTypes, true, guitarStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(guitarStringAllColumns, guitarStringPrimaryKeyColumns) {
		fields = guitarStringAllColumns
	} else {
		fields = strmangle.SetComplement(
			guitarStringAllColumns,
			guitarStringPrimaryKeyColumns,
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

	slice := GuitarStringSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGuitarStringsUpsert(t *testing.T) {
	t.Parallel()

	if len(guitarStringAllColumns) == len(guitarStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLGuitarStringUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := GuitarString{}
	if err = randomize.Struct(seed, &o, guitarStringDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GuitarString: %s", err)
	}

	count, err := GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, guitarStringDBTypes, false, guitarStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GuitarString struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GuitarString: %s", err)
	}

	count, err = GuitarStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
