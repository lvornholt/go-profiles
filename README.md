# Golang profile service 
This projects wants to supports other projects with profiles, to use the same code on different environsments just be setting a PROFILE environment varibale. Inspired by java spring profiles. This makes it easier to set up the service without define several environment variables on every system. 

## Different profiles
The application profile files needs to be named like application-<PROFILE-NAME>.yml. The default profile file application.yml will be used if the environment variable is not set.

## Example profile file
```
profile:
    name: development
application:
	port: 8081
    context: /context
    name: default-application
```

## Usage
Checkout this project with
```
go get github.com/lvornholt/go-profiles
```

## Example
```go
import (
    ...
	profile "github.com/lvornholt/go-profiles"
}

var name string

func init() {

	configName := profile.GetConfigValue("application.name")
	if len(configName) > 0 {
		name = configName
	}
}

func main() {
    ...
```