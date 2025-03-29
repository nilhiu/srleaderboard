---
title: SRLeaderboard API v1.0.0
language_tabs:
  - shell: Shell
language_clients:
  - shell: curl
toc_footers: []
includes: []
search: false
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="srleaderboard-api">SRLeaderboard API v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

OAS for SRLeaderboard's API

Base URLs:

* <a href="http://localhost/api">http://localhost/api</a>

License: <a href="https://opensource.org/license/bsd-3-clause">BSD 3-Clause</a>

# Authentication

* API Key (CookieSecurity)
    - Parameter Name: **jwt**, in: cookie. 

<h1 id="srleaderboard-api-runs">runs</h1>

Runs Management

## Get runs from the leaderboard

<a id="opIdGetRuns"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost/api/runs \
  -H 'Accept: application/json'

```

`GET /runs`

<h3 id="get-runs-from-the-leaderboard-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|offset|query|integer|false|The number of runs to skip before collecting the resulting list|
|limit|query|integer|false|The number of runs to return|

> Example responses

> 200 Response

```json
{
  "runs": [
    {
      "username": "string",
      "completion_time": 0
    }
  ],
  "amount": 0,
  "offset": 0,
  "limit": 0,
  "full_amount": 0
}
```

<h3 id="get-runs-from-the-leaderboard-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully returned the requested runs|[GetRunsResponse](#schemagetrunsresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request couldn't be processed due to client error|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|The server failed internally to process the request|None|

<aside class="success">
This operation does not require authentication
</aside>

## Add a run to the leaderboard and user database

<a id="opIdAddRun"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost/api/runs \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

`POST /runs`

> Body parameter

```json
{
  "time": "1h2m3s400ms"
}
```

<h3 id="add-a-run-to-the-leaderboard-and-user-database-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[AddRunRequest](#schemaaddrunrequest)|true|The information about the run|

> Example responses

> 201 Response

```json
{
  "username": "string",
  "placement": 0,
  "time": "1h2m3s400ms",
  "date_added": "2019-08-24T14:15:22Z"
}
```

<h3 id="add-a-run-to-the-leaderboard-and-user-database-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Successfully added the run|[AddRunResponse](#schemaaddrunresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request couldn't be processed due to client error|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|The server failed internally to process the request|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
CookieSecurity
</aside>

## Get user specific runs

<a id="opIdGetUserRuns"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost/api/runs/{user} \
  -H 'Accept: application/json'

```

`GET /runs/{user}`

<h3 id="get-user-specific-runs-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|user|path|string|true|The user for which to return the data about|
|offset|query|integer|false|The number of runs to skip before collecting the resulting list|
|limit|query|integer|false|The number of runs to return|

> Example responses

> 200 Response

```json
{
  "runs": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "completion_time": 0,
      "created_at": "2019-08-24T14:15:22Z"
    }
  ],
  "amount": 0,
  "offset": 0,
  "limit": 0,
  "full_amount": 0
}
```

<h3 id="get-user-specific-runs-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully returned the requested runs|[GetUserRunsResponse](#schemagetuserrunsresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request couldn't be processed due to client error|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|The specified resource couldn't be found|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|The server failed internally to process the request|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="srleaderboard-api-user">user</h1>

User Management

## Authenticate a registered user

<a id="opIdLogin"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost/api/auth/login \
  -H 'Content-Type: application/json'

```

`POST /auth/login`

> Body parameter

```json
{
  "username": "string",
  "password": "pa$$word"
}
```

<h3 id="authenticate-a-registered-user-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[LoginRequest](#schemaloginrequest)|true|The login form|

<h3 id="authenticate-a-registered-user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully authenticated the user|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request couldn't be processed due to client error|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|The user is unauthorized to make the request|None|

### Response Headers

|Status|Header|Type|Format|Description|
|---|---|---|---|---|
|200|Set-Cookie|string|jwt|none|

<aside class="success">
This operation does not require authentication
</aside>

## Register a user

<a id="opIdRegister"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost/api/auth/register \
  -H 'Content-Type: application/json'

```

`POST /auth/register`

> Body parameter

```json
{
  "username": "string",
  "email": "user@example.com",
  "password": "pa$$word"
}
```

<h3 id="register-a-user-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[RegisterRequest](#schemaregisterrequest)|true|The registration form|

