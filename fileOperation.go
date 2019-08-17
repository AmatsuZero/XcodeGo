package XcodeGo

import (
	"./Utils"
	"os"
	"path"
)

type FileOperationQueue struct {
	baseDirectory       string
	filesToWrite        map[string][]byte
	frameworksToCopy    map[string]string
	filesToDelete       []string
	directoriesToCreate []string
}

func NewXCFileOperationQueueWithBaseDirectory(baseDirectory string) FileOperationQueue {
	return FileOperationQueue{
		baseDirectory: baseDirectory,
	}
}

func (t FileOperationQueue) destinationPath(fileName string, folder string) string {
	return path.Join(t.baseDirectory, folder, fileName)
}

func (t FileOperationQueue) FileExistInDir(fileName string, dir string) bool {

	return Utils.IsFileExist(t.destinationPath(fileName, dir))
}

func (t *FileOperationQueue) EnqueueTextFile(fileName string, dir string, contents string) {
	data := []byte(contents)
	t.EnqueueDataFile(fileName, dir, data)
}

func (t *FileOperationQueue) EnqueueDataFile(fileName string, dir string, contents []byte) {
	key := t.destinationPath(fileName, dir)
	t.filesToWrite[key] = contents
}

func (t *FileOperationQueue) EnqueueFramework(filePath string, dir string) {
	dest := t.destinationPath(path.Base(filePath), dir)
	t.frameworksToCopy[dest] = filePath
}

func (t *FileOperationQueue) EnqueueDeletion(path string) {
	t.filesToDelete = append(t.filesToDelete, path)
}

func (t *FileOperationQueue) EnqueueDirect(name string, parentDir string) {
	dir := t.destinationPath(name, parentDir)
	t.directoriesToCreate = append(t.directoriesToCreate, dir)
}

func (t *FileOperationQueue) commitFileOperations() (e error) {
	if err := t.performFileWrites(); err != nil {
		e = err
	}
	if err := t.performCopyFrameworks(); err != nil {
		e = err
	}
	if err := t.performFileDeletions(); err != nil {
		e = err
	}
	if err := t.performCreateDirectories(); err != nil {
		e = err
	}
	return
}

func (t *FileOperationQueue) performFileWrites() error {
	for filePath, buf := range t.filesToWrite {
		fo, err := os.Create(filePath)
		if err != nil {
			return err
		}
		if _, err := fo.Write(buf[:]); err != nil {
			return err
		}
		if err = fo.Close(); err != nil {
			return err
		}
	}
	// 清除键值对
	t.filesToWrite = make(map[string][]byte)
}

func (t *FileOperationQueue) performCopyFrameworks() error {
	for destinationUrl, frameworkPath := range t.frameworksToCopy {
		if Utils.IsFileExist(destinationUrl) {
			if err := os.Remove(destinationUrl); err != nil {
				return err
			}
		}
		if err := Utils.CopyFile(frameworkPath, destinationUrl); err != nil {
			return err
		}
	}
	t.frameworksToCopy = make(map[string]string)
}

func (t *FileOperationQueue) performFileDeletions() error {
	for i := len(t.filesToDelete) - 1; i >= 0; i-- {
		fullPath := path.Join(t.baseDirectory, t.filesToDelete[i])
		if !Utils.IsFileExist(fullPath) {
			continue
		}
		if err := os.Remove(fullPath); err != nil {
			return err
		}
	}
}

func (t *FileOperationQueue) performCreateDirectories() error {
	for _, filePath := range t.directoriesToCreate {
		if !Utils.IsFileExist(filePath) {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
		}
	}
}
