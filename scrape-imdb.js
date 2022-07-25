require("dot-env");
const { PG_USER, PG_PASS } = process.env;
const { Sequelize, Model, DataTypes } = require("sequelize");
const Table = require("./table.js");

const sequelize = new Sequelize("imdb", PG_USER, PG_PASS, {
  host: "localhost",
  dialect: "postgres",
  logging: false,
});

const MODELS = {
  BASICS_NAME: sequelize.define("basics_name", {
    nconst: {
      type: DataTypes.TEXT,
      allowNull: false,
      primaryKey: true,
    },
    primaryName: {
      type: DataTypes.TEXT,
    },
    birthYear: {
      type: DataTypes.INTEGER,
    },
    deathYear: {
      type: DataTypes.INTEGER,
    },
    primaryProfession: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
    knownForTitles: {
      type: DataTypes.ARRAY(Sequelize.TEXT),
    },
  }),
  BASICS_TITLE: sequelize.define("basics_title", {
    tconst: {
      type: DataTypes.TEXT,
      allowNull: false,
      primaryKey: true,
    },
    titleType: {
      type: DataTypes.TEXT,
    },
    primaryTitle: {
      type: DataTypes.TEXT,
    },
    originalTitle: {
      type: DataTypes.TEXT,
    },
    isAdult: {
      type: DataTypes.BOOLEAN,
    },
    startYear: {
      type: DataTypes.INTEGER,
    },
    endYear: {
      type: DataTypes.INTEGER,
    },
    runtimeMinutes: {
      type: DataTypes.INTEGER,
    },
    genres: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
  }),
};

const tables = {
  basics_name: new Table("basics_name", {
    url: "https://datasets.imdbws.com/name.basics.tsv.gz",
    model: MODELS.BASICS_NAME,
  }),
  basics_title: new Table("basics_title", {
    url: "https://datasets.imdbws.com/title.basics.tsv.gz",
    model: MODELS.BASICS_TITLE,
  }),
  akas: {
    title: {
      url: "https://datasets.imdbws.com/title.akas.tsv.gz",
      model: null,
    },
  },
  crew: {
    title: {
      url: "https://datasets.imdbws.com/title.crew.tsv.gz",
      model: null,
    },
  },
  episode: {
    title: {
      url: "https://datasets.imdbws.com/title.episode.tsv.gz",
      model: null,
    },
  },
  principals: {
    title: {
      url: "https://datasets.imdbws.com/title.principals.tsv.gz",
      model: null,
    },
  },
  ratings: {
    title: {
      url: "https://datasets.imdbws.com/title.ratings.tsv.gz",
      model: null,
    },
  },
};

module.exports = {
  sequelize,
  tables,
  MODELS,
};
