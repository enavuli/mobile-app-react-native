import * as fs from 'fs';
import * as path from 'path';

function parseJsFile(filePath) {
  const content = fs.readFileSync(filePath, 'utf8');
  const file = /function\s+([a-zA-Z_$][a-zA-Z_$0-9]*)\s*\(([^)]*)\)\s*{([^}]*)}/g.exec(content);
  if (file) {
    const functionName = file[1];
    const params = file[2].split(',').map((param) => param.trim());
    const body = file[3];
    return { functionName, params, body };
  } else {
    return null;
  }
}

function parseJsProject(rootDir) {
  const files = [];
  const filesToParse = fs.readdirSync(rootDir).filter((file) => {
    const filePath = path.join(rootDir, file);
    return fs.statSync(filePath).isDirectory();
  }).map((file) => path.join(rootDir, file));
  filesToParse.forEach((file) => {
    const dirFiles = fs.readdirSync(file);
    dirFiles.filter((dirFile) => dirFile.endsWith('.js')).forEach((jsFile) => {
      const jsFilePath = path.join(file, jsFile);
      const parsedFile = parseJsFile(jsFilePath);
      if (parsedFile) {
        files.push(parsedFile);
      }
    });
  });
  return files;
}

export { parseJsFile, parseJsProject };