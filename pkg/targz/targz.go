package targz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Extract(r io.Reader, targetDir string) error {
	gzipReader, err := gzip.NewReader(r)

	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed\n", err)
	}

	tarReader := tar.NewReader(gzipReader)
	gzipReader.Close()

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		path := filepath.Join(targetDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg, tar.TypeGNUSparse:
			// tar.Next() will externally only iterate files, so we might have to create intermediate directories here
			if err := untarFile(tarReader, header, path); err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, path); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown file type: %s in %s", string(header.Typeflag), header.Name)
		}
	}

	return nil
}

func untarFile(tarReader *tar.Reader, header *tar.Header, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, header.FileInfo().Mode())
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, tarReader)

	return err
}
