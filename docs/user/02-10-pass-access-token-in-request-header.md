# Pass an Access Token in a Request Header

The Application Gateway service supports passing the access token directly in the request.

## Passing the Access Token

If the user is already authenticated to the target API, the access token can be passed in a custom `Access-token` header. The header's value is of the `Bearer {TOKEN}` or `Basic {TOKEN}` form. If the Application Gateway service detects that the custom header is present, instead of performing authentication steps, it removes the `Access-token` header and passes the received value in the `Authorization` header.
