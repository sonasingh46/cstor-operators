# cstor-operators
[![Go](https://github.com/openebs/cstor-operators/workflows/Go/badge.svg)](https://github.com/openebs/cstor-operators/actions)
[![Build Status](https://travis-ci.org/openebs/cstor-operators.svg?branch=master)](https://travis-ci.org/openebs/cstor-operators)
[![Slack](https://img.shields.io/badge/JOIN-SLACK-blue)](https://openebs-community.slack.com)

<img width="300" align="right" alt="OpenEBS Logo" src="https://raw.githubusercontent.com/cncf/artwork/master/projects/openebs/stacked/color/openebs-stacked-color.png" xmlns="http://www.w3.org/1999/html">

Collection of OpenEBS cStor Data Engine Operators. The data engine operators
are responsible for cStor pool and volume orchestration.


## Project Status

This project is under active development and is considered to be in beta state.

The current implementation supports the following for cStor pools and volumes:
1. Provisioning and De-provisioning of cStor pools.
2. Pool expansion by adding disk.
3. Disk replacement by removing a disk.
4. Volume replica scale up and scale down.
5. Volume resize.

The data engine operators works in conjunction with the cStor CSI driver to finally
provide a consumable volume for stateful workloads.

For more details on cStor-CSI drivers, please refer [here](https://github.com/openebs/cstor-csi)

Table of contents:
==================
- [Quickly deploy it on K8s and get started](docs/quick.md)
- [Pool Operations Tutorial](docs/tutorial/intro.md)
- [High level overview](#overview)
- [Issues and PRs](#issues-and-prs)
- [Contributing](#contributing)
- [Code of conduct](#code-of-conduct)

## Overview

At a high-level, cstor operators consists of following components.
- cspc-operator
- pool-manager
- cvc-operator
- volume-manager

An OpenEBS admin/user can use CSPC(CStorPoolCluster) API (YAML) to provision cStor pools in a Kubernetes cluster.
As the name suggests, CSPC can be used to create a cluster of cStor pools across Kubernetes nodes.
It is the job of **cspc-operator** to reconcile the CSPC obejct and provision CStorPoolInstance(s) as specified 
in the CSPC. A cStor pool is provisioned on node by utilising the disks attached to the node and is represented by 
CStorPoolInstance(CSPI) custom resource in a Kubernetes cluster. One has freedom to specify the disks that they
want to use for pool provisioning.

CSPC API comes with a variety of tunables and features and the API can be viewed for [here](https://github.com/openebs/api/blob/master/pkg/apis/cstor/v1/cstorpoolcluster.go)

Once a CSPC is created, cspc-operator provision CSPI CR and **pool-manager** deployment on each nodes where cStor pool should 
be created. The pool-manager deployment watches for its corresponding CSPI on the node and finally execute commands to
perform pool operations e.g pool provisioning.

**Note:** It is not recommended to modify the CSPI CR and pool-manager in the running cluster unless you know what you are 
trying to do. CSPC should be the only point of interaction.

## Issues And PRs
We consider issues also as a part of contribution to the project.
If you want to raise any issue for cstor-operators please do that at [openebs/openebs].
We are tracking issues for all the OpenEBS components at the same place.
If you are unsure about how to proceed, do not hesitate in communicating in 
OpenEBS community slack [channel]. 

See [contributing](#contributing) section to learn more about how to contribute to cstor-operators.


## Contributing

This [link](docs/developer-guide/start.md) contains all information about
building cstor-operators from source, contribution guide and whom to contact etc.

If you want to build the cstor-operators right away then the following is the step:

#### You have a working [Go environment] and [Docker environment].

```
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/openebs/cstor-operators
cd cstor-operators
make all.amd64
```

## Code of conduct

Please read the code of conduct [here](CODE_OF_CONDUCT.md)

[Docker environment]: https://docs.docker.com/engine
[Go environment]: https://golang.org/doc/install
[openebs/openebs]: https://github.com/openebs/openebs
[channel]: https://openebs-community.slack.com