/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/services"
)

type flagsViewResource struct {
}

var viewResourceFlags = flagsViewResource{}

var ViewCmd = &command.Command{
	Cmd: &cobra.Command{
		Use:     "view",
		Short:   "View a list of resource entities in the configuration / View the properties of a particular entity",
		Example: "flow config view account \nflow config view account <accountname>\nflow config view contract \nflow config view contract <contractname>\nflow config view deployment \nflow config view deployment <networkname>\nflow config view emulator \nflow config view emulator <emulatorname>\nflow config view network \nflow config view network <networkname>",
		Args:    cobra.MaximumNArgs(2),
	},
	Flags: &viewResourceFlags,
	RunS:  viewResource,
}

func viewResource(args []string,
	_ flowkit.ReaderWriter,
	globalFlags command.GlobalFlags,
	_ *services.Services,
	state *flowkit.State,
) (command.Result, error) {

	// If the user does not input any argument, then we print a message with usage example
	if len(args) == 0 {
		return &Result{
			result: "Please enter resource or resource + field arguments \nUsage examples: \n1) flow config view account --> shows the list of all accounts\n2) flow config view account <accountname> --> shows the properties of <accountname>",
		}, nil
	}

	// If the argument passed is "account".
	if args[0] == "account" {

		// Count variable for printing format.
		var count int = 0

		// IF CONDITION:
		// If there is one argument in the command i.e. command looks like --> flow config view account,
		// Then we print the list of all the present accounts in the configuration.
		// ELSE IF CONDITION:
		// If there are arguments == 2 i.e. command looks like --> flow config view account <accountname>,
		// Then we print all the details of the account "<accountname>".
		// 	If the <accountname> doesn't exist in the configuration, then we print "Account <accountname> does not exist".

		// Declaring a string variable to print our final list.
		var list string = ""

		if len(args) == 1 {
			for _, value := range *state.Accounts() {
				if count == 0 {
					list = list + string(value.Name())
					count = count + 1
				} else if count > 0 {
					list = list + ", " + string(value.Name())
				}
			}
			// Return the list of accounts.
			return &Result{
				result: fmt.Sprintf("List of Accounts: %s", list),
			}, nil
		} else if len(args) == 2 {
			for _, value := range state.Config().Accounts {
				if value.Name == args[1] {
					// Return the properties of a particular account.
					return &Result{
						result: fmt.Sprintf("Account Name: %s \nAddress: %s \nKey Properties: \nType: %s, Index: %d, Signature Algorithm: %s, Hash Algorithm: %s, Private Key: %s", value.Name, value.Address, value.Key.Type, value.Key.Index, value.Key.SigAlgo, value.Key.HashAlgo, value.Key.PrivateKey),
					}, nil
				}
			}
			return &Result{
				result: fmt.Sprintf("Account %s does not exist", args[1]),
			}, nil
		}
	}

	// If the argument passed is "network".
	if args[0] == "network" {

		// Count variable for printing format.
		var count int = 0

		// IF CONDITION:
		// If there is one argument in the command i.e. command looks like --> flow config view network,
		// Then we print the list of all the present accounts in the configuration.
		// ELSE IF CONDITION:
		// If there are arguments == 2 i.e. command looks like --> flow config view network <networkname>,
		// Then we print all the details of the network "<networkname>".
		// 	If the <networkname> doesn't exist in the configuration, then we print "Network <networkname> does not exist".

		// Declaring a string variable to print our final list.
		var list string = ""

		if len(args) == 1 {
			for _, value := range *state.Networks() {
				if count == 0 {
					list = list + string(value.Name)
					count = count + 1
				} else if count > 0 {
					list = list + ", " + string(value.Name)
				}
			}
			// Return the list of networks.
			return &Result{
				result: fmt.Sprintf("List of Networks: %s", list),
			}, nil
		} else if len(args) == 2 {
			for _, value := range state.Config().Networks {
				if value.Name == args[1] {
					// Return the properties of a particular network.
					return &Result{
						result: fmt.Sprintf("Network Name: %s \nHost: %s \n", value.Name, value.Host),
					}, nil
				}
			}
			return &Result{
				result: fmt.Sprintf("Network %s does not exist", args[1]),
			}, nil
		}
	}

	// If the argument passed is "contract".
	if args[0] == "contract" {

		// Count variable for printing format.
		var count int = 0

		// IF CONDITION:
		// If there is one argument in the command i.e. command looks like --> flow config view contract,
		// Then we print the list of all the present accounts in the configuration.
		// ELSE IF CONDITION:
		// If there are arguments == 2 i.e. command looks like --> flow config view contract <contractname>,
		// Then we print all the details of the contract "<contractname>".
		// 	If the <contractname> doesn't exist in the configuration, then we print "Contract <contractname> does not exist".

		// Declaring a string variable to print our final list.
		var list string = ""

		if len(args) == 1 {
			for _, value := range *state.Contracts() {
				if count == 0 {
					list = list + string(value.Name)
					count = count + 1
				} else if count > 0 {
					list = list + ", " + string(value.Name)
				}
			}
			// Return the list of contracts.
			return &Result{
				result: fmt.Sprintf("List of Contracts: %s", list),
			}, nil
		} else if len(args) == 2 {
			for _, value := range state.Config().Contracts {
				if value.Name == args[1] {
					// Return the properties of a particular contract.
					return &Result{
						result: fmt.Sprintf("Contract Name: %s \nSource: %s \nNetwork: %s \nAlias: %s", value.Name, value.Source, value.Network, value.Alias),
					}, nil
				}
			}
			return &Result{
				result: fmt.Sprintf("Contract %s does not exist", args[1]),
			}, nil
		}
	}

	// If the argument passed is "emulator".
	if args[0] == "emulator" {

		// Count variable for printing format.
		var count int = 0

		// IF CONDITION:
		// If there is one argument in the command i.e. command looks like --> flow config view emulator,
		// Then we print the list of all the present emulators in the configuration.
		// ELSE IF CONDITION:
		// If there are arguments == 2 i.e. command looks like --> flow config view emulator <emulatorname>,
		// Then we print all the details of the emulator "<emulatorname>".
		// 	If the <emulatorname> doesn't exist in the configuration, then we print "Emulator <emulatorname> does not exist".

		// Declaring a string variable to print our final list.
		var list string = ""

		if len(args) == 1 {
			for _, value := range state.Config().Emulators {
				if count == 0 {
					list = list + string(value.Name)
					count = count + 1
				} else if count > 0 {
					list = list + ", " + string(value.Name)
				}
			}
			// Return the list of emulators.
			return &Result{
				result: fmt.Sprintf("List of Emulators: %s", list),
			}, nil
		} else if len(args) == 2 {
			for _, value := range state.Config().Emulators {
				if value.Name == args[1] {
					// Return the properties of a particular emulator.
					return &Result{
						result: fmt.Sprintf("Emulator Name: %s \nPort: %d \nService Account: %s", value.Name, value.Port, value.ServiceAccount),
					}, nil
				}
			}
			return &Result{
				result: fmt.Sprintf("Emulator %s does not exist", args[1]),
			}, nil
		}
	}

	// If the argument passed is "deployment".
	if args[0] == "deployment" {

		// Count variable for printing format.
		var count int = 0

		// IF CONDITION:
		// If there is one argument in the command i.e. command looks like --> flow config view deployment,
		// Then we print the list of all the deployed networks in the configuration.
		// ELSE IF CONDITION:
		// If there are arguments == 2 i.e. command looks like --> flow config view deployment <networkname>,
		// Then we print all the details of the deployed network "<networkname>".
		// 	If the <networkname> isn't deployed in the configuration, then we print "Network <networkname> is not deployed".

		// Declaring a string variable to print our final list.
		var list string = ""

		if len(args) == 1 {
			for _, value := range *state.Deployments() {
				if count == 0 {
					list = list + string(value.Network)
					count = count + 1
				} else if count > 0 {
					list = list + ", " + string(value.Network)
				}
			}
			// Return the list of emulators.
			return &Result{
				result: fmt.Sprintf("List of Networks Deployed: %s", list),
			}, nil
		} else if len(args) == 2 {
			for _, value := range state.Config().Deployments {
				if value.Network == args[1] {
					// Return the properties of a particular deployed network.
					return &Result{
						result: fmt.Sprintf("Network Name: %s \nAccount: %s \nContracts with Cadence Value: %s", value.Network, value.Account, value.Contracts),
					}, nil
				}
			}
			return &Result{
				result: fmt.Sprintf("Network %s is not deployed", args[1]),
			}, nil
		}
	}

	// If user enters an invalid resource name, then we print a error response
	return &Result{
		result: "Invalid resource name given. \nValid resources: account, network, emulator, contract, deployment.",
	}, nil
}
