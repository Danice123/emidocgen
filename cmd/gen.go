package cmd

import (
	"emidocgen/package/ckp"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

func overFolderFiles(path string, f func(name string, path string) error) error {
	folders, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		err = f(folder.Name(), fmt.Sprintf("%s/%s", path, folder.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

var genCmd = &cobra.Command{
	Use:   "gen [data directory]",
	Short: "Read in a data directory and output a Emi Calc file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var data ckp.EmiCalcData
		err := overFolderFiles(args[0], func(name string, path string) error {
			switch name {
			case "pokemon":
				return overFolderFiles(path, func(name, path string) error {
					return overFolderFiles(path, func(name, path string) error {
						in, err := os.ReadFile(path)
						if err != nil {
							return err
						}
						var p ckp.Pokemon
						err = yaml.Unmarshal(in, &p)
						if err != nil {
							return err
						}
						data.Pokemon = append(data.Pokemon, p)
						return nil
					})
				})
			case "move":
				return overFolderFiles(path, func(name, path string) error {
					return overFolderFiles(path, func(name, path string) error {
						in, err := os.ReadFile(path)
						if err != nil {
							return err
						}
						var m ckp.PokeMove
						err = yaml.Unmarshal(in, &m)
						if err != nil {
							return err
						}
						data.Moves = append(data.Moves, m)
						return nil
					})
				})
			case "landmark":
				return overFolderFiles(path, func(name, path string) error {
					in, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					var l ckp.Landmark
					err = yaml.Unmarshal(in, &l)
					if err != nil {
						return err
					}
					data.Landmarks = append(data.Landmarks, l)
					return nil
				})
			case "trainer":
				return overFolderFiles(path, func(name, path string) error {
					return overFolderFiles(path, func(name, path string) error {
						in, err := os.ReadFile(path)
						if err != nil {
							return err
						}
						var t ckp.Trainer
						err = yaml.Unmarshal(in, &t)
						if err != nil {
							return err
						}
						data.Trainers = append(data.Trainers, t)
						return nil
					})
				})
			case "encounter_pools.yml":
				in, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var d map[string]interface{}
				err = yaml.Unmarshal(in, &d)
				if err != nil {
					return err
				}
				data.Pools = d["encounter_pools"]
			case "encounters.yml":
				in, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var d map[string][]interface{}
				err = yaml.Unmarshal(in, &d)
				if err != nil {
					return err
				}
				data.Encounters = d["encounters"]
			case "items.yml":
				in, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var d map[string][]ckp.Item
				err = yaml.Unmarshal(in, &d)
				if err != nil {
					return err
				}
				data.Items = d["items"]
			case "typeMatchups.yml":
				in, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var tm map[string][]ckp.TypeMatchup
				err = yaml.Unmarshal(in, &tm)
				if err != nil {
					return err
				}
				data.TypeMatchups = tm["matchups"]

			}
			return nil
		})
		if err != nil {
			return err
		}

		sort.Slice(data.Pokemon, func(i, j int) bool {
			return data.Pokemon[i].Pokedex < data.Pokemon[j].Pokedex
		})

		sort.Slice(data.Trainers, func(i, j int) bool {
			return data.Trainers[i].Order < data.Trainers[j].Order
		})

		o, err := json.Marshal(data)
		if err != nil {
			return err
		}

		var fm map[string]interface{}
		err = json.Unmarshal(o, &fm)
		if err != nil {
			return err
		}
		o, err = json.MarshalIndent(fm, "", "    ")
		if err != nil {
			return err
		}
		return os.WriteFile("output.json", o, 0777)
	},
}
