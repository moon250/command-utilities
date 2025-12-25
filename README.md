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


## createproject
This tool creates a project directory and includes the project in ``projects.json``.
To use it, build it using go
⚠️ Make sure to change ``prefix`` var in ``createproject.go`` according to your system
```shell
go build -o createproject ./createproject.go
```

## up, down and restart
These scripts are shortcuts for commands 
```bash
docker-compose up -d
docker-compose down
```
Simply run ``up`` or ``down`` to run the corresponding command.
If there is not any ``docker-compose.yml`` file in the current directory, the script will just abort.

Restart uses both ``down`` and ``up`` to fully recreate compose containers.
