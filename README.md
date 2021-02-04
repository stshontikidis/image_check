# Image Version Checker


This is a little piece of CI tooling I wrote as a way to start learning Go.  This could have easily been a python script and there is not much leverage of what makes Go great, it is pretty straight forward stuff.  I just wanted a way to know if a docker image was updated to kick off a build elsewhere.

### What it does
This will poll a repisitory on docker hub registry to pull latest manifest for a given tag and compare it against the last time it was ran, in essence if there is a change to the image.  It will then send a `workflow_dispatch` event to a github repo. I currently capture that event and rebuild an image via github actions that is based off the image I am polling.

## Usage
I think docker is probably the best way to implement this as I want to keep the process running indefinitely and docker these days is easier than to daemonize.  You will need to create a digest file to mount, this will act as permanent store for the current digest of the image, that way it can persist reboots and whatnot.

```bash
touch /path/to/digest.out
```
Be careful about not using a suffix as docker seems to always want to mount things as dirs. Might clean that up at some point

```bash
docker pull ghcr.io/stshontikidis/image_check
docker run -d \
    -e GH_REF=master \
    -e GH_REPO=stshontikidis/nextcloud_preview \
    -e GH_TOKEN=<PAT for repo> \
    -e GH_WORKFLOW=5323441 \
    -e DOCKER_REPO=library/nextcloud \
    -e DOCKER_TAG=stable \
    --name image_check ghcr.io/stshontikidis/image_check
```

Or you better yet, you should probably be using compose

```yaml
version: "3"

services:
  poller:
    image: ghcr.io/stshontikidis/image_check
    volumes:
      - /path/to/digest:/tmp/digest
    env_file:
      - .env
```

### Environment Variables
| Var | Description |
| --- | --- |
| `DOCKER_REPO` | The repo you want to poll on docker hub registry |
| `DOCKER_TAG` | The tag for image you are polling the status of |
| `GH_REPO` | The repo on github you want to send the workflow action to. |
| `GH_REF` | A The git reference for the workflow. The reference can be a branch or tag name.|
| `GH_WORKFLOW` | The ID or name of the workflow. You can find the ID using an [API](https://docs.github.com/en/rest/reference/actions#list-repository-workflows) via curl. |
| `GH_TOKEN` | [PAT](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) with repo scope for given github repo. |


Checking out the [workflow_dispatch](https://docs.github.com/en/rest/reference/actions#create-a-workflow-dispatch-event) API might help clearify smoe of the envs.

