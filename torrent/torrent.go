package torrent

import (
	"os"

	bencode "github.com/jackpal/bencode-go"
)
type bencodeInfo struct {
    Pieces      string `bencode:"pieces"`
    PieceLength int    `bencode:"piece length"`
    Length      int    `bencode:"length"`
    Name        string `bencode:"name"`
}

type bencodeFile struct {
    Announce string      `bencode:"announce"`
    Info     bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}
func Open(path string)  bencodeFile{
	file,err := os.Open(path)
	
	if err != nil{
		os.Exit(1)
	}
	defer file.Close()

	BFile := bencodeFile{}
	err = bencode.Unmarshal(file, &BFile)
	if err != nil {
		os.Exit(1)
	}
	return BFile
}


func DownloadFile()  {
	
}