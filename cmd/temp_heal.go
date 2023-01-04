package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Shikachuu/health-tracker/pkg/model"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

func NewTemporaryHealCommand() *cobra.Command {
	var boltDB *bolt.DB

	return &cobra.Command{
		Use:     "temporary-heal",
		Short:   "Add temporary health to a character",
		Aliases: []string{"th", "temp-heal"},
		Args:    cobra.ExactArgs(2),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			var err error
			boltDB, err = bolt.Open("./health.db", 0600, &bolt.Options{})
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			tHP, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			return boltDB.Update(func(tx *bolt.Tx) error {
				var character model.Character

				bucket := tx.Bucket([]byte("characters"))

				b := bucket.Get([]byte(args[0]))
				json.Unmarshal(b, &character)

				fmt.Println(character.TemporaryHeal(tHP))

				encChar, err := json.Marshal(character)
				if err != nil {
					return err
				}

				return bucket.Put([]byte(args[0]), encChar)
			})
		},
		PostRunE: func(_ *cobra.Command, _ []string) error {
			return boltDB.Close()
		},
	}
}
