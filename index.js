const { tables, sequelize, MODELS } = require("./scrape-imdb.js");
const axios = require("axios");
const gunzip = require("gunzip-file");
const Table = require("./table.js");
const path = require("path");
const {
  createReadStream,
  createWriteStream,
  existsSync,
  mkdirSync,
  readFile,
} = require("fs");
const CHUNK_SIZE = 200000;
const arrayFields = ["genres", "knownForTitles", "primaryProfession"];

let insertCount = 0;
let start;

const download = async ({ url, zippedFileName, fileName, model }) => {
  const zippedOutputLocationPath = path.join(".", "temp", zippedFileName);
  const outputLocationPath = path.join(".", "temp", fileName);
  const writer = createWriteStream(zippedOutputLocationPath);
  await axios({
    url: url,
    method: "GET",
    responseType: "stream",
  }).then(({ data }) => {
    return new Promise((resolve, reject) => {
      data.pipe(writer);
      let error = null;
      writer.on("error", (err) => {
        error = err;
        writer.close();
        reject(err);
      });
      writer.on("close", () => {
        if (!error) {
          resolve(true);
        }
      });
    });
  });

  await gunzip(zippedOutputLocationPath, outputLocationPath, () => {
    parseFileToCSV(outputLocationPath, model);
  });
};

const parseFileToCSV = async (filePath, model) => {
  let headers;
  let left;

  if (!start) {
    start = Date.now();
  }

  const stream = createReadStream(filePath, {
    highWaterMark: CHUNK_SIZE,
    encoding: "utf8",
  });

  stream.on("data", async (chunk) => {
    stream.pause();
    const chunkArr = chunk.split("\n");
    const { length } = chunkArr;
    const durationSeconds = (Date.now() - start) / 1000;
    const rate = insertCount / durationSeconds;
    console.log(
      `Inserting ${length} records. ${insertCount} total records. ${Math.round(
        rate
      )} records/sec. ${durationSeconds} sec.`
    );
    if (!headers) {
      headers = chunkArr.shift().split("\t");
    }
    if (left) {
      chunkArr[0] = left + chunkArr[0];
    }
    left = chunkArr.pop();

    for await (const rowString of chunkArr) {
      const row = Object.fromEntries(
        rowString.split("\t").map((cell, idx) => {
          const key = headers[idx];
          let value = cell;
          if (cell === "\\N") {
            value = null;
          } else if (key === "isAdult") {
            value = !!Number(value);
          } else if (arrayFields.includes(key)) {
            value = cell.split(",");
          }
          return [key, value];
        })
      );
      await model
        .create(row)
        .then(() => {
          insertCount++;
        })
        .catch((err) => {
          console.log("Error:", err);
          console.log(row);
        });
    }
    stream.resume();
  });

  stream.on("close", () => {
    console.log("closed", filePath);
  });
};

const runner = async () => {
  console.log("Seeding");
  await MODELS.BASICS_NAME.drop();
  await MODELS.BASICS_TITLE.drop();
  await sequelize.sync();
  if (!existsSync("temp")) {
    mkdirSync("temp");
  }
  console.log("Downloading records");
  const tablesToScrape = Object.values(tables).filter(
    (table) => table instanceof Table
  );

  tablesToScrape.forEach(download);
};

runner();
