# apirole

API for Role Management

### Requirement

- Create Collection

```
- casbin_rule
- role_user
- roles
```

- Insert document role `Admin` in collection `roles`

```json
{
    "_id": {
        "$oid": "5f95a41f2e94e13067a087e0"
    },
    "display": "Admin",
    "description": "All API managements",
    "createdat": {
        "$date": "2020-10-02T01:11:18.965Z"
    },
    "updatedat": {
        "$date": "2020-10-02T01:11:18.965Z"
    }    
}
```

- Insert document role `Anonymous` in collection `roles`

```json
{
    "_id": {
        "$oid": "5f95a61d2e94e13067a087e1"
    },
    "display": "Anonymous",
    "description": "Access API without authorization",
    "createdat": {
        "$date": "2020-10-02T01:11:18.965Z"
    },
    "updatedat": {
        "$date": "2020-10-02T01:11:18.965Z"
    }    
}
```