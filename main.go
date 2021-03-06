package main

import (
	"encoding/xml"
	"fmt"
	"github.com/strive-after/gojenkins/api"
	"github.com/strive-after/gojenkins/resource"
	"net/http"
)

var jenkins  *api.Jenkins






func main() {
	var req *http.Request
	//
	ar := api.NewJsonAPIRequest("get","/")
	//
	req, err := http.NewRequest(ar.Method, "http://39.105.114.198:8080/job/aaahblog/config.xml", ar.Payload)
	jenkins = api.NewJenkins("admin","123","http://39.105.114.198:8080","")
	req.SetBasicAuth(jenkins.User,jenkins.Password)
	responce,err := jenkins.Client.Do(req)
	if err != nil {
		fmt.Println(err,"err1")
		return
	}
	//data,err := ioutil.ReadAll(responce.Body)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(data))

	data := `<?xml version="1.0" encoding="UTF-8"?><flow-definition plugin="workflow-job@2.32">
	<actions>
	<org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.3.8"/>
	</actions>
	<description>Pipeline Job Sample</description>
	<keepDependencies>false</keepDependencies>
	<properties>
	<hudson.plugins.jira.JiraProjectProperty plugin="jira@3.0.7"/>
	<com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty plugin="gitlab-plugin@1.5.12">
	<gitLabConnection/>
	</com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty>
	</properties>
	<definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.69">
	<script>pipeline {
		agent any

		stages {
		stage('Build') {
		steps {
		sh 'echo Build stage ...'
	}
	}
		stage('Test'){
		steps {
		sh 'echo Test stage ...'
	}
	}
		stage('Deploy') {
		steps {
		sh 'echo Deploy stage ...'
	}
	}
	}
	}</script>
	<sandbox>false</sandbox>
	</definition>
	<triggers/>
	<disabled>false</disabled>
	</flow-definition>
`
	test := resource.FlowDefinition{}
	err := xml.Unmarshal([]byte(data),&test)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v",test)
}
