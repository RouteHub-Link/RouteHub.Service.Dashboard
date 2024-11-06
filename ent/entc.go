// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
	// This option is disabled in this example,
	// because the schema file is edited by the
	// internal/todo/ent example.
	//
	// entgql.WithSchemaPath("../todo.graphql"),
	//
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2024
		`,
	}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
