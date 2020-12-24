[![Build Status](https://travis-ci.org/massicer/Oh-My-Gate-IOT-executor.svg?branch=main)](https://travis-ci.org/massicer/Oh-My-Gate-IOT-executor)


# Oh My Gate Executor 

## Pipelines:
This repo has a set of pipeline in order to produce artifacts in an automatic way: [Docker hub image](https://hub.docker.com/repository/docker/massicer/oh-my-gate-iot-executor])

- on branch != main an image with tag `dev` is published
- on branch == main an image with tag `latest` is pubblished
- on tag an image with `tag-name` is published


## How to release
- `git checkout master`
- `git tag {your tag}`
- `git push --follow-tags`