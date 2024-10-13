package schemaorg

/*

DateSent is https://schema.org/dateSent

The date/time at which the message was sent.
*/
type DateSent string

/*

PaymentStatus is https://schema.org/paymentStatus

The status of payment; whether the invoice has been paid or not.
*/
type PaymentStatus string

/*

PermittedUsage is https://schema.org/permittedUsage

Indications regarding the permitted usage of the accommodation.
*/
type PermittedUsage string

/*

ScheduledTime is https://schema.org/scheduledTime

The time the object is scheduled to.
*/
type ScheduledTime string

/*

UploadDate is https://schema.org/uploadDate

Date (including time if available) when this media object was uploaded to this site.
*/
type UploadDate string

/*

AddressCountry is https://schema.org/addressCountry

The country. For example, USA. You can also provide the two-letter [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1).
*/
type AddressCountry string

/*

BroadcastSignalModulation is https://schema.org/broadcastSignalModulation

The modulation (e.g. FM, AM, etc) used by a particular broadcast service.
*/
type BroadcastSignalModulation string

/*

Issn is https://schema.org/issn

The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
*/
type Issn string

/*

DeathDate is https://schema.org/deathDate

Date of death.
*/
type DeathDate string

/*

CvdNumVentUse is https://schema.org/cvdNumVentUse

numventuse - MECHANICAL VENTILATORS IN USE: Total number of ventilators in use.
*/
type CvdNumVentUse float64

/*

ProprietaryName is https://schema.org/proprietaryName

Proprietary name given to the diet plan, typically by its originator or creator.
*/
type ProprietaryName string

/*

SeatRow is https://schema.org/seatRow

The row location of the reserved seat (e.g., B).
*/
type SeatRow string

/*

Supply is https://schema.org/supply

A sub-property of instrument. A supply consumed when performing instructions or a direction.
*/
type Supply string

/*

CvdNumICUBedsOcc is https://schema.org/cvdNumICUBedsOcc

numicubedsocc - ICU BED OCCUPANCY: Total number of staffed inpatient ICU beds that are occupied.
*/
type CvdNumICUBedsOcc float64

/*

TaxID is https://schema.org/taxID

The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
*/
type TaxID string

/*

RestPeriods is https://schema.org/restPeriods

How often one should break from the activity.
*/
type RestPeriods string

/*

InChIKey is https://schema.org/inChIKey

InChIKey is a hashed version of the full InChI (using the SHA-256 algorithm).
*/
type InChIKey string

/*

CvdNumC19OFMechVentPats is https://schema.org/cvdNumC19OFMechVentPats

numc19ofmechventpats - ED/OVERFLOW and VENTILATED: Patients with suspected or confirmed COVID-19 who are in the ED or any overflow location awaiting an inpatient bed and on a mechanical ventilator.
*/
type CvdNumC19OFMechVentPats float64

/*

EducationalRole is https://schema.org/educationalRole

An educationalRole of an EducationalAudience.
*/
type EducationalRole string

/*

ArticleSection is https://schema.org/articleSection

Articles may belong to one or more 'sections' in a magazine or newspaper, such as Sports, Lifestyle, etc.
*/
type ArticleSection string

/*

EncodingFormat is https://schema.org/encodingFormat

Media type typically expressed using a MIME format (see [IANA site](http://www.iana.org/assignments/media-types/media-types.xhtml) and [MDN reference](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)), e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.

In cases where a [[CreativeWork]] has several media type representations, [[encoding]] can be used to indicate each [[MediaObject]] alongside particular [[encodingFormat]] information.

Unregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry.
*/
type EncodingFormat string

/*

DatePublished is https://schema.org/datePublished

Date of first publication or broadcast. For example the date a [[CreativeWork]] was broadcast or a [[Certification]] was issued.
*/
type DatePublished string

/*

ValidUntil is https://schema.org/validUntil

The date when the item is no longer valid.
*/
type ValidUntil string

/*

AmountOfThisGood is https://schema.org/amountOfThisGood

The quantity of the goods included in the offer.
*/
type AmountOfThisGood float64

/*

DiseasePreventionInfo is https://schema.org/diseasePreventionInfo

Information about disease prevention.
*/
type DiseasePreventionInfo string

/*

Skills is https://schema.org/skills

A statement of knowledge, skill, ability, task or any other assertion expressing a competency that is desired or required to fulfill this role or to work in this occupation.
*/
type Skills string

/*

EthicsPolicy is https://schema.org/ethicsPolicy

Statement about ethics policy, e.g. of a [[NewsMediaOrganization]] regarding journalistic and publishing practices, or of a [[Restaurant]], a page describing food source policies. In the case of a [[NewsMediaOrganization]], an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
*/
type EthicsPolicy string

/*

ConstraintProperty is https://schema.org/constraintProperty

Indicates a property used as a constraint. For example, in the definition of a [[StatisticalVariable]]. The value is a property, either from within Schema.org or from other compatible (e.g. RDF) systems such as DataCommons.org or Wikidata.org. 
*/
type ConstraintProperty string

/*

ValueReference is https://schema.org/valueReference

A secondary value that provides additional information on the original value, e.g. a reference temperature or a type of measurement.
*/
type ValueReference string

/*

MonthsOfExperience is https://schema.org/monthsOfExperience

Indicates the minimal number of months of experience required for a position.
*/
type MonthsOfExperience float64

/*

ProgramPrerequisites is https://schema.org/programPrerequisites

Prerequisites for enrolling in the program.
*/
type ProgramPrerequisites string

/*

HonorificPrefix is https://schema.org/honorificPrefix

An honorific prefix preceding a Person's name such as Dr/Mrs/Mr.
*/
type HonorificPrefix string

/*

UnnamedSourcesPolicy is https://schema.org/unnamedSourcesPolicy

For an [[Organization]] (typically a [[NewsMediaOrganization]]), a statement about policy on use of unnamed sources and the decision process required.
*/
type UnnamedSourcesPolicy string

/*

Genre is https://schema.org/genre

Genre of the creative work, broadcast channel or group.
*/
type Genre string

/*

Polygon is https://schema.org/polygon

A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
*/
type Polygon string

/*

DownloadUrl is https://schema.org/downloadUrl

If the file can be downloaded, URL to download the binary.
*/
type DownloadUrl string

/*

DefaultValue is https://schema.org/defaultValue

The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it's an ID reference to one of the current values.
*/
type DefaultValue string

/*

CodeValue is https://schema.org/codeValue

A short textual code that uniquely identifies the value.
*/
type CodeValue string

/*

BroadcastDisplayName is https://schema.org/broadcastDisplayName

The name displayed in the channel guide. For many US affiliates, it is the network name.
*/
type BroadcastDisplayName string

/*

LegislationIdentifier is https://schema.org/legislationIdentifier

An identifier for the legislation. This can be either a string-based identifier, like the CELEX at EU level or the NOR in France, or a web-based, URL/URI identifier, like an ELI (European Legislation Identifier) or an URN-Lex.
*/
type LegislationIdentifier string

/*

IsFamilyFriendly is https://schema.org/isFamilyFriendly

Indicates whether this content is family friendly.
*/
type IsFamilyFriendly bool

/*

DiscountCode is https://schema.org/discountCode

Code used to redeem a discount.
*/
type DiscountCode string

/*

Percentile75 is https://schema.org/percentile75

The 75th percentile value.
*/
type Percentile75 float64

/*

BroadcastSubChannel is https://schema.org/broadcastSubChannel

The subchannel used for the broadcast.
*/
type BroadcastSubChannel string

/*

ArchivedAt is https://schema.org/archivedAt

Indicates a page or other link involved in archival of a [[CreativeWork]]. In the case of [[MediaReview]], the items in a [[MediaReviewItem]] may often become inaccessible, but be archived by archival, journalistic, activist, or law enforcement organizations. In such cases, the referenced page may not directly publish the content.
*/
type ArchivedAt string

/*

StageAsNumber is https://schema.org/stageAsNumber

The stage represented as a number, e.g. 3.
*/
type StageAsNumber float64

/*

StorageRequirements is https://schema.org/storageRequirements

Storage requirements (free space required).
*/
type StorageRequirements string

/*

Algorithm is https://schema.org/algorithm

The algorithm or rules to follow to compute the score.
*/
type Algorithm string

/*

RepeatCount is https://schema.org/repeatCount

Defines the number of times a recurring [[Event]] will take place.
*/
type RepeatCount float64

/*

AccessibilityFeature is https://schema.org/accessibilityFeature

Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility. Values should be drawn from the [approved vocabulary](https://www.w3.org/2021/a11y-discov-vocab/latest/#accessibilityFeature-vocabulary).
*/
type AccessibilityFeature string

/*

TextValue is https://schema.org/textValue

Text value being annotated.
*/
type TextValue string

/*

ProductReturnDays is https://schema.org/productReturnDays

The productReturnDays property indicates the number of days (from purchase) within which relevant product return policy is applicable.
*/
type ProductReturnDays float64

/*

Breadcrumb is https://schema.org/breadcrumb

A set of links that can help a user understand and navigate a website hierarchy.
*/
type Breadcrumb string

/*

Color is https://schema.org/color

The color of the product.
*/
type Color string

/*

SeatSection is https://schema.org/seatSection

The section location of the reserved seat (e.g. Orchestra).
*/
type SeatSection string

/*

JobStartDate is https://schema.org/jobStartDate

The date on which a successful applicant for this job would be expected to start work. Choose a specific date in the future or use the jobImmediateStart property to indicate the position is to be filled as soon as possible.
*/
type JobStartDate string

/*

ReturnPolicyCountry is https://schema.org/returnPolicyCountry

The country where the product has to be sent to for returns, for example "Ireland" using the [[name]] property of [[Country]]. You can also provide the two-letter [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1). Note that this can be different from the country where the product was originally shipped from or sent to.
*/
type ReturnPolicyCountry string

/*

CvdNumC19Died is https://schema.org/cvdNumC19Died

numc19died - DEATHS: Patients with suspected or confirmed COVID-19 who died in the hospital, ED, or any overflow location.
*/
type CvdNumC19Died float64

/*

Url is https://schema.org/url

URL of the item.
*/
type Url string

/*

EligibleRegion is https://schema.org/eligibleRegion

The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.\n\nSee also [[ineligibleRegion]].
    
*/
type EligibleRegion string

/*

MuscleAction is https://schema.org/muscleAction

The movement the muscle generates.
*/
type MuscleAction string

/*

ObservationDate is https://schema.org/observationDate

The observationDate of an [[Observation]].
*/
type ObservationDate string

/*

QuarantineGuidelines is https://schema.org/quarantineGuidelines

Guidelines about quarantine rules, e.g. in the context of a pandemic.
*/
type QuarantineGuidelines string

/*

PriceCurrency is https://schema.org/priceCurrency

The currency of the price, or a price component when attached to [[PriceSpecification]] and its subtypes.\n\nUse standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217), e.g. "USD"; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies, e.g. "BTC"; well known names for [Local Exchange Trading Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types, e.g. "Ithaca HOUR".
*/
type PriceCurrency string

/*

JobTitle is https://schema.org/jobTitle

The job title of the person (for example, Financial Manager).
*/
type JobTitle string

/*

Jurisdiction is https://schema.org/jurisdiction

Indicates a legal jurisdiction, e.g. of some legislation, or where some government service is based.
*/
type Jurisdiction string

/*

MaxPrice is https://schema.org/maxPrice

The highest price if the price is a range.
*/
type MaxPrice float64

/*

KnowsAbout is https://schema.org/knowsAbout

Of a [[Person]], and less typically of an [[Organization]], to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or relate this to educational content, events, objectives or [[JobPosting]] descriptions.
*/
type KnowsAbout string

/*

ContentType is https://schema.org/contentType

The supported content type(s) for an EntryPoint response.
*/
type ContentType string

/*

CatalogNumber is https://schema.org/catalogNumber

The catalog number for the release.
*/
type CatalogNumber string

/*

AcceptsReservations is https://schema.org/acceptsReservations

Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings ```Yes``` or ```No```.
*/
type AcceptsReservations string

/*

LegalName is https://schema.org/legalName

The official name of the organization, e.g. the registered company name.
*/
type LegalName string

/*

ArtEdition is https://schema.org/artEdition

The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, 'artEdition' refers to the total number of copies (in this example "20").
*/
type ArtEdition float64

/*

EstimatedSalary is https://schema.org/estimatedSalary

An estimated salary for a job posting or occupation, based on a variety of variables including, but not limited to industry, job title, and location. Estimated salaries  are often computed by outside organizations rather than the hiring organization, who may not have committed to the estimated value.
*/
type EstimatedSalary float64

/*

CollectionSize is https://schema.org/collectionSize

The number of items in the [[Collection]].
*/
type CollectionSize float64

/*

BroadcastTimezone is https://schema.org/broadcastTimezone

The timezone in [ISO 8601 format](http://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts.
*/
type BroadcastTimezone string

/*

SensoryRequirement is https://schema.org/sensoryRequirement

A description of any sensory requirements and levels necessary to function on the job, including hearing and vision. Defined terms such as those in O*net may be used, but note that there is no way to specify the level of ability as well as its nature when using a defined term.
*/
type SensoryRequirement string

/*

HealthPlanMarketingUrl is https://schema.org/healthPlanMarketingUrl

The URL that goes directly to the plan brochure for the specific standard plan or plan variation.
*/
type HealthPlanMarketingUrl string

/*

Pathophysiology is https://schema.org/pathophysiology

Changes in the normal mechanical, physical, and biochemical functions that are associated with this activity or condition.
*/
type Pathophysiology string

/*

SizeGroup is https://schema.org/sizeGroup

The size group (also known as "size type") for a product's size. Size groups are common in the fashion industry to define size segments and suggested audiences for wearable products. Multiple values can be combined, for example "men's big and tall", "petite maternity" or "regular".
*/
type SizeGroup string

/*

AuditDate is https://schema.org/auditDate

Date when a certification was last audited. See also  [gs1:certificationAuditDate](https://www.gs1.org/voc/certificationAuditDate).
*/
type AuditDate string

/*

ActionPlatform is https://schema.org/actionPlatform

The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
*/
type ActionPlatform string

/*

TissueSample is https://schema.org/tissueSample

The type of tissue sample required for the test.
*/
type TissueSample string

/*

OrderDate is https://schema.org/orderDate

Date order was placed.
*/
type OrderDate string

/*

Epidemiology is https://schema.org/epidemiology

The characteristics of associated patients, such as age, gender, race etc.
*/
type Epidemiology string

/*

LegislationDate is https://schema.org/legislationDate

The date of adoption or signature of the legislation. This is the date at which the text is officially aknowledged to be a legislation, even though it might not even be published or in force.
*/
type LegislationDate string

/*

StrengthUnit is https://schema.org/strengthUnit

The units of an active ingredient's strength, e.g. mg.
*/
type StrengthUnit string

/*

EducationRequirements is https://schema.org/educationRequirements

Educational background needed for the position or Occupation.
*/
type EducationRequirements string

/*

Incentives is https://schema.org/incentives

Description of bonus and commission compensation aspects of the job.
*/
type Incentives string

/*

RecipeCategory is https://schema.org/recipeCategory

The category of the recipe—for example, appetizer, entree, etc.
*/
type RecipeCategory string

/*

DrugUnit is https://schema.org/drugUnit

The unit in which the drug is measured, e.g. '5 mg tablet'.
*/
type DrugUnit string

/*

MolecularFormula is https://schema.org/molecularFormula

The empirical formula is the simplest whole number ratio of all the atoms in a molecule.
*/
type MolecularFormula string

/*

SecurityClearanceRequirement is https://schema.org/securityClearanceRequirement

A description of any security clearance requirements of the job.
*/
type SecurityClearanceRequirement string

/*

DataFeedElement is https://schema.org/dataFeedElement

An item within a data feed. Data feeds may have many elements.
*/
type DataFeedElement string

/*

RepresentativeOfPage is https://schema.org/representativeOfPage

Indicates whether this image is representative of the content of the page.
*/
type RepresentativeOfPage bool

/*

ExifData is https://schema.org/exifData

exif data for this object.
*/
type ExifData string

/*

SafetyConsideration is https://schema.org/safetyConsideration

Any potential safety concern associated with the supplement. May include interactions with other drugs and foods, pregnancy, breastfeeding, known adverse reactions, and documented efficacy of the supplement.
*/
type SafetyConsideration string

/*

TaxonomicRange is https://schema.org/taxonomicRange

The taxonomic grouping of the organism that expresses, encodes, or in some way related to the BioChemEntity.
*/
type TaxonomicRange string

/*

AdditionalNumberOfGuests is https://schema.org/additionalNumberOfGuests

If responding yes, the number of guests who will attend in addition to the invitee.
*/
type AdditionalNumberOfGuests float64

/*

TotalPrice is https://schema.org/totalPrice

The total price for the reservation or ticket, including applicable taxes, shipping, etc.\n\nUsage guidelines:\n\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
*/
type TotalPrice string

/*

NumberOfEpisodes is https://schema.org/numberOfEpisodes

The number of episodes in this season or series.
*/
type NumberOfEpisodes float64

/*

TaxonRank is https://schema.org/taxonRank

The taxonomic rank of this taxon given preferably as a URI from a controlled vocabulary – typically the ranks from TDWG TaxonRank ontology or equivalent Wikidata URIs.
*/
type TaxonRank string

/*

EndDate is https://schema.org/endDate

The end date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
*/
type EndDate string

/*

LeiCode is https://schema.org/leiCode

An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
*/
type LeiCode string

/*

NumberOfBeds is https://schema.org/numberOfBeds

The quantity of the given bed type available in the HotelRoom, Suite, House, or Apartment.
*/
type NumberOfBeds float64

/*

TickerSymbol is https://schema.org/tickerSymbol

The exchange traded instrument associated with a Corporation object. The tickerSymbol is expressed as an exchange and an instrument name separated by a space character. For the exchange component of the tickerSymbol attribute, we recommend using the controlled vocabulary of Market Identifier Codes (MIC) specified in ISO 15022.
*/
type TickerSymbol string

/*

DoesNotShip is https://schema.org/doesNotShip

Indicates when shipping to a particular [[shippingDestination]] is not available.
*/
type DoesNotShip bool

