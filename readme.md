The created kubernetes files are currently narrowed down to my personal usecase. I'm very open for Changerequests to make this cli useful for a broader audiency. 

## Usage Example

To initialize a new project do
> goflux init --component myService

This will generate a basic folder structure.

To create all files needed for a backend deployment do
> goflux backend --component myService --namespace myNamespace