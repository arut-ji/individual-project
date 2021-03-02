# Introduction

## Overview
[Overview of the project]

## Motivation
Infrastructure as Code (IaC) is an approach of managing, provisioning, and orchestrating computational 
re-sources using declarative domain-specific languages. Since such concepts describe infrastructure in 
a form of code, potential is shown that code smells could persist throughout its evolution overtime as 
in general purpose programming languages. Therefore, research to explore and define code smells is 
significant as it would help indicate problematic implementation - which could lead to infrastructure 
failure - in an early stage of software development lifecycle. In this study, Kubernetes is chosen to be 
explored regarding its extensive adoption in the industry. Furthermore, Kubernetes is a unique tool 
that is for container orchestration unlike ones in the previous studies in which they are for configuration 
management - Puppet, Chef, Ansible.

## Aims
This project will explore the existing catalogues of code smells in DevOps tooling - Puppet, Chef, Ansible, 
etc. - provided by the previous research. The smells will then be determined whether they are applicable 
with the Kubernetes manifests or not. Moreover, Kubernetes-specific best practices will be gathered and 
used to define additional potential smells.