/*

SizeSystem is https://schema.org/sizeSystem

The size system used to identify a product's size. Typically either a standard (for example, "GS1" or "ISO-EN13402"), country code (for example "US" or "JP"), or a measuring system (for example "Metric" or "Imperial").
*/
type SizeSystem string

/*

Name is https://schema.org/name

The name of the item.
*/
type Name string

/*

ExecutableLibraryName is https://schema.org/executableLibraryName

Library file name, e.g., mscorlib.dll, system.web.dll.
*/
type ExecutableLibraryName string

/*

Warning is https://schema.org/warning

Any FDA or other warnings about the drug (text or URL).
*/
type Warning string

/*

BroadcastFrequencyValue is https://schema.org/broadcastFrequencyValue

The frequency in MHz for a particular broadcast.
*/
type BroadcastFrequencyValue float64

/*

RecipeIngredient is https://schema.org/recipeIngredient

A single ingredient used in the recipe, e.g. sugar, flour or garlic.
*/
type RecipeIngredient string

/*

CvdFacilityCounty is https://schema.org/cvdFacilityCounty

Name of the County of the NHSN facility that this data record applies to. Use [[cvdFacilityId]] to identify the facility. To provide other details, [[healthcareReportingData]] can be used on a [[Hospital]] entry.
*/
type CvdFacilityCounty string

/*

TouristType is https://schema.org/touristType

Attraction suitable for type(s) of tourist. E.g. children, visitors from a particular country, etc. 
*/
type TouristType string

/*

LastReviewed is https://schema.org/lastReviewed

Date on which the content on this web page was last reviewed for accuracy and/or completeness.
*/
type LastReviewed string

/*

Preparation is https://schema.org/preparation

Typical preparation that a patient must undergo before having the procedure performed.
*/
type Preparation string

/*

NumberOfPages is https://schema.org/numberOfPages

The number of pages in the book.
*/
type NumberOfPages float64

/*

DateIssued is https://schema.org/dateIssued

The date the ticket was issued.
*/
type DateIssued string

/*

Gtin8 is https://schema.org/gtin8

The GTIN-8 code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
*/
type Gtin8 string

/*

OccupationalCategory is https://schema.org/occupationalCategory

A category describing the job, preferably using a term from a taxonomy such as [BLS O*NET-SOC](http://www.onetcenter.org/taxonomy.html), [ISCO-08](https://www.ilo.org/public/english/bureau/stat/isco/isco08/) or similar, with the property repeated for each applicable value. Ideally the taxonomy should be identified, and both the textual label and formal code for the category should be provided.\n
Note: for historical reasons, any textual label and formal code provided as a literal may be assumed to be from O*NET-SOC.
*/
type OccupationalCategory string

/*

ApplicationSuite is https://schema.org/applicationSuite

The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
*/
type ApplicationSuite string

/*

AccessMode is https://schema.org/accessMode

The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Values should be drawn from the [approved vocabulary](https://www.w3.org/2021/a11y-discov-vocab/latest/#accessMode-vocabulary).
*/
type AccessMode string

/*

MeasurementTechnique is https://schema.org/measurementTechnique

A technique, method or technology used in an [[Observation]], [[StatisticalVariable]] or [[Dataset]] (or [[DataDownload]], [[DataCatalog]]), corresponding to the method used for measuring the corresponding variable(s) (for datasets, described using [[variableMeasured]]; for [[Observation]], a [[StatisticalVariable]]). Often but not necessarily each [[variableMeasured]] will have an explicit representation as (or mapping to) an property such as those defined in Schema.org, or other RDF vocabularies and "knowledge graphs". In that case the subproperty of [[variableMeasured]] called [[measuredProperty]] is applicable.
    
The [[measurementTechnique]] property helps when extra clarification is needed about how a [[measuredProperty]] was measured. This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but can often serve as a high level summary for dataset discovery. 

For example, if [[variableMeasured]] is: molecule concentration, [[measurementTechnique]] could be: "mass spectrometry" or "nmr spectroscopy" or "colorimetry" or "immunofluorescence". If the [[variableMeasured]] is "depression rating", the [[measurementTechnique]] could be "Zung Scale" or "HAM-D" or "Beck Depression Inventory". 

If there are several [[variableMeasured]] properties recorded for some given data object, use a [[PropertyValue]] for each [[variableMeasured]] and attach the corresponding [[measurementTechnique]]. The value can also be from an enumeration, organized as a [[MeasurementMetholdEnumeration]].
*/
type MeasurementTechnique string

/*

Category is https://schema.org/category

A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
*/
type Category string

/*

EducationalUse is https://schema.org/educationalUse

The purpose of a work in the context of education; for example, 'assignment', 'group work'.
*/
type EducationalUse string

/*

AlternateName is https://schema.org/alternateName

An alias for the item.
*/
type AlternateName string

/*

InDefinedTermSet is https://schema.org/inDefinedTermSet

A [[DefinedTermSet]] that contains this term.
*/
type InDefinedTermSet string

/*

IsrcCode is https://schema.org/isrcCode

The International Standard Recording Code for the recording.
*/
type IsrcCode string

/*

MaximumAttendeeCapacity is https://schema.org/maximumAttendeeCapacity

The total number of individuals that may attend an event or venue.
*/
type MaximumAttendeeCapacity float64

/*

UnitCode is https://schema.org/unitCode

The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
*/
type UnitCode string

/*

GameEdition is https://schema.org/gameEdition

The edition of a video game.
*/
type GameEdition string

/*

Maps is https://schema.org/maps

A URL to a map of the place.
*/
type Maps string

/*

Gtin is https://schema.org/gtin

A Global Trade Item Number ([GTIN](https://www.gs1.org/standards/id-keys/gtin)). GTINs identify trade items, including products and services, using numeric identification codes.

A correct [[gtin]] value should be a valid GTIN, which means that it should be an all-numeric string of either 8, 12, 13 or 14 digits, or a "GS1 Digital Link" URL based on such a string. The numeric component should also have a [valid GS1 check digit](https://www.gs1.org/services/check-digit-calculator) and meet the other rules for valid GTINs. See also [GS1's GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) and [Wikipedia](https://en.wikipedia.org/wiki/Global_Trade_Item_Number) for more details. Left-padding of the gtin values is not required or encouraged. The [[gtin]] property generalizes the earlier [[gtin8]], [[gtin12]], [[gtin13]], and [[gtin14]] properties.

The GS1 [digital link specifications](https://www.gs1.org/standards/Digital-Link/) expresses GTINs as URLs (URIs, IRIs, etc.).
Digital Links should be populated into the [[hasGS1DigitalLink]] attribute.

Note also that this is a definition for how to include GTINs in Schema.org data, and not a definition of GTINs in general - see the GS1 documentation for authoritative details.
*/
type Gtin string

/*

CvdNumC19HospPats is https://schema.org/cvdNumC19HospPats

numc19hosppats - HOSPITALIZED: Patients currently hospitalized in an inpatient care location who have suspected or confirmed COVID-19.
*/
type CvdNumC19HospPats float64

/*

ContentUrl is https://schema.org/contentUrl

Actual bytes of the media object, for example the image file or video file.
*/
type ContentUrl string

/*

AccountId is https://schema.org/accountId

The identifier for the account the payment will be applied to.
*/
type AccountId string

/*

LodgingUnitType is https://schema.org/lodgingUnitType

Textual description of the unit type (including suite vs. room, size of bed, etc.).
*/
type LodgingUnitType string

/*

NonProprietaryName is https://schema.org/nonProprietaryName

The generic name of this drug or supplement.
*/
type NonProprietaryName string

/*

SoftwareRequirements is https://schema.org/softwareRequirements

Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (examples: DirectX, Java or .NET runtime).
*/
type SoftwareRequirements string

/*

Material is https://schema.org/material

A material that something is made from, e.g. leather, wool, cotton, paper.
*/
type Material string

/*

DoseUnit is https://schema.org/doseUnit

The unit of the dose, e.g. 'mg'.
*/
type DoseUnit string

/*

PropertyID is https://schema.org/propertyID

A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
(1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific ID of the property), or (3)
a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
*/
type PropertyID string

/*

ValidFrom is https://schema.org/validFrom

The date when the item becomes valid.
*/
type ValidFrom string

/*

MembershipPointsEarned is https://schema.org/membershipPointsEarned

The number of membership points earned by the member. If necessary, the unitText can be used to express the units the points are issued in. (E.g. stars, miles, etc.)
*/
type MembershipPointsEarned float64

/*

CertificationIdentification is https://schema.org/certificationIdentification

Identifier of a certification instance (as registered with an independent certification body). Typically this identifier can be used to consult and verify the certification instance. See also [gs1:certificationIdentification](https://www.gs1.org/voc/certificationIdentification).
*/
type CertificationIdentification string

/*

PassengerPriorityStatus is https://schema.org/passengerPriorityStatus

The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
*/
type PassengerPriorityStatus string

/*

LegalStatus is https://schema.org/legalStatus

The drug or supplement's legal status, including any controlled substance schedules that apply.
*/
type LegalStatus string

/*

FaxNumber is https://schema.org/faxNumber

The fax number.
*/
type FaxNumber string

/*

ItemListElement is https://schema.org/itemListElement

For itemListElement values, you can use simple strings (e.g. "Peter", "Paul", "Mary"), existing entities, or use ListItem.\n\nText values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.\n\nNote: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a 'position' property in such cases.
*/
type ItemListElement string

/*

BookingTime is https://schema.org/bookingTime

The date and time the reservation was booked.
*/
type BookingTime string

/*

CurrenciesAccepted is https://schema.org/currenciesAccepted

The currency accepted.\n\nUse standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217), e.g. "USD"; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies, e.g. "BTC"; well known names for [Local Exchange Trading Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types, e.g. "Ithaca HOUR".
*/
type CurrenciesAccepted string

/*

RelatedLink is https://schema.org/relatedLink

A link related to this web page, for example to other related web pages.
*/
type RelatedLink string

/*

DatePosted is https://schema.org/datePosted

Publication date of an online listing.
*/
type DatePosted string

/*

DietFeatures is https://schema.org/dietFeatures

Nutritional information specific to the dietary plan. May include dietary recommendations on what foods to avoid, what foods to consume, and specific alterations/deviations from the USDA or other regulatory body's approved dietary guidelines.
*/
type DietFeatures string

/*

ContactlessPayment is https://schema.org/contactlessPayment

A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
*/
type ContactlessPayment bool

/*

Masthead is https://schema.org/masthead

For a [[NewsMediaOrganization]], a link to the masthead page or a page listing top editorial management.
*/
type Masthead string

/*

ServiceType is https://schema.org/serviceType

The type of service being offered, e.g. veterans' benefits, emergency relief, etc.
*/
type ServiceType string

/*

FeesAndCommissionsSpecification is https://schema.org/feesAndCommissionsSpecification

Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
*/
type FeesAndCommissionsSpecification string

/*

NumberOfCredits is https://schema.org/numberOfCredits

The number of credits or units awarded by a Course or required to complete an EducationalOccupationalProgram.
*/
type NumberOfCredits float64

/*

ValueName is https://schema.org/valueName

Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML's input@name.
*/
type ValueName string

/*

AvailableLanguage is https://schema.org/availableLanguage

A language someone may use with or at the item, service or place. Please use one of the language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47). See also [[inLanguage]].
*/
type AvailableLanguage string

/*

Yield is https://schema.org/yield

The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
*/
type Yield string

/*

EmployerOverview is https://schema.org/employerOverview

A description of the employer, career opportunities and work environment for this position.
*/
type EmployerOverview string

/*

ClaimReviewed is https://schema.org/claimReviewed

A short summary of the specific claims reviewed in a ClaimReview.
*/
type ClaimReviewed string

/*

ClincalPharmacology is https://schema.org/clincalPharmacology

Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
*/
type ClincalPharmacology string

/*

ByMonthWeek is https://schema.org/byMonthWeek

Defines the week(s) of the month on which a recurring Event takes place. Specified as an Integer between 1-5. For clarity, byMonthWeek is best used in conjunction with byDay to indicate concepts like the first and third Mondays of a month.
*/
type ByMonthWeek float64

/*

ActiveIngredient is https://schema.org/activeIngredient

An active ingredient, typically chemical compounds and/or biologic substances.
*/
type ActiveIngredient string

/*

PostalCode is https://schema.org/postalCode

The postal code. For example, 94043.
*/
type PostalCode string

/*

PickupTime is https://schema.org/pickupTime

When a taxi will pick up a passenger or a rental car can be picked up.
*/
type PickupTime string

/*

Dependencies is https://schema.org/dependencies

Prerequisites needed to fulfill steps in article.
*/
type Dependencies string

/*

BookEdition is https://schema.org/bookEdition

The edition of the book.
*/
type BookEdition string

/*

NumberOfBedrooms is https://schema.org/numberOfBedrooms

The total integer number of bedrooms in a some [[Accommodation]], [[ApartmentComplex]] or [[FloorPlan]].
*/
type NumberOfBedrooms float64

/*

DepartureGate is https://schema.org/departureGate

Identifier of the flight's departure gate.
*/
type DepartureGate string

/*

GivenName is https://schema.org/givenName

Given name. In the U.S., the first name of a Person.
*/
type GivenName string

/*

AccommodationCategory is https://schema.org/accommodationCategory

Category of an [[Accommodation]], following real estate conventions, e.g. RESO (see [PropertySubType](https://ddwiki.reso.org/display/DDW17/PropertySubType+Field), and [PropertyType](https://ddwiki.reso.org/display/DDW17/PropertyType+Field) fields  for suggested values).
*/
type AccommodationCategory string

/*

InLanguage is https://schema.org/inLanguage

The language of the content or performance or used in an action. Please use one of the language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47). See also [[availableLanguage]].
*/
type InLanguage string

/*

DownPayment is https://schema.org/downPayment

a type of payment made in cash during the onset of the purchase of an expensive good/service. The payment typically represents only a percentage of the full purchase price.
*/
type DownPayment float64

/*

DateRead is https://schema.org/dateRead

The date/time at which the message has been read by the recipient if a single recipient exists.
*/
type DateRead string

/*

MusicCompositionForm is https://schema.org/musicCompositionForm

The type of composition (e.g. overture, sonata, symphony, etc.).
*/
type MusicCompositionForm string

/*

ContentSize is https://schema.org/contentSize

File size in (mega/kilo)bytes.
*/
type ContentSize string

/*

NegativeNotes is https://schema.org/negativeNotes

Provides negative considerations regarding something, most typically in pro/con lists for reviews (alongside [[positiveNotes]]). For symmetry 

In the case of a [[Review]], the property describes the [[itemReviewed]] from the perspective of the review; in the case of a [[Product]], the product itself is being described. Since product descriptions 
tend to emphasise positive claims, it may be relatively unusual to find [[negativeNotes]] used in this way. Nevertheless for the sake of symmetry, [[negativeNotes]] can be used on [[Product]].

The property values can be expressed either as unstructured text (repeated as necessary), or if ordered, as a list (in which case the most negative is at the beginning of the list).
*/
type NegativeNotes string

/*

BusNumber is https://schema.org/busNumber

The unique identifier for the bus.
*/
type BusNumber string

/*

WordCount is https://schema.org/wordCount

The number of words in the text of the Article.
*/
type WordCount float64

/*

ValueMinLength is https://schema.org/valueMinLength

Specifies the minimum allowed range for number of characters in a literal value.
*/
type ValueMinLength float64

/*

BirthDate is https://schema.org/birthDate

Date of birth.
*/
type BirthDate string

/*

Sku is https://schema.org/sku

The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
*/
type Sku string

/*

TemporalCoverage is https://schema.org/temporalCoverage

The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in [ISO 8601 time interval format](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals). In
      the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written "2011/2012"). Other forms of content, e.g. ScholarlyArticle, Book, TVSeries or TVEpisode, may indicate their temporalCoverage in broader terms - textually or via well-known URL.
      Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via "1939/1945".

Open-ended date ranges can be written with ".." in place of the end date. For example, "2015-11/.." indicates a range beginning in November 2015 and with no specified final date. This is tentative and might be updated in future when ISO 8601 is officially updated.
*/
type TemporalCoverage string

/*

ContactType is https://schema.org/contactType

A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
*/
type ContactType string

/*

ServesCuisine is https://schema.org/servesCuisine

The cuisine of the restaurant.
*/
type ServesCuisine string

/*

CvdFacilityId is https://schema.org/cvdFacilityId

Identifier of the NHSN facility that this data record applies to. Use [[cvdFacilityCounty]] to indicate the county. To provide other details, [[healthcareReportingData]] can be used on a [[Hospital]] entry.
*/
type CvdFacilityId string

/*

Image is https://schema.org/image

An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
*/
type Image string

/*

BranchCode is https://schema.org/branchCode

A short textual code (also called "store code") that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.\n\nFor example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code "3047" is a branchCode for a particular branch.
      
*/
type BranchCode string

/*

MealService is https://schema.org/mealService

Description of the meals that will be provided or available for purchase.
*/
type MealService string

/*

ScreenCount is https://schema.org/screenCount

The number of screens in the movie theater.
*/
type ScreenCount float64

/*

StreetAddress is https://schema.org/streetAddress

The street address. For example, 1600 Amphitheatre Pkwy.
*/
type StreetAddress string

/*

Percentile10 is https://schema.org/percentile10

The 10th percentile value.
*/
type Percentile10 float64

/*

AssociatedDisease is https://schema.org/associatedDisease

Disease associated to this BioChemEntity. Such disease can be a MedicalCondition or a URL. If you want to add an evidence supporting the association, please use PropertyValue.
*/
type AssociatedDisease string

/*

Sport is https://schema.org/sport

A type of sport (e.g. Baseball).
*/
type Sport string

/*

IsGift is https://schema.org/isGift

Indicates whether the offer was accepted as a gift for someone other than the buyer.
*/
type IsGift bool

/*

ProductSupported is https://schema.org/productSupported

The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. "iPhone") or a general category of products or services (e.g. "smartphones").
*/
type ProductSupported string

/*

GameAvailabilityType is https://schema.org/gameAvailabilityType

Indicates the availability type of the game content associated with this action, such as whether it is a full version or a demo.
*/
type GameAvailabilityType string

