/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"github.com/spf13/cobra"
)

// Coll this struct is used to implement the mongo.Collection interface
type Coll struct {
	name string
}

func (c *Coll) Collection() string {
	return c.name
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is used to create indexes, etc.",
	Long:  `This command is used to create indexes, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		coll, err := cmd.Flags().GetString("coll")
		if err != nil || coll == "" {
			_ = cmd.Usage()
			return
		}
		ope, err := cmd.Flags().GetString("ope")
		if err != nil || ope == "" {
			_ = cmd.Usage()
			return
		}

		if ope != "index" {
			// only support index now
			return
		}
		// create index
		indexs, err := cmd.Flags().GetStringSlice("indexs")
		if err != nil || len(indexs) == 0 {
			_ = cmd.Usage()
			return
		}

		// Handling index rules
		var rules []mongo.SortRule
		for _, index := range indexs {
			strs := strings.Split(index, "_")
			if len(strs) != 2 {
				panic("The format for the fields to create an index is 'field_name_sorting_rule(1 or -1)'")
			}
			typ, err := strconv.Atoi(strs[1])
			if err != nil || (typ != mongo.SortTypeASC && typ != mongo.SortTypeDESC) {
				panic("The format for the fields to create an index is 'field_name_sorting_rule(1 or -1)'")
			}

			rules = append(rules, mongo.SortRule{Typ: mongo.SortType(typ), Field: strs[0]})
		}
		name, err := conn.CreateIndex(context.Background(), &Coll{coll}, rules...)
		if err != nil {
			panic(err)
		}
		fmt.Println("Index created successfully, index name:", name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("coll", "", "The name of the collection to be operated")
	createCmd.Flags().String("ope", "", "The name of the operation to be performed, now only index")
	createCmd.Flags().StringSlice("indexs", []string{}, "The format for the fields to create an index is 'field_name_sorting_rule(1 or -1)'")
}
