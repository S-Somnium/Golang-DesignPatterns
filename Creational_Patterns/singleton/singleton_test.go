package singleton

import "testing"

func TestMusicBox(t *testing.T) {
	pcSpotify := GetMusicBox()
	pcSpotify.queueMusic("nirvana").queueMusic("imagine dragons").queueMusic("ACDC")
	smartPhoneSpotify := GetMusicBox()
	smartPhoneSpotify.queueMusic("avengend sevenfold")
	if pcSpotify.queue[len(pcSpotify.getQueue())-1] != "avengend sevenfold" {
		t.Fatal("Expected avenged sevenfold but got instead ", pcSpotify.queue[len(pcSpotify.getQueue())-1])
	}
}
