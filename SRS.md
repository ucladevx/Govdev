# Govdev

A highly configurable and scalable backend for managing user services.
Originally built to serve UCLA DevX's application portal needs. 
This document
serves as the software requirements specification needed for our club. It does
not suffice as a technical spec, rather one that specifies what Govdev should
functionally do.

- [Govdev](#govdev)
  * [Introduction](#introduction)
    + [Purpose](#purpose)
    + [Intended Audience](#intended-audience)
    + [Intended Use](#intended-use)
    + [Scope](#scope)
    + [Definitions/Acronyms](#definitions-acronyms)
  * [Description](#description)
    + [User needs](#user-needs)
  * [System Features/Requirements](#system-features-requirements)
    + [Functional Requirements](#functional-requirements)
    + [Nonfunctional Requirements](#nonfunctional-requirements)
  * [Change Management](#change-management)
  * [Document Approvals](#document-approvals)
  * [Supporting Information](#supporting-information)

## Introduction

### Purpose

This document
serves as the software requirements specification needed for our club. It does
not suffice as a technical spec, rather one that specifies what Govdev should
functionally do.

### Intended Audience

We have three main intended audiences: 

* DevX applicants
* DevX members
* DevX admins

### Intended Use

* DevX applicants:
    * DevX applicants will use this application portal to apply (quarterly) to
      get into DevX. They can create a profile, submit a resume, and answer
      questions in our applications.
* DevX members:
    * DevX members should belong to teams and be able to manage their personal
      profiles. Members might want to mark themselves as active or inactive,
      etc.
* DevX Admins:
    * DevX admins will want to be able to review applications, accept or reject
      applications and send out the corresponding emails, manage DevX member
      accounts, and manage content on the website. 
    * There is also a subset of admins that are owners, who have all admin
      privileges, but can also create other admins/owners. 


### Scope 

The scope of this project lies with managing membership applications, managing
DevX users, and user profiles. This does not deal specifically with DevX
teams and their projects, although links may be created to those projects.

### Definitions/Acronyms

## Description

### User needs

* DevX applicants

DevX applicants will require the ability to create a user profile, log in, fill
out questions, upload a resume/profile picture, and submit the application. Once
submitted, we can also let them know about the application results in the
application portal. Moreover, these results will be emailed to the user as well.
The application portal will give them instructions as to next steps (any further
interviews or going to demo night, etc).

* DevX Members

DevX members are applicants who have been accepted into DevX. Their main ability
is to be able to change their profile and customize it to their liking (profile
pic, description, etc). DevX members might also want to be able to mark
themselves as inactive or active.

* DevX Admins

The bulk of the work should go into the admin page. Admins will require a page
that allows them to manage user profiles (editing/review profiles to make sure
they are appropriate) and also manage applications. 

On the reviewing applications page, admins should be able to see a list of all
applicants in row format. Each row is split into different sections, such as
Name, year, age, role, resume, description, spotlight answers. Admins should be
able to filter applications by name, year, role, etc. Additionally, admins will
have the ability to score each entry (a number out of 5). It should be possible
to show the average score submitted by all the reviewers. Lastly, admins should
be able to choose accept or reject for an applicant.

For the applicants that are accepted or rejected, admins should be able to
trigger emails notifying the applicants of the decision. Admins should be able
to review/edit the email template before sending all the emails.

In regards to other users, admins should be able to edit or delete other users
from the system. Admins should also be able to send emails to all active DevX
users, eliminating the need to create a revolving email list. 

Lastly, there is a subset of admins denoted as owners, who can create other
admins and owners. Owners have all other abilities of admins.

## System Features/Requirements

### Functional Requirements

### Nonfunctional Requirements

* Sessions

Users should be logged in, and once they are logged in, the frontend will be
able to keep them logged in with sessions and refresh tokens.

* Security

This application should be secure, use HTTPS, and use general security
practices, such as not storing passwords in plain text.

* Performance

This application should be able to have decent performance, and not lag badly.
Doesn't have to be amazing, but no Chrome turtles allowed.

* Persistence

This application should be able to restart without any problems. This implies
the existance of some persistance storage. Application should be stateless.

## Change Management

Change management should be handled through communication and agreed upon
changes. Changes should be logged in a change log, and tagged with versions
(using semantic versioning).

## Document Approvals

Approved by Executive Board and primary developers of this application.

## Supporting Information

Initially written by Terrence Ho.

For a more technical document, view the [Technical Design Doc](ARCHITECTURE.md)
