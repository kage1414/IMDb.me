const countdown = require('countdown');
const { createReadStream, unlink } = require('fs');
let start;
let totalRecords = 0;
let insertedTotal = 0;

module.exports = {
  logProgress: () => {
    const durationMilliseconds = Date.now() - start;
    const rate = insertedTotal / durationMilliseconds;
    const recordsLeft = totalRecords - insertedTotal;
    const estimatedFinish = Math.round(
      Date.now() + recordsLeft / rate
    );
    const totalRecordsMessage = `Total records: ${totalRecords}`;
    const insertMessage = `Inserted ${insertedTotal} records of ${totalRecords} - ${Math.round(
      insertedTotal / totalRecords
    )}% - ${Math.round(
      rate * 1000
    )} records/sec - ${countdown(
      null,
      new Date(estimatedFinish)
    ).toString()} left`;
    console.clear();
    console.log(totalRecordsMessage);
    console.log(insertMessage);
  },
  startTimer: () => {
    if (!start) {
      start = Date.now();
    }
  },
  addToInsertTotal: (num) => {
    insertedTotal += num;
  },
  countTotalRecords: (
    filePath,
    { deleteFile, deleteFilePath }
  ) => {
    const stream = createReadStream(filePath, {
      encoding: 'utf8',
    });
    let headersRemoved;
    let left;
    stream.on('data', (chunk) => {
      const chunkArr = chunk.split('\n');
      if (!headersRemoved) {
        chunkArr.shift();
        headersRemoved = true;
      }
      if (left) {
        chunkArr[0] = left + chunkArr[0];
      }
      left = chunkArr.pop();
      totalRecords += chunkArr.length;
      module.exports.logProgress();
    });
    stream.on('close', () => {
      if (deleteFile) {
        unlink(filePath, () => {});
        unlink(deleteFilePath, () => {});
      }
    });
  },
};
