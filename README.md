<h1 align="center">
  <img width="300px" src=".github/assets/corteza_logo.svg" />
  <br />
  <p>Corteza Server</p>

  <div align="center">

  [![DockerHub Downloads Card](https://img.shields.io/docker/pulls/cortezaproject/corteza-server)](https://img.shields.io/docker/pulls/cortezaproject/corteza-server)
  [![Latest Version Card](https://img.shields.io/github/v/tag/cortezaproject/corteza-server?label=stable%20version)](https://img.shields.io/github/v/tag/cortezaproject/corteza-server?label=stable%20version)
  [![License Card](https://img.shields.io/github/license/cortezaproject/corteza-server)](https://img.shields.io/github/license/cortezaproject/corteza-server)

  [![Go Report Card](https://goreportcard.com/badge/github.com/cortezaproject/corteza-server)](https://goreportcard.com/report/github.com/cortezaproject/corteza-server)
  [![Build Status](https://drone.crust.tech/api/badges/cortezaproject/corteza/status.svg)](https://drone.crust.tech/cortezaproject/corteza)
  [![CodeCov Report Card](https://img.shields.io/codecov/c/github/cortezaproject/corteza-server)](https://img.shields.io/codecov/c/github/cortezaproject/corteza-server)

  </div>
</h1>

<br />

**Corteza server** is the **back-end** of the Corteza ecosystem.
The core logic is written in GO, using [go-chi](https://pkg.go.dev/github.com/go-chi/chi@v3.3.4+incompatible?utm_source=gopls) for the routing.
External communication occurs via the **REST API and web sockets**; internal communication happens via gRPC.

Non-binary data is handled by the store layer and stored inside a database; currently **MySQL and PostgreSQL**.
We will add support for additional databases in the future.
Binary data is handled locally or on a supported cloud provider using [MinIO](https://min.io/).

<br />

## What is Corteza?

<div align="center">
  <img style="max-height: 350px;" src=".github/assets/corteza_dashboard.png" />
</div>

Corteza is the only **100% free**, **open-source**, **standardized** and **enterprise-grade** Low-code platform.
It is developed entirely in the public domain and maintained by [Crust Technology](https://www.crust.tech/), its founder.

With Corteza, you can quickly **build scalable cloud applications** that are **integrable** with external services and **accessible (WCAG 2.1)**.

### Core features:

* quick setup,
* flexible and intuitive low-code configuration,
* powerful automation system using workflows and automation scripts,
* flexible reporting capabilities,
* secure RBAC access control system.

## Online demo

You can check out Corteza online by creating an account on our community instance https://latest.cortezaproject.org.

## Deploying Corteza

Refer to the [DevOps guide](https://docs.cortezaproject.org/corteza-docs/2021.3/devops-guide/index.html) for a complete guide on how to get Corteza up and running.
Additionally, we've provided some [video instructions](https://forum.cortezaproject.org/t/videos-on-how-to-set-up-corteza/91).

Quick references:

* [data backup and restore](https://docs.cortezaproject.org/corteza-docs/2021.3/devops-guide/maintenance/backups.html)
* [troubleshooting](https://docs.cortezaproject.org/corteza-docs/2021.3/devops-guide/maintenance/troubleshooting.html)
* [setting up an email relay](https://docs.cortezaproject.org/corteza-docs/2021.3/devops-guide/extension-requirements/email-relay.html)
* [setting up sink routes](https://docs.cortezaproject.org/corteza-docs/2021.3/devops-guide/extension-requirements/sink-route.html)

## Upgrading Corteza

[![Latest Version Card](https://img.shields.io/github/v/tag/cortezaproject/corteza-server?label=latest%20stable%20version)](https://img.shields.io/github/v/tag/cortezaproject/corteza-server?label=latest%20stable%20version)

Refer to the [changelog](https://docs.cortezaproject.org/corteza-docs/2021.3/changelog/index.html) and the [upgrade guide](https://docs.cortezaproject.org/corteza-docs/2021.3/upgrade-guide/index.html) to upgrade your Corteza instance.

## Using Corteza

Refer to the [End-User Guide](https://docs.cortezaproject.org/corteza-docs/2021.3/end-user-guide/index.html) to learn the built-in applications and features from the end-user perspective.

Quick references:

* [Corteza CRM](https://docs.cortezaproject.org/corteza-docs/2021.3/end-user-guide/crm/index.html)
* [Corteza Service Solution](https://docs.cortezaproject.org/corteza-docs/2021.3/end-user-guide/service-solution/index.html)

## Create with Corteza

Refer to the [Integrator Guide](https://docs.cortezaproject.org/corteza-docs/2021.3/integrator-guide/index.html) to learn how you can build on the core features to create virtually anything.

Quick references:

* [Corteza Compose configuration](https://docs.cortezaproject.org/corteza-docs/2021.3/integrator-guide/compose-configuration/index.html)
* automation using [workflows](https://docs.cortezaproject.org/corteza-docs/2021.3/integrator-guide/workflows/index.html) and [automation scripts](https://docs.cortezaproject.org/corteza-docs/2021.3/integrator-guide/automation-scripts/index.html)
* [using the REST API](https://docs.cortezaproject.org/corteza-docs/2021.3/integrator-guide/accessing-corteza/index.html),

## Contributing

Refer to the [Developer Guide/Corteza Server](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/index.html) document for details regarding the [development setup](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/index.html#_development_setup), the [project structure](https://docs.cortezaproject.org/corteza-docs/2021.3/developer-guide/corteza-server/structure.html), and the feature insight documents.

Refer to the [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines and code of conduct.

Quick command reference:

* run server: `make watch` or `make realize`
* run unit tests: `make test.unit`
* run integration tests: `make test.integration`
* run all tests `make test.all`
* run MailHog (local SMTP server): `make mailhog.up` (requires Docker)

## Community

Reach out to us on [our forum](https://forum.cortezaproject.org/).

## License

Corteza is released under the Apache-2.0 license.
Refer to the [LICENSE](LICENSE) file for additional information.
