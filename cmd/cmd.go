package cmd

import (
	"fmt"
	"github.com/danielmalka/dice-go/roller"
	"github.com/urfave/cli"
)

// Start inicia a aplicação em linha de comando
func Start() *cli.App {
	app := cli.NewApp()

	app.Name = "Dice Go"
	app.Usage = "Rolador de dados via linha de comando"
	app.Commands = []*cli.Command{
		{
			Name:  "dice",
			Usage: "Executa a rolagem dos dados",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "roll",
					Value: "1d6",
				},
			},
			Action: ExecuteRoller,
		},
	}

	return app
}

// ExecuteRoller chama o roller para processar a rolagem
func ExecuteRoller(c *cli.Context) error {
	req := c.String("roll")
	r := roller.RollDice(req)
	fmt.Println(r)
	return nil
}
