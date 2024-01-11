/*
 * Copyright (C) 2022 by Teradata Corporation.
 * All Rights Reserved.
 * TERADATA CORPORATION CONFIDENTIAL AND TRADE SECRET
 */

package teradata

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func DialectOptions() *goqu.SQLDialectOptions {
	opts := goqu.DefaultDialectOptions()

	// Set any Teradata specific SQL dialect options here
	// See an example dialect at https://github.com/doug-martin/goqu/blob/master/dialect/sqlserver/sqlserver.go
	opts.QuoteRune = '"'
	opts.BooleanOperatorLookup[exp.NeqOp] = []byte("<>")

	return opts
}

func init() {
	opts := DialectOptions()
	goqu.SetDefaultPrepared(true)
	goqu.RegisterDialect("teradata", opts)
}