/*

UserInteractionCount is https://schema.org/userInteractionCount

The number of interactions for the CreativeWork using the WebSite or SoftwareApplication.
*/
type UserInteractionCount float64

/*

VehicleIdentificationNumber is https://schema.org/vehicleIdentificationNumber

The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
*/
type VehicleIdentificationNumber string

/*

OwnedThrough is https://schema.org/ownedThrough

The date and time of giving up ownership on the product.
*/
type OwnedThrough string

/*

Logo is https://schema.org/logo

An associated logo.
*/
type Logo string

/*

NumChildren is https://schema.org/numChildren

The number of children staying in the unit.
*/
type NumChildren float64

/*

TravelBans is https://schema.org/travelBans

Information about travel bans, e.g. in the context of a pandemic.
*/
type TravelBans string

/*

Significance is https://schema.org/significance

The significance associated with the superficial anatomy; as an example, how characteristics of the superficial anatomy can suggest underlying medical conditions or courses of treatment.
*/
type Significance string

/*

ConditionsOfAccess is https://schema.org/conditionsOfAccess

Conditions that affect the availability of, or method(s) of access to, an item. Typically used for real world items such as an [[ArchiveComponent]] held by an [[ArchiveOrganization]]. This property is not suitable for use as a general Web access control mechanism. It is expressed only in natural language.\n\nFor example "Available by appointment from the Reading Room" or "Accessible only from logged-in accounts ". 
*/
type ConditionsOfAccess string

/*

AddressLocality is https://schema.org/addressLocality

The locality in which the street address is, and which is in the region. For example, Mountain View.
*/
type AddressLocality string

/*

Repetitions is https://schema.org/repetitions

Number of times one should repeat the activity.
*/
type Repetitions float64

/*

CostOrigin is https://schema.org/costOrigin

Additional details to capture the origin of the cost data. For example, 'Medicare Part B'.
*/
type CostOrigin string

/*

Responsibilities is https://schema.org/responsibilities

Responsibilities associated with this role or Occupation.
*/
type Responsibilities string

/*

CodeRepository is https://schema.org/codeRepository

Link to the repository where the un-compiled, human readable code and related code is located (SVN, GitHub, CodePlex).
*/
type CodeRepository string

/*

Season is https://schema.org/season

A season in a media series.
*/
type Season string

/*

NumberOfItems is https://schema.org/numberOfItems

The number of items in an ItemList. Note that some descriptions might not fully describe all items in a list (e.g., multi-page pagination); in such cases, the numberOfItems would be for the entire list.
*/
type NumberOfItems float64

/*

ApplicationCategory is https://schema.org/applicationCategory

Type of software application, e.g. 'Game, Multimedia'.
*/
type ApplicationCategory string

/*

PrintColumn is https://schema.org/printColumn

The number of the column in which the NewsArticle appears in the print edition.
*/
type PrintColumn string

/*

ByMonth is https://schema.org/byMonth

Defines the month(s) of the year on which a recurring [[Event]] takes place. Specified as an [[Integer]] between 1-12. January is 1.
*/
type ByMonth float64

/*

EmploymentType is https://schema.org/employmentType

Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
*/
type EmploymentType string

/*

StructuralClass is https://schema.org/structuralClass

The name given to how bone physically connects to each other.
*/
type StructuralClass string

/*

CostPerUnit is https://schema.org/costPerUnit

The cost per unit of the drug.
*/
type CostPerUnit float64

/*

DepartureTime is https://schema.org/departureTime

The expected departure time.
*/
type DepartureTime string

/*

Step is https://schema.org/step

A single step item (as HowToStep, text, document, video, etc.) or a HowToSection.
*/
type Step string

/*

UsesHealthPlanIdStandard is https://schema.org/usesHealthPlanIdStandard

The standard for interpreting the Plan ID. The preferred is "HIOS". See the Centers for Medicare & Medicaid Services for more details.
*/
type UsesHealthPlanIdStandard string

/*

NamedPosition is https://schema.org/namedPosition

A position played, performed or filled by a person or organization, as part of an organization. For example, an athlete in a SportsTeam might play in the position named 'Quarterback'.
*/
type NamedPosition string

/*

RuntimePlatform is https://schema.org/runtimePlatform

Runtime platform or script interpreter dependencies (example: Java v1, Python 2.3, .NET Framework 3.0).
*/
type RuntimePlatform string

/*

Intensity is https://schema.org/intensity

Quantitative measure gauging the degree of force involved in the exercise, for example, heartbeats per minute. May include the velocity of the movement.
*/
type Intensity string

/*

Speakable is https://schema.org/speakable

Indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the 'speakable' property serves to indicate the parts most likely to be generally useful for speech.

The *speakable* property can be repeated an arbitrary number of times, with three kinds of possible 'content-locator' values:

1.) *id-value* URL references - uses *id-value* of an element in the page being annotated. The simplest use of *speakable* has (potentially relative) URL values, referencing identified sections of the document concerned.

2.) CSS Selectors - addresses content in the annotated page, e.g. via class attribute. Use the [[cssSelector]] property.

3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the [[xpath]] property.


For more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this
we define a supporting type, [[SpeakableSpecification]]  which is defined to be a possible value of the *speakable* property.
         
*/
type Speakable string

/*

RequiredQuantity is https://schema.org/requiredQuantity

The required quantity of the item(s).
*/
type RequiredQuantity float64

/*

CvdNumBeds is https://schema.org/cvdNumBeds

numbeds - HOSPITAL INPATIENT BEDS: Inpatient beds, including all staffed, licensed, and overflow (surge) beds used for inpatients.
*/
type CvdNumBeds float64

/*

ReviewBody is https://schema.org/reviewBody

The actual body of the review.
*/
type ReviewBody string

/*

SeatNumber is https://schema.org/seatNumber

The location of the reserved seat (e.g., 27).
*/
type SeatNumber string

/*

CountriesNotSupported is https://schema.org/countriesNotSupported

Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
*/
type CountriesNotSupported string

/*

Temporal is https://schema.org/temporal

The "temporal" property can be used in cases where more specific properties
(e.g. [[temporalCoverage]], [[dateCreated]], [[dateModified]], [[datePublished]]) are not known to be appropriate.
*/
type Temporal string

/*

RoleName is https://schema.org/roleName

A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named 'inker', 'penciller', and 'letterer'; or an athlete in a SportsTeam might play in the position named 'Quarterback'.
*/
type RoleName string

/*

CashBack is https://schema.org/cashBack

A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
*/
type CashBack bool

/*

Artform is https://schema.org/artform

e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
*/
type Artform string

/*

IsPartOf is https://schema.org/isPartOf

Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of.
*/
type IsPartOf string

/*

CommentTime is https://schema.org/commentTime

The time at which the UserComment was made.
*/
type CommentTime string

/*

ActivityFrequency is https://schema.org/activityFrequency

How often one should engage in the activity.
*/
type ActivityFrequency string

/*

HasGS1DigitalLink is https://schema.org/hasGS1DigitalLink

The <a href="https://www.gs1.org/standards/gs1-digital-link">GS1 digital link</a> associated with the object. This URL should conform to the particular requirements of digital links. The link should only contain the Application Identifiers (AIs) that are relevant for the entity being annotated, for instance a [[Product]] or an [[Organization]], and for the correct granularity. In particular, for products:<ul><li>A Digital Link that contains a serial number (AI <code>21</code>) should only be present on instances of [[IndividualProduct]]</li><li>A Digital Link that contains a lot number (AI <code>10</code>) should be annotated as [[SomeProduct]] if only products from that lot are sold, or [[IndividualProduct]] if there is only a specific product.</li><li>A Digital Link that contains a global model number (AI <code>8013</code>)  should be attached to a [[Product]] or a [[ProductModel]].</li></ul> Other item types should be adapted similarly.
*/
type HasGS1DigitalLink string

/*

Ingredients is https://schema.org/ingredients

A single ingredient used in the recipe, e.g. sugar, flour or garlic.
*/
type Ingredients string

/*

MusicalKey is https://schema.org/musicalKey

The key, mode, or scale this composition uses.
*/
type MusicalKey string

/*

ExperienceInPlaceOfEducation is https://schema.org/experienceInPlaceOfEducation

Indicates whether a [[JobPosting]] will accept experience (as indicated by [[OccupationalExperienceRequirements]]) in place of its formal educational qualifications (as indicated by [[educationRequirements]]). If true, indicates that satisfying one of these requirements is sufficient.
*/
type ExperienceInPlaceOfEducation bool

/*

Abridged is https://schema.org/abridged

Indicates whether the book is an abridged edition.
*/
type Abridged bool

/*

AccessibilityControl is https://schema.org/accessibilityControl

Identifies input methods that are sufficient to fully control the described resource. Values should be drawn from the [approved vocabulary](https://www.w3.org/2021/a11y-discov-vocab/latest/#accessibilityControl-vocabulary).
*/
type AccessibilityControl string

/*

Description is https://schema.org/description

A description of the item.
*/
type Description string

/*

OwnedFrom is https://schema.org/ownedFrom

The date and time of obtaining the product.
*/
type OwnedFrom string

/*

HttpMethod is https://schema.org/httpMethod

An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
*/
type HttpMethod string

/*

VehicleSpecialUsage is https://schema.org/vehicleSpecialUsage

Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
*/
type VehicleSpecialUsage string

/*

Longitude is https://schema.org/longitude

The longitude of a location. For example ```-122.08585``` ([WGS 84](https://en.wikipedia.org/wiki/World_Geodetic_System)).
*/
type Longitude float64

/*

LodgingUnitDescription is https://schema.org/lodgingUnitDescription

A full description of the lodging unit.
*/
type LodgingUnitDescription string

/*

EmbeddedTextCaption is https://schema.org/embeddedTextCaption

Represents textual captioning from a [[MediaObject]], e.g. text of a 'meme'.
*/
type EmbeddedTextCaption string

/*

SignificantLinks is https://schema.org/significantLinks

The most significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
*/
type SignificantLinks string

/*

TermsPerYear is https://schema.org/termsPerYear

The number of times terms of study are offered per year. Semesters and quarters are common units for term. For example, if the student can only take 2 semesters for the program in one year, then termsPerYear should be 2.
*/
type TermsPerYear float64

/*

LinkRelationship is https://schema.org/linkRelationship

Indicates the relationship type of a Web link. 
*/
type LinkRelationship string

/*

MeetsEmissionStandard is https://schema.org/meetsEmissionStandard

Indicates that the vehicle meets the respective emission standard.
*/
type MeetsEmissionStandard string

/*

StrengthValue is https://schema.org/strengthValue

The value of an active ingredient's strength, e.g. 325.
*/
type StrengthValue float64

/*

HighPrice is https://schema.org/highPrice

The highest price of all offers available.\n\nUsage guidelines:\n\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
*/
type HighPrice float64

/*

HealthPlanId is https://schema.org/healthPlanId

The 14-character, HIOS-generated Plan ID number. (Plan IDs must be unique, even across different markets.)
*/
type HealthPlanId string

/*

HasMolecularFunction is https://schema.org/hasMolecularFunction

Molecular function performed by this BioChemEntity; please use PropertyValue if you want to include any evidence.
*/
type HasMolecularFunction string

/*

FamilyName is https://schema.org/familyName

Family name. In the U.S., the last name of a Person.
*/
type FamilyName string

/*

BillingStart is https://schema.org/billingStart

Specifies after how much time this price (or price component) becomes valid and billing starts. Can be used, for example, to model a price increase after the first year of a subscription. The unit of measurement is specified by the unitCode property.
*/
type BillingStart float64

/*

MaxValue is https://schema.org/maxValue

The upper value of some characteristic or property.
*/
type MaxValue float64

/*

ProgramName is https://schema.org/programName

The program providing the membership. It is preferable to use [:program](http://schema.org/program) instead.
*/
type ProgramName string

/*

Target is https://schema.org/target

Indicates a target EntryPoint, or url, for an Action.
*/
type Target string

/*

PregnancyWarning is https://schema.org/pregnancyWarning

Any precaution, guidance, contraindication, etc. related to this drug's use during pregnancy.
*/
type PregnancyWarning string

/*

CvdNumC19HOPats is https://schema.org/cvdNumC19HOPats

numc19hopats - HOSPITAL ONSET: Patients hospitalized in an NHSN inpatient care location with onset of suspected or confirmed COVID-19 14 or more days after hospitalization.
*/
type CvdNumC19HOPats float64

/*

ExperienceRequirements is https://schema.org/experienceRequirements

Description of skills and experience needed for the position or Occupation.
*/
type ExperienceRequirements string

/*

BankAccountType is https://schema.org/bankAccountType

The type of a bank account.
*/
type BankAccountType string

/*

Pattern is https://schema.org/pattern

A pattern that something has, for example 'polka dot', 'striped', 'Canadian flag'. Values are typically expressed as text, although links to controlled value schemes are also supported.
*/
type Pattern string

/*

PaymentDue is https://schema.org/paymentDue

The date that payment is due.
*/
type PaymentDue string

/*

GettingTestedInfo is https://schema.org/gettingTestedInfo

Information about getting tested (for a [[MedicalCondition]]), e.g. in the context of a pandemic.
*/
type GettingTestedInfo string

/*

HealthPlanCostSharing is https://schema.org/healthPlanCostSharing

The costs to the patient for services under this network or formulary.
*/
type HealthPlanCostSharing bool

/*

ExceptDate is https://schema.org/exceptDate

Defines a [[Date]] or [[DateTime]] during which a scheduled [[Event]] will not take place. The property allows exceptions to
      a [[Schedule]] to be specified. If an exception is specified as a [[DateTime]] then only the event that would have started at that specific date and time
      should be excluded from the schedule. If an exception is specified as a [[Date]] then any event that is scheduled for that 24 hour period should be
      excluded from the schedule. This allows a whole day to be excluded from the schedule without having to itemise every scheduled event.
*/
type ExceptDate string

/*

BillingIncrement is https://schema.org/billingIncrement

This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
*/
type BillingIncrement float64

/*

MaximumEnrollment is https://schema.org/maximumEnrollment

The maximum number of students who may be enrolled in the program.
*/
type MaximumEnrollment float64

/*

PlayerType is https://schema.org/playerType

Player type required&#x2014;for example, Flash or Silverlight.
*/
type PlayerType string

/*

HowPerformed is https://schema.org/howPerformed

How the procedure is performed.
*/
type HowPerformed string

/*

DirectApply is https://schema.org/directApply

Indicates whether an [[url]] that is associated with a [[JobPosting]] enables direct application for the job, via the posting website. A job posting is considered to have directApply of [[True]] if an application process for the specified job can be directly initiated via the url(s) given (noting that e.g. multiple internet domains might nevertheless be involved at an implementation level). A value of [[False]] is appropriate if there is no clear path to applying directly online for the specified job, navigating directly from the JobPosting url(s) supplied.
*/
type DirectApply bool

/*

OpeningHours is https://schema.org/openingHours

The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas ',' separating each day. Day or time ranges are specified using a hyphen '-'.\n\n* Days are specified using the following two-letter combinations: ```Mo```, ```Tu```, ```We```, ```Th```, ```Fr```, ```Sa```, ```Su```.\n* Times are specified using 24:00 format. For example, 3pm is specified as ```15:00```, 10am as ```10:00```. \n* Here is an example: <code>&lt;time itemprop="openingHours" datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;</code>.\n* If a business is open 7 days a week, then it can be specified as <code>&lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;</code>.
*/
type OpeningHours string

/*

ArrivalPlatform is https://schema.org/arrivalPlatform

The platform where the train arrives.
*/
type ArrivalPlatform string

/*

JobLocationType is https://schema.org/jobLocationType

A description of the job location (e.g. TELECOMMUTE for telecommute jobs).
*/
type JobLocationType string

/*

CheckoutTime is https://schema.org/checkoutTime

The latest someone may check out of a lodging establishment.
*/
type CheckoutTime string

/*

LearningResourceType is https://schema.org/learningResourceType

The predominant type or kind characterizing the learning resource. For example, 'presentation', 'handout'.
*/
type LearningResourceType string

/*

OriginalMediaContextDescription is https://schema.org/originalMediaContextDescription

Describes, in a [[MediaReview]] when dealing with [[DecontextualizedContent]], background information that can contribute to better interpretation of the [[MediaObject]].
*/
type OriginalMediaContextDescription string

/*

NumberOfAirbags is https://schema.org/numberOfAirbags

The number or type of airbags in the vehicle.
*/
type NumberOfAirbags float64

/*

EndOffset is https://schema.org/endOffset

The end time of the clip expressed as the number of seconds from the beginning of the work.
*/
type EndOffset float64

/*

IsResizable is https://schema.org/isResizable

Whether the 3DModel allows resizing. For example, room layout applications often do not allow 3DModel elements to be resized to reflect reality.
*/
type IsResizable bool

/*

SubStageSuffix is https://schema.org/subStageSuffix

The substage, e.g. 'a' for Stage IIIa.
*/
type SubStageSuffix string

/*

Opens is https://schema.org/opens

The opening hour of the place or service on the given day(s) of the week.
*/
type Opens string

/*

OfferCount is https://schema.org/offerCount

The number of offers for the product.
*/
type OfferCount float64

/*

PublicationType is https://schema.org/publicationType

The type of the medical article, taken from the US NLM MeSH publication type catalog. See also [MeSH documentation](http://www.nlm.nih.gov/mesh/pubtypes.html).
*/
type PublicationType string

/*

CommentText is https://schema.org/commentText

The text of the UserComment.
*/
type CommentText string

/*

JobBenefits is https://schema.org/jobBenefits

Description of benefits associated with the job.
*/
type JobBenefits string

/*

BroadcastServiceTier is https://schema.org/broadcastServiceTier

The type of service required to have access to the channel (e.g. Standard or Premium).
*/
type BroadcastServiceTier string

/*

LegislationJurisdiction is https://schema.org/legislationJurisdiction

The jurisdiction from which the legislation originates.
*/
type LegislationJurisdiction string

/*

PriceValidUntil is https://schema.org/priceValidUntil

The date after which the price is no longer available.
*/
type PriceValidUntil string

/*

Aspect is https://schema.org/aspect

An aspect of medical practice that is considered on the page, such as 'diagnosis', 'treatment', 'causes', 'prognosis', 'etiology', 'epidemiology', etc.
*/
type Aspect string

