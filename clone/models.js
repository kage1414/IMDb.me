require('dotenv').config();
const { IMDB_USER, IMDB_PASS, IMDB_DBMS, IMDB_HOST } =
  process.env;
const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize(
  'imdb',
  IMDB_USER,
  IMDB_PASS,
  {
    host: IMDB_HOST || 'localhost',
    dialect: IMDB_DBMS || 'postgres',
    logging: false,
  }
);

const models = {
  NAME: sequelize.define(
    'name',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
      nconst: {
        type: DataTypes.TEXT,
        allowNull: false,
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
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['nconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  TITLE: sequelize.define(
    'title',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
      tconst: {
        type: DataTypes.TEXT,
        allowNull: false,
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
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['tconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  AKAS: sequelize.define(
    'akas',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
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
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['titleId'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  CREW: sequelize.define(
    'crew',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
      tconst: {
        type: DataTypes.TEXT,
      },
      directors: {
        type: DataTypes.ARRAY(DataTypes.TEXT),
      },
      writers: {
        type: DataTypes.ARRAY(DataTypes.TEXT),
      },
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['tconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  EPISODE: sequelize.define(
    'episode',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
      tconst: {
        type: DataTypes.TEXT,
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
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['tconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  PRINCIPALS: sequelize.define(
    'principals',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
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
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['tconst', 'nconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
  RATINGS: sequelize.define(
    'ratings',
    {
      // id: {
      //   type: DataTypes.UUID,
      //   defaultValue: DataTypes.UUIDV4,
      //   primaryKey: true,
      // },
      tconst: {
        type: DataTypes.TEXT,
      },
      averageRating: {
        type: DataTypes.FLOAT,
      },
      numVotes: {
        type: DataTypes.INTEGER,
      },
      createdAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
      updatedAt: {
        type: DataTypes.DATE,
        allowNull: true,
      },
    },
    {
      indexes: [
        {
          unique: false,
          fields: ['tconst'],
        },
        {
          unique: true,
          fields: ['id'],
        },
      ],
    }
  ),
};

// models.NAME.hasMany(models.AKAS, {
//   foreignKey: 'titleId',
//   onDelete: 'cascade',
// });

// models.NAME.belongsTo(models.CREW, {
//   foreignKey: 'tconst',
//   onDelete: 'cascade',
// });

module.exports = {
  models,
  sequelize,
};
