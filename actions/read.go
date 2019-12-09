package actions

import (
	"bufio"
	"fmt"
	"os"

	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

const readFileName = "test"

func Read(c *cli.Context) error {
	fname := c.String("file")
	if fname == "" {
		fname = readFileName
	}
	file, _ := os.Open(fmt.Sprintf("data/%s.mid", fname))
	defer file.Close()

	midi, err := smfio.Read(bufio.NewReader(file))
	if err != nil {
		return err
	}

	track := midi.GetTrack(0)
	fmt.Println(midi.GetTracksNum())

	iter := track.GetIterator()
	for iter.MoveNext() {
		fmt.Println(iter.GetValue())
	}
	return nil
}
