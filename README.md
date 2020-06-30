![](./six910.png)

# Six910 (pronounced six nine ten)

[![Build Status](https://travis-ci.org/Ulbora/Six910.svg?branch=master)](https://travis-ci.org/Ulbora/Six910)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Six910&metric=alert_status)](https://sonarcloud.io/dashboard?id=Six910)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/Six910)](https://goreportcard.com/report/github.com/Ulbora/Six910)


### The Speeding Fast Shopping Cart with Dependency Injection
An open source ecommerce platform server for building online shopping cart systems; written in golang. Includes REST service backend including admin panel and switchable templates.

All components of Six910 are developed as Go modules and injected using dependency injection.



This project is intended to be and REST implementation of a hosted ecommerce solution. 

1. Users of the Six910, together with Six910 UI will be able to customize templates just like other hosted solutions.
2. Templates can either be written in Golang templating or use a JavaScript framework like Angular or React.
3. REST services will be used to expose all cart functionality.

## Template Designer
There will also be a template designer to make desiging templates much easier than it currently is with most hosted shopping cart solutions.

1. Users will be able to download there current template
2. Modify the template using the user's store URL
3. Upload the template back to the hosted site

## Database
### The database module in injected using dependency injection in the main func.
The default database for Six910 will be MySQL. The database module can be found [here](https://github.com/Ulbora/six910-mysql). The database interface, however is modular and can easily be switched out for any other database.

## Addon Applications
Six910 will allow third party developers to build addon applications that integrate into templates.

**We need Golang developers to work on this project.** If you are interested in contributing, send a **pull request** with your name added to the developer section of the readme.

## Developers and Contributors:

1. Ken Williamson
2. Your Name Here


