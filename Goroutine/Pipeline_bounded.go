/*
 md5 checksum 프로그램을 파이프라인으로 변경 (3단계)

 Pipeline_parallel 프로그램은 파일이 매우 많고 클 경우,
 가능한 것 보다 더 많은 메모리를 할당할 수 있다는 문제가 있기 때문에
 여기선 고정된 수의 고루틴만 생성함으로써 병렬로 읽을 파일 수를 제한한다.
*/

package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

/*
 파이프라인의 첫번째 스테이지.
 root로부터 파일 트리를 순회하여 일반 파일들의 path를 채널로 전송한다.
*/
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

/*
 파이프라인의 두번째 스테이지.
 paths 채널로부터 파일 이름을 받고, 처리한 result 값을 c채널로 보낸다.

 walkFiles와는 달리 여러개의 고루틴이 c채널에 접근하므로 여기서 c채널을 닫지 않고
 MD5All에서 모든 digester 고루틴이 끝났을 때 c채널을 닫도록 한다.
*/
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

/*
 파이프라인의 마지막 스테이지.
 walkFiles로부터 각 파일의 path값을 받아 digester에 넘겨주고
 digester로부터 각 파일의 result값을 받아서 map으로 만든다.
*/
func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c := make(chan result)

	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)

	paths, errc := walkFiles(done, root)

	// 정해놓은 상수만큼만(20번) digester 고루틴을 실행함
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	// 20개의 고루틴이 종료되면 c 채널을 닫음
	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)

	// 각 digester가 자신만의 output 채널을 생성하고 리턴하는 것이 아니라
	// c라는 하나의 채널로 결과값을 fan-in 함
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// 파일 트리 순회에 실패하진 않았는지 체크
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

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
