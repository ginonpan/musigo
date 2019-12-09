package actions

import (
	"bufio"
	"fmt"
	"os"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

const createFileName = "test"

func Create(c *cli.Context) error {
	file := c.String("output")
	if file == "" {
		file = createFileName
	}

	division, err := smf.NewDivision(960, smf.NOSMTPE)
	if err != nil {
		return err
	}

	midi, err := smf.NewSMF(smf.Format0, *division)
	if err != nil {
		return err
	}

	track := &smf.Track{}
	err = midi.AddTrack(track)
	if err != nil {
		return err
	}

	midiEventOne, err := smf.NewMIDIEvent(0, smf.NoteOnStatus, 0x00, 0x30, 0x50)
	if err != nil {
		return err
	}
	midiEventTwo, err := smf.NewMIDIEvent(10000, smf.NoteOnStatus, 0x00, 0x30, 0x00)
	if err != nil {
		return err
	}
	metaEventOne, err := smf.NewMetaEvent(0, smf.MetaEndOfTrack, []byte{})
	if err != nil {
		return err
	}

	// Add created events to track
	if err := track.AddEvent(midiEventOne); err != nil {
		return err
	}
	if err := track.AddEvent(midiEventTwo); err != nil {
		return err
	}
	if err := track.AddEvent(metaEventOne); err != nil {
		return err
	}

	outputMidi, err := os.Create(fmt.Sprintf("data/%s.mid", file))
	if err != nil {
		return err
	}
	defer outputMidi.Close()

	writer := bufio.NewWriter(outputMidi)
	smfio.Write(writer, midi)
	return writer.Flush()
}
