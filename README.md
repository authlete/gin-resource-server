Resource Server Implementation in Go
====================================

Overview
--------

This is a resource server implementation in Go. It supports a
[userinfo endpoint][UserInfoEndpoint] defined in
[OpenID Connect Core 1.0][OIDCCore] and includes an example of a protected
resource endpoint that accepts an access token in the way defined in
[2.1. Authorization Request Header Field][RFC6750_2_1] of [RFC 6750][RFC6750]
(The OAuth 2.0 Authorization Framework: Bearer Token Usage).

This implementation is written using [Gin][Gin] and
[authlete-go-gin][AuthleteGoGin] library which is an Anthlete's open-source
library for Gin.

To validate an access token presented by a client application, this resource
server makes an inquiry to the Authlete server. This means that this resource
server expects that the authorization server which has issued the access token
uses Authlete as a backend service. [gin-oauth-server][GinOAuthServer] is such
an authorization server implementation and it supports [OAuth 2.0][RFC6749]
and [OpenID Connect][OIDC].

License
-------

  Apache License, Version 2.0

Source Code
-----------

  <code>https://github.com/authlete/gin-resource-server</code>

About Authlete
--------------

[Authlete][Authlete] is a cloud service that provides an implementation of
OAuth 2.0 & OpenID Connect ([overview][AuthleteOverview]). You can easily get
the functionalities of OAuth 2.0 and OpenID Connect either by using the default
implementation provided by Authlete or by implementing your own authorization
server using [Authlete Web APIs][AuthleteAPI].

To use this resource server implementation, you need to get API credentials
from Authlete and set them in `authlete.toml`. The steps to get API credentials
are very easy. All you have to do is just to register your account
([sign up][AuthleteSignUp]). See [Getting Started][AuthleteGettingStarted] for
details.

How To Run
----------

1. Install authlete-go and authlete-go-gin libraries.

        $ go get github.com/authlete/authlete-go
        $ go get github.com/authlete/authlete-go-gin

2. Download the source code of this resource server implementation.

        $ git clone https://github.com/authlete/gin-resource-server.git
        $ cd gin-resource-server

3. Edit the configuration file to set the API credentials of yours.

        $ vi authlete.toml

4. Build the resource server.

        $ make

5. Start the resource server on `http://localhost:8081`.

        $ make run

Endpoints
---------

This implementation exposes endpoints as listed in the table below.

| Endpoint          | Path            |
|:------------------|:----------------|
| UserInfo Endpoint | `/api/userinfo` |
| Time Endpoint     | `/api/time`     |

#### UserInfo Endpoint

The userinfo endpoint is an implementation of the requirements described in
[5.3. UserInfo Endpoint][UserInfoEndpoint] of [OpenID Connect Core 1.0][OIDCCore].

The endpoint returns user information in JSON or JWT format, depending on the
configuration of the client application. If both `userinfo_signed_response_alg`
and `userinfo_encrypted_response_alg` of the metadata of the client application
are not specified, user information is returned as a plain JSON. Otherwise, it
is returned as a serialized JWT. Authlete provides you with a Web console
([Developer Console][DeveloperConsole]) to manage metadata of client applications.
As for metadata of client applications, see [2. Client Metadata][OIDCDynReg_Metadata]
of [OpenID Connect Dynamic Client Registration 1.0][OIDCDynReg].

User information returned from the endpoint contains [claims][OIDCCore_Claims]
of the user. In short, _claims_ are pieces of information about the user such
as a given name and an email address. Because Authlete does not manage user
data (although it supports OpenID Connect), you have to provide claim values.
It is achieved by implementing `UserInfoReqHandlerSpi` interface.

In this resource server implementation, `UserInfoReqHandlerSpiImpl` is an
example implementation of `UserInfoReqHandlerSpi` interface and it retrieves
claim values from the dummy user database (`user_management.go`).

#### Time Endpoint

The time endpoint implemented in this resource server is just an example of a
protected resource endpoint. Its main purpose is to show how to validate an
access token at a protected resource endpoint.

The path of the time endpoint is `/api/time`. The endpoint accepts an access
token in the way defined in [2.1. Authorization Request Header Field][RFC6750_2_1]
of [RFC 6750][RFC6750].

```
$ ACCESS_TOKEN=YOUR_ACCESS_TOKEN
$ curl -v http://localhost:8081/api/time \
       -H "Authorization: Bearer ${ACCESS_TOKEN}"
```

The time endpoint returns information about the current time in JSON format.
The following is an example response.

```
{
  "year":   2019,
  "month":  8,
  "day":    9,
  "hour":   14,
  "minute": 45,
  "second": 2
}
```

As for generic and Authlete-specific information regarding how to protect Web
APIs by OAuth 2.0 access tokens, see [Protected Resource][ProtectedResource].

See Also
--------

- [Authlete][Authlete] - Authlete Home Page
- [authlete-go][AuthleteGo] - Authlete Library for Go
- [authlete-go-gin][AuthleteGoGin] - Authlete Library for Gin
- [gin-oauth-server][GinOAuthServer] - Authorization Server Implementation

Contact
-------

Contact Form : https://www.authlete.com/contact/

| Purpose   | Email Address        |
|:----------|:---------------------|
| General   | info@authlete.com    |
| Sales     | sales@authlete.com   |
| PR        | pr@authlete.com      |
| Technical | support@authlete.com |

[Authlete]:               https://www.authlete.com/
[AuthleteAPI]:            https://docs.authlete.com/
[AuthleteGettingStarted]: https://www.authlete.com/developers/getting_started/
[AuthleteOverview]:       https://www.authlete.com/developers/overview/
[AuthleteGo]:             https://github.com/authlete/authlete-go/
[AuthleteGoGin]:          https://github.com/authlete/authlete-go-gin/
[AuthleteSignUp]:         https://so.authlete.com/accounts/signup
[DeveloperConsole]:       https://www.authlete.com/developers/cd_console/
[Gin]:                    https://github.com/gin-gonic/gin
[GinOAuthServer]:         https://github.com/authlete/gin-oauth-server/
[OIDC]:                   https://openid.net/connect/
[OIDCCore]:               https://openid.net/specs/openid-connect-core-1_0.html
[OIDCCore_Claims]:        https://openid.net/specs/openid-connect-core-1_0.html#Claims
[OIDCDynReg]:             https://openid.net/specs/openid-connect-registration-1_0.html
[OIDCDynReg_Metadata]:    https://openid.net/specs/openid-connect-registration-1_0.html#
[ProtectedResource]:      https://www.authlete.com/developers/definitive_guide/protected_resource/
[RFC6749]:                https://tools.ietf.org/html/rfc6749
[RFC6750]:                https://tools.ietf.org/html/rfc6750
[RFC6750_2_1]:            https://tools.ietf.org/html/rfc6750#section-2.1
[UserInfoEndpoint]:       https://openid.net/specs/openid-connect-core-1_0.html#UserInfo
