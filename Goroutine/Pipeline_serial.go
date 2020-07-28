/* md5 checksum 예시 */

package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

/*
 root로부터 파일 트리를 순회하며 모든 일반 파일들을 읽고 map을 리턴한다.
 map의 key는 파일의 path, value는 파일의 MD5 체크섬 값이다.
 디렉토리 탐색에 실패하거나 파일을 읽는 데 실패한다면 에러를 리턴한다.
*/
func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

/*
 매개변수로 받은 디렉토리 하위의 모든 일반 파일에 대해 MD5 체크섬을 계산하고,
 계산 후의 결과값을 path 이름별로 정렬하여 출력한다.
*/
func main() {
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}
