const exec = require('child_process').exec;
const myShellScript = exec('sh clone/create_database.sh');
myShellScript.stdout.on('data', (data) => {
  console.log(data);
});
myShellScript.stderr.on('data', (data) => {
  console.error(data);
});
