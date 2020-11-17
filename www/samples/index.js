import { Buffer } from "buffer";
import samples from "./samples.json";

export const getSamples = async ({ offset = 0, limit = 5 }) => {
  return samples.slice(offset, offset + limit).map(({ content, ...rest }) => ({
    ...rest,
    content: base64Decoding(content),
  }));
};

const base64Decoding = (data) => {
  let buff = new Buffer(data, "base64");
  return buff.toString("ascii");
};
