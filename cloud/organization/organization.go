package organization

import (
	http_context "context"
	"io"
	"strconv"

	"github.com/pkg/errors"

	astro "github.com/astronomer/astro-cli/astro-client"
	astrocore "github.com/astronomer/astro-cli/astro-client-core"
	"github.com/astronomer/astro-cli/cloud/auth"
	"github.com/astronomer/astro-cli/config"
	"github.com/astronomer/astro-cli/context"
	"github.com/astronomer/astro-cli/pkg/input"
	"github.com/astronomer/astro-cli/pkg/printutil"
)

var (
	errInvalidOrganizationKey  = errors.New("invalid organization selection")
	errInvalidOrganizationName = errors.New("invalid organization name")
	Login                      = auth.Login
	CheckUserSession           = auth.CheckUserSession
	FetchDomainAuthConfig      = auth.FetchDomainAuthConfig
)

func newTableOut() *printutil.Table {
	return &printutil.Table{
		Padding:        []int{44, 50},
		DynamicPadding: true,
		Header:         []string{"NAME", "ID"},
		ColorRowCode:   [2]string{"\033[1;32m", "\033[0m"},
	}
}

func ListOrganizations(coreClient astrocore.CoreClient) ([]astrocore.Organization, error) {
	resp, err := coreClient.ListOrganizationsWithResponse(http_context.Background())
	if err != nil {
		return nil, err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return nil, err
	}
	orgs := *resp.JSON200
	return orgs, nil
}

// List all Organizations
func List(out io.Writer, coreClient astrocore.CoreClient) error {
	c, err := config.GetCurrentContext()
	if err != nil {
		return err
	}
	or, err := ListOrganizations(coreClient)
	if err != nil {
		return errors.Wrap(err, astro.AstronomerConnectionErrMsg)
	}
	tab := newTableOut()
	for i := range or {
		name := or[i].Name
		organizationID := or[i].Id

		var color bool

		if c.Organization == or[i].Id {
			color = true
		}
		tab.AddRow([]string{name, organizationID}, color)
	}

	tab.Print(out)

	return nil
}

func getOrganizationSelection(out io.Writer, coreClient astrocore.CoreClient) (*astrocore.Organization, error) {
	tab := printutil.Table{
		Padding:        []int{5, 44, 50},
		DynamicPadding: true,
		Header:         []string{"#", "NAME", "ID"},
		ColorRowCode:   [2]string{"\033[1;32m", "\033[0m"},
	}

	var c config.Context
	c, err := config.GetCurrentContext()
	if err != nil {
		return nil, err
	}

	or, err := ListOrganizations(coreClient)
	if err != nil {
		return nil, err
	}

	deployMap := map[string]astrocore.Organization{}
	for i := range or {
		index := i + 1

		color := c.Organization == or[i].Id
		tab.AddRow([]string{strconv.Itoa(index), or[i].Name, or[i].Id}, color)

		deployMap[strconv.Itoa(index)] = or[i]
	}
	tab.Print(out)
	choice := input.Text("\n> ")
	selected, ok := deployMap[choice]
	if !ok {
		return nil, errInvalidOrganizationKey
	}

	return &selected, nil
}

func SwitchWithLogin(domain string, targetOrg *astrocore.Organization, astroClient astro.Client, coreClient astrocore.CoreClient, out io.Writer, shouldDisplayLoginLink bool) error {
	return Login(domain, targetOrg.AuthServiceId, "", astroClient, coreClient, out, shouldDisplayLoginLink)
}

func SwitchWithContext(domain string, targetOrg *astrocore.Organization, authConfig astro.AuthConfig, astroClient astro.Client, coreClient astrocore.CoreClient, out io.Writer) error {
	c, _ := context.GetCurrentContext()
	// reset org context
	_ = c.SetOrganizationContext(targetOrg.Id, targetOrg.ShortName)
	// need to reset all relevant keys because of https://github.com/spf13/viper/issues/1106 :shrug
	_ = c.SetContextKey("token", c.Token)
	_ = c.SetContextKey("refreshtoken", c.RefreshToken)
	_ = c.SetContextKey("user_email", c.UserEmail)
	c, _ = context.GetCurrentContext()
	// call check user session which will trigger workspace switcher flow
	return CheckUserSession(&c, authConfig, astroClient, coreClient, out)
}

// Switch switches organizations
func Switch(orgNameOrID string, astroClient astro.Client, coreClient astrocore.CoreClient, out io.Writer, shouldDisplayLoginLink bool) error {
	// get current context
	c, err := context.GetCurrentContext()
	if err != nil {
		return err
	}

	// get target org
	var targetOrg *astrocore.Organization
	if orgNameOrID == "" {
		targetOrg, err = getOrganizationSelection(out, coreClient)
		if err != nil {
			return err
		}
	} else {
		or, err := ListOrganizations(coreClient)
		if err != nil {
			return err
		}
		for i := range or {
			if or[i].Name == orgNameOrID {
				targetOrg = &or[i]
			}
			if or[i].Id == orgNameOrID {
				targetOrg = &or[i]
			}
		}
	}
	if targetOrg == nil {
		return errInvalidOrganizationName
	}
	// fetch auth config
	authConfig, err := FetchDomainAuthConfig(c.Domain)
	if err != nil {
		return err
	}
	if authConfig.AuthFlow == auth.AuthFlowIdentityFirst {
		return SwitchWithContext(c.Domain, targetOrg, authConfig, astroClient, coreClient, out)
	}
	return SwitchWithLogin(c.Domain, targetOrg, astroClient, coreClient, out, shouldDisplayLoginLink)
}

// Write the audit logs to the provided io.Writer.
func ExportAuditLogs(client astro.Client, out io.Writer, orgName string, earliest int) error {
	logStreamBuffer, err := client.GetOrganizationAuditLogs(orgName, earliest)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, logStreamBuffer)
	if err != nil {
		logStreamBuffer.Close()
		return err
	}
	logStreamBuffer.Close()
	return nil
}