/*

Map is https://schema.org/map

A URL to a map of the place.
*/
type Map string

/*

AssociatedPathophysiology is https://schema.org/associatedPathophysiology

If applicable, a description of the pathophysiology associated with the anatomical system, including potential abnormal changes in the mechanical, physical, and biochemical functions of the system.
*/
type AssociatedPathophysiology string

/*

ScheduleTimezone is https://schema.org/scheduleTimezone

Indicates the timezone for which the time(s) indicated in the [[Schedule]] are given. The value provided should be among those listed in the IANA Time Zone Database.
*/
type ScheduleTimezone string

/*

Assembly is https://schema.org/assembly

Library file name, e.g., mscorlib.dll, system.web.dll.
*/
type Assembly string

/*

OrderItemNumber is https://schema.org/orderItemNumber

The identifier of the order item.
*/
type OrderItemNumber string

/*

NumberOfSeasons is https://schema.org/numberOfSeasons

The number of seasons in this series.
*/
type NumberOfSeasons float64

/*

DateReceived is https://schema.org/dateReceived

The date/time the message was received if a single recipient exists.
*/
type DateReceived string

/*

StartTime is https://schema.org/startTime

The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. E.g. John wrote a book from *January* to December. For media, including audio and video, it's the time offset of the start of a clip within a larger file.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
*/
type StartTime string

/*

HealthPlanCoinsuranceOption is https://schema.org/healthPlanCoinsuranceOption

Whether the coinsurance applies before or after deductible, etc. TODO: Is this a closed set?
*/
type HealthPlanCoinsuranceOption string

/*

Rxcui is https://schema.org/rxcui

The RxCUI drug identifier from RXNORM.
*/
type Rxcui string

/*

AvailableThrough is https://schema.org/availableThrough

After this date, the item will no longer be available for pickup.
*/
type AvailableThrough string

/*

ProficiencyLevel is https://schema.org/proficiencyLevel

Proficiency needed for this content; expected values: 'Beginner', 'Expert'.
*/
type ProficiencyLevel string

/*

CopyrightYear is https://schema.org/copyrightYear

The year during which the claimed copyright for the CreativeWork was first asserted.
*/
type CopyrightYear float64

/*

InteractivityType is https://schema.org/interactivityType

The predominant mode of learning supported by the learning resource. Acceptable values are 'active', 'expositive', or 'mixed'.
*/
type InteractivityType string

/*

KnowsLanguage is https://schema.org/knowsLanguage

Of a [[Person]], and less typically of an [[Organization]], to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47).
*/
type KnowsLanguage string

/*

PhoneticText is https://schema.org/phoneticText

Representation of a text [[textValue]] using the specified [[speechToTextMarkup]]. For example the city name of Houston in IPA: /ˈhjuːstən/.
*/
type PhoneticText string

/*

TargetDescription is https://schema.org/targetDescription

The description of a node in an established educational framework.
*/
type TargetDescription string

/*

InChI is https://schema.org/inChI

Non-proprietary identifier for molecular entity that can be used in printed and electronic data sources thus enabling easier linking of diverse data compilations.
*/
type InChI string

/*

HasMenu is https://schema.org/hasMenu

Either the actual menu as a structured representation, as text, or a URL of the menu.
*/
type HasMenu string

/*

CvdNumICUBeds is https://schema.org/cvdNumICUBeds

numicubeds - ICU BEDS: Total number of staffed inpatient intensive care unit (ICU) beds.
*/
type CvdNumICUBeds float64

/*

Currency is https://schema.org/currency

The currency in which the monetary amount is expressed.\n\nUse standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217), e.g. "USD"; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies, e.g. "BTC"; well known names for [Local Exchange Trading Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types, e.g. "Ithaca HOUR".
*/
type Currency string

/*

ChildMinAge is https://schema.org/childMinAge

Minimal age of the child.
*/
type ChildMinAge float64

/*

RecommendationStrength is https://schema.org/recommendationStrength

Strength of the guideline's recommendation (e.g. 'class I').
*/
type RecommendationStrength string

/*

ByDay is https://schema.org/byDay

Defines the day(s) of the week on which a recurring [[Event]] takes place. May be specified using either [[DayOfWeek]], or alternatively [[Text]] conforming to iCal's syntax for byDay recurrence rules.
*/
type ByDay string

/*

KnownVehicleDamages is https://schema.org/knownVehicleDamages

A textual description of known damages, both repaired and unrepaired.
*/
type KnownVehicleDamages string

/*

Value is https://schema.org/value

The value of a [[QuantitativeValue]] (including [[Observation]]) or property value node.\n\n* For [[QuantitativeValue]] and [[MonetaryAmount]], the recommended type for values is 'Number'.\n* For [[PropertyValue]], it can be 'Text', 'Number', 'Boolean', or 'StructuredValue'.\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
*/
type Value bool

/*

Risks is https://schema.org/risks

Specific physiologic risks associated to the diet plan.
*/
type Risks string

/*

Discount is https://schema.org/discount

Any discount applied (to an Order).
*/
type Discount string

/*

ExerciseType is https://schema.org/exerciseType

Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
*/
type ExerciseType string

/*

Keywords is https://schema.org/keywords

Keywords or tags used to describe some item. Multiple textual entries in a keywords list are typically delimited by commas, or by repeating the property.
*/
type Keywords string

/*

NaturalProgression is https://schema.org/naturalProgression

The expected progression of the condition if it is not treated and allowed to progress naturally.
*/
type NaturalProgression string

/*

EvidenceOrigin is https://schema.org/evidenceOrigin

Source of the data used to formulate the guidance, e.g. RCT, consensus opinion, etc.
*/
type EvidenceOrigin string

/*

MinPrice is https://schema.org/minPrice

The lowest price if the price is a range.
*/
type MinPrice float64

/*

CopyrightNotice is https://schema.org/copyrightNotice

Text of a notice appropriate for describing the copyright aspects of this Creative Work, ideally indicating the owner of the copyright for the Work.
*/
type CopyrightNotice string

/*

EndTime is https://schema.org/endTime

The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. E.g. John wrote a book from January to *December*. For media, including audio and video, it's the time offset of the end of a clip within a larger file.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
*/
type EndTime string

/*

ArtworkSurface is https://schema.org/artworkSurface

The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc.
*/
type ArtworkSurface string

/*

ExpectedPrognosis is https://schema.org/expectedPrognosis

The likely outcome in either the short term or long term of the medical condition.
*/
type ExpectedPrognosis string

/*

PetsAllowed is https://schema.org/petsAllowed

Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
*/
type PetsAllowed bool

/*

EducationalProgramMode is https://schema.org/educationalProgramMode

Similar to courseMode, the medium or means of delivery of the program as a whole. The value may either be a text label (e.g. "online", "onsite" or "blended"; "synchronous" or "asynchronous"; "full-time" or "part-time") or a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous ).
*/
type EducationalProgramMode string

/*

FlightDistance is https://schema.org/flightDistance

The distance of the flight.
*/
type FlightDistance string

/*

PurchaseDate is https://schema.org/purchaseDate

The date the item, e.g. vehicle, was purchased by the current owner.
*/
type PurchaseDate string

/*

SeatingType is https://schema.org/seatingType

The type/class of the seat.
*/
type SeatingType string

/*

ClinicalPharmacology is https://schema.org/clinicalPharmacology

Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
*/
type ClinicalPharmacology string

/*

TypicalCreditsPerTerm is https://schema.org/typicalCreditsPerTerm

The number of credits or units a full-time student would be expected to take in 1 term however 'term' is defined by the institution.
*/
type TypicalCreditsPerTerm float64

/*

TargetPlatform is https://schema.org/targetPlatform

Type of app development: phone, Metro style, desktop, XBox, etc.
*/
type TargetPlatform string

/*

PrintSection is https://schema.org/printSection

If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
*/
type PrintSection string

/*

BeforeMedia is https://schema.org/beforeMedia

A media object representing the circumstances before performing this direction.
*/
type BeforeMedia string

/*

PaymentDueDate is https://schema.org/paymentDueDate

The date that payment is due.
*/
type PaymentDueDate string

/*

BillingDuration is https://schema.org/billingDuration

Specifies for how long this price (or price component) will be billed. Can be used, for example, to model the contractual duration of a subscription or payment plan. Type can be either a Duration or a Number (in which case the unit of measurement, for example month, is specified by the unitCode property).
*/
type BillingDuration float64

/*

FoundingDate is https://schema.org/foundingDate

The date that this organization was founded.
*/
type FoundingDate string

/*

CoursePrerequisites is https://schema.org/coursePrerequisites

Requirements for taking the Course. May be completion of another [[Course]] or a textual description like "permission of instructor". Requirements may be a pre-requisite competency, referenced using [[AlignmentObject]].
*/
type CoursePrerequisites string

/*

Version is https://schema.org/version

The version of the CreativeWork embodied by a specified resource.
*/
type Version string

/*

IsAccessibleForFree is https://schema.org/isAccessibleForFree

A flag to signal that the item, event, or place is accessible for free.
*/
type IsAccessibleForFree bool

/*

VideoFormat is https://schema.org/videoFormat

The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
*/
type VideoFormat string

/*

VerificationFactCheckingPolicy is https://schema.org/verificationFactCheckingPolicy

Disclosure about verification and fact-checking processes for a [[NewsMediaOrganization]] or other fact-checking [[Organization]].
*/
type VerificationFactCheckingPolicy string

/*

BiomechnicalClass is https://schema.org/biomechnicalClass

The biomechanical properties of the bone.
*/
type BiomechnicalClass string

/*

SecurityScreening is https://schema.org/securityScreening

The type of security screening the passenger is subject to.
*/
type SecurityScreening string

/*

ActionOption is https://schema.org/actionOption

A sub property of object. The options subject to this action.
*/
type ActionOption string

/*

ExchangeRateSpread is https://schema.org/exchangeRateSpread

The difference between the price at which a broker or other intermediary buys and sells foreign currency.
*/
type ExchangeRateSpread float64

/*

VatID is https://schema.org/vatID

The Value-added Tax ID of the organization or person.
*/
type VatID string

/*

VariableMeasured is https://schema.org/variableMeasured

The variableMeasured property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue, or more explicitly as a [[StatisticalVariable]].
*/
type VariableMeasured string

/*

RemainingAttendeeCapacity is https://schema.org/remainingAttendeeCapacity

The number of attendee places for an event that remain unallocated.
*/
type RemainingAttendeeCapacity float64

/*

MolecularWeight is https://schema.org/molecularWeight

This is the molecular weight of the entity being described, not of the parent. Units should be included in the form '&lt;Number&gt; &lt;unit&gt;', for example '12 amu' or as '&lt;QuantitativeValue&gt;.
*/
type MolecularWeight string

/*

IcaoCode is https://schema.org/icaoCode

ICAO identifier for an airport.
*/
type IcaoCode string

/*

EligibilityToWorkRequirement is https://schema.org/eligibilityToWorkRequirement

The legal requirements such as citizenship, visa and other documentation required for an applicant to this job.
*/
type EligibilityToWorkRequirement string

/*

CvdNumTotBeds is https://schema.org/cvdNumTotBeds

numtotbeds - ALL HOSPITAL BEDS: Total number of all inpatient and outpatient beds, including all staffed, ICU, licensed, and overflow (surge) beds used for inpatients or outpatients.
*/
type CvdNumTotBeds float64

/*

BroadcastFrequency is https://schema.org/broadcastFrequency

The frequency used for over-the-air broadcasts. Numeric values or simple ranges, e.g. 87-99. In addition a shortcut idiom is supported for frequencies of AM and FM radio channels, e.g. "87 FM".
*/
type BroadcastFrequency string

/*

SampleType is https://schema.org/sampleType

What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template.
*/
type SampleType string

/*

ReportNumber is https://schema.org/reportNumber

The number or other unique designator assigned to a Report by the publishing organization.
*/
type ReportNumber string

/*

AccessibilitySummary is https://schema.org/accessibilitySummary

A human-readable summary of specific accessibility features or deficiencies, consistent with the other accessibility metadata but expressing subtleties such as "short descriptions are present but long descriptions will be needed for non-visual users" or "short descriptions are present and no long descriptions are needed".
*/
type AccessibilitySummary string

/*

TargetName is https://schema.org/targetName

The name of a node in an established educational framework.
*/
type TargetName string

/*

IsInvolvedInBiologicalProcess is https://schema.org/isInvolvedInBiologicalProcess

Biological process this BioChemEntity is involved in; please use PropertyValue if you want to include any evidence.
*/
type IsInvolvedInBiologicalProcess string

/*

BenefitsSummaryUrl is https://schema.org/benefitsSummaryUrl

The URL that goes directly to the summary of benefits and coverage for the specific standard plan or plan variation.
*/
type BenefitsSummaryUrl string

/*

Caption is https://schema.org/caption

The caption for this object. For downloadable machine formats (closed caption, subtitles etc.) use MediaObject and indicate the [[encodingFormat]].
*/
type Caption string

/*

NumberOfLoanPayments is https://schema.org/numberOfLoanPayments

The number of payments contractually required at origination to repay the loan. For monthly paying loans this is the number of months from the contractual first payment date to the maturity date.
*/
type NumberOfLoanPayments float64

/*

Dateline is https://schema.org/dateline

A [dateline](https://en.wikipedia.org/wiki/Dateline) is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.

Structured representations of dateline-related information can also be expressed more explicitly using [[locationCreated]] (which represents where a work was created, e.g. where a news report was written).  For location depicted or described in the content, use [[contentLocation]].

Dateline summaries are oriented more towards human readers than towards automated processing, and can vary substantially. Some examples: "BEIRUT, Lebanon, June 2.", "Paris, France", "December 19, 2017 11:43AM Reporting from Washington", "Beijing/Moscow", "QUEZON CITY, Philippines".
      
*/
type Dateline string

/*

IsAvailableGenerically is https://schema.org/isAvailableGenerically

True if the drug is available in a generic form (regardless of name).
*/
type IsAvailableGenerically bool

/*

ReleaseNotes is https://schema.org/releaseNotes

Description of what changed in this version.
*/
type ReleaseNotes string

/*

PageStart is https://schema.org/pageStart

The page on which the work starts; for example "135" or "xiii".
*/
type PageStart float64

/*

TransmissionMethod is https://schema.org/transmissionMethod

How the disease spreads, either as a route or vector, for example 'direct contact', 'Aedes aegypti', etc.
*/
type TransmissionMethod string

/*

PositiveNotes is https://schema.org/positiveNotes

Provides positive considerations regarding something, for example product highlights or (alongside [[negativeNotes]]) pro/con lists for reviews.

In the case of a [[Review]], the property describes the [[itemReviewed]] from the perspective of the review; in the case of a [[Product]], the product itself is being described.

The property values can be expressed either as unstructured text (repeated as necessary), or if ordered, as a list (in which case the most positive is at the beginning of the list).
*/
type PositiveNotes string

/*

MobileUrl is https://schema.org/mobileUrl

The [[mobileUrl]] property is provided for specific situations in which data consumers need to determine whether one of several provided URLs is a dedicated 'mobile site'.

To discourage over-use, and reflecting intial usecases, the property is expected only on [[Product]] and [[Offer]], rather than [[Thing]]. The general trend in web technology is towards [responsive design](https://en.wikipedia.org/wiki/Responsive_web_design) in which content can be flexibly adapted to a wide range of browsing environments. Pages and sites referenced with the long-established [[url]] property should ideally also be usable on a wide variety of devices, including mobile phones. In most cases, it would be pointless and counter productive to attempt to update all [[url]] markup to use [[mobileUrl]] for more mobile-oriented pages. The property is intended for the case when items (primarily [[Product]] and [[Offer]]) have extra URLs hosted on an additional "mobile site" alongside the main one. It should not be taken as an endorsement of this publication style.
    
*/
type MobileUrl string

/*

EducationalFramework is https://schema.org/educationalFramework

The framework to which the resource being described is aligned.
*/
type EducationalFramework string

/*

EducationalLevel is https://schema.org/educationalLevel

The level in terms of progression through an educational or training context. Examples of educational levels include 'beginner', 'intermediate' or 'advanced', and formal sets of level indicators.
*/
type EducationalLevel string

/*

Elevation is https://schema.org/elevation

The elevation of a location ([WGS 84](https://en.wikipedia.org/wiki/World_Geodetic_System)). Values may be of the form 'NUMBER UNIT\_OF\_MEASUREMENT' (e.g., '1,000 m', '3,200 ft') while numbers alone should be assumed to be a value in meters.
*/
type Elevation float64

/*

ObservationPeriod is https://schema.org/observationPeriod

The length of time an Observation took place over. The format follows `P[0-9]*[Y|M|D|h|m|s]`. For example, P1Y is Period 1 Year, P3M is Period 3 Months, P3h is Period 3 hours.
*/
type ObservationPeriod string

/*

ProcessorRequirements is https://schema.org/processorRequirements

Processor architecture required to run the application (e.g. IA64).
*/
type ProcessorRequirements string

/*

ArrivalTerminal is https://schema.org/arrivalTerminal

Identifier of the flight's arrival terminal.
*/
type ArrivalTerminal string

/*

PhysicalRequirement is https://schema.org/physicalRequirement

A description of the types of physical activity associated with the job. Defined terms such as those in O*net may be used, but note that there is no way to specify the level of ability as well as its nature when using a defined term.
*/
type PhysicalRequirement string

/*

FeatureList is https://schema.org/featureList

Features or modules provided by this application (and possibly required by other applications).
*/
type FeatureList string

/*

RatingCount is https://schema.org/ratingCount

The count of total number of ratings.
*/
type RatingCount float64

/*

CourseCode is https://schema.org/courseCode

The identifier for the [[Course]] used by the course [[provider]] (e.g. CS101 or 6.001).
*/
type CourseCode string

/*

UsNPI is https://schema.org/usNPI

A <a href="https://en.wikipedia.org/wiki/National_Provider_Identifier">National Provider Identifier</a> (NPI) 
    is a unique 10-digit identification number issued to health care providers in the United States by the Centers for Medicare and Medicaid Services.
*/
type UsNPI string

