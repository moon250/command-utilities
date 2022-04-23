# Command Utilities
This repository contains some scripts that I have created to help me in my development environment. Note that I use zsh so you may need to modify some scripts to fit with your needs.

## openproject
This script is used to open a project, wherever I am. Just edit ``projectsFile`` var to link to a json file containing your projects.
```json
{
  "portfolio": "~/dev/portfolio",
  "secondproject": "~/dev/projectTwo"
}
```

Simply run 
```bash
$ openproject portfolio
```
to be redirected to ``~/dev/portfolio``
