const { tables } = require('./db.js');
const { sequelize, models } = require('./models.js');
const firstline = require('firstline');
const axios = require('axios');
const gunzip = require('gunzip-file');
const Table = require('./table.js');
const replace = require('replace-in-file');
const readline = require('readline');
const path = require('path');
const { Client } = require('pg');
const {
  startTimer,
  addToInsertTotal,
  countTotalRecords,
  logProgress,
} = require('./util.js');
const {
  createReadStream,
  createWriteStream,
  existsSync,
  mkdirSync,
  rmdirSync,
  unlink,
} = require('fs');
const arrayFields = [
  'genres',
  'knownForTitles',
  'primaryProfession',
  'attributes',
  'types',
  'directors',
  'writers',
];
const booleanFields = ['isAdult', 'isOriginalTitle'];

const importToPostgres = async (
  client,
  { model, name },
  filePath
) => {
  const fullPath = path.resolve(filePath);
  const firstLine = await firstline(filePath, {
    encoding: 'utf8',
  });
  const mutatedFirstLine = firstLine
    .split('\t')
    .map((ele) => `"${ele}"`)
    .join(',');
  const queryString = `COPY ${model.getTableName()}(${mutatedFirstLine}) FROM '${fullPath}' DELIMITER '\t' CSV HEADER`;
  console.log(queryString);
  await client.query(queryString);
  console.log('Imported', name);
  unlink(filePath, () => {});
};

const removeNulls = async (
  filePath,
  { deleteExtraFiles },
  cb
) => {
  const arrayIdx = [];
  const booleanIdx = [];
  let firstLine = await firstline(filePath, {
    encoding: 'utf8',
  });
  firstLine = firstLine.split('\t');
  firstLine.forEach((header, idx) => {
    if (arrayFields.includes(header)) {
      arrayIdx.push(idx);
    }
    if (booleanFields.includes(header)) {
      booleanIdx.push(idx);
    }
  });
  let newFilePath = filePath.split('.');
  newFilePath.splice(newFilePath.length - 1, 0, 'import');
  newFilePath = newFilePath.join('.');
  const read = createReadStream(filePath, {
    encoding: 'utf8',
  });
  const write = createWriteStream(newFilePath);
  const rl = readline.createInterface({
    input: read,
    crlfDelay: Infinity,
  });
  let onFirstLine = true;
  const beginLogging = false;
  for await (const line of rl) {
    const split = line.split('\t').map((cell, idx) => {
      if (cell === '\\N' || cell === '\\\\N') {
        cell = '';
      }
      if (booleanIdx.includes(idx) && !onFirstLine) {
        cell = !!cell;
      }
      if (arrayIdx.includes(idx) && !onFirstLine) {
        const array = cell.split(',');
        let arrayLiteral = JSON.stringify(array);
        arrayLiteral = arrayLiteral.replace('[', '{');
        arrayLiteral = arrayLiteral.replace(']', '}');
        return arrayLiteral;
      }
      return cell;
    });

    write.write(split.join('\t'));
    write.write('\n');
    if (onFirstLine) {
      onFirstLine = false;
    }
  }
  read.on('close', () => {
    write.close();
    if (deleteExtraFiles) {
      unlink(filePath, () => {});
    }
    cb(newFilePath);
  });
};

const download = async ({
  url,
  zippedFileName,
  fileName,
  model,
  cb,
}) => {
  const zippedOutputLocationPath = path.join(
    '.',
    'temp',
    zippedFileName
  );
  const outputLocationPath = path.join(
    '.',
    'temp',
    fileName
  );
  const writer = createWriteStream(
    zippedOutputLocationPath
  );
  await axios({
    url,
    method: 'GET',
    responseType: 'stream',
  }).then(({ data }) => {
    return new Promise((resolve, reject) => {
      data.pipe(writer);
      let error = null;
      writer.on('error', (err) => {
        error = err;
        writer.close();
        reject(err);
      });
      writer.on('close', () => {
        if (!error) {
          resolve(true);
        }
      });
    });
  });

  await gunzip(
    zippedOutputLocationPath,
    outputLocationPath,
    () => {
      unlink(zippedOutputLocationPath, () => {});
      removeNulls(
        outputLocationPath,
        {
          deleteExtraFiles: true,
        },
        cb
      );
    }
  );
};

const runner = async () => {
  console.log('Cloning IMDB Database');
  const client = new Client({
    user: process.env.IMDB_USER,
    host: 'localhost',
    database: 'imdb',
    password: process.env.IMDB_PASS,
  });
  await client.connect();
  await sequelize.sync();
  if (existsSync('temp')) {
    rmdirSync('temp', { recursive: true });
  }
  mkdirSync('temp');
  console.log('Downloading records');
  const tablesToScrape = Object.values(tables).filter(
    (table) => table instanceof Table
  );
  tablesToScrape.forEach(async (table) => {
    download({
      ...table,
      cb: (filePath) => {
        importToPostgres(client, table, filePath);
      },
    });
  });
};

runner();
