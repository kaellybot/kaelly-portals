# kaelly-portals 

[![CI](https://github.com/kaellybot/kaelly-portals/actions/workflows/ci.yml/badge.svg)](https://github.com/kaellybot/kaelly-portals/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kaellybot/kaelly-portals/branch/master/graph/badge.svg)](https://codecov.io/gh/kaellybot/kaelly-portals) 


Application to retrieve dimension portals from different sources, written in Go

This module is no longer relevant since [Osamodas dimension feature](https://youtu.be/B9RnkZAOdPM?si=8Pt3oOqQCiUlRWpW&t=157); portals are accessible through the Traveler tower after one personal hunt (one per dimension).

## Current supported sources

- [dofus-portals](https://dofus-portals.fr)

## Generate client boilerplate

```Bash
# CLI Installation
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Examples
oapi-codegen -package payloads -generate types,client,spec openapi.yaml > openapi.gen.go
oapi-codegen -package dofusportals -generate types,client,spec payloads/dofusportals/openapi.yaml > payloads/dofusportals/openapi.gen.go
```