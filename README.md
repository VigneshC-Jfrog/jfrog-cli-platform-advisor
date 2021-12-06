# platform-advisor

## About this plugin
This plugin will analyse the JFrog Platform instance and provide the non conformance against the best practices based on the predefines rules.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install hello-frog`

Installing a specific version:

`$ jfrog plugin install hello-frog@version`

Uninstalling a plugin

`$ jfrog plugin uninstall hello-frog`

## Usage
### Commands
* hello
    - Arguments:
        - addressee - The name of the person you would like to greet.
    - Flags:
        - shout: Makes output uppercase **[Default: false]**
        - repeat: Greets multiple times **[Default: 1]**
    - Example:
    ```
  $ jfrog hello-frog hello world --shout --repeat=2
  
  NEW GREETING: HELLO WORLD!
  NEW GREETING: HELLO WORLD!
  ```

### Environment variables
* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: New greeting: ]**

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).
