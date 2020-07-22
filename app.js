commands = [
  {
     "command":"echo {ENV_TEST} {ENV_POOL}",
     "session":false,
     "parameters":[
        {
           "name":"variable",
           "value":"helloworld;whoami"
        }
     ],
     "environment_variables":[
        {
           "name":"ENV_TEST",
           "value":"test"
        },
        {
           "name":"ENV_POOL",
           "value":"poolll"
        }
     ],
     "break":false
  },
  {
     "command":"echo {ENV_TEST}",
     "session":false,
     "parameters":[],
     "environment_variables":[
        {
           "name":"ENV_TEST",
           "value":"heey"
        },
     ],
     "break":false
  }
]
for(obj of commands) {
  let env_values = {}
  obj.environment_variables.map((obj) => {
    env_values[obj.name] = obj.value
  })
  const matches = obj.command.match(/\{([^}]+)\}/g)
  for(let i = 0; i < matches.length; i++){
    const matchString = matches[i].replace(/[{}]/g, '');
    obj.command = obj.command.replace(matches[i], env_values[matchString]);
  }
  console.log('Command: ', obj.command)
}
