# Command Utilities
This repository contains some scripts that I have created to help me in my development environment. Note that I use zsh so you may need to modify some scripts to fit with your needs.

## openproject
This script is used to open a project, wherever I am. Just edit ``projectsFile`` var to link to a json file containing your projects.
⚠️ It needs ``jq`` to be installed ! 
```bash
sudo apt-get update && sudo apt-get install jq
```

```json
{
  "portfolio": "~/dev/portfolio",
  "secondproject": "~/dev/projectTwo"
}
```

Simply run 
```bash
openproject portfolio
```
to be redirected to ``~/dev/portfolio``


## up, down and restart
Theses scripts are shortcuts for commands 
```bash
docker-compose up -d
docker-compose down
```
Simply run ``up`` or ``down`` to run the corresponding command.
If there is not any ``docker-compose.yml`` file in the current directory, the script will just abort.

Restart uses both ``down`` and ``up`` to fully recreate compose containers.
