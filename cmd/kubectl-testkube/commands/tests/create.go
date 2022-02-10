package tests

import (
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

// NewCreateTestsCmd is a command tp create new Script Custom Resource
func NewCreateTestsCmd() *cobra.Command {

	var (
		testName        string
		testNamespace   string
		testContentType string
		file            string
		executorType    string
		uri             string
		gitUri          string
		gitBranch       string
		gitPath         string
		gitUsername     string
		gitToken        string
		tags            []string
	)

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create new test",
		Long:    `Create new Test Custom Resource`,
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()

			client, _ := common.GetClient(cmd)
			test, _ := client.GetTest(testName, testNamespace)
			if testName == test.Name {
				ui.Failf("Test with name '%s' already exists in namespace %s", testName, testNamespace)
			}

			options, err := NewUpsertTestOptionsFromFlags(cmd, test)
			ui.ExitOnError("getting test options", err)

			test, err = client.CreateTest(options)
			ui.ExitOnError("creating test "+testName+" in namespace "+testNamespace, err)

			ui.Success("Test created", testNamespace, "/", testName)
		},
	}

	cmd.Flags().StringVarP(&testName, "name", "n", "", "unique test name - mandatory")
	cmd.Flags().StringVarP(&file, "file", "f", "", "test file - will be read from stdin if not specified")
	cmd.Flags().StringVarP(&testNamespace, "test-namespace", "", "testkube", "namespace where test will be created defaults to 'testkube' namespace")
	cmd.Flags().StringVarP(&testContentType, "test-content-type", "", "", "content type of test one of string|file-uri|git-file|git-dir")

	cmd.Flags().StringVarP(&executorType, "type", "t", "", "test type (defaults to postman/collection)")

	cmd.Flags().StringVarP(&uri, "uri", "", "", "URI of resource - will be loaded by http GET")
	cmd.Flags().StringVarP(&gitUri, "git-uri", "", "", "Git repository uri")
	cmd.Flags().StringVarP(&gitBranch, "git-branch", "", "", "if uri is git repository we can set additional branch parameter")
	cmd.Flags().StringVarP(&gitPath, "git-path", "", "", "if repository is big we need to define additional path to directory/file to checkout partially")
	cmd.Flags().StringVarP(&gitUsername, "git-username", "", "", "if git repository is private we can use username as an auth parameter")
	cmd.Flags().StringVarP(&gitToken, "git-token", "", "", "if git repository is private we can use token as an auth parameter")
	cmd.Flags().StringSliceVar(&tags, "tags", nil, "comma separated list of tags: --tags tag1,tag2,tag3")

	return cmd
}
