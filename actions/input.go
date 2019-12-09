package actions

import (
	"bufio"
	"fmt"
	"musigo/values"
	"os"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

const (
	onDeltaTime   = 800
	offDeltaTime  = 80
	ticks         = 960
	inputFileName = "test"
)

// Input Keys or Get Strings convert MIDI
func Input(c *cli.Context) error {
	var file string
	fmt.Println("Enter the output file name")
	fmt.Scan(&file)
	if file == "" {
		file = inputFileName
	}

	var str string
	fmt.Println("Enter the music you want to play")
	fmt.Scan(&str)

	var score []uint8
	for _, t := range str {
		score = append(score, values.Scale(fmt.Sprintf("%c", t)))
	}

	division, err := smf.NewDivision(ticks, smf.NOSMTPE)
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

	var list []*smf.MIDIEvent
	for i, t := range score {
		var d uint32
		if i != 0 {
			d = onDeltaTime
		}
		toneOn, err := smf.NewMIDIEvent(d, smf.NoteOnStatus, 0x00, t, 0x64)
		if err != nil {
			return err
		}
		list = append(list, toneOn)
		toneOff, err := smf.NewMIDIEvent(offDeltaTime, smf.NoteOffStatus, 0x00, t, 0x64)
		if err != nil {
			return err
		}
		list = append(list, toneOff)
	}

	for _, l := range list {
		if err := track.AddEvent(l); err != nil {
			return err
		}
	}

	metaEvent, err := smf.NewMetaEvent(21, smf.MetaEndOfTrack, []byte{})
	if err != nil {
		return err
	}
	if err := track.AddEvent(metaEvent); err != nil {
		return err
	}

	outputMidi, err := os.Create(fmt.Sprintf("./%s.mid", file))
	if err != nil {
		return err
	}
	defer outputMidi.Close()

	writer := bufio.NewWriter(outputMidi)
	smfio.Write(writer, midi)
	return writer.Flush()
}
