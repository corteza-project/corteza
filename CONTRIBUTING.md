# Contributing

Thank you for helping us make [our vision](https://cortezaproject.org/about/what-is-corteza/) a reality!
All contributions are welcome; from bug reports, codefixes, and new features!

## Ground Rules

Corteza projects are [Apache 2.0 licensed](LICENSE) and accept contributions via GitHub pull requests.

Cover the [terminology](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/release-cycle/index.html#_terminology) for the development process and versioning.

Cover the [Git and GitHub](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/release-cycle/index.html#_github) ground rules regarding branch naming and conventions.

When you wish to start working on a code contribution, assign yourself to a GitHub issue.
If there is no issue, create one beforehand.

A quick summary on [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/):

1. Separate subject from body with a blank line
2. Limit the subject line to 50 characters
3. Capitalize the subject line
4. Do not end the subject line with a period
5. Use the imperative mood in the subject line
6. Wrap the body at 72 characters
7. Use the body to explain what and why vs. how

## Core repositories

### Corteza Server

Corteza server is the back-end of the Corteza ecosystem.
The core logic is written in GO, using [go-chi](https://pkg.go.dev/github.com/go-chi/chi@v3.3.4+incompatible?utm_source=gopls) for the routing.

Communication between the Corteza server and web applications is done using the REST API and web sockets.
Communication between back-end services (Corteza server and Corredor) is done using gRPC.

The [Developer Guide/Corteza Server](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/index.html) covers the [development setup](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/index.html#_development_setup), the [project structure](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/structure.html), and the feature insight documents.


### Corteza Web Applications

The web applications are written in Vue.js and provide the user interface to interact with the entire system.
The repositories:

1. [corteza-webapp-one](https://github.com/cortezaproject/corteza-webapp-one)
2. [corteza-webapp-admin](https://github.com/cortezaproject/corteza-webapp-admin)
3. [corteza-webapp-compose](https://github.com/cortezaproject/corteza-webapp-compose)
4. [corteza-webapp-workflow](https://github.com/cortezaproject/corteza-webapp-workflow)

Communication between the Corteza server and web applications is done using the REST API and web sockets.

The [Developer Guide/Corteza Web Applications](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/web-applications/index.html) covers the [development setup](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/web-applications/index.html#_development_setup), the [project structure](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/web-applications/structure.html), and the feature insight documents.

### Documentation

The documentation is written in [AsciiDoc](https://asciidoc.org/) and compiled using [Antora](https://antora.org/).
The source code is available on the [GitHub cortezaproject/corteza-docs repository](https://github.com/cortezaproject/corteza-docs); the generated output is available on the [documentation page](http://docs.cortezaproject.org/).

The [Developer Guide/Documentation](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/documentation/index.html) covers the [conventions](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/documentation/index.html#_conventions), [writing guidelines](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/documentation/index.html#documentation-writing-guidelines), as well as some [examples](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/documentation/examples/index.html) to help you get started.

## Bug reporting

Please submit any bug reports on the **issues** section of the corresponding GitHub repository.
If you are unsure where to submit the issue, or you are unsure if this is a feature; reach out to us on [our forum](https://forum.cortezaproject.org/).

## Feature requests

Feature and improvement requests should be submitted on [our forum](https://forum.cortezaproject.org/).
Before opening a new topic, search around to see if there are any similar topics that would cover your case.

## DCO

By contributing to this project you agree to the Developer Certificate of Origin (DCO).
This document was created by the Linux Kernel community and is a simple statement that you, as a contributor, have the legal right to make the contribution.
See the [DCO](DCO) file for details.
