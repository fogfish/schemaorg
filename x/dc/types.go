//
// Copyright (C) 2023 - 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/schemaorg
//

package dc

import "time"

// An entity responsible for making contributions to the resource.
//
// http://purl.org/dc/elements/1.1/contributor
//
// Examples of a Contributor include a person, an organization, or a service.
// Typically, the name of a Contributor should be used to indicate the entity.
type Contributor string

// The spatial or temporal topic of the resource, the spatial applicability of
// the resource, or the jurisdiction under which the resource is relevant.
//
// http://purl.org/dc/elements/1.1/coverage
//
// Spatial topic and spatial applicability may be a named place or a location
// specified by its geographic coordinates. Temporal topic may be a named period,
// date, or date range. A jurisdiction may be a named administrative entity or
// a geographic place to which the resource applies. Recommended best practice
// is to use a controlled vocabulary such as the Thesaurus of Geographic Names
// [TGN]. Where appropriate, named places or time periods can be used in
// preference to numeric identifiers such as sets of coordinates or date ranges.
type Coverage string

// An entity primarily responsible for making the resource.
//
// http://purl.org/dc/elements/1.1/creator
//
// Examples of a Creator include a person, an organization, or a service.
// Typically, the name of a Creator should be used to indicate the entity.
type Creator string

// A point or period of time associated with an event in the lifecycle of the resource.
//
// http://purl.org/dc/elements/1.1/date
//
// Date may be used to express temporal information at any level of granularity.
// Recommended best practice is to use an encoding scheme, such as the W3CDTF
// profile of ISO 8601 [W3CDTF].
type Date time.Time

// An account of the resource.
//
// http://purl.org/dc/elements/1.1/description
//
// Description may include but is not limited to: an abstract, a table of
// contents, a graphical representation, or a free-text account of the resource.
type Description string

// The file format, physical medium, or dimensions of the resource.
//
// http://purl.org/dc/elements/1.1/format
//
// Examples of dimensions include size and duration. Recommended best practice
// is to use a controlled vocabulary such as the list of Internet Media Types [MIME].
type Format string

// An unambiguous reference to the resource within a given context.
//
// http://purl.org/dc/elements/1.1/identifier
//
// Recommended best practice is to identify the resource by means of a string
// conforming to a formal identification system.
type Identifier string

// A language of the resource.
//
// http://purl.org/dc/elements/1.1/language
//
// Recommended best practice is to use a controlled vocabulary such as RFC 4646 [RFC4646].
type Language string

// An entity responsible for making the resource available.
//
// http://purl.org/dc/elements/1.1/publisher
//
// Examples of a Publisher include a person, an organization, or a service.
// Typically, the name of a Publisher should be used to indicate the entity.
type Publisher string

// A related resource.
//
// http://purl.org/dc/elements/1.1/relation
//
// Recommended best practice is to identify the related resource by means of a
// string conforming to a formal identification system.
type Relation string

// Information about rights held in and over the resource.
//
// http://purl.org/dc/elements/1.1/rights
//
// Typically, rights information includes a statement about various property
// rights associated with the resource, including intellectual property rights.
type Rights string

// A related resource from which the described resource is derived.
//
// http://purl.org/dc/elements/1.1/source
//
// The described resource may be derived from the related resource in whole or
// in part. Recommended best practice is to identify the related resource by
// means of a string conforming to a formal identification system.
type Source string

// The topic of the resource.
//
// http://purl.org/dc/elements/1.1/subject
//
// Typically, the subject will be represented using keywords, key phrases, or
// classification codes. Recommended best practice is to use a controlled vocabulary.
type Subject string

// A name given to the resource.
//
// http://purl.org/dc/elements/1.1/title
//
// Typically, a Title will be a name by which the resource is formally known.
type Title string

// The nature or genre of the resource.
//
// http://purl.org/dc/elements/1.1/type
//
// Recommended best practice is to use a controlled vocabulary such as the DCMI
// Type Vocabulary [DCMITYPE]. To describe the file format, physical medium, or
// dimensions of the resource, use the Format element.
type Type string
