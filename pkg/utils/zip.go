package utils

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ZipFolder(origin string, zipDest string) error{
	// Create a buffer to write our archives to
    zipFile, err := os.Create(zipDest)
    if err != nil{
        return err
    }

    // Create a new zip archive
    zipWriter := zip.NewWriter(zipFile)

    // Iterate through every file in directory
    err = filepath.Walk(origin, func(path string, info fs.FileInfo, err error) error {
        if err != nil{
            return err
        }

        // Catch first iteration
        if path == origin{
            return nil
        }
        
        // Remove the destiny of path
        pathInZip := strings.Replace(path, strings.Replace(origin, "./", "", 1) + "/", "", 1)       

        // Check if path goes to a directoy
        if (info.IsDir()){
            _, err := zipWriter.Create(pathInZip + "/")
            if (err != nil){
                return err
            }   
            return nil
        }   
        // Create zip Writer
        zipFileWriter, err := zipWriter.Create(pathInZip)
        if (err != nil){
            return err
        }   
        

        // Open file
        fileDescriptor, err := os.Open(path)
        if (err != nil){
            return err
        }

        // Copy content of file to zipfile
        _, err = io.Copy(zipFileWriter, fileDescriptor)
        if (err != nil){
            return err
        }
        return nil
    })
    if err != nil{
        return err
    }

    err = zipWriter.Close()
    if err != nil{
        return err
    }

    err = zipFile.Close()
    if err != nil{
        return err
    }

    return nil
}

func UnzipToFolder(origin string, destination string) error{

    // Open zip file
    zipFile, err := zip.OpenReader(origin)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    // Iterate through the files in the archive,
    for _, fileInsideZip := range zipFile.File {
        if fileInsideZip.FileInfo().IsDir() {
            os.MkdirAll(filepath.Join(destination, fileInsideZip.Name), 0777)   
            log.Println(fileInsideZip.Open())
        } else {
            //Open file inside zip (content)
            rc, err := fileInsideZip.Open()
            if err != nil {
                rc.Close()
                return err
            }

            //Create file
            newfile, err := os.OpenFile(filepath.Join(destination, fileInsideZip.Name), os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0777)
            if err != nil{
                rc.Close()
                return err
            }

            _, err = io.Copy(newfile, rc)
            if err != nil {
                rc.Close()
                return err
            }
            rc.Close()
        }
    }
    return nil
}
