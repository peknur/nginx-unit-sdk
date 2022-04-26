# Nginx Unit SDK
[![Test](https://github.com/peknur/nginx-unit-sdk/actions/workflows/test.yml/badge.svg)](https://github.com/peknur/nginx-unit-sdk/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/peknur/nginx-unit-sdk/unit.svg)](https://pkg.go.dev/github.com/peknur/nginx-unit-sdk/unit)

An unofficial [Nginx Unit](https://www.nginx.com/products/nginx-unit/) SDK for Go.  
SDK is in alpha state and breaking changes might occur. Tested against Nginx Unit `v1.26`.

From the Nginx Unit authors:  
*NGINX Unit is a dynamic application server, capable of running beside NGINX Plus and NGINX Open Source or standalone. NGINX Unit supports a RESTful JSON API, deploys configuration changes without service disruptions, and runs apps built with multiple languages and frameworks. Designed from scratch around the needs of your distributed applications, it lays the foundation for your.*

Nginx Unit developer [documentation](https://unit.nginx.org/)

## Install
```
$ go get github.com/peknur/nginx-unit-sdk
```
## Example

```go
svc, err := unit.NewServiceFromURL("http://127.0.0.1:8080")
if err != nil {
	log.Fatal(err)
}
cfg := config.Config{
	Listeners: config.Listeners{
		"*:80": listener.Config{
			Pass: "routes/main",
		},
	},
	Routes: config.Routes{
		"main": []route.Config{
			{
				Action: &route.Action{
					Pass: "applications/go",
				},
			}},
	},
	Applications: config.Applications{
		"go": application.Config{
			Type:             application.TypeGo,
			Executable:       "app",
			WorkingDirectory: "/apps/go",
			User:             "www-data",
			Group:            "www-data",
		},
	},
	AccessLog: "/var/log/unit.log",
}

if err := svc.CreateConfig(context.Background(), cfg); err != nil {
	log.Fatal(err)
}
```
