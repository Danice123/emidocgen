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
			err = writeFile(fmt.Sprintf("data/pokemon/%s/%s.yml", p.Types[0], p.Name), p)
			if err != nil {
				return err
			}
		}

		mm := map[ckp.PokeType][]ckp.PokeMove{}
		for i, m := range data.Moves {
			m.Order = i
			mm[m.Type] = append(mm[m.Type], m)
		}

		err = os.MkdirAll("data/move", 0777)
		if err != nil {
			return err
		}
		for t, ml := range mm {
			err = writeFile(fmt.Sprintf("data/move/%s.yml", t), map[string][]ckp.PokeMove{"moves": ml})
			if err != nil {
				return err
			}
		}

		err = writeFile("data/landmarks.yml", map[string][]ckp.Landmark{"landmarks": data.Landmarks})
		if err != nil {
			return err
		}

		for i, t := range data.Trainers {
			t.Order = i
			err = os.MkdirAll(fmt.Sprintf("data/trainer/%s", t.Area), 0777)
			if err != nil {
				return err
			}
			err = writeFile(fmt.Sprintf("data/trainer/%s/%s.yml", t.Area, t.Name), t)
			if err != nil {
				return err
			}
		}

		err = writeFile("data/encounter_pools.yml", map[string]interface{}{"encounter_pools": data.Pools})
		if err != nil {
			return err
		}

		err = os.MkdirAll("data/encounter", 0777)
		if err != nil {
			return err
		}
		for _, e := range data.Encounters {
			err = writeFile(fmt.Sprintf("data/encounter/%s.yml", e.Area), e)
			if err != nil {
				return err
			}
		}

		err = writeFile("data/items.yml", map[string]interface{}{"items": data.Items})
		if err != nil {
			return err
		}

		err = writeFile("data/typeMatchups.yml", map[string]interface{}{"matchups": data.TypeMatchups})
		if err != nil {
			return err
		}

		return nil
	},
}

func writeFile(file string, data interface{}) error {
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
