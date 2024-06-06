const fs = require('fs');


// 生成一个5G的文件

const fileSize = 5 * 1024 * 1024 * 1024; // 5G

const file = fs.createWriteStream('largeFile.bin');

let totalBytesWritten = 0;

const writeData = () => {

  const data = Buffer.alloc(1024 * 1024); // 1MB
  file.write(data);
  totalBytesWritten += data.length;
}

const writeLoop = setInterval(() => {
  if (totalBytesWritten >= fileSize) {
    clearInterval(writeLoop);
    file.end();
    console.log('文件生成完成');
  } else {
    writeData();
  }
})