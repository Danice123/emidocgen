package cmd

import (
	"emidocgen/package/ckp"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use:   "parse [file]",
	Short: "Read an Emi calc file and organize it into a editable format",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := ckp.ParseEmiData(args[0])
		if err != nil {
			return err
		}

		for _, p := range data.Pokemon {
			err = os.MkdirAll(fmt.Sprintf("data/pokemon/%s", p.Types[0]), 0777)
			if err != nil {
				return err
			}
			o, err := yaml.Marshal(p)
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("data/pokemon/%s/%s.yml", p.Types[0], p.Name), o, 0777)
			if err != nil {
				return err
			}
		}

		for _, m := range data.Moves {
			err = os.MkdirAll(fmt.Sprintf("data/move/%s", m.Type), 0777)
			if err != nil {
				return err
			}
			o, err := yaml.Marshal(m)
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("data/move/%s/%s.yml", m.Type, m.Name), o, 0777)
			if err != nil {
				return err
			}
		}

		// for _, l := range data.Landmarks {
		// 	err = os.MkdirAll("data/landmark", 0777)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	o, err := yaml.Marshal(l)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	err = os.WriteFile(fmt.Sprintf("data/landmark/%s.yml", l.Name), o, 0777)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		for i, t := range data.Trainers {
			t.Order = i
			err = os.MkdirAll(fmt.Sprintf("data/trainer/%s", t.Area), 0777)
			if err != nil {
				return err
			}
			o, err := yaml.Marshal(t)
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("data/trainer/%s/%s.yml", t.Area, t.Name), o, 0777)
			if err != nil {
				return err
			}
		}

		err = writeExtra("data/encounter_pools.yml", map[string]interface{}{"encounter_pools": data.Pools})
		if err != nil {
			return err
		}

		err = writeExtra("data/encounters.yml", map[string]interface{}{"encounters": data.Encounters})
		if err != nil {
			return err
		}

		err = writeExtra("data/items.yml", map[string]interface{}{"items": data.Items})
		if err != nil {
			return err
		}

		err = writeExtra("data/typeMatchups.yml", map[string]interface{}{"matchups": data.TypeMatchups})
		if err != nil {
			return err
		}

		return nil
	},
}

func writeExtra(file string, data map[string]interface{}) error {
	o, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, o, 0777)
	if err != nil {
		return err
	}
	return nil
}
