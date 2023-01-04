package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/Shikachuu/health-tracker/pkg/model"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

func NewCharacterCommand() *cobra.Command {
	var boltDB *bolt.DB

	return &cobra.Command{
		Use:     "new-character",
		Short:   "Generate new character",
		Aliases: []string{"nc", "new", "n"},
		Args:    cobra.ExactArgs(2),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			var err error
			boltDB, err = bolt.Open("./health.db", 0600, &bolt.Options{})
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			health, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			encChar, err := json.Marshal(model.Character{
				Health:          health,
				TemporaryHealth: 0,
				HPLog:           []int{},
			})
			if err != nil {
				return err
			}

			return boltDB.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("characters"))
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
