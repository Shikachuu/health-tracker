package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Shikachuu/health-tracker/pkg/model"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

func NewGetMaxCommand() *cobra.Command {
	var boltDB *bolt.DB

	return &cobra.Command{
		Use:     "get-max",
		Short:   "Get character's max HP",
		Aliases: []string{"gm"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			var err error
			boltDB, err = bolt.Open("./health.db", 0600, &bolt.Options{})
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			return boltDB.View(func(tx *bolt.Tx) error {
				var character model.Character

				bucket := tx.Bucket([]byte("characters"))

				b := bucket.Get([]byte(args[0]))
				json.Unmarshal(b, &character)
				fmt.Println(character.Health)
				return nil
			})
		},
		PostRunE: func(_ *cobra.Command, _ []string) error {
			return boltDB.Close()
		},
	}
}
