# js-to-go-example
Same code from Javascript to Go.

### For go file, send POST request to localhost:3001 containing body:
```
{
    "commands": [
        {
            "command": "echo {ENV_TEST} {ENV_POOL}",
            "session": false,
            "parameter": null,
            "environment_variables": [
                {
                    "name": "ENV_TEST",
                    "value": "test"
                },
                {
                    "name": "ENV_POOL",
                    "value": "poolll"
                }
            ],
            "break": false
        },
        {
            "command": "echo {ENV_MOO}",
            "session": false,
            "parameter": null,
            "environment_variables": [
                {
                    "name": "ENV_MOO",
                    "value": "hello"
                }
            ],
            "break": false
        }
    ]
}

```
