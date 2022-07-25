class Table {
  constructor(name, props) {
    const { url, model } = props;
    this.name = name;
    this.url = url;
    this.model = model;
    this.zippedFileName = url.split("/")[url.split("/").length - 1];
    this.fileName = this.zippedFileName
      .split(".")
      .slice(0, this.zippedFileName.split(".").length - 1)
      .join(".");
  }
}

module.exports = Table;
