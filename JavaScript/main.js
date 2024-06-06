// @ts-check

const QRCode = require('qrcode');
const Jimp = require('jimp');
const sharp = require('sharp');
const fs = require('fs')

async function convertToMonochrome(inputPath, outputPath) {

  // 读取图像
  const image = await Jimp.read(inputPath);

  await sharp(inputPath, {
    raw: {
      // width: image.bitmap.width,
      // height: image.bitmap.height,
      width: 1024,
      height: 1024,
      channels: 1
    }
  }).toColourspace('b-w').toFile(outputPath);
}

(async () => {
  await QRCode.toFile('myqr.png', 'https://intelligent-book-wmm01-e.suanshubang.cc/ibserver/video.html?nodeId=661032', {
    width: 1024,
  })

  await convertToMonochrome('myqr.png', 'myqr_monochrome.png')
  await convertToMonochrome('myqr.png', 'myqr_monochrome.tif')

  fs.rmSync('myqr.png')
})()

