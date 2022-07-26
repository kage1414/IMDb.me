const { models } = require('./models.js');
const Table = require('./table.js');
const {
  NAME,
  TITLE,
  AKAS,
  CREW,
  EPISODE,
  PRINCIPALS,
  RATINGS,
} = models;

const tables = {
  basics_name: new Table('basics_name', {
    url: 'https://datasets.imdbws.com/name.basics.tsv.gz',
    model: NAME,
  }),
  basics_title: new Table('basics_title', {
    url: 'https://datasets.imdbws.com/title.basics.tsv.gz',
    model: TITLE,
  }),
  akas: new Table('akas', {
    url: 'https://datasets.imdbws.com/title.akas.tsv.gz',
    model: AKAS,
  }),
  crew: new Table('crew', {
    url: 'https://datasets.imdbws.com/title.crew.tsv.gz',
    model: CREW,
  }),
  episode: new Table('episode', {
    url: 'https://datasets.imdbws.com/title.episode.tsv.gz',
    model: EPISODE,
  }),
  principals: new Table('principals', {
    url: 'https://datasets.imdbws.com/title.principals.tsv.gz',
    model: PRINCIPALS,
  }),
  ratings: new Table('ratings', {
    url: 'https://datasets.imdbws.com/title.ratings.tsv.gz',
    model: RATINGS,
  }),
};

module.exports = {
  tables,
};
