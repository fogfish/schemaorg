# schema.org for golang

The library declares Go types of https://schema.org ontology.

[![Version](https://img.shields.io/github/v/tag/fogfish/schemaorg?label=version)](https://github.com/fogfish/schemaorg/releases)
[![Documentation](https://pkg.go.dev/badge/github.com/fogfish/schemaorg)](https://pkg.go.dev/github.com/fogfish/schemaorg)
[![Git Hub](https://img.shields.io/github/last-commit/fogfish/schemaorg.svg)](https://github.com/fogfish/schemaorg)


## Inspiration

> Schema.org is a collaborative, community activity with a mission to create, maintain, and promote schemas for structured data on the Internet, on web pages, in email messages, and beyond.

The common vocabulary layouts content into data structure interpretable by human and machines. It also guarantees an interoperability baseline for software in the distribute systems, in the absence of strong content negotiation techniques. Despite, the scheme.org elevates creation of web structured data markup schema, it also defines an ontology usable for data design in other context such as public RESTful interfaces, gRPC or GraphQL.

A core principle of the schema.org vocabulary is to indicate an entity type. As an example, the vocabulary defines a rich collection of `rdf:Property`, which are essential building blocks for product types.

```go
type Person struct {
  schemaorg.FamilyName `json:"schema:familyName"`
}
```


## Get started

The latest version of the library is available at its `main` branch. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

The library uses semver schema, the **minor version** identifies the https://schema.org release used to generate types. For example, the version v1.14.0 of the library is built from release 14.0. The **patch version** is just a sequential number, increased every time this library is fixed or updated, it do not have any correlation with https://schema.org release.  
  
**The minor version number** 

## Core domain types (product types)

Ontologies use `rdf:Property` to depict characteristic of entities, which differs significantly from Golang type system. Instead of defining product type in terms of named properties, ontologies describe properties in terms of types to which they apply, which allows to extends concepts, defining additional properties without the need to re-define the original description of types. 

`rdf:Property` is a relation between subject resources and object resources. The concept of property at Go programming languages does not correspond to the `rdf:Property`, with an exception that "properties" of structs has names and corresponding types.

```go
/*
The statement 
  Person schema:familyName "Doe"
*/
type Person struct {
  schemaorg.FamilyName `json:"schema:familyName"`
}
```

Type-level meta-programming helps us better reflect the `rdf:Property` concept but type-level programming is not supported at Golang. Therefore, this library reflects the `rdf:Property` to the type.

```go
/*
{
  "@id": "schema:familyName",
  "@type": "rdf:Property",
  "rdfs:comment": "Family name. In the U.S., the last name of a Person.",
  "rdfs:label": "familyName",
  ...
}
*/
type FamilyName string
```

So that, the application "assembles" core domain type as a product of the types defined by schema.org

```go
import "github.com/fogfish/schemaorg"

type Person struct {
  schemaorg.Identifier `json:"schema:@id"`
  schemaorg.GivenName  `json:"schema:givenName"`
  schemaorg.FamilyName `json:"schema:familyName"`
  schemaorg.Email      `json:"schema:email"`
}
```

## How To Contribute

The library is auto-generated from schema.org JSON-LD definition using [schemacli](https://github.com/fogfish/schemacli) utility.

Use the following commands to re-generate the type definitions

```bash
schemacli property -f schemaorg-all-http.jsonld > types.go
```

If you experience any issues with the library or requires a new feature, please let us know via [GitHub issues](https://github.com/fogfish/schemaorg/issue).


## License

[![See LICENSE](https://img.shields.io/github/license/fogfish/schemaorg.svg?style=for-the-badge)](LICENSE)
