/* md5 checksum 프로그램을 파이프라인으로 변경 (2단계) */

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
 root로부터 파일 트리를 새로운 고루틴에서 순회하고 result형 값을 채널로 전송한다.
 반환값은 result형과 error형 채널이며, error는 filepath.Walk 과정에서 나온 에러이다.

 filepath.Walk의 인자 함수에선 또다른 고루틴을 실행시켜 각 파일을 처리한다.
 그 후 done 채널을 체크하여, done 채널이 닫히면 즉시 순회를 종료한다.
*/
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
					// 각 일반 파일에 대한 md5 체크섬 값을 c 채널로 전송
				case <-done:
					// done 채널이 닫히면 어떤 처리도 하지 않음
				}
				wg.Done()
			}()
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		go func() {
			// c 채널에 대한 모든 전송이 종료되면 c 채널을 닫음
			wg.Wait()
			close(c)
		}()
		// walk 과정에서 발생한 에러를 errc 채널로 전송
		// errc 채널은 1 크기의 버퍼를 가지고 있기 때문에 select는 필요 없음
		errc <- err
	}()
	return c, errc
}

/*
 파이프라인의 두번째 스테이지.
 sumFiles로부터 각 파일의 result값을 받아서 map으로 만든다.
 함수가 종료되거나 에러가 발생해 일찍 중단하게 되면 defer를 통해 done 채널을 닫아
 이전 단계도 종료될 수 있도록 한다.
*/

func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
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
