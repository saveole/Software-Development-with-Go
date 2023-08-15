package main

import (
	"log"

	s "syscall"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid:", s.Getpid())
	log.Println("Getpgrp:", s.Getpgrp())
	log.Println("Getppid:", s.Getppid())
	log.Println("Gettid:", s.Gettid())

	_, err := s.Getcwd(c)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(c))

	var statfs = s.Statfs_t{}
	var total, used, free uint64;
	err = s.Statfs("/", &statfs)
	if err != nil {
		log.Fatalln(err)
	} else {
		total = statfs.Blocks * uint64(statfs.Bsize)
		free = statfs.Bfree * uint64(statfs.Bsize)
		used = total - free
	}
	log.Println("total:", total)
	log.Println("used:", used)
	log.Println("free:", free)

}