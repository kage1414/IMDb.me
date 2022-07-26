require('dot-env');
const {
  IMDB_DB_USER,
  IMDB_DB_PASS,
  IMDB_DB_DMS,
  IMDB_DB_HOST,
} = process.env;
const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize(
  'imdb',
  IMDB_DB_USER,
  IMDB_DB_PASS,
  {
    host: IMDB_DB_HOST || 'localhost',
    dialect: IMDB_DB_DMS || 'postgres',
    logging: false,
  }
);

const models = {
  NAME: sequelize.define('basics_name', {
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
  TITLE: sequelize.define('basics_title', {
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
  AKAS: sequelize.define('akas', {
    titleId: {
      type: DataTypes.TEXT,
    },
    ordering: {
      type: DataTypes.INTEGER,
    },
    title: {
      type: DataTypes.TEXT,
    },
    region: {
      type: DataTypes.TEXT,
    },
    language: {
      type: DataTypes.TEXT,
    },
    types: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
    attributes: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
    isOriginalTitle: {
      type: DataTypes.BOOLEAN,
    },
  }),
  CREW: sequelize.define('crew', {
    tconst: {
      type: DataTypes.TEXT,
      primaryKey: true,
    },
    directors: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
    writers: {
      type: DataTypes.ARRAY(DataTypes.TEXT),
    },
  }),
  EPISODE: sequelize.define('episode', {
    tconst: {
      type: DataTypes.TEXT,
      primaryKey: true,
    },
    parentTconst: {
      type: DataTypes.TEXT,
    },
    seasonNumber: {
      type: DataTypes.INTEGER,
    },
    episodeNumber: {
      type: DataTypes.INTEGER,
    },
  }),
  PRINCIPALS: sequelize.define('principals', {
    tconst: {
      type: DataTypes.TEXT,
    },
    ordering: {
      type: DataTypes.INTEGER,
    },
    nconst: {
      type: DataTypes.TEXT,
    },
    category: {
      type: DataTypes.TEXT,
    },
    job: {
      type: DataTypes.TEXT,
    },
    characters: {
      type: DataTypes.TEXT,
    },
  }),
  RATINGS: sequelize.define('ratings', {
    tconst: {
      type: DataTypes.TEXT,
    },
    averageRating: {
      type: DataTypes.FLOAT,
    },
    numVotes: {
      type: DataTypes.INTEGER,
    },
  }),
};

models.NAME.hasMany(models.AKAS, {
  foreignKey: 'titleId',
});

models.NAME.belongsTo(models.CREW, {
  foreignKey: 'tconst',
});

module.exports = {
  models,
  sequelize,
};