<h3 id="register-a-user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successfully registered the user|None|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|The request couldn't be processed due to client error|None|
|409|[Conflict](https://tools.ietf.org/html/rfc7231#section-6.5.8)|Failed to register the user as the username already exists|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|The server failed internally to process the request|None|

### Response Headers

|Status|Header|Type|Format|Description|
|---|---|---|---|---|
|200|Set-Cookie|string|jwt|none|

<aside class="success">
This operation does not require authentication
</aside>

## Log the user out

<a id="opIdLogout"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost/api/auth/logout

```

`POST /auth/logout`

<h3 id="log-the-user-out-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|Successfully logged the user out|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|The user is unauthorized to make the request|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
CookieSecurity
</aside>

# Schemas

<h2 id="tocS_RegisterRequest">RegisterRequest</h2>
<!-- backwards compatibility -->
<a id="schemaregisterrequest"></a>
<a id="schema_RegisterRequest"></a>
<a id="tocSregisterrequest"></a>
<a id="tocsregisterrequest"></a>

```json
{
  "username": "string",
  "email": "user@example.com",
  "password": "pa$$word"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|true|none|none|
|email|string(email)|true|none|none|
|password|string(password)|true|none|none|

<h2 id="tocS_LoginRequest">LoginRequest</h2>
<!-- backwards compatibility -->
<a id="schemaloginrequest"></a>
<a id="schema_LoginRequest"></a>
<a id="tocSloginrequest"></a>
<a id="tocsloginrequest"></a>

```json
{
  "username": "string",
  "password": "pa$$word"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|true|none|none|
|password|string(password)|true|none|none|

<h2 id="tocS_AddRunRequest">AddRunRequest</h2>
<!-- backwards compatibility -->
<a id="schemaaddrunrequest"></a>
<a id="schema_AddRunRequest"></a>
<a id="tocSaddrunrequest"></a>
<a id="tocsaddrunrequest"></a>

```json
{
  "time": "1h2m3s400ms"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|time|[Duration](#schemaduration)|true|none|none|

<h2 id="tocS_AddRunResponse">AddRunResponse</h2>
<!-- backwards compatibility -->
<a id="schemaaddrunresponse"></a>
<a id="schema_AddRunResponse"></a>
<a id="tocSaddrunresponse"></a>
<a id="tocsaddrunresponse"></a>

```json
{
  "username": "string",
  "placement": 0,
  "time": "1h2m3s400ms",
  "date_added": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|true|none|none|
|placement|integer(int64)|true|none|none|
|time|[Duration](#schemaduration)|true|none|none|
|date_added|string(date-time)|true|none|none|

<h2 id="tocS_GetRunsResponse">GetRunsResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetrunsresponse"></a>
<a id="schema_GetRunsResponse"></a>
<a id="tocSgetrunsresponse"></a>
<a id="tocsgetrunsresponse"></a>

```json
{
  "runs": [
    {
      "username": "string",
      "completion_time": 0
    }
  ],
  "amount": 0,
  "offset": 0,
  "limit": 0,
  "full_amount": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|runs|[[LeaderboardRun](#schemaleaderboardrun)]|true|none|none|
|amount|integer|true|none|none|
|offset|integer|true|none|none|
|limit|integer|true|none|none|
|full_amount|integer|true|none|none|

<h2 id="tocS_GetUserRunsResponse">GetUserRunsResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetuserrunsresponse"></a>
<a id="schema_GetUserRunsResponse"></a>
<a id="tocSgetuserrunsresponse"></a>
<a id="tocsgetuserrunsresponse"></a>

```json
{
  "runs": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "completion_time": 0,
      "created_at": "2019-08-24T14:15:22Z"
    }
  ],
  "amount": 0,
  "offset": 0,
  "limit": 0,
  "full_amount": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|runs|[[UserRun](#schemauserrun)]|true|none|none|
|amount|integer|true|none|none|
|offset|integer|true|none|none|
|limit|integer|true|none|none|
|full_amount|integer|true|none|none|

<h2 id="tocS_LeaderboardRun">LeaderboardRun</h2>
<!-- backwards compatibility -->
<a id="schemaleaderboardrun"></a>
<a id="schema_LeaderboardRun"></a>
<a id="tocSleaderboardrun"></a>
<a id="tocsleaderboardrun"></a>

```json
{
  "username": "string",
  "completion_time": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|true|none|none|
|completion_time|integer(int64)|true|none|none|

<h2 id="tocS_UserRun">UserRun</h2>
<!-- backwards compatibility -->
<a id="schemauserrun"></a>
<a id="schema_UserRun"></a>
<a id="tocSuserrun"></a>
<a id="tocsuserrun"></a>

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "completion_time": 0,
  "created_at": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string(uuid)|true|none|none|
|completion_time|integer(int64)|true|none|none|
|created_at|string(date-time)|true|none|none|

<h2 id="tocS_Duration">Duration</h2>
<!-- backwards compatibility -->
<a id="schemaduration"></a>
<a id="schema_Duration"></a>
<a id="tocSduration"></a>
<a id="tocsduration"></a>

```json
"1h2m3s400ms"

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|none|

