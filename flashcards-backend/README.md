# Flashcards backend

Edit ent/schema/\* for db models. Then run `go generate ./ent`

Edit graph/\*.gql for graphql. Then run the `go generate ./graph`

https://jmeubank.github.io/tdm-gcc/download/ required to run tests in windows

# Dockerfile

Taken from https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/

## Dockerfile requires DOCKER_BUILDKIT

On Linux, macOS, or using WSL 2 you can do this using the following command:
`$ export DOCKER_BUILDKIT=1`

On Windows for PowerShell you can use:
`\$env:DOCKER_BUILDKIT=1`

Or for command prompt:
`set DOCKER_BUILDKIT=1`
