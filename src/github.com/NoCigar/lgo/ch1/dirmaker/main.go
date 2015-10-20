package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	askAndCreateDirs()

}

func askAndCreateDirs() {

	var folderName string
	var n int

	folderName = scanDirName()
	n = scanDirCount()

	createDirWithPrefix(folderName, n)
}
//asks the user for the directory archetype to be created
func scanDirName() (folderName string) {
	fmt.Println("Gieb folder name: \n")
	_, err := fmt.Scanf("%s", &folderName)

	if err != nil {
		fmt.Println(err)
	}

	return folderName
}
//asks the user for the amount of directories to be created
func scanDirCount() (n int) {

	fmt.Print("Gieb the amount of folders you want to create," +
		"\n(During creation, existing folders will be skipped): ")

	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nThe number of folders to be created is: %d", n)

	return n
}
//creates directories in the selected target directory
func createDirWithPrefix(folderName string, n int) {

	var targetDir string

	var workingDir string = getWorkingDir()
	targetDir = workingDir + "\\" + folderName

	for i := 1; i <= n; i++ {
		newDir := targetDir + strconv.Itoa(i)
		createDir(newDir)
	}

}

//creates a single directory
func createDir(targetDir string) {

	err := os.Mkdir(targetDir, os.ModeDir)
	println("\nWP from create Dir: "+targetDir)
	if err != nil {
		fmt.Println(err)
	}
}
//returns the working(target) directory
func getWorkingDir() (dir string){

	var workingDir string

	fmt.Println("\nGieb working dir: ")
	fmt.Scanf("%s",&workingDir)
	return workingDir
}
