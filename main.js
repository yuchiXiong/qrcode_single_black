// @ts-check

const QRCode = require('qrcode');
const Jimp = require('jimp');
const sharp = require('sharp');

async function convertToMonochrome(inputPath, outputPath) {

  // 读取图像
  const image = await Jimp.read(inputPath);

  await sharp(inputPath, {
    raw: {
      width: image.bitmap.width,
      height: image.bitmap.height,
      channels: 1
    }
  }).toColourspace('b-w').toFile(outputPath);
}

(async () => {
  await QRCode.toFile('myqr.png', 'https://www.google.com')

  await convertToMonochrome('myqr.png', 'myqr_monochrome.png')
  await convertToMonochrome('myqr.png', 'myqr_monochrome.tif')

})()

