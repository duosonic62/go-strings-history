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

func testGuitars(t *testing.T) {
	t.Parallel()

	query := Guitars()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGuitarsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
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

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Guitars().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuitarSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuitarsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GuitarExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Guitar exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GuitarExists to return true, but got false.")
	}
}

func testGuitarsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	guitarFound, err := FindGuitar(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if guitarFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGuitarsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Guitars().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGuitarsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Guitars().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGuitarsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guitarOne := &Guitar{}
	guitarTwo := &Guitar{}
	if err = randomize.Struct(seed, guitarOne, guitarDBTypes, false, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}
	if err = randomize.Struct(seed, guitarTwo, guitarDBTypes, false, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guitarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guitarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Guitars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGuitarsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	guitarOne := &Guitar{}
	guitarTwo := &Guitar{}
	if err = randomize.Struct(seed, guitarOne, guitarDBTypes, false, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}
	if err = randomize.Struct(seed, guitarTwo, guitarDBTypes, false, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guitarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guitarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func guitarBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func guitarAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Guitar) error {
	*o = Guitar{}
	return nil
}

func testGuitarsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Guitar{}
	o := &Guitar{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, guitarDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Guitar object: %s", err)
	}

	AddGuitarHook(boil.BeforeInsertHook, guitarBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	guitarBeforeInsertHooks = []GuitarHook{}

	AddGuitarHook(boil.AfterInsertHook, guitarAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	guitarAfterInsertHooks = []GuitarHook{}

	AddGuitarHook(boil.AfterSelectHook, guitarAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	guitarAfterSelectHooks = []GuitarHook{}

	AddGuitarHook(boil.BeforeUpdateHook, guitarBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	guitarBeforeUpdateHooks = []GuitarHook{}

	AddGuitarHook(boil.AfterUpdateHook, guitarAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	guitarAfterUpdateHooks = []GuitarHook{}

	AddGuitarHook(boil.BeforeDeleteHook, guitarBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	guitarBeforeDeleteHooks = []GuitarHook{}

	AddGuitarHook(boil.AfterDeleteHook, guitarAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	guitarAfterDeleteHooks = []GuitarHook{}

	AddGuitarHook(boil.BeforeUpsertHook, guitarBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	guitarBeforeUpsertHooks = []GuitarHook{}

	AddGuitarHook(boil.AfterUpsertHook, guitarAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	guitarAfterUpsertHooks = []GuitarHook{}
}

func testGuitarsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuitarsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(guitarColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuitarToManyStringExchangeLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Guitar
	var b, c StringExchangeLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
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

	b.GuitarID = a.ID
	c.GuitarID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.StringExchangeLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.GuitarID == b.GuitarID {
			bFound = true
		}
		if v.GuitarID == c.GuitarID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GuitarSlice{&a}
	if err = a.L.LoadStringExchangeLogs(ctx, tx, false, (*[]*Guitar)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StringExchangeLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.StringExchangeLogs = nil
	if err = a.L.LoadStringExchangeLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.StringExchangeLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGuitarToManyAddOpStringExchangeLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Guitar
	var b, c, d, e StringExchangeLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guitarDBTypes, false, strmangle.SetComplement(guitarPrimaryKeyColumns, guitarColumnsWithoutDefault)...); err != nil {
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
		err = a.AddStringExchangeLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuitarID {
			t.Error("foreign key was wrong value", a.ID, first.GuitarID)
		}
		if a.ID != second.GuitarID {
			t.Error("foreign key was wrong value", a.ID, second.GuitarID)
		}

		if first.R.Guitar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guitar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.StringExchangeLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.StringExchangeLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.StringExchangeLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testGuitarToOneMemberUsingMember(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Guitar
	var foreign Member

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, guitarDBTypes, false, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, memberDBTypes, false, memberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Member struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MemberID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Member().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := GuitarSlice{&local}
	if err = local.L.LoadMember(ctx, tx, false, (*[]*Guitar)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Member = nil
	if err = local.L.LoadMember(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGuitarToOneSetOpMemberUsingMember(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Guitar
	var b, c Member

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guitarDBTypes, false, strmangle.SetComplement(guitarPrimaryKeyColumns, guitarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, memberDBTypes, false, strmangle.SetComplement(memberPrimaryKeyColumns, memberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Member{&b, &c} {
		err = a.SetMember(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Member != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Guitars[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MemberID != x.ID {
			t.Error("foreign key was wrong value", a.MemberID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MemberID))
		reflect.Indirect(reflect.ValueOf(&a.MemberID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MemberID != x.ID {
			t.Error("foreign key was wrong value", a.MemberID, x.ID)
		}
	}
}

func testGuitarsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
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

func testGuitarsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuitarSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGuitarsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Guitars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	guitarDBTypes = map[string]string{`ID`: `varchar`, `MemberID`: `varchar`, `Name`: `varchar`, `Description`: `varchar`, `BodyType`: `varchar`, `Maker`: `varchar`, `ImageURL`: `varchar`, `IsDeleted`: `tinyint`, `Version`: `int`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_             = bytes.MinRead
)

func testGuitarsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(guitarPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(guitarAllColumns) == len(guitarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGuitarsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(guitarAllColumns) == len(guitarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Guitar{}
	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guitarDBTypes, true, guitarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(guitarAllColumns, guitarPrimaryKeyColumns) {
		fields = guitarAllColumns
	} else {
		fields = strmangle.SetComplement(
			guitarAllColumns,
			guitarPrimaryKeyColumns,
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

	slice := GuitarSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGuitarsUpsert(t *testing.T) {
	t.Parallel()

	if len(guitarAllColumns) == len(guitarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLGuitarUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Guitar{}
	if err = randomize.Struct(seed, &o, guitarDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Guitar: %s", err)
	}

	count, err := Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, guitarDBTypes, false, guitarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guitar struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Guitar: %s", err)
	}

	count, err = Guitars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
