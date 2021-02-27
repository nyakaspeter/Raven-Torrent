# Get Server Information

Get server information.

**URL** : `/api/about`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type   | Description                                               |
| --------- | ------ | --------------------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful.               |
| `message` | string | Text message that contains the server's name and version. |

## Examples

**Request** :

`GET http://localhost:9000/api/about`

**Success Response** :

```json
{
  "success": true,
  "message": "White Raven Server v0.5.0"
}
```
