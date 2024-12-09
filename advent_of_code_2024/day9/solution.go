package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() DiskMap {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dm := DiskMap{}

	for scanner.Scan() {
		dm.content += scanner.Text()
	}

	return dm
}

type DiskMap struct{ content string }

// segment layout

type Layout struct{ content []*int }

func (dm DiskMap) asLayout() Layout {
	layout := Layout{}
	currentFileID := 0

	for idx, r := range dm.content {
		segmentLen, _ := strconv.Atoi(string(r))

		if idx%2 == 0 {
			// file
			id := currentFileID
			for i := 0; i < segmentLen; i++ {
				layout.content = append(layout.content, &id)
			}
			currentFileID++
		} else {
			// empty
			for i := 0; i < segmentLen; i++ {
				layout.content = append(layout.content, nil)
			}
		}
	}

	return layout
}

func (layout Layout) moveBlocks() {
	l, r := 0, len(layout.content)-1

	for {
		// skip filled blocks
		for (layout.content)[l] != nil {
			l++
		}

		// skip empty at the end
		for (layout.content)[r] == nil {
			r--
		}

		if l >= r {
			break
		}

		layout.content[l] = layout.content[r]
		layout.content[r] = nil
	}
}

func (layout Layout) print() {
	for _, v := range layout.content {
		if v == nil {
			fmt.Print(".")
		} else {
			fmt.Print(*v)
		}
	}
	fmt.Println()
}

func (layout Layout) checksum() int64 {
	var sum int64

	for i, val := range layout.content {
		if val == nil {
			continue
		}
		sum += int64(i * (*val))
	}
	return sum
}

// File layout

type FileLayoutSegment struct {
	fileId *int
	length int
}

func (fls FileLayoutSegment) isFile() bool {
	return fls.fileId != nil
}

type FileLayout struct{ content []FileLayoutSegment }

func (dm DiskMap) asFileLayout() FileLayout {
	layout := FileLayout{}
	currentFileID := 0

	for idx, r := range dm.content {
		segmentLen, _ := strconv.Atoi(string(r))

		fls := FileLayoutSegment{length: segmentLen}

		if idx%2 == 0 {
			// file
			fileId := currentFileID
			fls.fileId = &fileId
			currentFileID++
		}

		layout.content = append(layout.content, fls)
	}

	return layout
}

func (layout FileLayout) moveFiles() {
	r := len(layout.content) - 1

	// layout.print()

	for {
		// find rightmost file
		for r >= 0 && !layout.content[r].isFile() {
			r--
		}

		if r < 0 {
			break
		}

		fileToMove := layout.content[r]

		// find the block that fits
		l := 0
		for l < r {
			if layout.content[l].isFile() == false && layout.content[l].length >= fileToMove.length {
				segment := layout.content[l]
				remainingFreeSpace := segment.length - fileToMove.length

				if remainingFreeSpace == 0 {
					// file fits perfectly
					layout.content[l] = layout.content[r]
					layout.content[r].fileId = nil
				} else {
					// split segment
					layout.content[l].fileId = layout.content[r].fileId
					layout.content[l].length = layout.content[r].length
					layout.content[r].fileId = nil

					layout.content = append(layout.content[:l+1], layout.content[l:]...)
					layout.content[l+1] = FileLayoutSegment{length: remainingFreeSpace}
				}

				break
			}

			l++
		}

		r--
	}
}

func (layout FileLayout) print() {
	for _, s := range layout.content {
		if s.isFile() {
			for i := 0; i < s.length; i++ {
				fmt.Print(*s.fileId)
			}
		} else {
			for i := 0; i < s.length; i++ {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}

func (layout FileLayout) checksum() int64 {
	var sum int64

	currentIdx := 0
	for _, val := range layout.content {
		for i := 0; i < val.length; i++ {
			if val.isFile() {
				sum += int64(currentIdx * *val.fileId)
			}
			currentIdx++
		}
	}
	return sum
}

func main() {
	diskMap := getInputData()

	layout := diskMap.asLayout()
	layout.moveBlocks()
	fmt.Println("Part 1 solution is", layout.checksum())

	fileLayout := diskMap.asFileLayout()
	fileLayout.moveFiles()
	fmt.Println("Part 2 solution is", fileLayout.checksum())
}
