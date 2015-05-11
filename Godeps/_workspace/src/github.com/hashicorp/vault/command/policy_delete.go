package command

import (
	"fmt"
	"strings"
)

// PolicyDeleteCommand is a Command that enables a new endpoint.
type PolicyDeleteCommand struct {
	Meta
}

func (c *PolicyDeleteCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("policy-delete", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		flags.Usage()
		c.Ui.Error(fmt.Sprintf(
			"\npolicy-delete expects exactly one argument"))
		return 1
	}

	client, err := c.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error initializing client: %s", err))
		return 2
	}

	name := args[0]
	if err := client.Sys().DeletePolicy(name); err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error: %s", err))
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Policy '%s' deleted.", name))
	return 0
}

func (c *PolicyDeleteCommand) Synopsis() string {
	return "Delete a policy from the server"
}

func (c *PolicyDeleteCommand) Help() string {
	helpText := `
Usage: vault policy-delete [options] name

  Delete a policy with the given name.

  One the policy is deleted, all users associated with the policy will
  be affected immediately. When a user is associated with a policy that
  doesn't exist, it is identical to not being associated with that policy.

General Options:

  -address=addr           The address of the Vault server.

  -ca-cert=path           Path to a PEM encoded CA cert file to use to
                          verify the Vault server SSL certificate.

  -ca-path=path           Path to a directory of PEM encoded CA cert files
                          to verify the Vault server SSL certificate. If both
                          -ca-cert and -ca-path are specified, -ca-path is used.

  -tls-skip-verify        Do not verify TLS certificate. This is highly
                          not recommended. This is especially not recommended
                          for unsealing a vault.

`
	return strings.TrimSpace(helpText)
}
