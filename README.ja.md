リソースサーバー実装 (Go)
=======================

概要
----

これはリソースサーバーの Go 実装です。 [OpenID Connect Core 1.0][OIDCCore]
で定義されている[ユーザー情報エンドポイント][UserInfoEndpoint]をサポートし、また、[RFC 6750][RFC6750]
(The OAuth 2.0 Authorization Framework: Bearer Token Usage)
に定義されている方法でアクセストークンを受け取る保護リソースエンドポイントの例も含んでいます。

この実装は [Go][Go] と、[authlete-go-gin][AuthleteGoGin]
ライブラリを用いて書かれています。

クライアントアプリケーションが提示したアクセストークンの有効性を調べるため、このリソースサーバーは
Authlete サーバーに問い合わせをおこないます。
これはつまり、このリソースサーバーは、アクセストークンを発行した認可サーバーが
Authlete をバックエンドサービスとして使用していることを期待していることを意味します。
[gin-oauth-server][GinOAuthServer] はそのような認可サーバーの実装であり、[OAuth 2.0][RFC6749]
と [OpenID Connect][OIDC] をサポートしています。

ライセンス
----------

  Apache License, Version 2.0

ソースコード
------------

  <code>https://github.com/authlete/gin-resource-server</code>

Authlete について
-----------------

[Authlete][Authlete] (オースリート) は、OAuth 2.0 & OpenID Connect
の実装をクラウドで提供するサービスです ([概説][AuthleteOverview])。 Authlete
が提供するデフォルト実装を使うことにより、もしくは [Authlete Web API][AuthleteAPI]
を用いて認可サーバーを自分で実装することにより、OAuth 2.0 と OpenID Connect
の機能を簡単に実現できます。

このリソースサーバーの実装を使うには、Authlete から API クレデンシャルズを取得し、`authlete.toml`
に設定する必要があります。 API クレデンシャルズを取得する手順はとても簡単です。
単にアカウントを登録するだけで済みます ([サインアップ][AuthleteSignUp])。
詳細は[クイックガイド][AuthleteGettingStarted]を参照してください。

実行方法
--------

1. authlete-go ライブラリと authlete-go-gin ライブラリをインストールします。

        $ go get github.com/authlete/authlete-go
        $ go get github.com/authlete/authlete-go-gin

2. このリソースサーバーの実装をダウンロードします。

        $ git clone https://github.com/authlete/gin-resource-server.git
        $ cd gin-resource-server

3. 設定ファイルを編集して API クレデンシャルズをセットします。

        $ vi authlete.toml

4. リソースサーバーをビルドします。

        $ make

5. `http://localhost:8081`　でリソースサーバーを起動します。

        $ make run

エンドポイント
--------------

この実装は、下表に示すエンドポイントを公開します。

| エンドポイント             | パス            |
|:---------------------------|:----------------|
| ユーザー情報エンドポイント | `/api/userinfo` |
| 時刻エンドポイント         | `/api/time`     |

#### ユーザー情報エンドポイント

ユーザー情報エンドポイントは、[OpenID Connect Core 1.0][OIDCCore] の
[5.3. UserInfo Endpoint][UserInfoEndpoint] に記述されている要求事項を実装したものです。

このエンドポイントは、クライアントアプリケーションの設定に応じて、ユーザー情報を
JSON 形式もしくは JWT 形式で返します。 クライアントアプリケーションのメタデータの
`userinfo_signed_response_alg` と `userinfo_encrypted_response_alg`
の両方とも指定されていなければ、ユーザー情報は素の JSON で返されます。
そうでない場合は、シリアライズされた JWT で返されます。 Authlete
はクライアントアプリケーションのメタデータを管理するための Web コンソール
([デベロッパー・コンソール][DeveloperConsole]) を提供しています。
クライアントアプリケーションのメタデータについては、
[OpenID Connect Dynamic Client Registration 1.0][OIDCDynReg] の
[2. Client Metadata][OIDCDynReg_Metadata] を参照してください。

エンドポイントから返されるユーザー情報には、ユーザーの[クレーム][OIDCCore_Claims]が含まれています。
手短に言うと、クレームとは、名前やメールアドレスなどの、ユーザーに関する情報です。Authlete は
(OpenID Connect をサポートしているにもかかわらず)
ユーザーデータを管理しないので、あなたがクレーム値を提供しなければなりません。
これは、`UserInfoReqHandlerSpi` インターフェースを実装することでおこないます。

このリソースサーバーの実装では、`UserInfoReqHandlerSpiImpl` が `UserInfoReqHandlerSpi`
インターフェースの実装で、ダミーのユーザーデータベース (`user_management.go`)
からクレーム値を取り出しています。

#### 時刻エンドポイント

このリソースサーバーに実装されているカントリーエンドポイントは、
保護リソースエンドポイントの一例に過ぎません。
主な目的は、保護リソースエンドポイントにおけるアクセストークンの有効性の確認方法を示すことです。

時刻エンドポイントのパスは `/api/time` です。
このエンドポイントは [RFC 6750][RFC6750] の [2.1. Authorization Request Header Field][RFC6750_2_1]
で定義されている方法でアクセストークンを受け付けます。

```
$ ACCESS_TOKEN=YOUR_ACCESS_TOKEN
$ curl -v http://localhost:8081/api/time \
       -H "Authorization: Bearer ${ACCESS_TOKEN}"
```

時刻エンドポイントは、現在時刻に関する情報を JSON で返します。
下記はレスポンスの例です。

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

Web API を OAuth 2.0 のアクセストークンで保護する方法に関する一般的な情報および
Authlete 固有の情報については、[Protected Resource][ProtectedResource]
を参照してください。

その他の情報
------------

- [Authlete][Authlete] - Authlete ホームページ
- [authlete-go][AuthleteGo] - Go 用 Authlete ライブラリ
- [authlete-go-gin][AuthleteGoGin] - Gin 用 Authlete ライブラリ
- [gin-oauth-server][GinOAuthServer] - 認可サーバーの実装

コンタクト
----------

コンタクトフォーム : https://www.authlete.com/ja/contact/

| 目的 | メールアドレス       |
|:-----|:---------------------|
| 一般 | info@authlete.com    |
| 営業 | sales@authlete.com   |
| 広報 | pr@authlete.com      |
| 技術 | support@authlete.com |

[Authlete]:               https://www.authlete.com/ja/
[AuthleteAPI]:            https://docs.authlete.com/
[AuthleteGettingStarted]: https://www.authlete.com/ja/developers/getting_started/
[AuthleteOverview]:       https://www.authlete.com/ja/developers/overview/
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