/*

ShippingLabel is https://schema.org/shippingLabel

Label to match an [[OfferShippingDetails]] with a [[ShippingRateSettings]] (within the context of a [[shippingSettingsLink]] cross-reference).
*/
type ShippingLabel string

/*

TargetUrl is https://schema.org/targetUrl

The URL of a node in an established educational framework.
*/
type TargetUrl string

/*

LoanPaymentFrequency is https://schema.org/loanPaymentFrequency

Frequency of payments due, i.e. number of months between payments. This is defined as a frequency, i.e. the reciprocal of a period of time.
*/
type LoanPaymentFrequency float64

/*

BusName is https://schema.org/busName

The name of the bus (e.g. Bolt Express).
*/
type BusName string

/*

InStoreReturnsOffered is https://schema.org/inStoreReturnsOffered

Are in-store returns offered? (For more advanced return methods use the [[returnMethod]] property.)
*/
type InStoreReturnsOffered bool

/*

Benefits is https://schema.org/benefits

Description of benefits associated with the job.
*/
type Benefits string

/*

FinancialAidEligible is https://schema.org/financialAidEligible

A financial aid type or program which students may use to pay for tuition or fees associated with the program.
*/
type FinancialAidEligible string

/*

Sha256 is https://schema.org/sha256

The [SHA-2](https://en.wikipedia.org/wiki/SHA-2) SHA256 hash of the content of the item. For example, a zero-length input has value 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'.
*/
type Sha256 string

/*

ValueAddedTaxIncluded is https://schema.org/valueAddedTaxIncluded

Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
*/
type ValueAddedTaxIncluded bool

/*

Latitude is https://schema.org/latitude

The latitude of a location. For example ```37.42242``` ([WGS 84](https://en.wikipedia.org/wiki/World_Geodetic_System)).
*/
type Latitude string

/*

VideoFrameSize is https://schema.org/videoFrameSize

The frame size of the video.
*/
type VideoFrameSize string

/*

Title is https://schema.org/title

The title of the job.
*/
type Title string

/*

IssueNumber is https://schema.org/issueNumber

Identifies the issue of publication; for example, "iii" or "2".
*/
type IssueNumber float64

/*

EstimatedFlightDuration is https://schema.org/estimatedFlightDuration

The estimated time the flight will take.
*/
type EstimatedFlightDuration string

/*

ProgrammingModel is https://schema.org/programmingModel

Indicates whether API is managed or unmanaged.
*/
type ProgrammingModel string

/*

IsLocatedInSubcellularLocation is https://schema.org/isLocatedInSubcellularLocation

Subcellular location where this BioChemEntity is located; please use PropertyValue if you want to include any evidence.
*/
type IsLocatedInSubcellularLocation string

/*

BrowserRequirements is https://schema.org/browserRequirements

Specifies browser requirements in human-readable text. For example, 'requires HTML5 support'.
*/
type BrowserRequirements string

/*

PrescribingInfo is https://schema.org/prescribingInfo

Link to prescribing information for the drug.
*/
type PrescribingInfo string

/*

IsicV4 is https://schema.org/isicV4

The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
*/
type IsicV4 string

/*

VolumeNumber is https://schema.org/volumeNumber

Identifies the volume of publication or multi-part work; for example, "iii" or "2".
*/
type VolumeNumber float64

/*

Iso6523Code is https://schema.org/iso6523Code

An organization identifier as defined in [ISO 6523(-1)](https://en.wikipedia.org/wiki/ISO/IEC_6523). The identifier should be in the `XXXX:YYYYYY:ZZZ` or `XXXX:YYYYYY`format. Where `XXXX` is a 4 digit _ICD_ (International Code Designator), `YYYYYY` is an _OID_ (Organization Identifier) with all formatting characters (dots, dashes, spaces) removed with a maximal length of 35 characters, and `ZZZ` is an optional OPI (Organization Part Identifier) with a maximum length of 35 characters. The various components (ICD, OID, OPI) are joined with a colon character (ASCII `0x3a`). Note that many existing organization identifiers defined as attributes like [leiCode](http://schema.org/leiCode) (`0199`), [duns](http://schema.org/duns) (`0060`) or [GLN](http://schema.org/globalLocationNumber) (`0088`) can be expressed using ISO-6523. If possible, ISO-6523 codes should be preferred to populating [vatID](http://schema.org/vatID) or [taxID](http://schema.org/taxID), as ISO identifiers are less ambiguous.
*/
type Iso6523Code string

/*

PriceType is https://schema.org/priceType

Defines the type of a price specified for an offered product, for example a list price, a (temporary) sale price or a manufacturer suggested retail price. If multiple prices are specified for an offer the [[priceType]] property can be used to identify the type of each such specified price. The value of priceType can be specified as a value from enumeration PriceTypeEnumeration or as a free form text string for price types that are not already predefined in PriceTypeEnumeration.
*/
type PriceType string

/*

CommentCount is https://schema.org/commentCount

The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
*/
type CommentCount float64

/*

FlightNumber is https://schema.org/flightNumber

The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is 'UA', the flightNumber is 'UA110'.
*/
type FlightNumber string

/*

Email is https://schema.org/email

Email address.
*/
type Email string

/*

PostOfficeBoxNumber is https://schema.org/postOfficeBoxNumber

The post office box number for PO box addresses.
*/
type PostOfficeBoxNumber string

/*

SchoolClosuresInfo is https://schema.org/schoolClosuresInfo

Information about school closures.
*/
type SchoolClosuresInfo string

/*

ItemListOrder is https://schema.org/itemListOrder

Type of ordering (e.g. Ascending, Descending, Unordered).
*/
type ItemListOrder string

/*

SalaryCurrency is https://schema.org/salaryCurrency

The currency (coded using [ISO 4217](http://en.wikipedia.org/wiki/ISO_4217)) used for the main salary information in this job posting or for this employee.
*/
type SalaryCurrency string

/*

DiscountCurrency is https://schema.org/discountCurrency

The currency of the discount.\n\nUse standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217), e.g. "USD"; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies, e.g. "BTC"; well known names for [Local Exchange Trading Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types, e.g. "Ithaca HOUR".
*/
type DiscountCurrency string

/*

AlternativeHeadline is https://schema.org/alternativeHeadline

A secondary title of the CreativeWork.
*/
type AlternativeHeadline string

/*

Industry is https://schema.org/industry

The industry associated with the job position.
*/
type Industry string

/*

ServingSize is https://schema.org/servingSize

The serving size, in terms of the number of volume or mass.
*/
type ServingSize string

/*

Contraindication is https://schema.org/contraindication

A contraindication for this therapy.
*/
type Contraindication string

/*

MinValue is https://schema.org/minValue

The lower value of some characteristic or property.
*/
type MinValue float64

/*

IsUnlabelledFallback is https://schema.org/isUnlabelledFallback

This can be marked 'true' to indicate that some published [[DeliveryTimeSettings]] or [[ShippingRateSettings]] are intended to apply to all [[OfferShippingDetails]] published by the same merchant, when referenced by a [[shippingSettingsLink]] in those settings. It is not meaningful to use a 'true' value for this property alongside a transitTimeLabel (for [[DeliveryTimeSettings]]) or shippingLabel (for [[ShippingRateSettings]]), since this property is for use with unlabelled settings.
*/
type IsUnlabelledFallback bool

/*

CreativeWorkStatus is https://schema.org/creativeWorkStatus

The status of a creative work in terms of its stage in a lifecycle. Example terms include Incomplete, Draft, Published, Obsolete. Some organizations define a set of terms for the stages of their publication lifecycle.
*/
type CreativeWorkStatus string

/*

RequiredGender is https://schema.org/requiredGender

Audiences defined by a person's gender.
*/
type RequiredGender string

/*

ChildMaxAge is https://schema.org/childMaxAge

Maximal age of the child.
*/
type ChildMaxAge float64

/*

LayoutImage is https://schema.org/layoutImage

A schematic image showing the floorplan layout.
*/
type LayoutImage string

/*

InfectiousAgent is https://schema.org/infectiousAgent

The actual infectious agent, such as a specific bacterium.
*/
type InfectiousAgent string

/*

SuggestedGender is https://schema.org/suggestedGender

The suggested gender of the intended person or audience, for example "male", "female", or "unisex".
*/
type SuggestedGender string

/*

TypeOfBed is https://schema.org/typeOfBed

The type of bed to which the BedDetail refers, i.e. the type of bed available in the quantity indicated by quantity.
*/
type TypeOfBed string

/*

ServiceUrl is https://schema.org/serviceUrl

The website to access the service.
*/
type ServiceUrl string

/*

Median is https://schema.org/median

The median value.
*/
type Median float64

/*

DepartureTerminal is https://schema.org/departureTerminal

Identifier of the flight's departure terminal.
*/
type DepartureTerminal string

/*

CvdNumVent is https://schema.org/cvdNumVent

numvent - MECHANICAL VENTILATORS: Total number of ventilators available.
*/
type CvdNumVent float64

/*

SignificantLink is https://schema.org/significantLink

One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
*/
type SignificantLink string

/*

MaximumPhysicalAttendeeCapacity is https://schema.org/maximumPhysicalAttendeeCapacity

The maximum physical attendee capacity of an [[Event]] whose [[eventAttendanceMode]] is [[OfflineEventAttendanceMode]] (or the offline aspects, in the case of a [[MixedEventAttendanceMode]]). 
*/
type MaximumPhysicalAttendeeCapacity float64

/*

PublicAccess is https://schema.org/publicAccess

A flag to signal that the [[Place]] is open to public visitors.  If this property is omitted there is no assumed default boolean value.
*/
type PublicAccess bool

/*

ApplicationSubCategory is https://schema.org/applicationSubCategory

Subcategory of the application, e.g. 'Arcade Game'.
*/
type ApplicationSubCategory string

/*

Utterances is https://schema.org/utterances

Text of an utterances (spoken words, lyrics etc.) that occurs at a certain section of a media object, represented as a [[HyperTocEntry]].
*/
type Utterances string

/*

PaymentAccepted is https://schema.org/paymentAccepted

Cash, Credit Card, Cryptocurrency, Local Exchange Tradings System, etc.
*/
type PaymentAccepted string

/*

TotalJobOpenings is https://schema.org/totalJobOpenings

The number of positions open for this job posting. Use a positive integer. Do not use if the number of positions is unclear or not known.
*/
type TotalJobOpenings float64

/*

Line is https://schema.org/line

A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
*/
type Line string

/*

ReviewCount is https://schema.org/reviewCount

The count of total number of reviews.
*/
type ReviewCount float64

/*

CourseMode is https://schema.org/courseMode

The medium or means of delivery of the course instance or the mode of study, either as a text label (e.g. "online", "onsite" or "blended"; "synchronous" or "asynchronous"; "full-time" or "part-time") or as a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous).
*/
type CourseMode string

/*

Frequency is https://schema.org/frequency

How often the dose is taken, e.g. 'daily'.
*/
type Frequency string

/*

EmissionsCO2 is https://schema.org/emissionsCO2

The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put "g/km" into the unitText property of that value, since there is no UN/CEFACT Common Code for "g/km".
*/
type EmissionsCO2 float64

/*

Aircraft is https://schema.org/aircraft

The kind of aircraft (e.g., "Boeing 747").
*/
type Aircraft string

/*

Percentile90 is https://schema.org/percentile90

The 90th percentile value.
*/
type Percentile90 float64

/*

AlcoholWarning is https://schema.org/alcoholWarning

Any precaution, guidance, contraindication, etc. related to consumption of alcohol while taking this drug.
*/
type AlcoholWarning string

/*

MechanismOfAction is https://schema.org/mechanismOfAction

The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
*/
type MechanismOfAction string

/*

UnitText is https://schema.org/unitText

A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
<a href='unitCode'>unitCode</a>.
*/
type UnitText string

/*

RestockingFee is https://schema.org/restockingFee

Use [[MonetaryAmount]] to specify a fixed restocking fee for product returns, or use [[Number]] to specify a percentage of the product price paid by the customer.
*/
type RestockingFee float64

/*

TypicalAgeRange is https://schema.org/typicalAgeRange

The typical expected age range, e.g. '7-9', '11-'.
*/
type TypicalAgeRange string

/*

EngineType is https://schema.org/engineType

The type of engine or engines powering the vehicle.
*/
type EngineType string

/*

GlobalLocationNumber is https://schema.org/globalLocationNumber

The [Global Location Number](http://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
*/
type GlobalLocationNumber string

/*

EncodingType is https://schema.org/encodingType

The supported encoding type(s) for an EntryPoint request.
*/
type EncodingType string

/*

Menu is https://schema.org/menu

Either the actual menu as a structured representation, as text, or a URL of the menu.
*/
type Menu string

/*

Followup is https://schema.org/followup

Typical or recommended followup care after the procedure is performed.
*/
type Followup string

/*

FileSize is https://schema.org/fileSize

Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
*/
type FileSize string

/*

Gtin14 is https://schema.org/gtin14

The GTIN-14 code of the product, or the product to which the offer refers. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
*/
type Gtin14 string

/*

OffersPrescriptionByMail is https://schema.org/offersPrescriptionByMail

Whether prescriptions can be delivered by mail.
*/
type OffersPrescriptionByMail bool

/*

RecipeInstructions is https://schema.org/recipeInstructions

A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
*/
type RecipeInstructions string

/*

TermsOfService is https://schema.org/termsOfService

Human-readable terms of service documentation.
*/
type TermsOfService string

/*

Slogan is https://schema.org/slogan

A slogan or motto associated with the item.
*/
type Slogan string

/*

RecipeYield is https://schema.org/recipeYield

The quantity produced by the recipe (for example, number of people served, number of servings, etc).
*/
type RecipeYield string

/*

GamePlatform is https://schema.org/gamePlatform

The electronic systems used to play <a href="http://en.wikipedia.org/wiki/Category:Video_game_platforms">video games</a>.
*/
type GamePlatform string

/*

DoseValue is https://schema.org/doseValue

The value of the dose, e.g. 500.
*/
type DoseValue float64

/*

MonoisotopicMolecularWeight is https://schema.org/monoisotopicMolecularWeight

The monoisotopic mass is the sum of the masses of the atoms in a molecule using the unbound, ground-state, rest mass of the principal (most abundant) isotope for each element instead of the isotopic average mass. Please include the units in the form '&lt;Number&gt; &lt;unit&gt;', for example '770.230488 g/mol' or as '&lt;QuantitativeValue&gt;.
*/
type MonoisotopicMolecularWeight string

/*

GuidelineDate is https://schema.org/guidelineDate

Date on which this guideline's recommendation was made.
*/
type GuidelineDate string

/*

DeparturePlatform is https://schema.org/departurePlatform

The platform from which the train departs.
*/
type DeparturePlatform string

/*

ContentRating is https://schema.org/contentRating

Official rating of a piece of content&#x2014;for example, 'MPAA PG-13'.
*/
type ContentRating string

/*

MaximumVirtualAttendeeCapacity is https://schema.org/maximumVirtualAttendeeCapacity

The maximum virtual attendee capacity of an [[Event]] whose [[eventAttendanceMode]] is [[OnlineEventAttendanceMode]] (or the online aspects, in the case of a [[MixedEventAttendanceMode]]). 
*/
type MaximumVirtualAttendeeCapacity float64

/*

ValuePattern is https://schema.org/valuePattern

Specifies a regular expression for testing literal values according to the HTML spec.
*/
type ValuePattern string

/*

BreastfeedingWarning is https://schema.org/breastfeedingWarning

Any precaution, guidance, contraindication, etc. related to this drug's use by breastfeeding mothers.
*/
type BreastfeedingWarning string

/*

CoverageStartTime is https://schema.org/coverageStartTime

The time when the live blog will begin covering the Event. Note that coverage may begin before the Event's start time. The LiveBlogPosting may also be created before coverage begins.
*/
type CoverageStartTime string

/*

SuggestedMaxAge is https://schema.org/suggestedMaxAge

Maximum recommended age in years for the audience or user.
*/
type SuggestedMaxAge float64

/*

HealthPlanCopayOption is https://schema.org/healthPlanCopayOption

Whether the copay is before or after deductible, etc. TODO: Is this a closed set?
*/
type HealthPlanCopayOption string

/*

ByMonthDay is https://schema.org/byMonthDay

Defines the day(s) of the month on which a recurring [[Event]] takes place. Specified as an [[Integer]] between 1-31.
*/
type ByMonthDay float64

/*

AccessibilityHazard is https://schema.org/accessibilityHazard

A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3. Values should be drawn from the [approved vocabulary](https://www.w3.org/2021/a11y-discov-vocab/latest/#accessibilityHazard-vocabulary).
*/
type AccessibilityHazard string

/*

DateVehicleFirstRegistered is https://schema.org/dateVehicleFirstRegistered

The date of the first registration of the vehicle with the respective public authorities.
*/
type DateVehicleFirstRegistered string

/*

PostalCodePrefix is https://schema.org/postalCodePrefix

A defined range of postal codes indicated by a common textual prefix. Used for non-numeric systems such as UK.
*/
type PostalCodePrefix string

/*

AnnualPercentageRate is https://schema.org/annualPercentageRate

The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
*/
type AnnualPercentageRate float64

/*

Text is https://schema.org/text

The textual content of this CreativeWork.
*/
type Text string

/*

ProgrammingLanguage is https://schema.org/programmingLanguage

The computer programming language.
*/
type ProgrammingLanguage string

/*

HasDriveThroughService is https://schema.org/hasDriveThroughService

Indicates whether some facility (e.g. [[FoodEstablishment]], [[CovidTestingFacility]]) offers a service that can be used by driving through in a car. In the case of [[CovidTestingFacility]] such facilities could potentially help with social distancing from other potentially-infected users.
*/
type HasDriveThroughService bool

/*

VehicleInteriorColor is https://schema.org/vehicleInteriorColor

The color or color combination of the interior of the vehicle.
*/
type VehicleInteriorColor string

/*

DuringMedia is https://schema.org/duringMedia

A media object representing the circumstances while performing this direction.
*/
type DuringMedia string

/*

SerialNumber is https://schema.org/serialNumber

The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
*/
type SerialNumber string

/*

ProductReturnLink is https://schema.org/productReturnLink

Indicates a Web page or service by URL, for product return.
*/
type ProductReturnLink string

