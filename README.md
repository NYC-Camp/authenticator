authenticator
=============

The authenticator application that manages user authentication and oAuth authorization (but not resources).

Functionality Required:

- [x] Create User
- [x] User Storage
- [ ] Generic Error Handler
- [ ] Middleware Layers
    - [x] session handling
- [ ] 404 Page Handler
    - [ ] 404 Page Template
- [x] Page templating
    - [x] Create page templates package
    - [x] Expose page template interface
    - [x] Implement page template interface
- [ ] Register for a new account
    - [x] Create Registration Form
    - [ ] Add password strength tool
    - [x] Add username & email address duplicate checker
    - [ ] Create api endpoints to check for usernames and email addresses
    - [ ] Add captcha (if necessary)
    - [x] Add flash messages for information
- [ ] Log into an account
    - [x] Create Login form
    - [x] Add flash messages for information
    - [x] Session Based Flash Messages
- [x] Logout of an account
    - [x] Create Login endpoint
    - [x] Add flash message and redirect to login page
- [ ] Reset password
    - [ ] Create Reset Password form
    - [ ] Add reset password link to login form
- [ ] Verify email address
    - [ ] Add endpoint to verify email
    - [ ] Add template for verification page
- [ ] Add CSRF token functionality
- [ ] OAuth2 functionality
    - [ ] Implement implicit flow
    - [ ] Implement token verification endpoint
    - [ ] Add a page where users can administer their authorizations
- [ ] Create combined login and registration form
- [ ] Migrate application to Angular front end
- [ ] Add flood controls and rate limiting
- [ ] Email Notices functionality
    - [ ] Add SendGrid functionality
    - [ ] Account Registration notice
    - [ ] Password Reset notice
- [ ] Add a basic profile page
    - [ ] Add functionality to update the user's password
    - [ ] Add functionality to update the user's email address
    - [ ] Add functionality to update the user's username
- [ ] Basic user management interface
    - [ ] Create new users
    - [ ] Disable user accounts
    - [ ] Delete user accounts

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
