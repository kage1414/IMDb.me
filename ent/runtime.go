// Code generated by ent, DO NOT EDIT.

package ent

import (
	"imdb-db/ent/akas"
	"imdb-db/ent/episode"
	"imdb-db/ent/name"
	"imdb-db/ent/principals"
	"imdb-db/ent/ratings"
	"imdb-db/ent/schema"
	"imdb-db/ent/title"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	akasFields := schema.Akas{}.Fields()
	_ = akasFields
	// akasDescOrdering is the schema descriptor for ordering field.
	akasDescOrdering := akasFields[1].Descriptor()
	// akas.OrderingValidator is a validator for the "ordering" field. It is called by the builders before save.
	akas.OrderingValidator = akasDescOrdering.Validators[0].(func(int) error)
	episodeFields := schema.Episode{}.Fields()
	_ = episodeFields
	// episodeDescSeasonNumber is the schema descriptor for seasonNumber field.
	episodeDescSeasonNumber := episodeFields[2].Descriptor()
	// episode.SeasonNumberValidator is a validator for the "seasonNumber" field. It is called by the builders before save.
	episode.SeasonNumberValidator = episodeDescSeasonNumber.Validators[0].(func(int) error)
	// episodeDescEpisodeNumber is the schema descriptor for episodeNumber field.
	episodeDescEpisodeNumber := episodeFields[3].Descriptor()
	// episode.EpisodeNumberValidator is a validator for the "episodeNumber" field. It is called by the builders before save.
	episode.EpisodeNumberValidator = episodeDescEpisodeNumber.Validators[0].(func(int) error)
	nameFields := schema.Name{}.Fields()
	_ = nameFields
	// nameDescBirthYear is the schema descriptor for birthYear field.
	nameDescBirthYear := nameFields[2].Descriptor()
	// name.BirthYearValidator is a validator for the "birthYear" field. It is called by the builders before save.
	name.BirthYearValidator = nameDescBirthYear.Validators[0].(func(int) error)
	// nameDescDeathYear is the schema descriptor for deathYear field.
	nameDescDeathYear := nameFields[3].Descriptor()
	// name.DeathYearValidator is a validator for the "deathYear" field. It is called by the builders before save.
	name.DeathYearValidator = nameDescDeathYear.Validators[0].(func(int) error)
	principalsFields := schema.Principals{}.Fields()
	_ = principalsFields
	// principalsDescOrdering is the schema descriptor for ordering field.
	principalsDescOrdering := principalsFields[1].Descriptor()
	// principals.OrderingValidator is a validator for the "ordering" field. It is called by the builders before save.
	principals.OrderingValidator = principalsDescOrdering.Validators[0].(func(int) error)
	ratingsFields := schema.Ratings{}.Fields()
	_ = ratingsFields
	// ratingsDescAverageRating is the schema descriptor for averageRating field.
	ratingsDescAverageRating := ratingsFields[1].Descriptor()
	// ratings.AverageRatingValidator is a validator for the "averageRating" field. It is called by the builders before save.
	ratings.AverageRatingValidator = ratingsDescAverageRating.Validators[0].(func(float64) error)
	titleFields := schema.Title{}.Fields()
	_ = titleFields
	// titleDescIsAdult is the schema descriptor for isAdult field.
	titleDescIsAdult := titleFields[4].Descriptor()
	// title.DefaultIsAdult holds the default value on creation for the isAdult field.
	title.DefaultIsAdult = titleDescIsAdult.Default.(bool)
	// titleDescStartYear is the schema descriptor for startYear field.
	titleDescStartYear := titleFields[5].Descriptor()
	// title.StartYearValidator is a validator for the "startYear" field. It is called by the builders before save.
	title.StartYearValidator = titleDescStartYear.Validators[0].(func(int) error)
	// titleDescEndYear is the schema descriptor for endYear field.
	titleDescEndYear := titleFields[6].Descriptor()
	// title.EndYearValidator is a validator for the "endYear" field. It is called by the builders before save.
	title.EndYearValidator = titleDescEndYear.Validators[0].(func(int) error)
	// titleDescRuntimeMinutes is the schema descriptor for runtimeMinutes field.
	titleDescRuntimeMinutes := titleFields[7].Descriptor()
	// title.RuntimeMinutesValidator is a validator for the "runtimeMinutes" field. It is called by the builders before save.
	title.RuntimeMinutesValidator = titleDescRuntimeMinutes.Validators[0].(func(int) error)
}
