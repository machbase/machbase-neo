---
layout: page
title: Releases
permalink: /releases/
nav_order: 2
parent: Welcome to machbase
---

# Releases

## Latest version {{ site.latest_version }}

Choose the latest version for your platform.

| OS         | Architecture   |  Download |
|:-----------|:---------------|:----------|
| Linux      | arm64          | [machbase-neo-{{ site.latest_version }}-linux-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-linux-arm64.zip)   |
| Linux      | x64            | [machbase-neo-{{ site.latest_version }}-linux-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-linux-amd64.zip)   |
| Linux      | arm32          | [machbase-neo-{{ site.latest_version }}-linux-arm32.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-linux-arm32.zip)   |
| macOS      | arm64          | [machbase-neo-{{ site.latest_version }}-darwin-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-darwin-arm64.zip) |
| macOS      | x64            | [machbase-neo-{{ site.latest_version }}-darwin-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-darwin-amd64.zip) |
| Windows[^1] | x64     | [machbase-neo-{{ site.latest_version }}-windows-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/machbase-neo-{{ site.latest_version }}-windows-amd64.zip)[^2] |


<!--
### GUI Launcher for macOS

The (_experimental_) GUI releases for macOS users.

| OS         | Architecture   |  Download |
|:-----------|:---------------|:----------|
| macOS      | Apple          | [neow-{{ site.latest_version }}-darwin-arm64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/neow-{{ site.latest_version }}-darwin-arm64.zip)|
| macOS      | Intel          | [neow-{{ site.latest_version }}-darwin-amd64.zip]({{ site.releases_url }}/download/{{ site.latest_version }}/neow-{{ site.latest_version }}-darwin-amd64.zip)|
-->

## What's Changed {{ site.latest_version }}

[Changes](https://github.com/machbase/neo-server/releases/tag/{{ site.latest_version }})

## Previous releases

Find previously released versions in [here](https://github.com/machbase/neo-server/releases).

{: .note-title}
> Edge vs. Fog in the previsous releases before v1.5.0
>
> If you plan to run Machbase-neo on a small device such as Raspberry Pi, select the Edge edition.
> For machines with larger memory and more CPU cores, such as a personal workstation or server, choose the Fog edition.
>
> Since v1.5.0, the editions are integrated into the single "standard" edition.


--------------

[^1]: Windows Fall 2018 or newer versions.
[^2]: Windows release includes both of the `machbase-neo` and `neow` (GUI frontend) executables.
{: .fs-1}
