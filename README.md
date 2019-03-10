# Go Report Card Drone Plugin
[![Build Status](https://drone.lddsb.com/api/badges/lddsb/drone-goreportcard/status.svg)](https://drone.lddsb.com/lddsb/drone-goreportcard) 

[Drone CI](https://drone.io) plugin to auto refresh your project status in [go report card](https://goreportcard.com)
### Usage

```yaml
kind: pipeline
name: default

steps:
...
- name: goreportcard
  pull: always
  image: lddsb/drone-goreportcard
  
```