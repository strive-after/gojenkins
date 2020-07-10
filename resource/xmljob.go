package resource

import "encoding/xml"

type FlowDefinition struct {
	XMLName xml.Name `xml:"flow-definition"`
	Text    string   `xml:",chardata"`
	Plugin  string   `xml:"plugin,attr"`
	Actions struct {
		Text                                                                  string `xml:",chardata"`
		OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
	} `xml:"actions"`
	Description      string `xml:"description"`      // Pipeline Job Sample
	KeepDependencies string `xml:"keepDependencies"` // false
	Properties       struct {
		Text                                 string `xml:",chardata"`
		HudsonPluginsJiraJiraProjectProperty struct {
			Text   string `xml:",chardata"`
			Plugin string `xml:"plugin,attr"`
		} `xml:"hudson.plugins.jira.JiraProjectProperty"`
		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
			Text             string `xml:",chardata"`
			Plugin           string `xml:"plugin,attr"`
			GitLabConnection string `xml:"gitLabConnection"`
		} `xml:"com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty"`
	} `xml:"properties"`
	Definition struct {
		Text    string `xml:",chardata"`
		Class   string `xml:"class,attr"`
		Plugin  string `xml:"plugin,attr"`
		Script  string `xml:"script"`  // pipeline { 		agent any  	...
		Sandbox string `xml:"sandbox"` // false
	} `xml:"definition"`
	Triggers string `xml:"triggers"`
	Disabled string `xml:"disabled"` // false
}

//

//import "encoding/xml"
//
//type XmlJob struct {
//	XMLName xml.Name `xml:"flow-definition"`
//	Text    string   `xml:",chardata"`
//	Plugin  string   `xml:"plugin,attr"`
//	Actions Actions `xml:"actions"`
//	Description      string `xml:"description"`      // Pipeline Job Sample
//	KeepDependencies string `xml:"keepDependencies"` // false
//	Properties       struct {
//		Text                                                         string                                                       `xml:",chardata"`
//		HudsonPluginsJiraJiraProjectProperty                         HudsonPluginsJiraJiraProjectProperty                         `xml:"hudson.plugins.jira.JiraProjectProperty"`
//		ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty `xml:"properties"`
//	}
//	Definition Definition `xml:"definition"`
//	Triggers string `xml:"triggers"`
//	Disabled string `xml:"disabled"` // false
//}
//
//type Actions struct {
//	Text                                                                  string `xml:",chardata"`
//	OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction"`
//}
//
//
//type OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction struct {
//	Text   string `xml:",chardata"`
//	Plugin string `xml:"plugin,attr"`
//}
//
//type  HudsonPluginsJiraJiraProjectProperty struct {
//	Text   string `xml:",chardata"`
//	Plugin string `xml:"plugin,attr"`
//}
//
//type  ComDabsquaredGitlabjenkinsConnectionGitLabConnectionProperty struct {
//	Text             string `xml:",chardata"`
//	Plugin           string `xml:"plugin,attr"`
//	GitLabConnection string `xml:"gitLabConnection"`
//}
//
//type Definition struct {
//	Text    string `xml:",chardata"`
//	Class   string `xml:"class,attr"`
//	Plugin  string `xml:"plugin,attr"`
//	Script  string `xml:"script"`  // pipeline { 		agent any  	...
//	Sandbox string `xml:"sandbox"` // false
//}
//
//
//func NewXmlJob() *XmlJob {
//	return &XmlJob{}
//}
