# Design Proposal Testmachinery V2

The new testmachinery design should fit into new requirements and to support the following workflow.
All requirements and needs are chosen to satify such a workflow and to just extend the current testmachinery and not rewrite everything.

<p align="center">
  <img alt= "testrunner overview" src="https://github.com/gardener/test-infra/raw/v2/docs/v2/V2Workflow.png">
</p>

## Requirements / Needs

* Nested Testruns
* Testrun namespaces for resources and environment like shared folder or kubeconfigs
* Inheritance of resources form parent templates
* TestrunTemplates or Testruns that are not executed (ignore annotation)
* different locations | versions for specific runs

## Redesign

Modules and elements that need to be redesigned:
* Kubeconfigs from secret for more security

##  Proposal

### Testrun
A new testrun yaml could introduce a new testflow type called `template` or `testrun` besides the old types `name` and `label`.
Then there other templates can referenced if these templates are already deployed on the cluster as CRD.

A new testflow could look like this(without any configuration).
It is also planned that templates can have configuration like other testdefinitions.
```yaml
testflow:
- - name: create-gardener
- - template: create-aws-shoot
  - template: create-gcp-shoot
- - name: upgrade
- - template: test-aws-shoot
  - template: test-gcp-shoot
  - template: create-and-test-aws-shoot
  - template: create-and-test-gcp-shoot
- - name: delete-gardener
```

And the status like this:
```yaml
status:
    steps:
    - - name: create-gardener

    - - template: create-aws-shoot
        steps:
        - name: create-shoot
        - label: test...
        - name: delete shoot
        -
      - template: create-gcp-shoot
        steps:
        - name: create-shoot

    - - name: upgarde

    - - template: test-aws-shoot
        steps:
        - label: test...
        - name: delete shoot
      - template: test-gcp-shoot
        steps:
        - label: test...
        - name: delete shoot

      - template: create-and-test-aws-shoot
        steps:
        - name: create-shoot
        - label: test...
        - name: delete shoot

      - template: create-and-test-gcp-shoot
        steps:
        - name: create-shoot
        - label: test...
        - name: delete shoot

    - - name: delete-gardener
```