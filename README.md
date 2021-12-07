# platform-advisor

## About this plugin
This plugin will analyse the JFrog Platform instance and provide the non conformance against the best practices based on the predefined rules.

## Installation with JFrog CLI

Installing the latest version:
`$ jfrog plugin install platform-advisor`

Installing a specific version:
`$ jfrog plugin install platform-advisor@version`

Uninstalling a plugin
`$ jfrog plugin uninstall platform-advisor`

## Usage
### Commands

To Publish the plugin:
* jfrog plugin publish platform-advisor 1.0.0

To view the Security Advise report:
* jfrog platform-advisor adv security

To view the Performance Advise report:
* jfrog platform-advisor adv performance

In order to view both Security and Performance Advise report:
* jfrog platform-advisor adv all

### Environment variables
* export JFROG_CLI_PLUGINS_SERVER=local
* export JFROG_CLI_PLUGINS_REPO=cli-plugins

