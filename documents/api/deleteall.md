# Delete all running torrents

Delete all active torrents.

**URL** : `/api/deleteall`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Error Response

**Code** : `404 (Not Found)`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/deleteall`

**Success Response** :

```json
{
  "success": true,
  "message": "All torrents have been deleted."
}
```

**Error Response** :

```json
{
  "success": false,
  "message": "No active torrents found."
}
```