/*

ProviderMobility is https://schema.org/providerMobility

Indicates the mobility of a provided service (e.g. 'static', 'dynamic').
*/
type ProviderMobility string

/*

Amount is https://schema.org/amount

The amount of money.
*/
type Amount float64

/*

NormalRange is https://schema.org/normalRange

Range of acceptable values for a typical patient, when applicable.
*/
type NormalRange string

/*

Free is https://schema.org/free

A flag to signal that the item, event, or place is accessible for free.
*/
type Free bool

/*

Identifier is https://schema.org/identifier

The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
        
*/
type Identifier string

/*

PublishingPrinciples is https://schema.org/publishingPrinciples

The publishingPrinciples property indicates (typically via [[URL]]) a document describing the editorial principles of an [[Organization]] (or individual, e.g. a [[Person]] writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a [[CreativeWork]] (e.g. [[NewsArticle]]) the principles are those of the party primarily responsible for the creation of the [[CreativeWork]].

While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a [[funder]]) can be expressed using schema.org terminology.

*/
type PublishingPrinciples string

/*

UrlTemplate is https://schema.org/urlTemplate

An url template (RFC6570) that will be used to construct the target of the execution of the action.
*/
type UrlTemplate string

/*

PostOp is https://schema.org/postOp

A description of the postoperative procedures, care, and/or followups for this device.
*/
type PostOp string

/*

LabelDetails is https://schema.org/labelDetails

Link to the drug's label details.
*/
type LabelDetails string

/*

CvdNumC19OverflowPats is https://schema.org/cvdNumC19OverflowPats

numc19overflowpats - ED/OVERFLOW: Patients with suspected or confirmed COVID-19 who are in the ED or any overflow location awaiting an inpatient bed.
*/
type CvdNumC19OverflowPats float64

/*

Award is https://schema.org/award

An award won by or for this item.
*/
type Award string

/*

DatasetTimeInterval is https://schema.org/datasetTimeInterval

The range of temporal applicability of a dataset, e.g. for a 2011 census dataset, the year 2011 (in ISO 8601 time interval format).
*/
type DatasetTimeInterval string

/*

MonthlyMinimumRepaymentAmount is https://schema.org/monthlyMinimumRepaymentAmount

The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
*/
type MonthlyMinimumRepaymentAmount float64

/*

VariesBy is https://schema.org/variesBy

Indicates the property or properties by which the variants in a [[ProductGroup]] vary, e.g. their size, color etc. Schema.org properties can be referenced by their short name e.g. "color"; terms defined elsewhere can be referenced with their URIs.
*/
type VariesBy string

/*

MainEntityOfPage is https://schema.org/mainEntityOfPage

Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
*/
type MainEntityOfPage string

/*

Option is https://schema.org/option

A sub property of object. The options subject to this action.
*/
type Option string

/*

Citation is https://schema.org/citation

A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
*/
type Citation string

/*

NumAdults is https://schema.org/numAdults

The number of adults staying in the unit.
*/
type NumAdults float64

/*

DiversityStaffingReport is https://schema.org/diversityStaffingReport

For an [[Organization]] (often but not necessarily a [[NewsMediaOrganization]]), a report on staffing diversity issues. In a news context this might be for example ASNE or RTDNA (US) reports, or self-reported.
*/
type DiversityStaffingReport string

/*

PassengerSequenceNumber is https://schema.org/passengerSequenceNumber

The passenger's sequence number as assigned by the airline.
*/
type PassengerSequenceNumber string

/*

StatType is https://schema.org/statType

Indicates the kind of statistic represented by a [[StatisticalVariable]], e.g. mean, count etc. The value of statType is a property, either from within Schema.org (e.g. [[median]], [[marginOfError]], [[maxValue]], [[minValue]]) or from other compatible (e.g. RDF) systems such as DataCommons.org or Wikidata.org. 
*/
type StatType string

/*

PublicTransportClosuresInfo is https://schema.org/publicTransportClosuresInfo

Information about public transport closures.
*/
type PublicTransportClosuresInfo string

/*

Xpath is https://schema.org/xpath

An XPath, e.g. of a [[SpeakableSpecification]] or [[WebPageElement]]. In the latter case, multiple matches within a page can constitute a single conceptual "Web page element".
*/
type Xpath string

/*

AudienceType is https://schema.org/audienceType

The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
*/
type AudienceType string

/*

SoftwareVersion is https://schema.org/softwareVersion

Version of the software instance.
*/
type SoftwareVersion string

/*

NumberOfFullBathrooms is https://schema.org/numberOfFullBathrooms

Number of full bathrooms - The total number of full and ¾ bathrooms in an [[Accommodation]]. This corresponds to the [BathroomsFull field in RESO](https://ddwiki.reso.org/display/DDW17/BathroomsFull+Field).
*/
type NumberOfFullBathrooms float64

/*

CoverageEndTime is https://schema.org/coverageEndTime

The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
*/
type CoverageEndTime string

/*

PossibleComplication is https://schema.org/possibleComplication

A possible unexpected and unfavorable evolution of a medical condition. Complications may include worsening of the signs or symptoms of the disease, extension of the condition to other organ systems, etc.
*/
type PossibleComplication string

/*

CallSign is https://schema.org/callSign

A [callsign](https://en.wikipedia.org/wiki/Call_sign), as used in broadcasting and radio communications to identify people, radio and TV stations, or vehicles.
*/
type CallSign string

/*

WebFeed is https://schema.org/webFeed

The URL for a feed, e.g. associated with a podcast series, blog, or series of date-stamped updates. This is usually RSS or Atom.
*/
type WebFeed string

/*

VehicleSeatingCapacity is https://schema.org/vehicleSeatingCapacity

The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.\n\nTypical unit code(s): C62 for persons.
*/
type VehicleSeatingCapacity float64

/*

Runtime is https://schema.org/runtime

Runtime platform or script interpreter dependencies (example: Java v1, Python 2.3, .NET Framework 3.0).
*/
type Runtime string

/*

RecourseLoan is https://schema.org/recourseLoan

The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
*/
type RecourseLoan bool

/*

ContentReferenceTime is https://schema.org/contentReferenceTime

The specific time described by a creative work, for works (e.g. articles, video objects etc.) that emphasise a particular moment within an Event.
*/
type ContentReferenceTime string

/*

MerchantReturnDays is https://schema.org/merchantReturnDays

Specifies either a fixed return date or the number of days (from the delivery date) that a product can be returned. Used when the [[returnPolicyCategory]] property is specified as [[MerchantReturnFiniteReturnWindow]].
*/
type MerchantReturnDays float64

/*

TicketToken is https://schema.org/ticketToken

Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
*/
type TicketToken string

/*

ScheduledPaymentDate is https://schema.org/scheduledPaymentDate

The date the invoice is scheduled to be paid.
*/
type ScheduledPaymentDate string

/*

ShippingSettingsLink is https://schema.org/shippingSettingsLink

Link to a page containing [[ShippingRateSettings]] and [[DeliveryTimeSettings]] details.
*/
type ShippingSettingsLink string

/*

Query is https://schema.org/query

A sub property of instrument. The query used on this action.
*/
type Query string

/*

AfterMedia is https://schema.org/afterMedia

A media object representing the circumstances after performing this direction.
*/
type AfterMedia string

/*

VariantCover is https://schema.org/variantCover

A description of the variant cover
    	for the issue, if the issue is a variant printing. For example, "Bryan Hitch
    	Variant Cover" or "2nd Printing Variant".
*/
type VariantCover string

/*

Bitrate is https://schema.org/bitrate

The bitrate of the media object.
*/
type Bitrate string

/*

NumberOfPartialBathrooms is https://schema.org/numberOfPartialBathrooms

Number of partial bathrooms - The total number of half and ¼ bathrooms in an [[Accommodation]]. This corresponds to the [BathroomsPartial field in RESO](https://ddwiki.reso.org/display/DDW17/BathroomsPartial+Field). 
*/
type NumberOfPartialBathrooms float64

/*

SdLicense is https://schema.org/sdLicense

A license document that applies to this structured data, typically indicated by URL.
*/
type SdLicense string

/*

IsLiveBroadcast is https://schema.org/isLiveBroadcast

True if the broadcast is of a live event.
*/
type IsLiveBroadcast bool

/*

EstimatedCost is https://schema.org/estimatedCost

The estimated cost of the supply or supplies consumed when performing instructions.
*/
type EstimatedCost string

/*

ModifiedTime is https://schema.org/modifiedTime

The date and time the reservation was modified.
*/
type ModifiedTime string

/*

OwnershipFundingInfo is https://schema.org/ownershipFundingInfo

For an [[Organization]] (often but not necessarily a [[NewsMediaOrganization]]), a description of organizational ownership structure; funding and grants. In a news/media setting, this is with particular reference to editorial independence.   Note that the [[funder]] is also available and can be used to make basic funder information machine-readable.
*/
type OwnershipFundingInfo string

/*

Gtin13 is https://schema.org/gtin13

The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceding zero. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
*/
type Gtin13 string

/*

Duns is https://schema.org/duns

The Dun & Bradstreet DUNS number for identifying an organization or business person.
*/
type Duns string

/*

Nsn is https://schema.org/nsn

Indicates the [NATO stock number](https://en.wikipedia.org/wiki/NATO_Stock_Number) (nsn) of a [[Product]]. 
*/
type Nsn string

/*

ProductGroupID is https://schema.org/productGroupID

Indicates a textual identifier for a ProductGroup.
*/
type ProductGroupID string

/*

InProductGroupWithID is https://schema.org/inProductGroupWithID

Indicates the [[productGroupID]] for a [[ProductGroup]] that this product [[isVariantOf]]. 
*/
type InProductGroupWithID string

/*

DissolutionDate is https://schema.org/dissolutionDate

The date that this organization was dissolved.
*/
type DissolutionDate string

/*

MeasurementMethod is https://schema.org/measurementMethod

A subproperty of [[measurementTechnique]] that can be used for specifying specific methods, in particular via [[MeasurementMethodEnum]].
*/
type MeasurementMethod string

/*

DisambiguatingDescription is https://schema.org/disambiguatingDescription

A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
*/
type DisambiguatingDescription string

/*

Location is https://schema.org/location

The location of, for example, where an event is happening, where an organization is located, or where an action takes place.
*/
type Location string

/*

Circle is https://schema.org/circle

A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
*/
type Circle string

/*

BroadcastChannelId is https://schema.org/broadcastChannelId

The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
*/
type BroadcastChannelId string

/*

EduQuestionType is https://schema.org/eduQuestionType

For questions that are part of learning resources (e.g. Quiz), eduQuestionType indicates the format of question being given. Example: "Multiple choice", "Open ended", "Flashcard".
*/
type EduQuestionType string

/*

Size is https://schema.org/size

A standardized size of a product or creative work, specified either through a simple textual string (for example 'XL', '32Wx34L'), a  QuantitativeValue with a unitCode, or a comprehensive and structured [[SizeSpecification]]; in other cases, the [[width]], [[height]], [[depth]] and [[weight]] properties may be more applicable. 
*/
type Size string

/*

LegislationType is https://schema.org/legislationType

The type of the legislation. Examples of values are "law", "act", "directive", "decree", "regulation", "statutory instrument", "loi organique", "règlement grand-ducal", etc., depending on the country.
*/
type LegislationType string

/*

IsProprietary is https://schema.org/isProprietary

True if this item's name is a proprietary/brand name (vs. generic name).
*/
type IsProprietary bool

/*

Gtin12 is https://schema.org/gtin12

The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
*/
type Gtin12 string

/*

Model is https://schema.org/model

The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
*/
type Model string

/*

AreaServed is https://schema.org/areaServed

The geographic area where a service or offered item is provided.
*/
type AreaServed string

/*

Transcript is https://schema.org/transcript

If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
*/
type Transcript string

/*

ApplicableCountry is https://schema.org/applicableCountry

A country where a particular merchant return policy applies to, for example the two-letter ISO 3166-1 alpha-2 country code.
*/
type ApplicableCountry string

/*

Box is https://schema.org/box

A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
*/
type Box string

/*

TrackingUrl is https://schema.org/trackingUrl

Tracking url for the parcel delivery.
*/
type TrackingUrl string

/*

AdditionalType is https://schema.org/additionalType

An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. Typically the value is a URI-identified RDF class, and in this case corresponds to the
    use of rdf:type in RDF. Text values can be used sparingly, for cases where useful information can be added without their being an appropriate schema to reference. In the case of text values, the class label should follow the schema.org <a href="http://schema.org/docs/styleguide.html">style guide</a>.
*/
type AdditionalType string

/*

HealthPlanDrugOption is https://schema.org/healthPlanDrugOption

TODO.
*/
type HealthPlanDrugOption string

/*

MaterialExtent is https://schema.org/materialExtent

The quantity of the materials being described or an expression of the physical space they occupy.
*/
type MaterialExtent string

/*

WorstRating is https://schema.org/worstRating

The lowest value allowed in this rating system.
*/
type WorstRating float64

/*

ReleaseDate is https://schema.org/releaseDate

The release date of a product or product model. This can be used to distinguish the exact variant of a product.
*/
type ReleaseDate string

/*

StartOffset is https://schema.org/startOffset

The start time of the clip expressed as the number of seconds from the beginning of the work.
*/
type StartOffset float64

/*

IsAcceptingNewPatients is https://schema.org/isAcceptingNewPatients

Whether the provider is accepting new patients.
*/
type IsAcceptingNewPatients bool

/*

MathExpression is https://schema.org/mathExpression

A mathematical expression (e.g. 'x^2-3x=0') that may be solved for a specific variable, simplified, or transformed. This can take many formats, e.g. LaTeX, Ascii-Math, or math as you would write with a keyboard.
*/
type MathExpression string

/*

NumberOfAxles is https://schema.org/numberOfAxles

The number of axles.\n\nTypical unit code(s): C62.
*/
type NumberOfAxles float64

/*

PaymentUrl is https://schema.org/paymentUrl

The URL for sending a payment.
*/
type PaymentUrl string

/*

ValueMaxLength is https://schema.org/valueMaxLength

Specifies the allowed range for number of characters in a literal value.
*/
type ValueMaxLength float64

/*

Awards is https://schema.org/awards

Awards won by or for this item.
*/
type Awards string

/*

PreOp is https://schema.org/preOp

A description of the workup, testing, and other preparations required before implanting this device.
*/
type PreOp string

/*

UpvoteCount is https://schema.org/upvoteCount

The number of upvotes this question, answer or comment has received from the community.
*/
type UpvoteCount float64

/*

Device is https://schema.org/device

Device required to run the application. Used in cases where a specific make/model is required to run the application.
*/
type Device string

/*

HealthPlanCoinsuranceRate is https://schema.org/healthPlanCoinsuranceRate

The rate of coinsurance expressed as a number between 0.0 and 1.0.
*/
type HealthPlanCoinsuranceRate float64

/*

ConfirmationNumber is https://schema.org/confirmationNumber

A number that confirms the given order or payment has been received.
*/
type ConfirmationNumber string

/*

MembershipNumber is https://schema.org/membershipNumber

A unique identifier for the membership.
*/
type MembershipNumber string

/*

ReadonlyValue is https://schema.org/readonlyValue

Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a "hidden" input in an HTML form.
*/
type ReadonlyValue bool

/*

OriginalMediaLink is https://schema.org/originalMediaLink

Link to the page containing an original version of the content, or directly to an online copy of the original [[MediaObject]] content, e.g. video file.
*/
type OriginalMediaLink string

/*

NumTracks is https://schema.org/numTracks

The number of tracks in this album or playlist.
*/
type NumTracks float64

/*

NoBylinesPolicy is https://schema.org/noBylinesPolicy

For a [[NewsMediaOrganization]] or other news-related [[Organization]], a statement explaining when authors of articles are not named in bylines.
*/
type NoBylinesPolicy string

/*

MissionCoveragePrioritiesPolicy is https://schema.org/missionCoveragePrioritiesPolicy

For a [[NewsMediaOrganization]], a statement on coverage priorities, including any public agenda or stance on issues.
*/
type MissionCoveragePrioritiesPolicy string

/*

HealthPlanNetworkTier is https://schema.org/healthPlanNetworkTier

The tier(s) for this network.
*/
type HealthPlanNetworkTier string

/*

RequiredMinAge is https://schema.org/requiredMinAge

Audiences defined by a person's minimum age.
*/
type RequiredMinAge float64

/*

Isbn is https://schema.org/isbn

The ISBN of the book.
*/
type Isbn string

/*

CharacterName is https://schema.org/characterName

The name of a character played in some acting or performing role, i.e. in a PerformanceRole.
*/
type CharacterName string

/*

AnswerCount is https://schema.org/answerCount

The number of answers this question has received.
*/
type AnswerCount float64

/*

CheckinTime is https://schema.org/checkinTime

The earliest someone may check into a lodging establishment.
*/
type CheckinTime string

/*

ParentTaxon is https://schema.org/parentTaxon

Closest parent taxon of the taxon in question.
*/
type ParentTaxon string

/*

IupacName is https://schema.org/iupacName

Systematic method of naming chemical compounds as recommended by the International Union of Pure and Applied Chemistry (IUPAC).
*/
type IupacName string

/*

Surface is https://schema.org/surface

A material used as a surface in some artwork, e.g. Canvas, Paper, Wood, Board, etc.
*/
type Surface string

/*

EpisodeNumber is https://schema.org/episodeNumber

Position of the episode within an ordered group of episodes.
*/
type EpisodeNumber string

/*

ReviewAspect is https://schema.org/reviewAspect

This Review or Rating is relevant to this part or facet of the itemReviewed.
*/
type ReviewAspect string

/*

WorkHours is https://schema.org/workHours

The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
*/
type WorkHours string

/*

BeneficiaryBank is https://schema.org/beneficiaryBank

A bank or bank’s branch, financial institution or international financial institution operating the beneficiary’s bank account or releasing funds for the beneficiary.
*/
type BeneficiaryBank string

/*

Screenshot is https://schema.org/screenshot

A link to a screenshot image of the app.
*/
type Screenshot string

/*

ProductID is https://schema.org/productID

The product identifier, such as ISBN. For example: ``` meta itemprop="productID" content="isbn:123-456-789" ```.
*/
type ProductID string

