The created kubernetes files are currently narrowed down to my personal usecase. I'm very open for Changerequests to make this cli useful for a broader audiency. 

## Usage Example

```sh
NAME:
   goflux - Used to automatically generate flux files

USAGE:
   goflux [global options] command [command options] [arguments...]

COMMANDS:
   init     Initialize new project
   backend  Create files for a backend service
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

To initialize a new project do
> goflux init --component myService

This will generate a basic folder structure.

To create all files needed for a backend deployment do
> goflux backend --component myService --namespace myNamespace