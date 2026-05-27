package domain

import (
	"errors"

	"github.com/spf13/cobra"
)

// DomainCmd manages domains
var DomainCmd = &cobra.Command{
	Use:     "domain",
	Aliases: []string{"domains"},
	Short:   "Details of Civo domains",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

var domainRecordCmd = &cobra.Command{
	Use:     "record",
	Aliases: []string{"records"},
	Short:   "Details of Civo domains records",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("command is required")
	},
}

func init() {

	DomainCmd.AddCommand(domainListCmd)
	DomainCmd.AddCommand(domainCreateCmd)
	DomainCmd.AddCommand(domainRemoveCmd)

	// Domains record cmd
	DomainCmd.AddCommand(domainRecordCmd)
	domainRecordCmd.AddCommand(domainRecordListCmd)
	domainRecordCmd.AddCommand(domainRecordCreateCmd)
	domainRecordCmd.AddCommand(domainRecordShowCmd)
	domainRecordCmd.AddCommand(domainRecordRemoveCmd)
	domainRecordCmd.AddCommand(domainRecordUpdateCmd)

	/*
		Flags for domain record create cmd
	*/
	domainRecordCreateCmd.Flags().StringVarP(&recordName, "name", "n", "", "the name of the record")
	domainRecordCreateCmd.Flags().StringVarP(&recordType, "type", "e", "", "type of the record (A, CNAME, TXT, SRV, MX, NS)")
	domainRecordCreateCmd.Flags().StringVarP(&recordValue, "value", "v", "", "the value of the record")
	domainRecordCreateCmd.Flags().IntVarP(&recordTTL, "ttl", "t", 600, "The TTL of the record")
	domainRecordCreateCmd.Flags().IntVarP(&recordPriority, "priority", "p", 0, "the priority of record only for SRV and MX record")

	/*
		Flags for domain record update cmd
	*/
	domainRecordUpdateCmd.Flags().StringVarP(&updateRecordName, "name", "n", "", "the name of the record")
	domainRecordUpdateCmd.Flags().StringVarP(&updateRecordType, "type", "e", "", "type of the record (A, CNAME, TXT, SRV, MX, NS)")
	domainRecordUpdateCmd.Flags().StringVarP(&updateRecordValue, "value", "v", "", "the value of the record")
	domainRecordUpdateCmd.Flags().IntVarP(&updateRecordTTL, "ttl", "t", 0, "The TTL of the record")
	domainRecordUpdateCmd.Flags().IntVarP(&updateRecordPriority, "priority", "p", 0, "the priority of record only for SRV and MX record")

}