/*

ModelDate is https://schema.org/modelDate

The release date of a vehicle model (often used to differentiate versions of the same make and model).
*/
type ModelDate string

/*

PartySize is https://schema.org/partySize

Number of people the reservation should accommodate.
*/
type PartySize float64

/*

DosageForm is https://schema.org/dosageForm

A dosage form in which this drug/supplement is available, e.g. 'tablet', 'suspension', 'injection'.
*/
type DosageForm string

/*

SpecialCommitments is https://schema.org/specialCommitments

Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
*/
type SpecialCommitments string

/*

RatingExplanation is https://schema.org/ratingExplanation

A short explanation (e.g. one to two sentences) providing background context and other information that led to the conclusion expressed in the rating. This is particularly applicable to ratings associated with "fact check" markup using [[ClaimReview]].
*/
type RatingExplanation string

/*

DiseaseSpreadStatistics is https://schema.org/diseaseSpreadStatistics

Statistical information about the spread of a disease, either as [[WebContent]], or
  described directly as a [[Dataset]], or the specific [[Observation]]s in the dataset. When a [[WebContent]] URL is
  provided, the page indicated might also contain more such markup.
*/
type DiseaseSpreadStatistics string

/*

FloorLevel is https://schema.org/floorLevel

The floor level for an [[Accommodation]] in a multi-storey building. Since counting
  systems [vary internationally](https://en.wikipedia.org/wiki/Storey#Consecutive_number_floor_designations), the local system should be used where possible.
*/
type FloorLevel string

/*

ValueRequired is https://schema.org/valueRequired

Whether the property must be filled in to complete the action.  Default is false.
*/
type ValueRequired bool

/*

CvdNumC19MechVentPats is https://schema.org/cvdNumC19MechVentPats

numc19mechventpats - HOSPITALIZED and VENTILATED: Patients hospitalized in an NHSN inpatient care location who have suspected or confirmed COVID-19 and are on a mechanical ventilator.
*/
type CvdNumC19MechVentPats float64

/*

AcquireLicensePage is https://schema.org/acquireLicensePage

Indicates a page documenting how licenses can be purchased or otherwise acquired, for the current item.
*/
type AcquireLicensePage string

/*

CorrectionsPolicy is https://schema.org/correctionsPolicy

For an [[Organization]] (e.g. [[NewsMediaOrganization]]), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
*/
type CorrectionsPolicy string

/*

Smiles is https://schema.org/smiles

A specification in form of a line notation for describing the structure of chemical species using short ASCII strings.  Double bond stereochemistry \ indicators may need to be escaped in the string in formats where the backslash is an escape character.
*/
type Smiles string

/*

AccessCode is https://schema.org/accessCode

Password, PIN, or access code needed for delivery (e.g. from a locker).
*/
type AccessCode string

/*

AssemblyVersion is https://schema.org/assemblyVersion

Associated product/technology version. E.g., .NET Framework 4.5.
*/
type AssemblyVersion string

/*

Position is https://schema.org/position

The position of an item in a series or sequence of items.
*/
type Position float64

/*

HealthPlanDrugTier is https://schema.org/healthPlanDrugTier

The tier(s) of drugs offered by this formulary or insurance plan.
*/
type HealthPlanDrugTier string

/*

PlayersOnline is https://schema.org/playersOnline

Number of players on the server.
*/
type PlayersOnline float64

/*

BodyLocation is https://schema.org/bodyLocation

Location in the body of the anatomical structure.
*/
type BodyLocation string

/*

Asin is https://schema.org/asin

An Amazon Standard Identification Number (ASIN) is a 10-character alphanumeric unique identifier assigned by Amazon.com and its partners for product identification within the Amazon organization (summary from [Wikipedia](https://en.wikipedia.org/wiki/Amazon_Standard_Identification_Number)'s article).

Note also that this is a definition for how to include ASINs in Schema.org data, and not a definition of ASINs in general - see documentation from Amazon for authoritative details.
ASINs are most commonly encoded as text strings, but the [asin] property supports URL/URI as potential values too.
*/
type Asin string

/*

LoanType is https://schema.org/loanType

The type of a loan or credit.
*/
type LoanType string

/*

IncentiveCompensation is https://schema.org/incentiveCompensation

Description of bonus and commission compensation aspects of the job.
*/
type IncentiveCompensation string

/*

ProductionDate is https://schema.org/productionDate

The date of production of the item, e.g. vehicle.
*/
type ProductionDate string

/*

Price is https://schema.org/price

The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.\n\nUsage guidelines:\n\n* Use the [[priceCurrency]] property (with standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217), e.g. "USD"; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies, e.g. "BTC"; well known names for [Local Exchange Trading Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types, e.g. "Ithaca HOUR") instead of including [ambiguous symbols](http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as '$' in the value.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.\n* Note that both [RDFa](http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a "content=" attribute for publishing simple machine-readable values alongside more human-friendly formatting.\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.
      
*/
type Price float64

/*

ValidThrough is https://schema.org/validThrough

The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
*/
type ValidThrough string

/*

ThumbnailUrl is https://schema.org/thumbnailUrl

A thumbnail image relevant to the Thing.
*/
type ThumbnailUrl string

/*

Procedure is https://schema.org/procedure

A description of the procedure involved in setting up, using, and/or installing the device.
*/
type Procedure string

/*

Pagination is https://schema.org/pagination

Any description of pages that is not separated into pageStart and pageEnd; for example, "1-6, 9, 55" or "10-12, 46-49".
*/
type Pagination string

/*

MerchantReturnLink is https://schema.org/merchantReturnLink

Specifies a Web page or service by URL, for product returns.
*/
type MerchantReturnLink string

/*

CostCurrency is https://schema.org/costCurrency

The currency (in 3-letter) of the drug cost. See: http://en.wikipedia.org/wiki/ISO_4217. 
*/
type CostCurrency string

/*

AvailabilityEnds is https://schema.org/availabilityEnds

The end of the availability of the product or service included in the offer.
*/
type AvailabilityEnds string

/*

AcrissCode is https://schema.org/acrissCode

The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
*/
type AcrissCode string

/*

Naics is https://schema.org/naics

The North American Industry Classification System (NAICS) code for a particular organization or business person.
*/
type Naics string

/*

PriceRange is https://schema.org/priceRange

The price range of the business, for example ```$$$```.
*/
type PriceRange string

/*

InterestRate is https://schema.org/interestRate

The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
*/
type InterestRate float64

/*

DiversityPolicy is https://schema.org/diversityPolicy

Statement on diversity policy by an [[Organization]] e.g. a [[NewsMediaOrganization]]. For a [[NewsMediaOrganization]], a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
*/
type DiversityPolicy string

/*

Gender is https://schema.org/gender

Gender of something, typically a [[Person]], but possibly also fictional characters, animals, etc. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender. The [[gender]] property can also be used in an extended sense to cover e.g. the gender of sports teams. As with the gender of individuals, we do not try to enumerate all possibilities. A mixed-gender [[SportsTeam]] can be indicated with a text value of "Mixed".
*/
type Gender string

/*

HasRepresentation is https://schema.org/hasRepresentation

A common representation such as a protein sequence or chemical structure for this entity. For images use schema.org/image.
*/
type HasRepresentation string

/*

ClipNumber is https://schema.org/clipNumber

Position of the clip within an ordered group of clips.
*/
type ClipNumber string

/*

Qualifications is https://schema.org/qualifications

Specific qualifications required for this role or Occupation.
*/
type Qualifications string

/*

MultipleValues is https://schema.org/multipleValues

Whether multiple values are allowed for the property.  Default is false.
*/
type MultipleValues bool

/*

PrescriptionStatus is https://schema.org/prescriptionStatus

Indicates the status of drug prescription, e.g. local catalogs classifications or whether the drug is available by prescription or over-the-counter, etc.
*/
type PrescriptionStatus string

/*

TicketNumber is https://schema.org/ticketNumber

The unique identifier for the ticket.
*/
type TicketNumber string

/*

GeoRadius is https://schema.org/geoRadius

Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
*/
type GeoRadius float64

/*

CourseWorkload is https://schema.org/courseWorkload

The amount of work expected of students taking the course, often provided as a figure per week or per month, and may be broken down by type. For example, "2 hours of lectures, 1 hour of lab work and 3 hours of independent study per week".
*/
type CourseWorkload string

/*

ItemLocation is https://schema.org/itemLocation

Current location of the item.
*/
type ItemLocation string

/*

CompetencyRequired is https://schema.org/competencyRequired

Knowledge, skill, ability or personal attribute that must be demonstrated by a person or other entity in order to do something such as earn an Educational Occupational Credential or understand a LearningResource.
*/
type CompetencyRequired string

/*

Documentation is https://schema.org/documentation

Further documentation describing the Web API in more detail.
*/
type Documentation string

/*

StepValue is https://schema.org/stepValue

The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
*/
type StepValue float64

/*

ExpectedArrivalFrom is https://schema.org/expectedArrivalFrom

The earliest date the package may arrive.
*/
type ExpectedArrivalFrom string

/*

ArtMedium is https://schema.org/artMedium

The material used. (E.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
*/
type ArtMedium string

/*

VehicleTransmission is https://schema.org/vehicleTransmission

The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) ("gearbox" for cars).
*/
type VehicleTransmission string

/*

AdministrationRoute is https://schema.org/administrationRoute

A route by which this drug may be administered, e.g. 'oral'.
*/
type AdministrationRoute string

/*

ApplicationDeadline is https://schema.org/applicationDeadline

The date at which the program stops collecting applications for the next enrollment cycle.
*/
type ApplicationDeadline string

/*

CountryOfAssembly is https://schema.org/countryOfAssembly

The place where the product was assembled.
*/
type CountryOfAssembly string

/*

HealthPlanPharmacyCategory is https://schema.org/healthPlanPharmacyCategory

The category or type of pharmacy associated with this cost sharing.
*/
type HealthPlanPharmacyCategory string

/*

IswcCode is https://schema.org/iswcCode

The International Standard Musical Work Code for the composition.
*/
type IswcCode string

/*

CheckoutPageURLTemplate is https://schema.org/checkoutPageURLTemplate

A URL template (RFC 6570) for a checkout page for an offer. This approach allows merchants to specify a URL for online checkout of the offered product, by interpolating parameters such as the logged in user ID, product ID, quantity, discount code etc. Parameter naming and standardization are not specified here.
*/
type CheckoutPageURLTemplate string

/*

License is https://schema.org/license

A license document that applies to this content, typically indicated by URL.
*/
type License string

/*

ChildTaxon is https://schema.org/childTaxon

Closest child taxa of the taxon in question.
*/
type ChildTaxon string

/*

IataCode is https://schema.org/iataCode

IATA identifier for an airline or airport.
*/
type IataCode string

/*

EmbedUrl is https://schema.org/embedUrl

A URL pointing to a player for a specific video. In general, this is the information in the ```src``` element of an ```embed``` tag and should not be the same as the content of the ```loc``` tag.
*/
type EmbedUrl string

/*

VehicleModelDate is https://schema.org/vehicleModelDate

The release date of a vehicle model (often used to differentiate versions of the same make and model).
*/
type VehicleModelDate string

/*

SdDatePublished is https://schema.org/sdDatePublished

Indicates the date on which the current structured data was generated / published. Typically used alongside [[sdPublisher]].
*/
type SdDatePublished string

/*

Permissions is https://schema.org/permissions

Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
*/
type Permissions string

/*

Requirements is https://schema.org/requirements

Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (examples: DirectX, Java or .NET runtime).
*/
type Requirements string

/*

NumberOfForwardGears is https://schema.org/numberOfForwardGears

The total number of forward gears available for the transmission system of the vehicle.\n\nTypical unit code(s): C62.
*/
type NumberOfForwardGears float64

/*

Steps is https://schema.org/steps

A single step item (as HowToStep, text, document, video, etc.) or a HowToSection (originally misnamed 'steps'; 'step' is preferred).
*/
type Steps string

/*

TitleEIDR is https://schema.org/titleEIDR

An [EIDR](https://eidr.org/) (Entertainment Identifier Registry) [[identifier]] representing at the most general/abstract level, a work of film or television.

For example, the motion picture known as "Ghostbusters" has a titleEIDR of  "10.5240/7EC7-228A-510A-053E-CBB8-J". This title (or work) may have several variants, which EIDR calls "edits". See [[editEIDR]].

Since schema.org types like [[Movie]], [[TVEpisode]], [[TVSeason]], and [[TVSeries]] can be used for both works and their multiple expressions, it is possible to use [[titleEIDR]] alone (for a general description), or alongside [[editEIDR]] for a more edit-specific description.

*/
type TitleEIDR string

/*

BestRating is https://schema.org/bestRating

The highest value allowed in this rating system.
*/
type BestRating float64

/*

VehicleInteriorType is https://schema.org/vehicleInteriorType

The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
*/
type VehicleInteriorType string

/*

PrintEdition is https://schema.org/printEdition

The edition of the print product in which the NewsArticle appears.
*/
type PrintEdition string

/*

Assesses is https://schema.org/assesses

The item being described is intended to assess the competency or learning outcome defined by the referenced term.
*/
type Assesses string

/*

PostalCodeBegin is https://schema.org/postalCodeBegin

First postal code in a range (included).
*/
type PostalCodeBegin string

/*

Teaches is https://schema.org/teaches

The item being described is intended to help a person learn the competency or learning outcome defined by the referenced term.
*/
type Teaches string

/*

NumberOfRooms is https://schema.org/numberOfRooms

The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.
Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
*/
type NumberOfRooms float64

/*

Bed is https://schema.org/bed

The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
      If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
*/
type Bed string

/*

Correction is https://schema.org/correction

Indicates a correction to a [[CreativeWork]], either via a [[CorrectionComment]], textually or in another document.
*/
type Correction string

/*

YearBuilt is https://schema.org/yearBuilt

The year an [[Accommodation]] was constructed. This corresponds to the [YearBuilt field in RESO](https://ddwiki.reso.org/display/DDW17/YearBuilt+Field). 
*/
type YearBuilt float64

/*

InstallUrl is https://schema.org/installUrl

URL at which the app may be installed, if different from the URL of the item.
*/
type InstallUrl string

/*

OrderNumber is https://schema.org/orderNumber

The identifier of the transaction.
*/
type OrderNumber string

/*

CvdNumBedsOcc is https://schema.org/cvdNumBedsOcc

numbedsocc - HOSPITAL INPATIENT BED OCCUPANCY: Total number of staffed inpatient beds that are occupied.
*/
type CvdNumBedsOcc float64

/*

DoorTime is https://schema.org/doorTime

The time admission will commence.
*/
type DoorTime string

/*

MemoryRequirements is https://schema.org/memoryRequirements

Minimum memory requirements.
*/
type MemoryRequirements string

/*

VariablesMeasured is https://schema.org/variablesMeasured

Originally named [[variablesMeasured]], the [[variableMeasured]] property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue.
*/
type VariablesMeasured string

/*

CountriesSupported is https://schema.org/countriesSupported

Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
*/
type CountriesSupported string

/*

AvailabilityStarts is https://schema.org/availabilityStarts

The beginning of the availability of the product or service included in the offer.
*/
type AvailabilityStarts string

/*

DateCreated is https://schema.org/dateCreated

The date on which the CreativeWork was created or the item was added to a DataFeed.
*/
type DateCreated string

/*

FuelType is https://schema.org/fuelType

The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
*/
type FuelType string

/*

CutoffTime is https://schema.org/cutoffTime

Order cutoff time allows merchants to describe the time after which they will no longer process orders received on that day. For orders processed after cutoff time, one day gets added to the delivery time estimate. This property is expected to be most typically used via the [[ShippingRateSettings]] publication pattern. The time is indicated using the ISO-8601 Time format, e.g. "23:30:00-05:00" would represent 6:30 pm Eastern Standard Time (EST) which is 5 hours behind Coordinated Universal Time (UTC).
*/
type CutoffTime string

/*

NewsUpdatesAndGuidelines is https://schema.org/newsUpdatesAndGuidelines

Indicates a page with news updates and guidelines. This could often be (but is not required to be) the main page containing [[SpecialAnnouncement]] markup on a site.
*/
type NewsUpdatesAndGuidelines string

/*

ProgramType is https://schema.org/programType

The type of educational or occupational program. For example, classroom, internship, alternance, etc.
*/
type ProgramType string

/*

CssSelector is https://schema.org/cssSelector

A CSS selector, e.g. of a [[SpeakableSpecification]] or [[WebPageElement]]. In the latter case, multiple matches within a page can constitute a single conceptual "Web page element".
*/
type CssSelector string

/*

PreviousStartDate is https://schema.org/previousStartDate

Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
*/
type PreviousStartDate string

/*

Expires is https://schema.org/expires

Date the content expires and is no longer useful or available. For example a [[VideoObject]] or [[NewsArticle]] whose availability or relevance is time-limited, a [[ClaimReview]] fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date, or a [[Certification]] the validity has expired.
*/
type Expires string

/*

HasTierRequirement is https://schema.org/hasTierRequirement

A requirement for a user to join a membership tier, for example: a CreditCard if the tier requires sign up for a credit card, A UnitPriceSpecification if the user is required to pay a (periodic) fee, or a MonetaryAmount if the user needs to spend a minimum amount to join the tier. If a tier is free to join then this property does not need to be specified.
*/
type HasTierRequirement string

/*

LowPrice is https://schema.org/lowPrice

The lowest price of all offers available.\n\nUsage guidelines:\n\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
*/
type LowPrice float64

/*

AccessibilityAPI is https://schema.org/accessibilityAPI

Indicates that the resource is compatible with the referenced accessibility API. Values should be drawn from the [approved vocabulary](https://www.w3.org/2021/a11y-discov-vocab/latest/#accessibilityAPI-vocabulary).
*/
type AccessibilityAPI string

/*

VideoQuality is https://schema.org/videoQuality

The quality of the video.
*/
type VideoQuality string

/*

VehicleConfiguration is https://schema.org/vehicleConfiguration

A short text indicating the configuration of the vehicle, e.g. '5dr hatchback ST 2.5 MT 225 hp' or 'limited edition'.
*/
type VehicleConfiguration string

