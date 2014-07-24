authenticator
=============

The authenticator application that manages user authentication and oAuth authorization (but not resources).

Functionality Required:

- [x] Create User
- [x] Save Users into Database
- [ ] Register for a new account
- [ ] Log into an account
- [ ] Reset password
- [ ] Verify email address
- [ ] OAuth2 functionality

All the functionality for user actions should be contained in libuser, but
the templates for HTML should be in the main authenticator repository. There
should be small pieces of functionality that glue the library to the templates
and handle any middleware needs, such as adding captcha.

__This system is not meant to handle user profile information__. This type of
information must live in the applications (partnership site, content engine,
etc..) and be exposed in a way those applications see fit. This system only
handles the authentication of users and the control of OAuth2 tokens and
authorizations.

Requirements
------------
After cloning into your Go Workspace, run go install to grab all the
dependencies.
If you want to use the database migration scripts, download
[goose](https://bitbucket.org/liamstask/goose)

Setup
-----
If you're using [goose](https://bitbucket.org/liamstask/goose) setup your
db/dbconf.yml file with your db connection:

```yaml
development:
    driver: mymysql
    open: authenticator/authenticator/authenticator
```

Reasoning
---------
The NYC Camp Web Properties are implemented as a set of diverse web
applications, using many different types of technologies. A traditional
authentication and authorization model will not function properly with this type
of setup. To allow for Single Sign On and create a strong barrier between
authentication server information and app information, we've decided to build
this system in Go. The long term goal is to migrate this system into Drupal when
possible, but the current Drupal resources are simply not mature enough to handle the
needs of this type of system. To handle the needs of both mobile applications,
single page applications, and separated resource servers, OAuth2 with Token
verification is a necessity. For the time being we don't have the bandwidth to
build out a hypermedia API and this type of functionality.
