---
layout: page
title: Releases
permalink: /releases/
nav_order: 2
parent: Welcome to machbase
---

# Releases

{: .note-title}
> Edge vs. Fog
>
> If you plan to run Machbase-neo on a small device such as Raspberry Pi, select the Edge edition.
> For machines with larger memory and more CPU cores, such as a personal workstation or server, choose the Fog edition.

## Latest version {{ site.latest_version }}

### Fog Edition

The Fog edition is for desktop and server-grade machines.

| OS         | Architecture   |  Download |
|:-----------|:---------------|:----------|
| Linux      | arm64          | [machbase-neo-fog-{{ site.latest_version }}-linux-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-fog-{{ site.latest_version }}-linux-arm64.zip)   |
| Linux      | x64            | [machbase-neo-fog-{{ site.latest_version }}-linux-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-fog-{{ site.latest_version }}-linux-amd64.zip)   |
| macOS      | arm64          | [machbase-neo-fog-{{ site.latest_version }}-darwin-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-fog-{{ site.latest_version }}-darwin-arm64.zip) |
| macOS      | x64            | [machbase-neo-fog-{{ site.latest_version }}-darwin-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-fog-{{ site.latest_version }}-darwin-amd64.zip) |


### Edge Edition

The Edge edition is optimized for small machines with limited resources, such as the Raspberry Pi.

| OS         | Architecture   |  Download |
|:-----------|:---------------|:----------|
| Linux      | arm64          | [machbase-neo-edge-{{ site.latest_version }}-linux-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-edge-{{ site.latest_version }}-linux-arm64.zip)   |
| Linux      | x64            | [machbase-neo-edge-{{ site.latest_version }}-linux-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-edge-{{ site.latest_version }}-linux-amd64.zip)   |
| Linux      | armv6l, armv7l | [machbase-neo-edge-{{ site.latest_version }}-linux-arm32.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-edge-{{ site.latest_version }}-linux-arm32.zip)   |
| macOS      | arm64          | [machbase-neo-edge-{{ site.latest_version }}-darwin-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-edge-{{ site.latest_version }}-darwin-arm64.zip) |
| macOS      | x64            | [machbase-neo-edge-{{ site.latest_version }}-darwin-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-edge-{{ site.latest_version }}-darwin-amd64.zip) |

## What's Changed {{ site.latest_version }}

### 🎉 Features
* inputstream spi by @OutOfBedlam in https://github.com/machbase/neo-server/pull/5
*  missing machbase bind host machbase/neo#17 by @OutOfBedlam in https://github.com/machbase/neo-server/pull/6
* fix missing elapse by @OutOfBedlam in https://github.com/machbase/neo-server/pull/7
* msg rearrange by @OutOfBedlam in https://github.com/machbase/neo-server/pull/8
* Refactoring neo-shell by @OutOfBedlam in https://github.com/machbase/neo-server/pull/9
* ssh key based user auth machbase/neo#13 by @OutOfBedlam in https://github.com/machbase/neo-server/pull/10
* linux arm32 machbase/neo#22 by @OutOfBedlam in https://github.com/machbase/neo-server/pull/11

**Full Changelog**: https://github.com/machbase/neo-server/compare/v0.2.2...v0.3.0

## Previous releases

Find previously released versions in [here](https://github.com/machbase/neo-server/releases).