/*

OccupationalCredentialAwarded is https://schema.org/occupationalCredentialAwarded

A description of the qualification, award, certificate, diploma or other occupational credential awarded as a consequence of successful completion of this course or program.
*/
type OccupationalCredentialAwarded string

/*

BoardingGroup is https://schema.org/boardingGroup

The airline-specific indicator of boarding order / preference.
*/
type BoardingGroup string

/*

HealthPlanNetworkId is https://schema.org/healthPlanNetworkId

Name or unique ID of network. (Networks are often reused across different insurance plans.)
*/
type HealthPlanNetworkId string

/*

Status is https://schema.org/status

The status of the study (enumerated).
*/
type Status string

/*

Colleague is https://schema.org/colleague

A colleague of the person.
*/
type Colleague string

/*

TargetPopulation is https://schema.org/targetPopulation

Characteristics of the population for which this is intended, or which typically uses it, e.g. 'adults'.
*/
type TargetPopulation string

/*

Overdosage is https://schema.org/overdosage

Any information related to overdose on a drug, including signs or symptoms, treatments, contact information for emergency response.
*/
type Overdosage string

/*

PostalCodeEnd is https://schema.org/postalCodeEnd

Last postal code in the range (included). Needs to be after [[postalCodeBegin]].
*/
type PostalCodeEnd string

/*

AdditionalVariable is https://schema.org/additionalVariable

Any additional component of the exercise prescription that may need to be articulated to the patient. This may include the order of exercises, the number of repetitions of movement, quantitative distance, progressions over time, etc.
*/
type AdditionalVariable string

/*

TrackingNumber is https://schema.org/trackingNumber

Shipper tracking number.
*/
type TrackingNumber string

/*

DateModified is https://schema.org/dateModified

The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
*/
type DateModified string

/*

CvdCollectionDate is https://schema.org/cvdCollectionDate

collectiondate - Date for which patient counts are reported.
*/
type CvdCollectionDate string

/*

TourBookingPage is https://schema.org/tourBookingPage

A page providing information on how to book a tour of some [[Place]], such as an [[Accommodation]] or [[ApartmentComplex]] in a real estate setting, as well as other kinds of tours as appropriate.
*/
type TourBookingPage string

/*

HonorificSuffix is https://schema.org/honorificSuffix

An honorific suffix following a Person's name such as M.D./PhD/MSCSW.
*/
type HonorificSuffix string

/*

EditEIDR is https://schema.org/editEIDR

An [EIDR](https://eidr.org/) (Entertainment Identifier Registry) [[identifier]] representing a specific edit / edition for a work of film or television.

For example, the motion picture known as "Ghostbusters" whose [[titleEIDR]] is "10.5240/7EC7-228A-510A-053E-CBB8-J" has several edits, e.g. "10.5240/1F2A-E1C5-680A-14C6-E76B-I" and "10.5240/8A35-3BEE-6497-5D12-9E4F-3".

Since schema.org types like [[Movie]] and [[TVEpisode]] can be used for both works and their multiple expressions, it is possible to use [[titleEIDR]] alone (for a general description), or alongside [[editEIDR]] for a more edit-specific description.

*/
type EditEIDR string

/*

SameAs is https://schema.org/sameAs

URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
*/
type SameAs string

/*

TimeOfDay is https://schema.org/timeOfDay

The time of day the program normally runs. For example, "evenings".
*/
type TimeOfDay string

/*

NumberOfPreviousOwners is https://schema.org/numberOfPreviousOwners

The number of owners of the vehicle, including the current one.\n\nTypical unit code(s): C62.
*/
type NumberOfPreviousOwners float64

/*

RequiresSubscription is https://schema.org/requiresSubscription

Indicates if use of the media require a subscription  (either paid or free). Allowed values are ```true``` or ```false``` (note that an earlier version had 'yes', 'no').
*/
type RequiresSubscription bool

/*

AddressRegion is https://schema.org/addressRegion

The region in which the locality is, and which is in the country. For example, California or another appropriate first-level [Administrative division](https://en.wikipedia.org/wiki/List_of_administrative_divisions_by_country).
*/
type AddressRegion string

/*

RenegotiableLoan is https://schema.org/renegotiableLoan

Whether the terms for payment of interest can be renegotiated during the life of the loan.
*/
type RenegotiableLoan bool

/*

Abstract is https://schema.org/abstract

An abstract is a short description that summarizes a [[CreativeWork]].
*/
type Abstract string

/*

OperatingSystem is https://schema.org/operatingSystem

Operating systems supported (Windows 7, OS X 10.6, Android 1.6).
*/
type OperatingSystem string

/*

TermCode is https://schema.org/termCode

A code that identifies this [[DefinedTerm]] within a [[DefinedTermSet]].
*/
type TermCode string

/*

DriveWheelConfiguration is https://schema.org/driveWheelConfiguration

The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle's engine via the drivetrain.
*/
type DriveWheelConfiguration string

/*

UsageInfo is https://schema.org/usageInfo

The schema.org [[usageInfo]] property indicates further information about a [[CreativeWork]]. This property is applicable both to works that are freely available and to those that require payment or other transactions. It can reference additional information, e.g. community expectations on preferred linking and citation conventions, as well as purchasing details. For something that can be commercially licensed, usageInfo can provide detailed, resource-specific information about licensing options.

This property can be used alongside the license property which indicates license(s) applicable to some piece of content. The usageInfo property can provide information about other licensing options, e.g. acquiring commercial usage rights for an image that is also available under non-commercial creative commons licenses.
*/
type UsageInfo string

/*

IsBasedOn is https://schema.org/isBasedOn

A resource from which this work is derived or from which it is a modification or adaptation.
*/
type IsBasedOn string

/*

PhysiologicalBenefits is https://schema.org/physiologicalBenefits

Specific physiologic benefits associated to the plan.
*/
type PhysiologicalBenefits string

/*

ReplyToUrl is https://schema.org/replyToUrl

The URL at which a reply may be posted to the specified UserComment.
*/
type ReplyToUrl string

/*

HasMap is https://schema.org/hasMap

A URL to a map of the place.
*/
type HasMap string

/*

CreditText is https://schema.org/creditText

Text that can be used to credit person(s) and/or organization(s) associated with a published Creative Work.
*/
type CreditText string

/*

CodingSystem is https://schema.org/codingSystem

The coding system, e.g. 'ICD-10'.
*/
type CodingSystem string

/*

RecipeCuisine is https://schema.org/recipeCuisine

The cuisine of the recipe (for example, French or Ethiopian).
*/
type RecipeCuisine string

/*

DateDeleted is https://schema.org/dateDeleted

The datetime the item was removed from the DataFeed.
*/
type DateDeleted string

/*

PageEnd is https://schema.org/pageEnd

The page on which the work ends; for example "138" or "xvi".
*/
type PageEnd float64

/*

AlignmentType is https://schema.org/alignmentType

A category of alignment between the learning resource and the framework node. Recommended values include: 'requires', 'textComplexity', 'readingLevel', and 'educationalSubject'.
*/
type AlignmentType string

/*

WebCheckinTime is https://schema.org/webCheckinTime

The time when a passenger can check into the flight online.
*/
type WebCheckinTime string

/*

DropoffTime is https://schema.org/dropoffTime

When a rental car can be dropped off.
*/
type DropoffTime string

/*

SubtitleLanguage is https://schema.org/subtitleLanguage

Languages in which subtitles/captions are available, in [IETF BCP 47 standard format](http://tools.ietf.org/html/bcp47).
*/
type SubtitleLanguage string

/*

Percentile25 is https://schema.org/percentile25

The 25th percentile value.
*/
type Percentile25 float64

/*

PrintPage is https://schema.org/printPage

If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
*/
type PrintPage string

/*

LegislationDateVersion is https://schema.org/legislationDateVersion

The point-in-time at which the provided description of the legislation is valid (e.g.: when looking at the law on the 2016-04-07 (= dateVersion), I get the consolidation of 2015-04-12 of the "National Insurance Contributions Act 2015")
*/
type LegislationDateVersion string

/*

RepeatFrequency is https://schema.org/repeatFrequency

Defines the frequency at which [[Event]]s will occur according to a schedule [[Schedule]]. The intervals between
      events should be defined as a [[Duration]] of time.
*/
type RepeatFrequency string

/*

RequiredMaxAge is https://schema.org/requiredMaxAge

Audiences defined by a person's maximum age.
*/
type RequiredMaxAge float64

/*

SuggestedMinAge is https://schema.org/suggestedMinAge

Minimum recommended age in years for the audience or user.
*/
type SuggestedMinAge float64

/*

NumberOfBathroomsTotal is https://schema.org/numberOfBathroomsTotal

The total integer number of bathrooms in some [[Accommodation]], following real estate conventions as [documented in RESO](https://ddwiki.reso.org/display/DDW17/BathroomsTotalInteger+Field): "The simple sum of the number of bathrooms. For example for a property with two Full Bathrooms and one Half Bathroom, the Bathrooms Total Integer will be 3.". See also [[numberOfRooms]].
*/
type NumberOfBathroomsTotal float64

/*

TrainNumber is https://schema.org/trainNumber

The unique identifier for the train.
*/
type TrainNumber string

/*

DownvoteCount is https://schema.org/downvoteCount

The number of downvotes this question, answer or comment has received from the community.
*/
type DownvoteCount float64

/*

SmokingAllowed is https://schema.org/smokingAllowed

Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
*/
type SmokingAllowed bool

/*

ReservationId is https://schema.org/reservationId

A unique identifier for the reservation.
*/
type ReservationId string

/*

ColorSwatch is https://schema.org/colorSwatch

A color swatch image, visualizing the color of a [[Product]]. Should match the textual description specified in the [[color]] property. This can be a URL or a fully described ImageObject.
*/
type ColorSwatch string

/*

IsBasedOnUrl is https://schema.org/isBasedOnUrl

A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
*/
type IsBasedOnUrl string

/*

CookingMethod is https://schema.org/cookingMethod

The method of cooking, such as Frying, Steaming, ...
*/
type CookingMethod string

/*

Telephone is https://schema.org/telephone

The telephone number.
*/
type Telephone string

/*

SeatingCapacity is https://schema.org/seatingCapacity

The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.\n\nTypical unit code(s): C62 for persons.
*/
type SeatingCapacity float64

/*

RatingValue is https://schema.org/ratingValue

The rating for the content.\n\nUsage guidelines:\n\n* Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similar Unicode symbols.\n* Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
*/
type RatingValue float64

/*

OrderQuantity is https://schema.org/orderQuantity

The number of the item ordered. If the property is not set, assume the quantity is one.
*/
type OrderQuantity float64

/*

NumberedPosition is https://schema.org/numberedPosition

A number associated with a role in an organization, for example, the number on an athlete's jersey.
*/
type NumberedPosition float64

/*

InCodeSet is https://schema.org/inCodeSet

A [[CategoryCodeSet]] that contains this category code.
*/
type InCodeSet string

/*

DomiciledMortgage is https://schema.org/domiciledMortgage

Whether borrower is a resident of the jurisdiction where the property is located.
*/
type DomiciledMortgage bool

/*

BaseSalary is https://schema.org/baseSalary

The base salary of the job or of an employee in an EmployeeRole.
*/
type BaseSalary float64

/*

ActionableFeedbackPolicy is https://schema.org/actionableFeedbackPolicy

For a [[NewsMediaOrganization]] or other news-related [[Organization]], a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
*/
type ActionableFeedbackPolicy string

/*

BodyType is https://schema.org/bodyType

Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
*/
type BodyType string

/*

StartDate is https://schema.org/startDate

The start date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
*/
type StartDate string

/*

NumConstraints is https://schema.org/numConstraints

Indicates the number of constraints property values defined for a particular [[ConstraintNode]] such as [[StatisticalVariable]]. This helps applications understand if they have access to a sufficiently complete description of a [[StatisticalVariable]] or other construct that is defined using properties on template-style nodes.
*/
type NumConstraints float64

/*

HasBioPolymerSequence is https://schema.org/hasBioPolymerSequence

A symbolic representation of a BioChemEntity. For example, a nucleotide sequence of a Gene or an amino acid sequence of a Protein.
*/
type HasBioPolymerSequence string

/*

TotalHistoricalEnrollment is https://schema.org/totalHistoricalEnrollment

The total number of students that have enrolled in the history of the course.
*/
type TotalHistoricalEnrollment float64

/*

RequiredCollateral is https://schema.org/requiredCollateral

Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
*/
type RequiredCollateral string

/*

TrainName is https://schema.org/trainName

The name of the train (e.g. The Orient Express).
*/
type TrainName string

/*

NumberOfDoors is https://schema.org/numberOfDoors

The number of doors.\n\nTypical unit code(s): C62.
*/
type NumberOfDoors float64

/*

ArticleBody is https://schema.org/articleBody

The actual body of the article.
*/
type ArticleBody string

/*

Mpn is https://schema.org/mpn

The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
*/
type Mpn string

/*

JobImmediateStart is https://schema.org/jobImmediateStart

An indicator as to whether a position is available for an immediate start.
*/
type JobImmediateStart bool

/*

ArrivalTime is https://schema.org/arrivalTime

The expected arrival time.
*/
type ArrivalTime string

/*

Address is https://schema.org/address

Physical address of the item.
*/
type Address string

/*

AvailableOnDevice is https://schema.org/availableOnDevice

Device required to run the application. Used in cases where a specific make/model is required to run the application.
*/
type AvailableOnDevice string

/*

IneligibleRegion is https://schema.org/ineligibleRegion

The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.\n\nSee also [[eligibleRegion]].
      
*/
type IneligibleRegion string

/*

ChemicalComposition is https://schema.org/chemicalComposition

The chemical composition describes the identity and relative ratio of the chemical elements that make up the substance.
*/
type ChemicalComposition string

/*

FoodWarning is https://schema.org/foodWarning

Any precaution, guidance, contraindication, etc. related to consumption of specific foods while taking this drug.
*/
type FoodWarning string

/*

AdditionalName is https://schema.org/additionalName

An additional name for a Person, can be used for a middle name.
*/
type AdditionalName string

/*

Backstory is https://schema.org/backstory

For an [[Article]], typically a [[NewsArticle]], the backstory property provides a textual summary giving a brief explanation of why and how an article was created. In a journalistic setting this could include information about reporting process, methods, interviews, data sources, etc.
*/
type Backstory string

/*

CredentialCategory is https://schema.org/credentialCategory

The category or type of credential being described, for example "degree”, “certificate”, “badge”, or more specific term.
*/
type CredentialCategory string

/*

FileFormat is https://schema.org/fileFormat

Media type, typically MIME format (see [IANA site](http://www.iana.org/assignments/media-types/media-types.xhtml)) of the content, e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, 'encoding' can be used to indicate each MediaObject alongside particular fileFormat information. Unregistered or niche file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia entry.
*/
type FileFormat string

/*

ExpectedArrivalUntil is https://schema.org/expectedArrivalUntil

The latest date the package may arrive.
*/
type ExpectedArrivalUntil string

/*

InSupportOf is https://schema.org/inSupportOf

Qualification, candidature, degree, application that Thesis supports.
*/
type InSupportOf string

/*

SpeechToTextMarkup is https://schema.org/speechToTextMarkup

Form of markup used. eg. [SSML](https://www.w3.org/TR/speech-synthesis11) or [IPA](https://www.wikidata.org/wiki/Property:P898).
*/
type SpeechToTextMarkup string

/*

TransitTimeLabel is https://schema.org/transitTimeLabel

Label to match an [[OfferShippingDetails]] with a [[DeliveryTimeSettings]] (within the context of a [[shippingSettingsLink]] cross-reference).
*/
type TransitTimeLabel string

/*

EducationalCredentialAwarded is https://schema.org/educationalCredentialAwarded

A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course or program.
*/
type EducationalCredentialAwarded string

/*

CodeSampleType is https://schema.org/codeSampleType

What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template.
*/
type CodeSampleType string

/*

Tool is https://schema.org/tool

A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
*/
type Tool string

/*

ExpertConsiderations is https://schema.org/expertConsiderations

Medical expert advice related to the plan.
*/
type ExpertConsiderations string

/*

AvailableFrom is https://schema.org/availableFrom

When the item is available for pickup from the store, locker, etc.
*/
type AvailableFrom string

/*

ApplicationStartDate is https://schema.org/applicationStartDate

The date at which the program begins collecting applications for the next enrollment cycle.
*/
type ApplicationStartDate string

/*

ArrivalGate is https://schema.org/arrivalGate

Identifier of the flight's arrival gate.
*/
type ArrivalGate string

/*

FunctionalClass is https://schema.org/functionalClass

The degree of mobility the joint allows.
*/
type FunctionalClass string

/*

DiscussionUrl is https://schema.org/discussionUrl

A link to the page containing the comments of the CreativeWork.
*/
type DiscussionUrl string

/*

Headline is https://schema.org/headline

Headline of the article.
*/
type Headline string

/*

SeasonNumber is https://schema.org/seasonNumber

Position of the season within an ordered group of seasons.
*/
type SeasonNumber float64

/*

CarrierRequirements is https://schema.org/carrierRequirements

Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
*/
type CarrierRequirements string

/*

SchemaVersion is https://schema.org/schemaVersion

Indicates (by URL or string) a particular version of a schema used in some CreativeWork. This property was created primarily to
    indicate the use of a specific schema.org release, e.g. ```10.0``` as a simple string, or more explicitly via URL, ```http://schema.org/docs/releases.html#v10.0```. There may be situations in which other schemas might usefully be referenced this way, e.g. ```http://dublincore.org/specifications/dublin-core/dces/1999-07-02/``` but this has not been carefully explored in the community.
*/
type SchemaVersion string

/*

PaymentMethodId is https://schema.org/paymentMethodId

An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
*/
type PaymentMethodId string

/*

GameLocation is https://schema.org/gameLocation

Real or fictional location of the game (or part of game).
*/
type GameLocation string

/*

CountryOfLastProcessing is https://schema.org/countryOfLastProcessing

The place where the item (typically [[Product]]) was last processed and tested before importation.
*/
type CountryOfLastProcessing string

/*

Closes is https://schema.org/closes

The closing hour of the place or service on the given day(s) of the week.
*/
type Closes string
