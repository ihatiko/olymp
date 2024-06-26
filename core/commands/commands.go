package commands

import (
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/ihatiko/olymp/components/clients/tech"
	"github.com/ihatiko/olymp/core/iface"
	_ "github.com/ihatiko/olymp/core/store"
	"github.com/ihatiko/olymp/core/utils"
	tC "github.com/ihatiko/olymp/infrastucture/components/utils/config"
	"github.com/spf13/cobra"
)

type deployment = func() (*cobra.Command, error)

func WithDeployment[Deployment iface.IDeployment]() deployment {
	return func() (*cobra.Command, error) {

		return &cobra.Command{
			Use: utils.ParseTypeName[Deployment](),
			Run: func(cmd *cobra.Command, args []string) {
				d := new(Deployment)

				defer func() {
					if r := recover(); r != nil {
						stack := string(debug.Stack())
						name := reflect.TypeOf(*d).String()
						fmt.Println(fmt.Sprintf("Recovered in core (Run) [%s] \n error: %s", name, stack))
					}
				}()
				err := tC.ToConfig(d)
				if err != nil {
					fmt.Println("Error in config:", err)
					return
				}
				commandName := utils.ParseTypeName[Deployment]()
				os.Setenv("TECH_SERVICE_COMMAND", commandName)
				app := (*d).Dep()
				rApp := reflect.ValueOf(app)
				var collectErrors []string
				for i := 0; i < rApp.NumField(); i++ {
					if rApp.Field(i).IsZero() {
						msg := fmt.Sprintf("empty field %s %s", rApp.Type().Field(i).Name, rApp.Type().Field(i).Type)
						collectErrors = append(collectErrors, msg)
					}
				}
				if len(collectErrors) != 0 {
					name := reflect.TypeOf(*d).String()
					fmt.Println(fmt.Sprintf("Error construct deployment [%s] with command %s", name, commandName))
					fmt.Println("-----------------------")
					fmt.Println(strings.Join(collectErrors, "\n"))
					fmt.Println("-----------------------")
					os.Exit(1)
				}
				app.Run()
			},
		}, nil
	}
}

func WithApp(operators ...deployment) {
	cmd := new(cobra.Command)
	var (
		err error
		c   *cobra.Command
	)
	for _, d := range operators {
		c, err = d()
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		cmd.AddCommand(c)
	}
	Compile(cmd, err)
}

func Compile(rootCommand *cobra.Command, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if os.Args[1] == "-test.v" {
			arg = os.Getenv("TEST_COMMAND")
		}
		rootCommand.SetArgs([]string{arg})
		err := tech.Use(arg)
		if err != nil {
			os.Exit(1)
		}
	}

	err = rootCommand.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
